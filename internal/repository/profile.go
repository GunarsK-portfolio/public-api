package repository

import (
	"github.com/GunarsK-portfolio/portfolio-common/utils"
	"github.com/GunarsK-portfolio/public-api/internal/models"
)

func (r *repository) GetProfile() (*models.Profile, error) {
	var profile models.Profile
	err := r.db.
		Preload("AvatarFile").
		Preload("ResumeFile").
		First(&profile).Error
	if err != nil {
		return nil, err
	}

	// Populate file URLs using helper
	utils.PopulateFileURL(profile.AvatarFile, r.filesAPIURL)
	utils.PopulateFileURL(profile.ResumeFile, r.filesAPIURL)

	return &profile, nil
}
