package main

import (
	"fmt"
	"go-backend/internal/config"
	"go-backend/internal/database"
	"go-backend/internal/logger"
	"go-backend/internal/middleware"
	"go-backend/internal/router"
	"net/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// Initialize logger
	logger.InitLogger(&cfg.Logging)
	logger.Info("Starting application...")

	// Initialize database
	if err := database.InitDatabase(&cfg.Database); err != nil {
		logger.Fatal("Failed to initialize database: ", err)
	}
	defer database.CloseDatabase()

	// Setup routes
	r := router.SetupRoutes()

	// Apply CORS middleware
	handler := middleware.CORSMiddleware()(r)

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	logger.Info("Server starting on ", addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Fatal("Server failed to start: ", err)
	}
}