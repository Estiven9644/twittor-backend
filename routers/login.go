package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Estiven9644/twittor-backend/bd"
	"github.com/Estiven9644/twittor-backend/jwt"
	"github.com/Estiven9644/twittor-backend/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o Contraseña inválidos ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token correspondiente"+err.Error(), 500)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp) // aquí hacemos el encoder de la respuesta y está va a hacer la respuesta que demos.

	/* como grabar en cookies en el navegador desde el backend*/
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
