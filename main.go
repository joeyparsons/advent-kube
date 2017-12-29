package main

import (
	"github.com/joeyparsons/advent-kube/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Starting the service...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router()

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, r))

}
