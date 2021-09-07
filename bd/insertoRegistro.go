package bd

import (
	"context"
	"time"

	"github.com/Estiven9644/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // me cancela el contexto de arriba pero de ultimo por el defer

	db := MongoCN.Database("twittor") //el mongo CN ya lo reconoce por que está en el mismo package
	col := db.Collection("usuarios")  // acá se conecta a la colección de usuario dentro de la database twittor

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
