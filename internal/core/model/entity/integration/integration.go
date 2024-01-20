package integration

import (
	"gorm.io/gorm"
	"time"
)

type Integration struct {
	ID             string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrganizationID string         `gorm:"type:uuid;not null"`
	ProjectID      string         `gorm:"type:uuid;not null"`
	UserID         string         `gorm:"type:uuid;not null"`
	Level          string         `gorm:"type:varchar(255);not null"` //PROJECT, ORGANIZATION, ACCOUNT (For users)
	Integration    string         `gorm:"type:varchar(255);not null"`
	IsActive       bool           `gorm:"type:boolean;default:true"`
	Settings       string         `gorm:"type:jsonb;"`
	CreatedBy      string         `gorm:"type:uuid;not null"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
