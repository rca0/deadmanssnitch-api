package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rca0/deadmanssnitch-api/deadmanssnitch"
)

func CreateSnitch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func GetSnitches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Not implemented yet :)")
	}
}

func GetSnitch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Not implemented yet :)")
	}
}

func UpdateSnitch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Not implemented yet :)")
	}
}

func DeleteSnitch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Not implemented yet :)")
	}
}
