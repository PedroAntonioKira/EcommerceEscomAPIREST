package bd

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"strconv"

	//"strings"
	//"errors"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Comienza Registro de Producto")

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return 0, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer Db.Close()

	//Declaramos la sentencia SQL para insertar el Producto.
	sentencia := "INSERT INTO products (Prod_Title "

	//Preguntamos si cada uno de los campos de la estructura nos vino como dato para incluirlo o no

	//Descripción del producto
	if len(p.ProdDescrition) > 0 {
		sentencia += ", Prod_Description"
	}
	//Precio del Producto
	if p.ProdPrice > 0 {
		sentencia += ", Prod_Price"
	}
	//Identificador de la Categoria del Producto
	if p.ProdCategId > 0 {
		sentencia += ", Prod_CategoryId"
	}
	//Stock del Producto
	if p.ProdStock > 0 {
		sentencia += ", Prod_Stock"
	}
	//Ruta del Producto
	if len(p.ProdPath) > 0 {
		sentencia += ", Prod_Path"
	}

	sentencia += ") VALUES ('" + tools.EscapeString(p.ProdTitle) + "'"

	//Preguntamos de nuevo si cada uno de los campos de la estructura nos vino como dato para incluirlo o no

	//Descripción del producto
	if len(p.ProdDescrition) > 0 {
		sentencia += ",'" + tools.EscapeString(p.ProdDescrition) + "'"
	}
	//Precio del Producto
	if p.ProdPrice > 0 {
		sentencia += ", " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}
	//Identificador de la Categoria del Producto
	if p.ProdCategId > 0 {
		sentencia += ", " + strconv.Itoa(p.ProdCategId)
	}
	//Stock del Producto
	if p.ProdStock > 0 {
		sentencia += ", " + strconv.Itoa(p.ProdStock)
	}
	//Ruta del Producto
	if len(p.ProdPath) > 0 {
		sentencia += ", '" + tools.EscapeString(p.ProdPath) + "'"
	}

	//Cerramos la consulta SQL
	sentencia += ")"

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

	fmt.Println("Insert Product > Ejecución Exitosa")

	return LastInsertId, nil

}
