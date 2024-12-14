package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
}

const (
	ErrorLoadEnv = "Error loading .env file"
)

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(ErrorLoadEnv, err)
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
