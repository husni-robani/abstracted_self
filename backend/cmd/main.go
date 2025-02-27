package main

import (
	"github.com/husni-robani/abstracted_self/backend/internal/config"
	"github.com/husni-robani/abstracted_self/backend/internal/db"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error.Println(err)
	}

	dbConn, err := db.Connect(cfg)
	if err != nil {
		logger.Error.Println(err)
	}

	defer dbConn.Close()
}