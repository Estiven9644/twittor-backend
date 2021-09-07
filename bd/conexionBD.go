package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://Estiven9644Escobar:Parque%251088339047@clustertwittor.rkrxn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConectarBD me permite conectarme a Mongo*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions) // se hace el contexto desde la librería para la memoria
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printf(" \n Conexión exitosa con la base de datos \n")
	return client
}

func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
