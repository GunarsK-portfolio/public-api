package repository

import "github.com/GunarsK-portfolio/public-api/internal/models"

func (r *repository) GetAllCertifications() ([]models.Certification, error) {
	var certifications []models.Certification
	err := r.db.Order("issue_date DESC").Find(&certifications).Error
	return certifications, err
}
