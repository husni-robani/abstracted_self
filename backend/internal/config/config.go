package config

import (
	"os"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		logger.Error.Printf("failed to load ENV: %v", err)
		os.Exit(0)
		return nil, err
	}

	err = godotenv.Load(".env.user_credential")
	if err != nil {
		logger.Error.Printf("failed load USER ENV: %v", err)
		os.Exit(0)
		return nil, err
	}

	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
	}, nil
}