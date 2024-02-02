package organization

import (
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	ID                    string             `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name                  string             `gorm:"column:name;type:varchar(255)"`
	CreatedBy             string             `gorm:"column:created_by;type:uuid;not null"`
	Slug                  string             `gorm:"column:slug;type:varchar(255)"`
	BillingEmail          string             `gorm:"column:billing_email;type:varchar(255)"`
	URL                   string             `gorm:"column:url;type:varchar(255)"`
	About                 string             `gorm:"column:about;type:text"`
	LogoURL               string             `gorm:"column:logo_url;type:varchar(255)"`
	BackgroundImageURL    string             `gorm:"column:background_image_url;type:varchar(255)"`
	Plan                  string             `gorm:"column:plan;type:varchar(255)"`
	CurrentPeriodStartsAt *time.Time         `gorm:"column:current_period_starts_at"`
	CurrentPeriodEndsAt   *time.Time         `gorm:"column:current_period_ends_at"`
	SubscriptionID        string             `gorm:"column:subscription_id;type:varchar(255)"`
	Status                string             `gorm:"column:status;type:varchar(255)"`
	Timezone              string             `gorm:"column:timezone;type:varchar(255)"`
	CreatedAt             time.Time          `gorm:"column:created_at;column:created_at;not null"`
	UpdatedAt             time.Time          `gorm:"column:updated_at;column:updated_at;not null"`
	DeletedAt             gorm.DeletedAt     `gorm:"column:deleted_at;column:deleted_at;index"`
	UserOrganizations     []UserOrganization `gorm:"foreignKey:OrganizationID"`
}

func (Organization) TableName() string {
	return "Organizations"
}
