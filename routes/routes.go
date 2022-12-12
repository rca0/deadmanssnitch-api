package routes

import (
	"github.com/gorilla/mux"
	"github.com/rca0/deadmanssnitch-api/controllers"
)

func SnitchRoute(router *mux.Router) {
	// create new snitch
	router.HandleFunc("/api/snitch", controllers.CreateSnitch()).Methods("POST")

	// list snitches
	router.HandleFunc("/api/snitch", controllers.GetSnitches()).Methods("GET")

	// get snitch
	router.HandleFunc("/api/snitch/{token}", controllers.GetSnitch()).Methods("GET")

	// to update a snitch, send PATCH request
	router.HandleFunc("/api/snitch/{token}", controllers.GetSnitch()).Methods("PATCH")

	// delete snitch
	router.HandleFunc("/api/snitch/{token}", controllers.GetSnitch()).Methods("DELETE")
}
