package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Estiven9644/twittor-backend/bd"
	"github.com/Estiven9644/twittor-backend/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar decodificar el body"+err.Error(), 500)
		return
	}

	registro := models.GraboTweet{
		UserId:  IDUsuario, // viene de ValidoJWT y procesoToken cuando se desencrypta el JWT
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(), // tal cual la fecha de hoy con todo
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, intente de nuevo"+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
