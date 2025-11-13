package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/public-api/internal/models"
)

func (r *repository) GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.db.WithContext(ctx).Order("start_date DESC").Find(&experiences).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all work experience: %w", err)
	}
	return experiences, nil
}
