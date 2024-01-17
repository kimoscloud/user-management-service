package application

import (
	"gorm.io/gorm"
	"time"
)

type Application struct {
	gorm.Model
	ID                    string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	About                 string         `gorm:"type:text"`
	Name                  string         `gorm:"type:varchar(255)"`
	Slug                  string         `gorm:"type:varchar(255)"`
	LogoURL               string         `gorm:"type:varchar(255)"`
	AllowsJiraIntegration bool           `gorm:"type:boolean;default:false"`
	JiraProjectName       string         `gorm:"type:varchar(255)"`
	JiraProjectKey        string         `gorm:"type:varchar(255)"`
	IsPrivateRepo         bool           `gorm:"type:boolean;default:false"`
	Status                string         `gorm:"type:varchar(255);default:'pending'"`
	TemplateID            string         `gorm:"type:uuid;not null"`
	ProjectId             string         `gorm:"type:uuid;not null"`
	IsActive              bool           `gorm:"type:boolean;default:true"`
	CreatedBy             string         `gorm:"type:uuid;not null"`
	CreatedAt             time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt             time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt             gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
