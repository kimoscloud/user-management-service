package organization

import (
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	ID                    string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name                  string `gorm:"type:varchar(255)"`
	CreatedBy             string `gorm:"type:uuid;not null"`
	Slug                  string `gorm:"type:varchar(255)"`
	BillingEmail          string `gorm:"type:varchar(255)"`
	URL                   string `gorm:"type:varchar(255)"`
	About                 string `gorm:"type:text"`
	LogoURL               string `gorm:"type:varchar(255)"`
	BackgroundImageURL    string `gorm:"type:varchar(255)"`
	Plan                  string `gorm:"type:varchar(255)"`
	CurrentPeriodStartsAt *time.Time
	CurrentPeriodEndsAt   *time.Time
	SubscriptionID        string             `gorm:"type:varchar(255)"`
	Status                string             `gorm:"type:varchar(255)"`
	Timezone              string             `gorm:"type:varchar(255)"`
	CreatedAt             time.Time          `gorm:"column:created_at;not null"`
	UpdatedAt             time.Time          `gorm:"column:updated_at;not null"`
	DeletedAt             gorm.DeletedAt     `gorm:"column:deleted_at;index"`
	UserOrganizations     []UserOrganization `gorm:"foreignKey:OrganizationID"`
}

func (Organization) TableName() string {
	return "Organizations"
}
