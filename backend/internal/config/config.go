package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	SSL        string
}

func LoadConfig() (*Config, error) {
	if os.Getenv("RUNNING_IN_DOCKER") != "true" {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("failed to load .env: %w", err)
		}

		if err := godotenv.Load(".env.access_credentials"); err != nil {
			return nil, fmt.Errorf("failed to load .env.access_credentials: %w", err)
		}
	}

	cfg := &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSL:        os.Getenv("SSL"),
	}

	var missing []string
	if cfg.ServerPort == "" {
		missing = append(missing, "SERVER_PORT")
	}
	if cfg.DBHost == "" {
		missing = append(missing, "DB_HOST")
	}
	if cfg.DBPort == "" {
		missing = append(missing, "DB_PORT")
	}
	if cfg.DBUser == "" {
		missing = append(missing, "DB_USER")
	}
	if cfg.DBPassword == "" {
		missing = append(missing, "DB_PASSWORD")
	}
	if cfg.DBName == "" {
		missing = append(missing, "DB_NAME")
	}
	// Used directly via os.Getenv elsewhere — fail fast here instead of
	// cryptic errors later (nil paths, open redirects, writes to /app).
	for _, key := range []string{
		"PROFILE_DB_PATH",
		"ALLOWED_ORIGIN_AUTHORIZED_ROUTES",
		"IMAGES_STORAGE_PATH",
		"DOCUMENTS_STORAGE_PATH",
		"ICONS_STORAGE_PATH",
		"ACCESS_KEY",
		"AUTH_KEY",
	} {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required env vars: %s", strings.Join(missing, ", "))
	}

	return cfg, nil
}
