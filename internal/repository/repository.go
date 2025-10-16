package repository

import (
	"github.com/GunarsK-portfolio/public-api/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetProfile() (*models.Profile, error)
	GetAllWorkExperience() ([]models.WorkExperience, error)
	GetAllCertifications() ([]models.Certification, error)
	GetAllSkills() ([]models.Skill, error)
	GetAllProjects() ([]models.PortfolioProject, error)
	GetProjectByID(id int64) (*models.PortfolioProject, error)
	GetAllMiniatureProjects() ([]models.MiniatureProject, error)
	GetMiniatureProjectByID(id int64) (*models.MiniatureProject, error)
	GetAllMiniatureThemes() ([]models.MiniatureTheme, error)
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
