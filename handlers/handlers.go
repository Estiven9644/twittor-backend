package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/Estiven9644/twittor-backend/middlewares"
	"github.com/Estiven9644/twittor-backend/routers"
)

func Manejadores() {
	router := mux.NewRouter() //manejar el http

	router.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	log.Println(PORT)

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Println("Starting server in : " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
