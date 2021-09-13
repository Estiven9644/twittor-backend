package middlewares

import (
	"net/http"

	"github.com/Estiven9644/twittor-backend/routers"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc { // esto ingresa siempre a los middlewares y devuelve el mismo tipo
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization")) // envio token a proceso token
		if err != nil {
			http.Error(w, "Error en el token"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
