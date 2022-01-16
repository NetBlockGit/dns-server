package env

import (
	"log"
	"os"
)

func Get(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("env variable %v is required", key)
	}
	return val
}
