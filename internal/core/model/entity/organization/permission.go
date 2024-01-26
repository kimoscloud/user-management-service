package organization

import (
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	ID           string         `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string         `gorm:"column:name;type:varchar(255)"`
	Description  string         `gorm:"column:description;type:varchar(255)"`
	InternalName string         `gorm:"column:internal_name;type:varchar(255)"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Permission) TableName() string {
	return "Permissions"
}
