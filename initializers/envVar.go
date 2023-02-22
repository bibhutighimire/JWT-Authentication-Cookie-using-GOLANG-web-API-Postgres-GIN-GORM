package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
