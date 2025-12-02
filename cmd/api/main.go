package main

import (
	"log"
	"os"
	"strconv"
	"time"

	commondb "github.com/GunarsK-portfolio/portfolio-common/database"
	"github.com/GunarsK-portfolio/portfolio-common/health"
	"github.com/GunarsK-portfolio/portfolio-common/logger"
	"github.com/GunarsK-portfolio/portfolio-common/metrics"
	"github.com/GunarsK-portfolio/portfolio-common/server"
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

	// Initialize structured logger
	appLogger := logger.New(logger.Config{
		Level:       os.Getenv("LOG_LEVEL"),
		Format:      os.Getenv("LOG_FORMAT"),
		ServiceName: "public-api",
		AddSource:   os.Getenv("LOG_SOURCE") == "true",
	})

	appLogger.Info("Starting public API", "version", "1.0")

	// Initialize Prometheus metrics
	metricsCollector := metrics.New(metrics.Config{
		ServiceName: "public",
		Namespace:   "portfolio",
	})

	// Connect to database
	//nolint:staticcheck // Embedded field name required due to ambiguous fields
	db, err := commondb.Connect(commondb.PostgresConfig{
		Host:     cfg.DatabaseConfig.Host,
		Port:     strconv.Itoa(cfg.DatabaseConfig.Port),
		User:     cfg.DatabaseConfig.User,
		Password: cfg.DatabaseConfig.Password,
		DBName:   cfg.DatabaseConfig.Name,
		SSLMode:  cfg.DatabaseConfig.SSLMode,
	})
	if err != nil {
		appLogger.Error("Failed to connect to database", "error", err)
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if closeErr := commondb.CloseDB(db); closeErr != nil {
			appLogger.Error("Failed to close database", "error", closeErr)
		}
	}()
	appLogger.Info("Database connection established")

	// Initialize health aggregator
	healthAgg := health.NewAggregator(3 * time.Second)
	healthAgg.Register(health.NewPostgresChecker(db))

	// Initialize repository
	repo := repository.New(db, cfg.FilesAPIURL)

	// Initialize handlers
	handler := handlers.New(repo)

	// Setup router with custom middleware
	router := gin.New()
	router.Use(logger.Recovery(appLogger))
	router.Use(logger.RequestLogger(appLogger))
	router.Use(metricsCollector.Middleware())

	// Setup routes
	routes.Setup(router, handler, cfg, metricsCollector, healthAgg)

	// Start server with graceful shutdown
	appLogger.Info("Public API ready", "port", cfg.ServiceConfig.Port, "environment", os.Getenv("ENVIRONMENT"))

	serverCfg := server.DefaultConfig(strconv.Itoa(cfg.ServiceConfig.Port))
	if err := server.Run(router, serverCfg, appLogger); err != nil {
		appLogger.Error("Server error", "error", err)
		log.Fatal("Server error:", err)
	}
}
