package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
