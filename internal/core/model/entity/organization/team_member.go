package organization

import (
	"gorm.io/gorm"
	"time"
)

type TeamMember struct {
	ID        string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TeamID    string         `gorm:"type:uuid;not null"`
	UserID    string         `gorm:"type:uuid;not null"`
	RoleID    string         `gorm:"type:uuid;not null"`
	IsActive  bool           `gorm:"type:boolean;default:true"`
	Status    string         `gorm:"type:varchar(255);default:'pending'"`
	InvitedAt time.Time      `gorm:"column:invited_at;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
