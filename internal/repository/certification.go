package repository

import (
	"context"

	"github.com/GunarsK-portfolio/public-api/internal/models"
)

func (r *repository) GetAllCertifications(ctx context.Context) ([]models.Certification, error) {
	var certifications []models.Certification
	err := r.db.WithContext(ctx).Order("issue_date DESC").Find(&certifications).Error
	return certifications, err
}
