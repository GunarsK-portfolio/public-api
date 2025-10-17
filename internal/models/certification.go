package models

import "time"

type Certification struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Issuer        string    `json:"issuer"`
	IssueDate     string    `json:"issueDate" gorm:"column:issue_date"`
	ExpiryDate    *string   `json:"expiryDate,omitempty" gorm:"column:expiry_date"`
	CredentialID  string    `json:"credentialId,omitempty" gorm:"column:credential_id"`
	CredentialURL string    `json:"credentialUrl" gorm:"column:credential_url"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (Certification) TableName() string {
	return "portfolio.certifications"
}
