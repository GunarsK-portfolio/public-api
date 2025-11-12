package repository

import (
	"context"

	"github.com/GunarsK-portfolio/public-api/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetProfile(ctx context.Context) (*models.Profile, error)
	GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error)
	GetAllCertifications(ctx context.Context) ([]models.Certification, error)
	GetAllSkills(ctx context.Context) ([]models.Skill, error)
	GetAllProjects(ctx context.Context) ([]models.PortfolioProject, error)
	GetProjectByID(ctx context.Context, id int64) (*models.PortfolioProject, error)
	GetAllMiniatureProjects(ctx context.Context) ([]models.MiniatureProject, error)
	GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error)
	GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error)
}

type repository struct {
	db          *gorm.DB
	filesAPIURL string
}

func New(db *gorm.DB, filesAPIURL string) Repository {
	return &repository{
		db:          db,
		filesAPIURL: filesAPIURL,
	}
}
