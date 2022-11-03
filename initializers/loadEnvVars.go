package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// if .env exists (local development) load it
func LoadEnvVars() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return
}
