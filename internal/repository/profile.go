package repository

import (
	"fmt"

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

	// Construct avatar URL
	if profile.AvatarFile != nil {
		profile.AvatarURL = fmt.Sprintf("%s/files/%s/%s", r.filesAPIURL, profile.AvatarFile.FileType, profile.AvatarFile.S3Key)
	}

	// Construct resume URL
	if profile.ResumeFile != nil {
		profile.ResumeURL = fmt.Sprintf("%s/files/%s/%s", r.filesAPIURL, profile.ResumeFile.FileType, profile.ResumeFile.S3Key)
	}

	return &profile, nil
}
