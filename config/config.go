package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PORT string

func LoadConfigEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
}
