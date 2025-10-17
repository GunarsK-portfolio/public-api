package repository

import "github.com/GunarsK-portfolio/public-api/internal/models"

func (r *repository) GetAllSkills() ([]models.Skill, error) {
	var skills []models.Skill
	err := r.db.
		Preload("SkillType").
		Where("is_visible = ?", true).
		Order("skill_type_id ASC, display_order ASC").
		Find(&skills).Error

	// Populate the Type field from SkillType.Name
	for i := range skills {
		if skills[i].SkillType != nil {
			skills[i].Type = skills[i].SkillType.Name
		}
	}

	return skills, err
}
