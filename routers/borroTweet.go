package routers

import (
	"net/http"

	"github.com/Estiven9644/twittor-backend/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario) // este IDUsuario es global del procesoToken

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
