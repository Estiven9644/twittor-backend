package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Estiven9644/twittor-backend/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusInternalServerError)
		return
	}

	openFile, err := os.Open("upload/banners/" + perfil.Banner)

	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile) // una vez abrido se lo envia al response writer

	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}
