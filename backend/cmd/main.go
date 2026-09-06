package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/husni-robani/abstracted_self/backend/internal/config"
	"github.com/husni-robani/abstracted_self/backend/internal/db"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/routes"
)

func main() {
	// load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error.Fatalf("invalid configuration: %v", err)
	}

	// connect to the database
	dbConn, err := db.Connect(cfg)
	if err != nil {
		logger.Error.Fatalf("failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// setup routes
	router := routes.SetupRouter(dbConn)

	// start the server with graceful shutdown so redeploys don't drop requests
	srv := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Info.Printf("Server is running on port %v\n", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	logger.Info.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error.Fatalf("Forced shutdown: %v", err)
	}
	logger.Info.Println("Server stopped")
}
