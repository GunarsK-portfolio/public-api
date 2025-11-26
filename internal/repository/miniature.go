package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/portfolio-common/utils"
	"github.com/GunarsK-portfolio/public-api/internal/models"
	"gorm.io/gorm"
)

func (r *repository) GetAllMiniatureProjects(ctx context.Context) ([]models.MiniatureProject, error) {
	var projects []models.MiniatureProject
	err := r.db.WithContext(ctx).
		Preload("MiniatureFiles.File").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		}).
		Order("display_order ASC, completed_date DESC").
		Find(&projects).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all miniature projects: %w", err)
	}

	// Convert MiniatureFiles to Images for frontend
	for i := range projects {
		projects[i].Images = make([]models.Image, len(projects[i].MiniatureFiles))
		for j, file := range projects[i].MiniatureFiles {
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

	return projects, nil
}

func (r *repository) GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error) {
	var project models.MiniatureProject
	err := r.db.WithContext(ctx).
		Preload("Theme").
		Preload("MiniatureFiles.File").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		}).
		First(&project, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get miniature project by id %d: %w", id, err)
	}

	// Convert MiniatureFiles to Images for frontend
	project.Images = make([]models.Image, len(project.MiniatureFiles))
	for j, file := range project.MiniatureFiles {
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
		Preload("CoverImageFile").
		Preload("Miniatures.MiniatureFiles.File").
		Preload("Miniatures.MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC")
		}).
		Preload("Miniatures", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC, completed_date DESC")
		}).
		Order("display_order ASC").
		Find(&themes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all miniature themes: %w", err)
	}

	// Populate URLs
	for i := range themes {
		utils.PopulateFileURL(themes[i].CoverImageFile, r.filesAPIURL)

		// Convert MiniatureFiles to Images for each miniature
		for j := range themes[i].Miniatures {
			miniFiles := themes[i].Miniatures[j].MiniatureFiles
			themes[i].Miniatures[j].Images = make([]models.Image, len(miniFiles))
			for k, file := range miniFiles {
				url := ""
				if file.File != nil {
					url = utils.BuildFileURL(r.filesAPIURL, file.File.FileType, file.File.S3Key)
				}
				themes[i].Miniatures[j].Images[k] = models.Image{
					ID:      file.ID,
					URL:     url,
					Caption: file.Caption,
				}
			}
		}
	}

	return themes, nil
}
