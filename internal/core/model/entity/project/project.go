package project

import (
	"gorm.io/gorm"
	"time"
)

// Project it's a something that englobes applications.
// One project can have many applications.
// One project can have many teams (owners)
// One user can have specific role (Belongs to a team in a project or not)
type Project struct {
	ID             string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string         `gorm:"type:varchar(255);not null"`
	Description    string         `gorm:"type:text"`
	Slug           string         `gorm:"type:varchar(255);not null"`
	ImageUrl       string         `gorm:"type:varchar(255);"`
	CreatedBy      string         `gorm:"type:uuid;not null"`
	IsActive       bool           `gorm:"type:boolean;default:true"`
	OrganizationID string         `gorm:"type:uuid;"` // If the project belongs to an organization
	UserID         string         `gorm:"type:uuid;"` // if the project belongs to a user
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Project) TableName() string {
	return "Projects"
}
