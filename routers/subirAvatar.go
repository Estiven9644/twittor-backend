package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Estiven9644/twittor-backend/bd"
	"github.com/Estiven9644/twittor-backend/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")

	if err != nil {
		http.Error(w, "Error al formatear la imagen"+err.Error(), http.StatusInternalServerError)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]

	var archivo string = "upload/avatars/" + IDUsuario + "." + extension // es para saber que id es el que se está subiendo

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen !"+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error copiando la imagen !"+err.Error(), http.StatusInternalServerError)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar en la BD!"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
