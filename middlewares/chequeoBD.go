package middlewares

import (
	"net/http"

	"github.com/Estiven9644/twittor-backend/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
