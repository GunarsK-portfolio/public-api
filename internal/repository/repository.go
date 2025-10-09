package repository

import (
	"github.com/GunarsK-portfolio/public-api/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetProfile() (*models.Profile, error)
	GetAllWorkExperience() ([]models.WorkExperience, error)
	GetAllCertifications() ([]models.Certification, error)
	GetAllMiniatureProjects() ([]models.MiniatureProject, error)
	GetMiniatureProjectByID(id int64) (*models.MiniatureProject, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetProfile() (*models.Profile, error) {
	var profile models.Profile
	err := r.db.First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *repository) GetAllWorkExperience() ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.db.Order("display_order ASC, start_date DESC").Find(&experiences).Error
	return experiences, err
}

func (r *repository) GetAllCertifications() ([]models.Certification, error) {
	var certifications []models.Certification
	err := r.db.Order("display_order ASC, issue_date DESC").Find(&certifications).Error
	return certifications, err
}

func (r *repository) GetAllMiniatureProjects() ([]models.MiniatureProject, error) {
	var projects []models.MiniatureProject
	err := r.db.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("display_order ASC")
	}).Order("display_order ASC, completed_date DESC").Find(&projects).Error
	return projects, err
}

func (r *repository) GetMiniatureProjectByID(id int64) (*models.MiniatureProject, error) {
	var project models.MiniatureProject
	err := r.db.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("display_order ASC")
	}).First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}
