package repository

import (
	"context"
	"fmt"

	"github.com/GunarsK-portfolio/portfolio-common/utils"
	"github.com/GunarsK-portfolio/public-api/internal/models"
	"gorm.io/gorm"
)

// convertMiniatureFilesToImages transforms MiniatureFiles into frontend Image structures
func (r *repository) convertMiniatureFilesToImages(miniFiles []models.MiniatureFile) []models.Image {
	images := make([]models.Image, len(miniFiles))
	for i, file := range miniFiles {
		url := ""
		if file.File != nil {
			url = utils.BuildFileURL(r.filesAPIURL, file.File.FileType, file.File.S3Key)
		}
		images[i] = models.Image{
			ID:      file.ID,
			URL:     url,
			Caption: file.Caption,
		}
	}
	return images
}

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
		projects[i].Images = r.convertMiniatureFilesToImages(projects[i].MiniatureFiles)
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
	project.Images = r.convertMiniatureFilesToImages(project.MiniatureFiles)

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
			themes[i].Miniatures[j].Images = r.convertMiniatureFilesToImages(themes[i].Miniatures[j].MiniatureFiles)
		}
	}

	return themes, nil
}
