package repository

import "github.com/GunarsK-portfolio/public-api/internal/models"

func (r *repository) GetAllWorkExperience() ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.db.Order("start_date DESC").Find(&experiences).Error
	return experiences, err
}
