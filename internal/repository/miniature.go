package repository

import (
	"context"

	"github.com/GunarsK-portfolio/portfolio-common/utils"
	"github.com/GunarsK-portfolio/public-api/internal/models"
	"gorm.io/gorm"
)

func (r *repository) GetAllMiniatureProjects(ctx context.Context) ([]models.MiniatureProject, error) {
	var projects []models.MiniatureProject
	err := r.db.WithContext(ctx).
		Preload("Files.File").
		Preload("Files", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		}).
		Order("display_order ASC, completed_date DESC").
		Find(&projects).Error

	// Convert Files to Images for frontend
	for i := range projects {
		projects[i].Images = make([]models.Image, len(projects[i].Files))
		for j, file := range projects[i].Files {
			url := ""
			if file.File != nil {
				url = utils.BuildFileURL(r.filesAPIURL, file.File.FileType, file.File.S3Key)
			}
			projects[i].Images[j] = models.Image{
				ID:      file.ID,
				URL:     url,
				Caption: file.Caption,
			}
		}
	}

	return projects, err
}

func (r *repository) GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error) {
	var project models.MiniatureProject
	err := r.db.WithContext(ctx).
		Preload("Files.File").
		Preload("Files", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		}).
		First(&project, id).Error
	if err != nil {
		return nil, err
	}

	// Convert Files to Images for frontend
	project.Images = make([]models.Image, len(project.Files))
	for j, file := range project.Files {
		url := ""
		if file.File != nil {
			url = utils.BuildFileURL(r.filesAPIURL, file.File.FileType, file.File.S3Key)
		}
		project.Images[j] = models.Image{
			ID:      file.ID,
			URL:     url,
			Caption: file.Caption,
		}
	}

	return &project, nil
}

func (r *repository) GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error) {
	var themes []models.MiniatureTheme
	err := r.db.WithContext(ctx).
		Preload("Miniatures", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC, completed_date DESC")
		}).
		Order("display_order ASC").
		Find(&themes).Error
	return themes, err
}
