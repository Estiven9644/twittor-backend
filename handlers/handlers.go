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
	router.HandleFunc("/modificarperfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leotweets", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminartweet", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/subiravatar", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirbanner", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obteneravatar", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/obtenerbanner", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.ObtenerBanner))).Methods("GET")

	PORT := os.Getenv("PORT")
	log.Println(PORT)

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Println("Starting server in : " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
