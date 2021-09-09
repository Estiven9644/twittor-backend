package jwt

import (
	"time"

	"github.com/Estiven9644/twittor-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("EstivenElQueEstaAprendiendoGo") // Clave de encriptación para jwt

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(), // expiración del token de 24 hora
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //algoritmo de encryptación del jwt
	tokenStr, err := token.SignedString(miClave)                // acá lo encrypto con mi clave

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
