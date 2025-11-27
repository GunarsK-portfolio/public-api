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
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC, miniatures.miniature_files.id ASC")
		}).
		Preload("MiniatureFiles.File").
		Preload("Techniques.Technique").
		Preload("Paints.Paint").
		Order("display_order ASC, id ASC").
		Find(&projects).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all miniature projects: %w", err)
	}

	// Convert MiniatureFiles to Images for frontend
	for i := range projects {
		projects[i].Images = utils.ConvertMiniatureFilesToImages(projects[i].MiniatureFiles, r.filesAPIURL)
	}

	return projects, nil
}

func (r *repository) GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error) {
	var project models.MiniatureProject
	err := r.db.WithContext(ctx).
		Preload("Theme").
		Preload("MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC, miniatures.miniature_files.id ASC")
		}).
		Preload("MiniatureFiles.File").
		Preload("Techniques.Technique").
		Preload("Paints.Paint").
		First(&project, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get miniature project by id %d: %w", id, err)
	}

	// Convert MiniatureFiles to Images for frontend
	project.Images = utils.ConvertMiniatureFilesToImages(project.MiniatureFiles, r.filesAPIURL)

	return &project, nil
}

func (r *repository) GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error) {
	var themes []models.MiniatureTheme
	err := r.db.WithContext(ctx).
		Preload("CoverImageFile").
		Order("display_order ASC, id ASC").
		Find(&themes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all miniature themes: %w", err)
	}

	// Populate cover image URLs
	for i := range themes {
		utils.PopulateFileURL(themes[i].CoverImageFile, r.filesAPIURL)
	}

	return themes, nil
}

func (r *repository) GetMiniatureThemeByID(ctx context.Context, id int64) (*models.MiniatureTheme, error) {
	var theme models.MiniatureTheme
	err := r.db.WithContext(ctx).
		Preload("CoverImageFile").
		Preload("Miniatures", func(db *gorm.DB) *gorm.DB {
			return db.Order("display_order ASC, id ASC")
		}).
		Preload("Miniatures.MiniatureFiles", func(db *gorm.DB) *gorm.DB {
			return db.Order("miniatures.miniature_files.display_order ASC, miniatures.miniature_files.id ASC")
		}).
		Preload("Miniatures.MiniatureFiles.File").
		Preload("Miniatures.Techniques.Technique").
		Preload("Miniatures.Paints.Paint").
		First(&theme, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get miniature theme by id %d: %w", id, err)
	}

	// Populate cover image URL
	utils.PopulateFileURL(theme.CoverImageFile, r.filesAPIURL)

	// Convert MiniatureFiles to Images for each miniature
	for i := range theme.Miniatures {
		theme.Miniatures[i].Images = utils.ConvertMiniatureFilesToImages(theme.Miniatures[i].MiniatureFiles, r.filesAPIURL)
	}

	return &theme, nil
}
