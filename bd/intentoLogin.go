package bd

import (
	"github.com/Estiven9644/twittor-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado { // si esto es igual a false entonces returne así
		return usu, false
	}

	passwordBytes := []byte(password)  // está no viene encriptada
	passwordBD := []byte(usu.Password) // está ya viene encryptada de la base de datos

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
