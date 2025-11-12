package routes

import (
	"github.com/GunarsK-portfolio/portfolio-common/metrics"
	common "github.com/GunarsK-portfolio/portfolio-common/middleware"
	"github.com/GunarsK-portfolio/public-api/internal/config"
	"github.com/GunarsK-portfolio/public-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.Engine, handler *handlers.Handler, cfg *config.Config, metricsCollector *metrics.Metrics) {
	// Security middleware with CORS validation (read-only public access)
	securityMiddleware := common.NewSecurityMiddleware(
		cfg.AllowedOrigins,
		"GET,OPTIONS",
		"Content-Type",
		false,
	)
	router.Use(securityMiddleware.Apply())

	// Health check
	router.GET("/health", handler.HealthCheck)

	// Metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
