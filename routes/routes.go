package routes

import (
	"github.com/gorilla/mux"
	"github.com/rca0/deadmanssnitch-api/controllers"
)

func SnitchRoute(router *mux.Router) {
	// create new snitch
	router.HandleFunc("/api/snitches", controllers.CreateSnitch()).Methods("POST")

	// get snitch
	router.HandleFunc("/api/snitches/{token}", controllers.GetSnitch()).Methods("GET")

	// to update a snitch, send PATCH request
	router.HandleFunc("/api/snitches/{token}", controllers.UpdateSnitch()).Methods("PATCH")

	// delete snitch
	router.HandleFunc("/api/snitches/{token}", controllers.DeleteSnitch()).Methods("DELETE")

	// list snitches
	router.HandleFunc("/api/snitches", controllers.GetSnitches()).Methods("GET")
}
