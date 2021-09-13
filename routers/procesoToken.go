package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Estiven9644/twittor-backend/bd"
	"github.com/Estiven9644/twittor-backend/models"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("EstivenElQueEstaAprendiendoGo")
	claims := &models.Claim{} // JWT debe tomarlo así para el puntero

	splitToken := strings.Split(tk, "Bearer") // hacer el split y lo deja en un vector

	if len(splitToken) != 2 { // 2 por que sería bearer y el jwt si no entonces hay un error
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1]) // TrimSpace quita los espacios

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) { // está es la validación propia del token y lo devuelve en el objeto claims
		return miClave, nil
	})

	if err == nil { // si no hay error pasamos al siguiente middleware en el cual verificamos la base de datos
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid { // valida desde el token si es valido
		return claims, false, string(""), err
	}

	return claims, false, string(""), err
}
