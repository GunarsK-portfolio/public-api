package models

import "time"

type WorkExperience struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Company      string    `json:"company"`
	Position     string    `json:"position"`
	Description  string    `json:"description"`
	StartDate    string    `json:"startDate" gorm:"column:start_date"`
	EndDate      *string   `json:"endDate" gorm:"column:end_date"`
	IsCurrent    bool      `json:"isCurrent" gorm:"column:is_current"`
	DisplayOrder int       `json:"displayOrder,omitempty" gorm:"column:display_order"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (WorkExperience) TableName() string {
	return "portfolio.work_experience"
}
