package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // me cancela el contexto de arriba pero de ultimo por el defer

	db := MongoCN.Database("twittor") //el mongo CN ya lo reconoce por que está en el mismo package
	col := db.Collection("tweet")     // acá se conecta a la colección de usuario dentro de la database twittor

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)

	return err
}
