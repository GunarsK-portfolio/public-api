package handlers

import (
	"net/http"

	commonHandlers "github.com/GunarsK-portfolio/portfolio-common/handlers"
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
	skills, err := h.repo.GetAllSkills(c.Request.Context())
	if err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to fetch skills")
		return
	}
	c.JSON(http.StatusOK, skills)
}
