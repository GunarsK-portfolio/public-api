package handlers

import (
	"net/http"

	commonHandlers "github.com/GunarsK-portfolio/portfolio-common/handlers"
	"github.com/GunarsK-portfolio/public-api/internal/models"
	"github.com/gin-gonic/gin"
)

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
	var profile *models.Profile
	profile, err := h.repo.GetProfile(c.Request.Context())
	if err != nil {
		commonHandlers.HandleRepositoryError(c, err, "profile not found", "failed to fetch profile")
		return
	}
	c.JSON(http.StatusOK, profile)
}
