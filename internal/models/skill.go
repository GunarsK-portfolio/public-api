package models

import "time"

type SkillType struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	DisplayOrder int       `json:"displayOrder,omitempty" gorm:"column:display_order"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (SkillType) TableName() string {
	return "portfolio.cl_skill_types"
}

type Skill struct {
	ID           int64      `json:"id" gorm:"primaryKey"`
	Skill        string     `json:"skill"`
	SkillTypeID  int64      `json:"skillTypeId,omitempty" gorm:"column:skill_type_id"`
	SkillType    *SkillType `json:"-" gorm:"foreignKey:SkillTypeID"` // Don't include in JSON by default
	Type         string     `json:"type" gorm:"-"`                   // Computed field for frontend
	IsVisible    bool       `json:"isVisible" gorm:"column:is_visible"`
	DisplayOrder int        `json:"displayOrder,omitempty" gorm:"column:display_order"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"column:updated_at"`
}

func (Skill) TableName() string {
	return "portfolio.skills"
}
