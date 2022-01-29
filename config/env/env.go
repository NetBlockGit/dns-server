package env

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	Init()
}

func Init() {
	if os.Getenv("SKIP_ENV_FILE") != "true" {
		// Load environment variables from .env file
		var (
			_, b, _, _ = runtime.Caller(0)
			basepath   = filepath.Join(b, "..", "..", "..")
		)
		err := godotenv.Load(filepath.Join(basepath, ".env"))
		if err != nil {
			log.Fatalf("failed to load env file, error: %v", err.Error())
		}

	}
}
