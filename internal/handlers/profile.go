package handlers

import (
	"net/http"

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
	profile, err := h.repo.GetProfile()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}
