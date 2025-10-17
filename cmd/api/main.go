package main

import (
	"fmt"
	"log"

	_ "github.com/GunarsK-portfolio/public-api/docs"
	"github.com/GunarsK-portfolio/public-api/internal/config"
	"github.com/GunarsK-portfolio/public-api/internal/database"
	"github.com/GunarsK-portfolio/public-api/internal/handlers"
	"github.com/GunarsK-portfolio/public-api/internal/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repository
	repo := repository.New(db, cfg.FilesAPIURL)

	// Initialize handlers
	handler := handlers.New(repo)

	// Setup router
	router := gin.Default()

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check
	router.GET("/api/v1/health", handler.HealthCheck)

	// API routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/profile", handler.GetProfile)
		v1.GET("/experience", handler.GetWorkExperience)
		v1.GET("/certifications", handler.GetCertifications)
		v1.GET("/skills", handler.GetSkills)
		v1.GET("/projects", handler.GetProjects)
		v1.GET("/projects/:id", handler.GetProjectByID)
		v1.GET("/miniatures", handler.GetMiniatures)
		v1.GET("/miniatures/:id", handler.GetMiniatureByID)
		v1.GET("/miniatures/themes", handler.GetMiniatureThemes)
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	log.Printf("Starting public API on port %s", cfg.Port)
	if err := router.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
