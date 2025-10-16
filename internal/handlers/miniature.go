package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// GetMiniatureThemes godoc
// @Summary Get all miniature themes with projects
// @Description Get list of all miniature themes with their associated projects
// @Tags miniatures
// @Produce json
// @Success 200 {array} models.MiniatureTheme
// @Failure 500 {object} map[string]string
// @Router /miniatures/themes [get]
func (h *Handler) GetMiniatureThemes(c *gin.Context) {
	themes, err := h.repo.GetAllMiniatureThemes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch miniature themes"})
		return
	}
	c.JSON(http.StatusOK, themes)
}
