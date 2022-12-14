package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rca0/deadmanssnitch-api/api"
)

func CreateSnitch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var snitch api.Snitch
		err := json.NewDecoder(r.Body).Decode(&snitch)
		if err != nil {
			log.Printf("[x] invalid request payload: %s", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		defer r.Body.Close()

		client, err := api.NewClient(&api.Config{
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
		client, err := api.NewClient(&api.Config{
			ApiKey: os.Getenv("DEADMANSSNITCH_APIKEY"),
		})
		if err != nil {
			fmt.Printf("[x] error when create new snitch client: %s", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		resp, err := client.Snitch.GetSnitches()
		if err != nil {
			log.Printf("[x] could not get any snitch: %s", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// HTTP 302 - Found
		w.WriteHeader(http.StatusFound)
		fmt.Println("[x] wow! look what i found...")
		for i, v := range *resp {
			fmt.Printf("[x] %d Snitch: %s", i+1, v.Name)
		}
	}
}

func GetSnitch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		token := params["token"]
		client, err := api.NewClient(&api.Config{
			ApiKey: os.Getenv("DEADMANSSNITCH_APIKEY"),
		})
		if err != nil {
			fmt.Printf("[x] error when create new snitch client: %s", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		resp, err := client.Snitch.GetSnitch(token)
		if err != nil {
			log.Printf("[x] could not find the snitch: %s", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusFound)
		fmt.Println("[x] wow! look what i found...")
		fmt.Printf("[x] Snitch: %s", resp.Name)
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
