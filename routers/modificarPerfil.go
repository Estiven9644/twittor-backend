package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Estiven9644/twittor-backend/bd"
	"github.com/Estiven9644/twittor-backend/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario) // IDUsuario se saca de proceso token es una variable global la cual utilizó acá

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el  registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
