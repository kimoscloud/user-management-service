package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID               string         `gorm:"column:id;primaryKey"`
	Name             string         `gorm:"column:name"`
	Email            string         `gorm:"column:email;uniqueIndex"`
	LastLogin        time.Time      `gorm:"column:last_login"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index"`
	BadLoginAttempts int            `gorm:"column:bad_login_attempts"`
	IsLocked         bool           `gorm:"column:is_locked"`
	IsSocialLogin    bool           `gorm:"column:is_social_login"`
	Hash             string         `gorm:"column:hash"`
}

func (User) TableName() string {
	return "Users"
}
