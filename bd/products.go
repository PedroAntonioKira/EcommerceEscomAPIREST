package bd

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//"errors"

	//"strings"

	//importaciones externas (descargadas)
	//"github.com/aws/aws-sdk-go-v2/internal/strings"
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

func UpdateProduct(p models.Product) error {
	fmt.Println(" Comienza Update Product")

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer Db.Close()

	//Declaramos la sentencia SQL para insertar el Producto.
	sentencia := "Update products SET "

	//Verificamos si nos pasaron como parametro a actualizar el titulo del producto.
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Title", "S", 0, 0, p.ProdTitle)
	//Verificamos si nos pasaron como parametro a actualizar la descripción del producto.
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Description", "S", 0, 0, p.ProdDescrition)
	//Verificamos si nos pasaron como parametro a actualizar el precio del producto.
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Price", "F", 0, p.ProdPrice, "")
	//Verificamos si nos pasaron como parametro a actualizar la categoria del producto.
	sentencia = tools.ArmoSentencia(sentencia, "Prod_CategoryId", "N", p.ProdCategId, 0, "")
	//Verificamos si nos pasaron como parametro a actualizar el stock del producto.
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Stock", "N", p.ProdStock, 0, "")
	//Verificamos si nos pasaron como parametro a actualizar la ruta del producto.
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Path", "S", 0, 0, p.ProdPath)

	//Terminamos la sentencia indicando el id del registro que se va actualizar
	sentencia += "WHERE Prod_Id = " + strconv.Itoa(p.ProdId)

	//Ejecutamos la sentencia SQL
	_, err = Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Produc > Ejecución Exitosa!")

	return nil

}

func DeleteProduct(id int) error {

	fmt.Println(" Comienza Delete Product")

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer Db.Close()

	//Declaramos la sentencia SQL para insertar el Producto.
	sentencia := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(id)

	//Ejecutamos la sentencia SQL
	_, err = Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Produc > Ejecución Exitosa!")

	return nil
}

func SelectProduct(p models.Product, choice string, page int, pageSize int, orderType string, orderField string) (models.ProductResp, error) {

	fmt.Println("Comienza SelectProduct")

	var Resp models.ProductResp
	var Prod []models.Product // Con esto devolveremos una colección de estructuras

	//Nos conectamos a la base de datos
	err := DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return Resp, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer Db.Close()

	var sentencia string
	var sentenciaCount string
	var where, limit string

	sentencia = "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Path, Prod_CategoryId, Prod_Stock FROM products"
	sentenciaCount = "SELECT count(*) as registros FROM products"

	switch choice {
	case "P":
		where = " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)
	case "S":
		where = " WHERE UCASE(CONCAT(Prod_Title, Prod_Description)) LIKE '%" + strings.ToUpper(p.ProdSearch) + "%' "
	case "C":
		where = " WHERE Prod_CategoryId = " + strconv.Itoa(p.ProdCategId)
	case "U":
		where = " WHERE UCASE(Prod_Path) LIKE '%" + strings.ToUpper(p.ProdPath) + "%' "
	case "K":
		join := " JOIN category ON Prod_CategoryId = Categ_Id AND Categ_Path LIKE '%" + strings.ToUpper(p.ProdCategPath) + "%' "
		sentencia += join
		sentenciaCount += join
	}

	sentenciaCount += where

	var rows *sql.Rows

	var err003 error
	rows, err003 = Db.Query(sentenciaCount)

	if err003 != nil {
		fmt.Println("Ocurrio un error al ejecutar la consulta en select producto:")
		fmt.Println(err003)
		return Resp, err003
	}

	fmt.Println("Se realizo Correctamente la consulta de Sentencia Acount:")
	fmt.Println(rows)
	defer rows.Close()

	rows.Next()

	fmt.Println("Prueba 02 Se realizo Correctamente la consulta de Sentencia Acount:")
	fmt.Println(rows)

	var regi sql.NullInt32

	err003 = rows.Scan(&regi)

	fmt.Println("Mostramos el error al guardar los ROWS en REGI:")
	fmt.Println(err003)
	registros := int(regi.Int32)

	fmt.Println("Mostramos EL REGISTRO al guardar REGI int32:")
	fmt.Println(registros)

	if page > 0 {
		if registros > pageSize {
			fmt.Println("Si estamos entrando A LIMIT IF:")
			limit = " LIMIT " + strconv.Itoa(pageSize)
			if page > 1 {
				offset := pageSize * (page - 1)
				limit += " OFFSET " + strconv.Itoa(offset)
			}
			fmt.Println(limit)

		} else {
			limit = ""
		}
	}

	var orderBy string
	if len(orderField) > 0 {
		switch orderField {
		case "I":
			fmt.Println("Si nos mandaron La LETRA:")
			fmt.Println(orderField)
			orderBy = " ORDER BY Prod_Id "
		case "T":
			fmt.Println("Si nos mandaron La LETRA:")
			fmt.Println(orderField)
			orderBy = " ORDER BY Prod_Title "
		case "D":
			fmt.Println("Si nos mandaron La LETRA:")
			fmt.Println(orderField)
			orderBy = " ORDER BY Prod_Description "
		case "F":
			fmt.Println("Si nos mandaron La LETRA:")
			fmt.Println(orderField)
			orderBy = " ORDER BY Prod_CreatedAt "
		case "P":
			fmt.Println("Si nos mandaron La LETRA:")
			fmt.Println(orderField)
			orderBy = " ORDER BY Prod_Price "
		case "S":
			fmt.Println("Si nos mandaron La LETRA:")
			fmt.Println(orderField)
			orderBy = " ORDER BY Prod_Stock "
		case "C":
			fmt.Println("Si nos mandaron La LETRA:")
			fmt.Println(orderField)
			orderBy = " ORDER BY Prod_CategoryId "
		}
		if orderType == "D" {
			orderBy += " DESC "
		}
	}

	sentencia += where + orderBy + limit

	fmt.Println(sentencia)

	rows, err = Db.Query(sentencia)

	for rows.Next() {
		var p models.Product
		var ProdId sql.NullInt32
		var ProdTitle sql.NullString
		var ProdDescription sql.NullString
		var ProdCreateAt sql.NullString // Cambiamos de NullTime a NullString
		var ProdUpdated sql.NullString
		var ProdPrice sql.NullFloat64
		var ProdPath sql.NullString
		var ProdCategId sql.NullInt32
		var ProdStock sql.NullInt32

		err := rows.Scan(&ProdId, &ProdTitle, &ProdDescription, &ProdCreateAt, &ProdUpdated, &ProdPrice, &ProdPath, &ProdCategId, &ProdStock)
		if err != nil {
			return Resp, err
		}

		p.ProdId = int(ProdId.Int32)
		p.ProdTitle = ProdTitle.String
		p.ProdDescrition = ProdDescription.String

		// Convertir el string de la fecha a time.Time
		if ProdCreateAt.Valid {
			p.ProdCreateAt = ProdCreateAt.String
		}

		// Asignar el valor de ProdUpdated.Time directamente a p.ProdUpdated si ProdUpdated es válido
		if ProdUpdated.Valid {
			p.ProdUpdated = ProdUpdated.String
		}

		p.ProdPrice = ProdPrice.Float64
		p.ProdPath = ProdPath.String
		p.ProdCategId = int(ProdCategId.Int32)
		p.ProdStock = int(ProdStock.Int32)

		Prod = append(Prod, p)
	}

	Resp.TotalItems = registros
	Resp.Data = Prod

	fmt.Println(" Select Product > Ejecución Exitosa !")

	return Resp, nil
}
