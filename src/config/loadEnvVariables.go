package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {

	err := godotenv.Load("/etc/secrets/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
