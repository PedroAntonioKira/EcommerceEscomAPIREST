package bd

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza Registro de InsertCategory")

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return 0, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer Db.Close()

	//Declaramos la sentencia SQL para insertar la categoria
	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	//Nos ayudara a guardar el resultado cuando ejecutemos la sentencia SQL (trae filas afectadas y ultima inserción)
	var result sql.Result

	//Ejecutamos la sentencia SQL
	result, err = Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	//nos regresa el ultimo ID insertado en la base
	LastInsertId, err2 := result.LastInsertId()

	//Verificamos no exista un error al haber preguntado cual era el ultimo ID insertado
	if err2 != nil {
		fmt.Println(err2.Error())
		return 0, err
	}

	fmt.Println("Insert Category > Ejecución Exitosa")

	return LastInsertId, nil

}

func UpdateCategory(c models.Category) error {

	fmt.Println("Comienza Registro de UpdateCategory")

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer Db.Close()

	//Declaramos la sentencia SQL para insertar la categoria
	sentencia := "UPDATE category SET "

	//Verificamos si estamos recibiendo "nombre de la cateria" para actualizar
	if len(c.CategName) > 0 {
		sentencia += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}

	//Verificamos si estamos recibiendo "ruta de la cateria" para actualizar
	if len(c.CategPath) > 0 {
		//Verificamos si previamente ya le habiamos establecido un "nombre de la cateria"
		if !strings.HasSuffix(sentencia, "SET ") {
			//En caso de no termine con "SET", entonces almacenamos una coma para separar las sentencias.
			sentencia += ", "
		}
		sentencia += "Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}

	sentencia += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	//Ejecutamos la sentencia SQL
	_, err = Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecución Exitosa !")

	return nil

}

func DeleteCategory(id int) error {

	fmt.Println("Comienza Registro de DeleteCategory")

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer Db.Close()

	//Declaramos la sentencia SQL para insertar la categoria
	sentencia := "DELETE FROM category WHERE Categ_Id = " + strconv.Itoa(id)

	//Ejecutamos la sentencia SQL
	_, err = Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category > Ejecución Exitosa !")

	return nil
}

func SelectCategories(CategId int, Slug string) ([]models.Category, error) {
	fmt.Println("Comienza SelectCategories")

	//Creamos la variable que almacenara cada registro devuelto de cada categoria de la base de datos.
	var Categ []models.Category

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return Categ, err
	}

	//Declaramos la sentencia SQL para insertar la categoria
	sentencia := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category"

	//Validamos si nos pidieron buscar por un ID en particular en caso que se haya especificado.
	if CategId > 0 {
		sentencia += " WHERE Categ_Id = " + strconv.Itoa(CategId)
	} else {
		// Validamos que nos permita buscar por una ruta en particular en caso que se haya especificado.
		if len(Slug) > 0 {
			sentencia += " WHERE Categ_Path LIKE '%" + Slug + "%'"
		}
	}

	//Imprimimos la sentencia SQL
	fmt.Println(sentencia)

	var rows *sql.Rows

	rows, err = Db.Query(sentencia)

	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err := rows.Scan(&categId, &categName, &categPath)

		//Validamos que no haya surgido ningun error y en caso de existir abortamos y salimos regresando el error
		if err != nil {
			return Categ, err
		}

		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String

		Categ = append(Categ, c)

	}

	fmt.Println("Select Categories > Ejecución Exitosa")

	return Categ, nil
}
