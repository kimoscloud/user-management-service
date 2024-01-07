package organization

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	gorm.Model
	ID             string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string         `gorm:"type:varchar(255)"`
	Description    string         `gorm:"type:varchar(255)"`
	Editable       bool           `gorm:"type:boolean;default:true"`
	OrganizationID string         `gorm:"type:uuid;not null"`
	UserID         string         `gorm:"type:uuid;not null"`
	Permissions    []Permission   `gorm:"many2many:role_permissions;"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Role) TableName() string {
	return "Roles"
}
