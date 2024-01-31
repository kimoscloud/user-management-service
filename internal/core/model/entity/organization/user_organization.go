package organization

import (
	"gorm.io/gorm"
	"time"
)

type UserOrganization struct {
	ID             string         `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         string         `gorm:"column:user_id;type:uuid;not null"`
	OrganizationID string         `gorm:"column:organization_id;type:uuid;not null"`
	IsActive       bool           `gorm:"column:is_active;type:boolean;default:true"`
	Role           Role           `gorm:"foreignKey:RoleID"`
	RoleID         string         `gorm:"column:role_id;type:uuid;not null"`
	Status         string         `gorm:"column:status;type:varchar(255);default:'pending'"`
	InvitedAt      time.Time      `gorm:"column:invited_at;not null"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (UserOrganization) TableName() string {
	return "Organization_Users"
}

func (o UserOrganization) hasPermission(permission string) bool {
	if o.Role.ID == "" || o.Role.Permissions == nil {
		return false
	}
	for _, p := range o.Role.Permissions {
		if p.ID == permission {
			return true
		}
	}
	return false
}

func (o UserOrganization) CheckIfOrgUserHasPermissions(permissions []string) bool {
	for _, permission := range permissions {
		if !o.hasPermission(permission) {
			return false
		}
	}
	return true
}
