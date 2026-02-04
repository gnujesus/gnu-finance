package config

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type Config struct {
	ApiKey string
}

func Init() Config {
	err := env.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{ApiKey: os.Getenv("API_KEY")}
}
