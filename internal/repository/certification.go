package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/public-api/internal/models"
)

func (r *repository) GetAllCertifications(ctx context.Context) ([]models.Certification, error) {
	var certifications []models.Certification
	err := r.db.WithContext(ctx).Order("issue_date DESC").Find(&certifications).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all certifications: %w", err)
	}
	return certifications, nil
}
