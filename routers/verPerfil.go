package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Estiven9644/twittor-backend/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar consultar el ID"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil) // está siempre es la respuesta
}
