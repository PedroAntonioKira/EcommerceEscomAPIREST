package routers

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

// Declaramos función para insertar los productos
func InsertProduct(body string, User string) (int, string) {

	// Creamos la variable que tiene la estructura de todo lo que conforma a un producto.
	var t models.Product

	// Decodificamos lo que nos viene en el endpoint (json) para guardarlo en la estructura.
	err := json.Unmarshal([]byte(body), &t)

	// Verificamos que no haya existido algun error al decodificar el json a la estructura
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	// Verificamos que nos hayan pasado el titulo del Producto
	if len(t.ProdTitle) == 0 {
		return 400, "Se debe especificar el Nombre (Title) del Producto"
	}

	// Verificamos si User Is Admin
	isAdmin, msg := bd.UserIsAdmin(User)

	// Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	// Insertamos el producto
	result, err2 := bd.InsertProduct(t)

	// Verificamos que no haya existido un error al insertar el producto.
	if err2 != nil {
		return 400, "Ocurrio Un error al intentar realizar el registro del producto " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ ProductID: " + strconv.Itoa(int(result)) + "}"
}

// Declaramos función para la actualización (Update) de la información de un producto.
func UpdateProduct(body string, User string, id int) (int, string) {
	//creamos la variable de la estructura que almacenará todo lo del producto relacionado
	var t models.Product

	//Decodificamos el json que nos mandan en el endpoint en la estructura del producto para poder guardarla.
	err := json.Unmarshal([]byte(body), &t)

	//Verificamos que no tengamos un error al decodificar la información en la estructura.
	if err != nil {
		return 400, "Error en los datos recibidos con el error: " + err.Error()
	}

	//Verificamos si User Is Admin
	isAdmin, msg := bd.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	// el id del producto lo asignamos que nos pasan como parametro
	t.ProdId = id

	//Actualizamos el Producto.
	err2 := bd.UpdateProduct(t)

	//Verificamos no exista un error al momento en que actualizamos el producto.
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del producto" + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update OK Producto"

}

// Declaramos función para el borrado (Delete) de la información de un producto.

func DeleteProduct(User string, id int) (int, string) {

	//Verificamos si User Is Admin
	isAdmin, msg := bd.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	//Actualizamos el Producto.
	err2 := bd.DeleteProduct(id)

	//Verificamos no exista un error al momento en que actualizamos el producto.
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el DELETE del producto" + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Delete Product OK"
}

func SelectProduct(request events.APIGatewayProxyRequest) (int, string) {

	// Variable donde almacenaremos la información recibida del producto.
	var t models.Product

	// Variable para almacenar el error en caso de existir.
	var err error

	// "page" nos ayudará a saber el numero de pagina de productos y "pagesize" sera la cantidad de productos que entraran por pagina.
	var page, pageSize int

	// variables para ordenar "orderField" ordena por el campo especificado (Descripción, Precio, etc), "ordenType" ordena de forma ascendente o descendente.
	var orderType, orderField string

	//Variable que nos convierte a JSON los datos que manda el endpoint
	param := request.QueryStringParameters

	//Guarfamos los datos que vienen en el Json en variables.
	page, _ = strconv.Atoi(param["page"])

	pageSize, _ = strconv.Atoi(param["pageSize"])

	orderType = param["orderType"] // (D = Descendente) , (A o Nil = Ascendente).

	orderField = param["orderField"] // 'I' id, 'T' Title, 'D' Description, 'F' Created At,
	// 'P' Price, 'C' CategId, 'S' Stock

	//Valido que especifique al menos un tipo ordenamiento (en caso de existir)
	if !strings.Contains("ITDFPCS", orderField) {
		orderField = ""
	}

	var choice string

	if len(param["prodId"]) > 0 {
		fmt.Println("SI LLEGO EL prodId, SE ACTUALIZA")
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}

	if len(param["search"]) > 0 {
		fmt.Println("SI LLEGO EL search, SE ACTUALIZA")
		choice = "S"
		t.ProdSearch, _ = param["search"]
	}

	if len(param["categId"]) > 0 {
		fmt.Println("SI LLEGO EL categId, SE ACTUALIZA")
		choice = "C"
		t.ProdCategId, _ = strconv.Atoi(param["categId"])
	}

	if len(param["slug"]) > 0 {
		fmt.Println("SI LLEGO EL slug, SE ACTUALIZA")
		choice = "U"
		t.ProdPath, _ = param["slug"]
	}

	if len(param["slugCateg"]) > 0 {
		fmt.Println("SI LLEGO EL slugCateg, SE ACTUALIZA")
		choice = "K"
		t.ProdCategPath, _ = param["slugCateg"]
	}

	fmt.Println(param)

	result, err := bd.SelectProduct(t, choice, page, pageSize, orderType, orderField)

	if err != nil {
		return 400, "Ocurrió un error al intentar capturar los resultados de la búsqueda de tipo '" + choice + "' en productos" + err.Error()
	}

	Product, err2 := json.Marshal(result)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON la busqueda de productos"
	}

	return 200, string(Product)

}
