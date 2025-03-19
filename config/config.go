package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads  from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

// GetEnv retrieves
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
