package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rca0/deadmanssnitch-api/deadmanssnitch"
)

func main() {
	client, err := deadmanssnitch.NewClient(&deadmanssnitch.Config{
		ApiKey: os.Getenv("DEADMANSSNITCH_APIKEY"),
	})

	if err != nil {
		log.Fatal(err)
	}

	snitch, resp, err := client.Snitch.NewSnitch(&deadmanssnitch.Snitch{
		Name:     "testing-api",
		Interval: "daily",
	})
	if err != nil {
		log.Fatal(err)
	}

	if resp.Response.StatusCode == 200 {
		fmt.Printf("Successfully created [%s]", snitch.Name)
	}
}
