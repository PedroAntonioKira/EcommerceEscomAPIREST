package handlers

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	"strconv"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/auth"
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/routers"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayProxyRequest) (int, string) {

	fmt.Println("Voy a procesar " + path + " > " + method)

	id := request.PathParameters["id"]

	fmt.Println("El ID tiene:")
	fmt.Println(id)
	fmt.Println("Se imprimio ID :")

	//Prueba 02
	id02 := request.PathParameters["Id"]

	fmt.Println("El ID 02 tiene:")
	fmt.Println(id02)
	fmt.Println("Se imprimio ID 02:")

	idn, _ := strconv.Atoi(id)

	fmt.Println("Mostramos ID: " + id)
	//validamos la autorización del token
	isOk, statusCode, user := validoAuthorization(path, method, headers)

	fmt.Println("El IsOk: ")
	fmt.Println(isOk)

	//Verificamos que la autorización no tenga problemas
	if !isOk {
		return statusCode, user
	}

	fmt.Println("Llegamos hasta aqui:")
	fmt.Println("Path1: " + path[0:19])
	fmt.Println("Path2: " + path[16:20])

	//Validamos y analizamos que nos viene en el path

	switch path[16:20] {
	case "user":
		fmt.Println("Entramos a User")
	//	return ProcesoUsers(body, path, method, user, id, request)
	case "prod":
		fmt.Println("Entramos a Products")
	//	return ProcesoProducts(body, path, method, user, idn, request)
	case "stoc":
		fmt.Println("Entramos a Stock")
	//	return ProcesoStock(body, path, method, user, idn, request)
	case "addr":
		fmt.Println("Entramos a Address")
	//	return ProcesoAddress(body, path, method, user, idn, request)
	case "cate":
		fmt.Println("Entramos a Category")
		return ProcesoCategory(body, path, method, user, idn, request)
	case "orde":
		fmt.Println("Entramos a Order")
		//	return ProcesoOrder(body, path, method, user, idn, request)
	}

	return 400, "Method Invalid loquisimo04"
}

func validoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}

	// Recibimos el token que viene en el headers
	token := headers["Authorization"]
	fmt.Println("Imprimiremos el token que nos dan authorization: ")
	fmt.Println(token)
	fmt.Println("Ya imprimimos el token")

	// Recibimos el token que viene en el headers 02
	token02 := headers["Authorization"]
	fmt.Println("Imprimiremos el token02 que nos dan authorization: ")
	fmt.Println(token02)
	fmt.Println("Ya imprimimos el token02")

	//verificamos que si hayamos recibido la autorización de "authorization"
	if len(token) == 0 {
		return false, 401, "Token Requerido"
	}

	fmt.Println("Se empezará a validar token con VALIDOTOKEN()")
	//Si nos llego el token correctamente validamos el token sea correcto
	todoOK, err, msg := auth.ValidoToken(token)

	// si algo no estuvo bien, verificamos que fue lo que fallo en el token
	if !todoOK {
		//Verificamos si fallo la verificació del token porque existio un error
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token porque: " + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK Yei")

	return true, 200, msg
}

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoStock(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	}
	return 400, "Method Invalid 22"
}

func ProcesoOrder(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}
