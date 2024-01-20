package environment

import "time"

type Environment struct {
	ID            string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name          string    `gorm:"type:varchar(255)"`
	Description   string    `gorm:"type:text"`
	ApplicationID string    `gorm:"type:uuid;not null"`
	IsActive      bool      `gorm:"type:boolean;default:true"`
	CreatedBy     string    `gorm:"type:uuid;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null"`
	DeletedAt     time.Time `gorm:"column:deleted_at;index"`
}
