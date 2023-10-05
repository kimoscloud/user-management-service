package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                         string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName                  string         `gorm:"column:first_name"`
	LastName                   string         `gorm:"column:last_name"`
	AcceptTermsAndConditions   bool           `gorm:"column:accept_terms_and_conditions"`
	AcceptTermsAndConditionsAt time.Time      `gorm:"column:accept_terms_and_conditions_at"`
	PhotoUrl                   string         `gorm:"column:photo_url"`
	Phone                      string         `gorm:"column:phone"`
	Timezone                   string         `gorm:"column:timezone"`
	Email                      string         `gorm:"column:email;uniqueIndex"`
	LastLogin                  time.Time      `gorm:"column:last_login"`
	CreatedAt                  time.Time      `gorm:"column:created_at"`
	UpdatedAt                  time.Time      `gorm:"column:updated_at"`
	EmailVerifiedAt            time.Time      `gorm:"column:email_verified_at"`
	DeletedAt                  gorm.DeletedAt `gorm:"column:deleted_at;index"`
	BadLoginAttempts           int            `gorm:"column:bad_attempts"`
	IsLocked                   bool           `gorm:"column:is_locked"`
	Hash                       string         `gorm:"column:hash"`
}

func (User) TableName() string {
	return "Users"
}
