package routers

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"

	//"fmt"

	//importaciones externas (descargadas)
	//"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func InsertCategory(body string, User string) (int, string) {
	//declaramos una variable con la estructura que tiene una categoria de acuerdo a como esta en nuestra base de datos.
	var t models.Category

	//descontraturamos lo que viene en el body en la estructura
	err := json.Unmarshal([]byte(body), &t)

	//verificamos que no haya un error (que venga en un formato tipo json)
	if err != nil {
		return 400, "Error en los datos recibidos"
	}

	//verificamos que en el json recibido tegamos el campo categName (nombre de la categoria)
	if len(t.CategName) == 0 {
		return 400, "debe especificar el Nombre (Title) de la Categoría"
	}

	//verificamos que en el json recibido tegamos el campo pathName (ruta de la categoria)
	if len(t.CategPath) == 0 {
		return 400, "debe especificar el Path (Ruta) de la Categoría"
	}

	//Verificamos si User Is Admin
	isAdmin, msg := bd.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	//Se realiza el registro de la categoria
	result, err2 := bd.InsertCategory(t)

	//Verificamos que no exista un error al intentar realizar el registro de la categoria
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la categoria " + t.CategName + " > " + err2.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"

}

func UpdateCategory(body string, User string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)

	//Verificamos que no tengamos un error al decodificar la información en la estructura.
	if err != nil {
		return 400, "Error en los datos recibidos con el error: " + err.Error()
	}

	//Verificamos que nos mande la información
	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Debe especificar CategName y/o CategPath para actualizar"
	}

	//Verificamos si User Is Admin
	isAdmin, msg := bd.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	// el id de la categoria lo asignamos que nos pasan como parametro
	t.CategID = id

	err2 := bd.UpdateCategory(t)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realziar el UPDATE de la categoria" + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update OK"

}

func DeleteCategory(body string, User string, id int) (int, string) {

	// Validamos que nos hayan pasado un Id valido
	if id == 0 {
		return 400, "Debe de especificar el ID de la categoria a borrar."
	}

	//Verificamos si User Is Admin
	isAdmin, msg := bd.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	//Eliminamos la categoria que corresponde al id
	err := bd.DeleteCategory(id)

	// Validamos que no haya surgido un error al eliminar la categoria
	if err != nil {
		return 400, "Ocurrió un error al intentar realziar el DELETE de la categoria" + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Delete OK"
}
