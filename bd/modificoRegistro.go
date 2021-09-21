package bd

import (
	"context"
	"time"

	"github.com/Estiven9644/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // me cancela el contexto de arriba pero de ultimo por el defer

	db := MongoCN.Database("twittor") //el mongo CN ya lo reconoce por que está en el mismo package
	col := db.Collection("usuarios")  // acá se conecta a la colección de usuario dentro de la database twittor

	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}

	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}

	if len(u.Email) > 0 {
		registro["email"] = u.Email
	}

	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}

	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	} // de la 58 a la 62 con el filtro se revisa que el _id sea igual al objID que estamos pidiendo con la instrucción "$eq"

	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
