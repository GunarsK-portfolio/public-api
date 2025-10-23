package main

import (
	"fmt"
	"log"

	commondb "github.com/GunarsK-portfolio/portfolio-common/database"
	_ "github.com/GunarsK-portfolio/public-api/docs"
	"github.com/GunarsK-portfolio/public-api/internal/config"
	"github.com/GunarsK-portfolio/public-api/internal/handlers"
	"github.com/GunarsK-portfolio/public-api/internal/repository"
	"github.com/GunarsK-portfolio/public-api/internal/routes"
	"github.com/gin-gonic/gin"
)

// @title Portfolio Public API
// @version 1.0
// @description Public read-only API for portfolio data
// @host localhost:8082
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := commondb.Connect(commondb.PostgresConfig{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  "disable",
		TimeZone: "UTC",
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repository
	repo := repository.New(db, cfg.FilesAPIURL)

	// Initialize handlers
	handler := handlers.New(repo)

	// Setup router
	router := gin.Default()

	// Setup routes
	routes.Setup(router, handler)

	// Start server
	log.Printf("Starting public API on port %s", cfg.Port)
	if err := router.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
