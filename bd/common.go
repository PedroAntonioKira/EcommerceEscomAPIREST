package bd

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"os"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/secretsmanager"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretsmanager.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))

	//Verificamos que todo este bien y no se produzca un error

	//Verificamos que nos conectamos a la bd
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()

	//Verificamos que podamos enviar ping (haya conección) a la bd
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión Exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "ecommerceEscom" //Nombre que aparece en HeidiSQL.
	//dsName
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		dbUser,
		authToken,
		dbEndpoint,
		dbName)

	fmt.Println(dsn)

	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza UserIsAdmin")

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return false, err.Error()
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer Db.Close()

	//Declaramos la sentencia SQL para buscar el usuario
	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"

	//Imprimimos la sentencia para ver la misma en cloudwatch y poder servir de ayuda para depurar el codigo.
	fmt.Println(sentencia)

	//Vamos a ejecutar la sentencia SQL
	rows, err := Db.Query(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		return false, err.Error()
	}

	//Variable para guardar la info del registro devuelto a traves de la consulta.
	var valor string

	//Nos posicionamos en el primer registro devuelto por la consulta
	rows.Next()

	//Leeremos los datos a traves de un scan y lo debemos guardar en una variable
	rows.Scan(&valor)

	fmt.Println("UserAdmin > Ejecución Exitosa !! - valor devuelto " + valor)

	//Verifiamos si valor nos indica que es admin
	if valor == "1" {
		return true, ""
	}

	return false, "User is not Admin"
}
