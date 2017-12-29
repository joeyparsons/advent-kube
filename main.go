package main

import (
	"github.com/joeyparsons/advent-kube/handlers"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the service...")

	router := handlers.Router()

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))

}
