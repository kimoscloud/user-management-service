package entity

import (
	"gorm.io/gorm"
	"time"
)

type Tenant struct {
	ID        string         `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string         `gorm:"column:name"`
	Subdomain string         `gorm:"column:subdomain"`
	IsActive  bool           `gorm:"column:is_active;not null;default:true"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
