package utils

// utils package has methods concerning common I/O operations to set up the environment.
import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file and returns the error if any.
func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}
}
