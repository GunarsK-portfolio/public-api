package repository

import (
	"fmt"
	"github.com/GunarsK-portfolio/public-api/internal/models"
	"gorm.io/gorm"
)

func (r *repository) GetAllProjects() ([]models.PortfolioProject, error) {
	var projects []models.PortfolioProject
	err := r.db.
		Preload("ImageFile").
		Preload("Technologies", func(db *gorm.DB) *gorm.DB {
			return db.Preload("SkillType").Order("portfolio.skills.display_order ASC")
		}).
		Order("featured DESC, display_order ASC, start_date DESC").
		Find(&projects).Error

	// Populate Technology.Type field and construct image URLs
	for i := range projects {
		// Construct image URL
		if projects[i].ImageFile != nil {
			projects[i].ImageURL = fmt.Sprintf("%s/files/%s/%s", r.filesAPIURL, projects[i].ImageFile.FileType, projects[i].ImageFile.S3Key)
		}

		// Populate technology types
		for j := range projects[i].Technologies {
			if projects[i].Technologies[j].SkillType != nil {
				projects[i].Technologies[j].Type = projects[i].Technologies[j].SkillType.Name
			}
		}
	}

	return projects, err
}

func (r *repository) GetProjectByID(id int64) (*models.PortfolioProject, error) {
	var project models.PortfolioProject
	err := r.db.
		Preload("ImageFile").
		Preload("Technologies", func(db *gorm.DB) *gorm.DB {
			return db.Preload("SkillType").Order("portfolio.skills.display_order ASC")
		}).
		First(&project, id).Error
	if err != nil {
		return nil, err
	}

	// Construct image URL
	if project.ImageFile != nil {
		project.ImageURL = fmt.Sprintf("%s/files/%s/%s", r.filesAPIURL, project.ImageFile.FileType, project.ImageFile.S3Key)
	}

	// Populate Technology.Type field
	for j := range project.Technologies {
		if project.Technologies[j].SkillType != nil {
			project.Technologies[j].Type = project.Technologies[j].SkillType.Name
		}
	}

	return &project, nil
}
