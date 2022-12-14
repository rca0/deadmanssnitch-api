package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitLoadConfigs() {
	// load env vars from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("[x] try to load environment variable - ERR: %s", err)
	}
}

func GetEnv(d string) (string, bool) {
	data := os.Getenv(d)
	if data == "" {
		return "", false
	}

	return data, true
}
