package entity

import (
	"gorm.io/gorm"
	"time"
)

type Technology struct {
	ID          string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string         `gorm:"type:varchar(255)"`
	Description string         `gorm:"type:varchar(255)"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
