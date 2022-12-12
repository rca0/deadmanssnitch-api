package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rca0/deadmanssnitch-api/deadmanssnitch"
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

	router.HandleFunc("/api", CreateSnitch).Methods("POST")

	log.Printf("Starting server http://0.0.0.0:%s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateSnitch(w http.ResponseWriter, r *http.Request) {
	var snitch deadmanssnitch.Snitch
	err := json.NewDecoder(r.Body).Decode(&snitch)
	if err != nil {
		log.Printf("[x] invalid request payload: %s", err)
		return
	}
	defer r.Body.Close()

	client, err := deadmanssnitch.NewClient(&deadmanssnitch.Config{
		ApiKey: os.Getenv("DEADMANSSNITCH_APIKEY"),
	})
	if err != nil {
		fmt.Printf("[x] error when create new snitch client: %s", err)
		return
	}

	resp, err := client.Snitch.NewSnitch(&snitch)
	if err != nil {
		log.Printf("[x] could not create new snitch: %s", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	fmt.Printf("Successfully created [%s]", resp.Name)
}

func GetSnitch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet :)")
}

func GetOneSnitch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet :)")
}

func UpdateSnitch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet :)")
}

func DeleteSnitch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet :)")
}
