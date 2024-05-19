package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var devMode = true

func GetEnv(key, fallback string) string {
	if devMode {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error loading .env file: %s\n", err)
		}
	}

	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
