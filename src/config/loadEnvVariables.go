package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {

	err := godotenv.Load("src/config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
