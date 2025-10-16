package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
