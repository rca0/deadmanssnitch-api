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
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		client, err := api.NewClient(&api.Config{
			ApiKey: os.Getenv("DEADMANSSNITCH_APIKEY"),
		})
		if err != nil {
			fmt.Printf("[x] error when create new snitch client: %s", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		resp, err := client.Snitch.NewSnitch(&snitch)
		if err != nil {
			log.Printf("[x] could not create new snitch: %s", err)
			return
		}

		snitchResp, err := json.Marshal(resp)
		if err != nil {
			log.Printf("[x] could not marshal the Snitch Response: %s", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(snitchResp)

		log.Printf("[x] Successfully created: %s", snitchResp)
	}
}

func GetSnitches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		snitches := []api.SnitchResponse{}
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
		for _, v := range *resp {
			snitches = append(snitches, v)
		}

		snitchResponses, err := json.Marshal(snitches)
		if err != nil {
			log.Printf("[x] could not marshal the Snitches Responses: %s", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(snitchResponses)
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

		snitchResp, err := json.Marshal(resp)
		if err != nil {
			log.Printf("[x] could not marshal the Snitch Response: %s", err)
			return
		}

		w.WriteHeader(http.StatusFound)
		w.Write(snitchResp)
	}
}

func UpdateSnitch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func DeleteSnitch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
