package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rca0/deadmanssnitch-api/config"
	"github.com/rca0/deadmanssnitch-api/routes"
)

func main() {
	config.InitLoadConfigs()
	port, ok := config.GetEnv("PORT")
	if !ok {
		port = "8000"
	}

	_, ok = config.GetEnv("DEADMANSSNITCH_APIKEY")
	if !ok {
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
