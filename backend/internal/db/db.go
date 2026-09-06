package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/husni-robani/abstracted_self/backend/internal/config"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable password=%s", cfg.DBUser, cfg.DBName, cfg.DBHost, cfg.DBPort, cfg.DBPassword)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	const maxRetries = 5
	for i := 0; i < maxRetries; i++ {
		if err := db.Ping(); err == nil {
			break
		} else if i == maxRetries-1 {
			db.Close()
			return nil, fmt.Errorf("database unreachable after %d retries: %w", maxRetries, err)
		}

		logger.Info.Printf("Database not connected yet (%d/%d), retry in 3s...", i+1, maxRetries)
		time.Sleep(time.Second * 3)
	}

	logger.Info.Println("Database Connected!")

	return db, nil
}
