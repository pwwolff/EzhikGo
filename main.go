package main

import (
	"log"
	"net/http"

	"github.com/pwwolff/EzhikGo/config"
	"github.com/pwwolff/EzhikGo/store"

	"github.com/gorilla/handlers"
)

func main() {
	port := config.GetConfig().Port

	router := store.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
