package organization

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID             string         `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string         `gorm:"column:name;type:varchar(255)"`
	Description    string         `gorm:"column:description;type:varchar(255)"`
	Editable       bool           `gorm:"column:editable;type:boolean;default:true"`
	OrganizationID string         `gorm:"column:organization_id;type:uuid;not null"`
	Organization   Organization   `gorm:"foreignKey:OrganizationID"`
	CreatedBy      string         `gorm:"column:created_by;type:uuid;not null"`
	Permissions    []Permission   `gorm:"many2many:Role_Permissions;foreignKey:ID;joinForeignKey:role_id;References:ID;joinReferences:permission_id"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Role) TableName() string {
	return "Roles"
}
