package handlers

import (
	"net/http"
	"strconv"

	commonHandlers "github.com/GunarsK-portfolio/portfolio-common/handlers"
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
	projects, err := h.repo.GetAllMiniatureProjects(c.Request.Context())
	if err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to fetch miniature projects")
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
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid id")
		return
	}

	project, err := h.repo.GetMiniatureProjectByID(c.Request.Context(), id)
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "miniature project not found", "failed to fetch miniature project")
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
	themes, err := h.repo.GetAllMiniatureThemes(c.Request.Context())
	if err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to fetch miniature themes")
		return
	}
	c.JSON(http.StatusOK, themes)
}
