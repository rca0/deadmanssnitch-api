package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rca0/deadmanssnitch-api/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	apikey := os.Getenv("DEADMANSSNITCH_APIKEY")
	if apikey == "" {
		log.Fatal("You should define the DEADMANSSNITCH_APIKEY environment variable, set and try again... :)")
	}

	router := mux.NewRouter()

	routes.SnitchRoute(router)

	log.Printf("Starting server http://0.0.0.0:%s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
