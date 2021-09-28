package bd

import (
	"context"
	"log"
	"time"

	"github.com/Estiven9644/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // me cancela el contexto de arriba pero de ultimo por el defer

	db := MongoCN.Database("twittor") //el mongo CN ya lo reconoce por que está en el mismo package
	col := db.Collection("tweet")     // acá se conecta a la colección de usuario dentro de la database twittor

	var results []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find() //para hacer un find en mongo

	opciones.SetLimit(20) // el limit para el limite que me quiero traer

	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) // me organiza según  la fecha en orden descendente

	opciones.SetSkip((pagina - 1) * 20) //se le resta de a 1 por la pagina para ir trayendo de a 20

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) { // el next hace el for con el primero y el siguiente dentre de él así lo va recorriendo todo y se crea el context.TODO() por crear un contexto vacío
		var registro models.DevuelvoTweets

		err := cursor.Decode(&registro)

		if err != nil {
			return results, false
		}

		results = append(results, &registro)
	}

	return results, true
}
