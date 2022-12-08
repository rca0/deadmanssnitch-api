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

	resp, err := client.Snitch.NewSnitch(&deadmanssnitch.Snitch{
		Name:     "testing-api",
		Interval: "daily",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created [%s]", resp.Name)
}
