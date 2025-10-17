package models

import "time"

type PortfolioProject struct {
	ID              int64        `json:"id" gorm:"primaryKey"`
	Title           string       `json:"title"`
	Category        string       `json:"category"`
	Description     string       `json:"description"`
	LongDescription string       `json:"longDescription" gorm:"column:long_description"`
	ImageFileID     *int64       `json:"-" gorm:"column:image_file_id"`
	ImageFile       *StorageFile `json:"-" gorm:"foreignKey:ImageFileID"`
	ImageURL        string       `json:"imageUrl" gorm:"-"`                         // Computed field
	GithubURL       string       `json:"githubUrl" gorm:"column:github_url"`
	LiveURL         string       `json:"liveUrl" gorm:"column:live_url"`
	StartDate       *string      `json:"startDate" gorm:"column:start_date"`
	EndDate         *string      `json:"endDate" gorm:"column:end_date"`
	IsOngoing       bool         `json:"isOngoing" gorm:"column:is_ongoing"`
	TeamSize        *int         `json:"teamSize" gorm:"column:team_size"`
	Role            string       `json:"role"`
	Featured        bool         `json:"featured"`
	Features        []string     `json:"features" gorm:"type:jsonb"`
	Challenges      []string     `json:"challenges" gorm:"type:jsonb"`
	Learnings       []string     `json:"learnings" gorm:"type:jsonb"`
	DisplayOrder    int          `json:"displayOrder,omitempty" gorm:"column:display_order"`
	Technologies    []Skill      `json:"technologies,omitempty" gorm:"many2many:portfolio.project_technologies;joinForeignKey:ProjectID;joinReferences:SkillID"`
	CreatedAt       time.Time    `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt       time.Time    `json:"updatedAt" gorm:"column:updated_at"`
}

func (PortfolioProject) TableName() string {
	return "portfolio.portfolio_projects"
}
