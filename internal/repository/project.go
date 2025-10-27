package repository

import (
	"github.com/GunarsK-portfolio/portfolio-common/utils"
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
	if err != nil {
		return nil, err
	}

	// Populate Technology.Type field and image URLs
	for i := range projects {
		// Populate image URL using helper
		utils.PopulateFileURL(projects[i].ImageFile, r.filesAPIURL)

		// Populate technology types
		for j := range projects[i].Technologies {
			if projects[i].Technologies[j].SkillType != nil {
				projects[i].Technologies[j].Type = projects[i].Technologies[j].SkillType.Name
			}
		}
	}

	return projects, nil
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

	// Populate image URL using helper
	utils.PopulateFileURL(project.ImageFile, r.filesAPIURL)

	// Populate Technology.Type field
	for j := range project.Technologies {
		if project.Technologies[j].SkillType != nil {
			project.Technologies[j].Type = project.Technologies[j].SkillType.Name
		}
	}

	return &project, nil
}
