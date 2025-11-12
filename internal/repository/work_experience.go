package repository

import (
	"context"

	"github.com/GunarsK-portfolio/public-api/internal/models"
)

func (r *repository) GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.db.WithContext(ctx).Order("start_date DESC").Find(&experiences).Error
	return experiences, err
}
