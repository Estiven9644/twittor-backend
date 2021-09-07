package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Estiven9644/twittor-backend/bd"
	"github.com/Estiven9644/twittor-backend/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario                      //acá cargo la estructura del usuario
	err := json.NewDecoder(r.Body).Decode(&t) //me carga lo que venga en el body a la estructura de t

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated) //acá devuelvo por status en mi w que se creo un status (el httpStatusCreated es una constante con el 201)

}
