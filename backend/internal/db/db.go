package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/husni-robani/abstracted_self/backend/internal/config"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		return nil, err
	}

	log.Default().Printf("Database Connected!")

	return db, nil
}