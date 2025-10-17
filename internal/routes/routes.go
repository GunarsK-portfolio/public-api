package routes

import (
	"github.com/GunarsK-portfolio/public-api/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.Engine, handler *handlers.Handler) {
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
}
