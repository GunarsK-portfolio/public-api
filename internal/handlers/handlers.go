package handlers

import (
	"net/http"
	"strconv"

	"github.com/GunarsK-portfolio/public-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

// GetProfile godoc
// @Summary Get profile information
// @Description Get the portfolio owner's profile information
// @Tags profile
// @Produce json
// @Success 200 {object} models.Profile
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	profile, err := h.repo.GetProfile()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// GetWorkExperience godoc
// @Summary Get all work experience
// @Description Get list of all work experience entries
// @Tags experience
// @Produce json
// @Success 200 {array} models.WorkExperience
// @Failure 500 {object} map[string]string
// @Router /experience [get]
func (h *Handler) GetWorkExperience(c *gin.Context) {
	experiences, err := h.repo.GetAllWorkExperience()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch work experience"})
		return
	}
	c.JSON(http.StatusOK, experiences)
}

// GetCertifications godoc
// @Summary Get all certifications
// @Description Get list of all certifications
// @Tags certifications
// @Produce json
// @Success 200 {array} models.Certification
// @Failure 500 {object} map[string]string
// @Router /certifications [get]
func (h *Handler) GetCertifications(c *gin.Context) {
	certifications, err := h.repo.GetAllCertifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch certifications"})
		return
	}
	c.JSON(http.StatusOK, certifications)
}

// GetMiniatures godoc
// @Summary Get all miniature projects
// @Description Get list of all miniature painting projects with images
// @Tags miniatures
// @Produce json
// @Success 200 {array} models.MiniatureProject
// @Failure 500 {object} map[string]string
// @Router /miniatures [get]
func (h *Handler) GetMiniatures(c *gin.Context) {
	projects, err := h.repo.GetAllMiniatureProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch miniature projects"})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// GetMiniatureByID godoc
// @Summary Get miniature project by ID
// @Description Get detailed information about a specific miniature project
// @Tags miniatures
// @Produce json
// @Param id path int true "Miniature Project ID"
// @Success 200 {object} models.MiniatureProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /miniatures/{id} [get]
func (h *Handler) GetMiniatureByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	project, err := h.repo.GetMiniatureProjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "miniature project not found"})
		return
	}
	c.JSON(http.StatusOK, project)
}

// HealthCheck godoc
// @Summary Health check
// @Description Check if service is healthy
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
