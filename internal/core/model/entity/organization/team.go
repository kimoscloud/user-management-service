package organization

import (
	"gorm.io/gorm"
	"time"
)

type Team struct {
	ID             string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string         `gorm:"type:varchar(255)"`
	ImageUrl       string         `gorm:"type:varchar(255)"`
	Slug           string         `gorm:"type:varchar(255)"`
	About          string         `gorm:"type:text"`
	OrganizationID string         `gorm:"type:uuid;not null"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Team) TableName() string {
	return "Teams"
}
