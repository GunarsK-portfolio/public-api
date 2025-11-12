package handlers

import (
	"net/http"
	"strconv"

	commonHandlers "github.com/GunarsK-portfolio/portfolio-common/handlers"
	"github.com/gin-gonic/gin"
)

// GetProjects godoc
// @Summary Get all portfolio projects
// @Description Get list of all portfolio projects with technologies
// @Tags projects
// @Produce json
// @Success 200 {array} models.PortfolioProject
// @Failure 500 {object} map[string]string
// @Router /projects [get]
func (h *Handler) GetProjects(c *gin.Context) {
	projects, err := h.repo.GetAllProjects(c.Request.Context())
	if err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to fetch projects")
		return
	}
	c.JSON(http.StatusOK, projects)
}

// GetProjectByID godoc
// @Summary Get portfolio project by ID
// @Description Get detailed information about a specific portfolio project
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} models.PortfolioProject
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /projects/{id} [get]
func (h *Handler) GetProjectByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		commonHandlers.RespondError(c, http.StatusBadRequest, "invalid id")
		return
	}

	project, err := h.repo.GetProjectByID(c.Request.Context(), id)
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "project not found", "failed to fetch project")
		return
	}
	c.JSON(http.StatusOK, project)
}
