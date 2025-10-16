package models

import "time"

type Profile struct {
	ID           int64        `json:"id" gorm:"primaryKey"`
	FullName     string       `json:"name"`                // Frontend expects "name"
	Title        string       `json:"title"`
	Bio          string       `json:"tagline"`             // Frontend expects "tagline" for short bio
	Email        string       `json:"email"`
	Phone        string       `json:"phone,omitempty"`     // Optional
	Location     string       `json:"location,omitempty"`  // Optional
	AvatarFileID *int64       `json:"-" gorm:"column:avatar_file_id"`
	AvatarFile   *StorageFile `json:"-" gorm:"foreignKey:AvatarFileID"`
	AvatarURL    string       `json:"avatarUrl" gorm:"-"`  // Computed field
	ResumeFileID *int64       `json:"-" gorm:"column:resume_file_id"`
	ResumeFile   *StorageFile `json:"-" gorm:"foreignKey:ResumeFileID"`
	ResumeURL    string       `json:"resumeUrl,omitempty" gorm:"-"`  // Computed field
	CreatedAt    time.Time    `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time    `json:"updatedAt" gorm:"column:updated_at"`
}

func (Profile) TableName() string {
	return "portfolio.profile"
}
