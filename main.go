package main

import (
	"github.com/joeyparsons/advent-kube/handlers"
	"github.com/joeyparsons/advent-kube/version"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router()

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, r))

}
