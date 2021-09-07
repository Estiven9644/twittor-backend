package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string) (string, error) {
	costo := 8 // esto hace que se encrypte a las 2 a la 8

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo) //esto devuelve y recibe un slide de bytes entonces por eso toca enviarlo así conviertiendo mi pass en un slide bytes y con el costo

	return string(bytes), err
}
