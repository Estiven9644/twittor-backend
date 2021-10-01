package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Estiven9644/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, pagina int) ([]*models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // me cancela el contexto de arriba pero de ultimo por el defer

	db := MongoCN.Database("twittor") //el mongo CN ya lo reconoce por que está en el mismo package
	col := db.Collection("relacion")  // acá se conecta a la colección de usuario dentro de la database twittor

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	var results []*models.DevuelvoTweetsSeguidores

	cursor, err := col.Aggregate(ctx, condiciones)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	err = cursor.All(ctx, &results)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	return results, true
}
