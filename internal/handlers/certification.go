package handlers

import (
	"net/http"

	commonHandlers "github.com/GunarsK-portfolio/portfolio-common/handlers"
	"github.com/GunarsK-portfolio/public-api/internal/models"
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
	var certifications []models.Certification
	certifications, err := h.repo.GetAllCertifications(c.Request.Context())
	if err != nil {
		commonHandlers.LogAndRespondError(c, http.StatusInternalServerError, err, "failed to fetch certifications")
		return
	}
	c.JSON(http.StatusOK, certifications)
}
