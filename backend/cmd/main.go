package main

import (
	"log"

	"github.com/husni-robani/abstracted_self/backend/internal/config"
	"github.com/husni-robani/abstracted_self/backend/internal/db"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Default().Printf("load config error: %v\n", err)
		return
	}

	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Default().Printf("database connection error: %v\n", err)
		return
	}

	defer dbConn.Close()
}