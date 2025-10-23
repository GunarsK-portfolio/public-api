package models

import "time"

type MiniatureTheme struct {
	ID           int64              `json:"id" gorm:"primaryKey"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	CoverImageID *int64             `json:"coverImageId,omitempty" gorm:"column:cover_image_id"`
	CoverImage   string             `json:"coverImage" gorm:"-"` // Computed URL field
	DisplayOrder int                `json:"displayOrder,omitempty" gorm:"column:display_order"`
	Miniatures   []MiniatureProject `json:"miniatures,omitempty" gorm:"foreignKey:ThemeID"`
	CreatedAt    time.Time          `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time          `json:"updatedAt" gorm:"column:updated_at"`
}

func (MiniatureTheme) TableName() string {
	return "miniatures.miniature_themes"
}

type MiniatureProject struct {
	ID            int64           `json:"id" gorm:"primaryKey"`
	ThemeID       *int64          `json:"themeId,omitempty" gorm:"column:theme_id"`
	Title         string          `json:"name"` // Frontend expects "name"
	Description   string          `json:"description"`
	CompletedDate *string         `json:"completedDate" gorm:"column:completed_date"`
	Scale         string          `json:"scale"`
	Manufacturer  string          `json:"manufacturer"`
	TimeSpent     *float64        `json:"timeSpent,omitempty" gorm:"column:time_spent"` // Hours as decimal
	Difficulty    string          `json:"difficulty"`
	DisplayOrder  int             `json:"displayOrder,omitempty" gorm:"column:display_order"`
	Techniques    []string        `json:"techniques,omitempty" gorm:"-"`          // Loaded separately from junction table
	Paints        []Paint         `json:"paints,omitempty" gorm:"-"`              // Loaded separately from junction table
	Files         []MiniatureFile `json:"-" gorm:"foreignKey:MiniatureProjectID"` // Database relation
	Images        []Image         `json:"images,omitempty" gorm:"-"`              // Computed for frontend
	CreatedAt     time.Time       `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time       `json:"updatedAt" gorm:"column:updated_at"`
}

func (MiniatureProject) TableName() string {
	return "miniatures.miniature_projects"
}

// MiniatureFile is the junction table between miniature projects and storage files
type MiniatureFile struct {
	ID                 int64        `json:"id" gorm:"primaryKey"`
	MiniatureProjectID int64        `json:"-" gorm:"column:miniature_project_id"`
	FileID             int64        `json:"-" gorm:"column:file_id"`
	Caption            string       `json:"caption"`
	DisplayOrder       int          `json:"displayOrder,omitempty" gorm:"column:display_order"`
	File               *StorageFile `json:"-" gorm:"foreignKey:FileID"` // Reference to actual file
	CreatedAt          time.Time    `json:"createdAt" gorm:"column:created_at"`
}

func (MiniatureFile) TableName() string {
	return "miniatures.miniature_files"
}

// StorageFile represents files stored in S3/MinIO
type StorageFile struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	S3Key    string `json:"-" gorm:"column:s3_key"`
	S3Bucket string `json:"-" gorm:"column:s3_bucket"`
	FileName string `json:"fileName" gorm:"column:file_name"`
	FileSize int64  `json:"fileSize" gorm:"column:file_size"`
	MimeType string `json:"mimeType" gorm:"column:mime_type"`
	FileType string `json:"fileType" gorm:"column:file_type"`
}

func (StorageFile) TableName() string {
	return "storage.files"
}

// Image is the simplified view for frontend
type Image struct {
	ID      int64  `json:"id"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

// Paint represents a paint used in a miniature project
type Paint struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Color        string `json:"color"` // Hex color code
}
