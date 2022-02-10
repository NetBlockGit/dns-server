package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init(path string) {
	if os.Getenv("SKIP_ENV_FILE") != "true" {
		// Load environment variables from .env file
		err := godotenv.Load(path)
		if err != nil {
			log.Fatalf("failed to load env file, error: %v", err.Error())
		}
	}
}
