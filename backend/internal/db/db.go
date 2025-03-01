package db

import (
	"database/sql"
	"fmt"

	"github.com/husni-robani/abstracted_self/backend/internal/config"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s host=%s sslmode=disable password=%s", cfg.DBUser, cfg.DBName, cfg.DBHost, cfg.DBPassword)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info.Println("Database Connected!")

	return db, nil
}