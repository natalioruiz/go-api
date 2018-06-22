package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/natalioruiz/go-api/album"
)

func main() {
	// get the application port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// initialize router
	router := album.NewRouter()

	// enable API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// launch server
	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
