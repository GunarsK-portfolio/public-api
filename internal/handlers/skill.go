package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSkills godoc
// @Summary Get all skills
// @Description Get list of all visible skills organized by category
// @Tags skills
// @Produce json
// @Success 200 {array} models.Skill
// @Failure 500 {object} map[string]string
// @Router /skills [get]
func (h *Handler) GetSkills(c *gin.Context) {
	skills, err := h.repo.GetAllSkills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch skills"})
		return
	}
	c.JSON(http.StatusOK, skills)
}
