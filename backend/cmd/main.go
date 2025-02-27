package main

import (
	"github.com/husni-robani/abstracted_self/backend/internal/config"
	"github.com/husni-robani/abstracted_self/backend/internal/db"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/routes"
)

func main() {
	// load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error.Println(err)
	}

	// connect to the database
	dbConn, err := db.Connect(cfg)
	if err != nil {
		logger.Error.Println(err)
	}
	defer dbConn.Close()

	// setup routes
	router := routes.SetupRouter(dbConn)

	// start the server
	logger.Info.Printf("Server is running on port %v\n", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		logger.Error.Fatalf("Failed to start server: %v", err)
	}
}