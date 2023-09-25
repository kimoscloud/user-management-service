package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID               string         `json:"id" gorm:"column:id;primaryKey"`
	Name             string         `json:"name" gorm:"column:name"`
	Email            string         `json:"email" gorm:"column:email;uniqueIndex"`
	LastLogin        time.Time      `json:"lastLogin" gorm:"column:last_login"`
	CreatedAt        time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index"`
	BadLoginAttempts int            `json:"badLoginAttempts" gorm:"column:bad_login_attempts"`
	IsLocked         bool           `json:"isLocked" gorm:"column:is_locked"`
	IsSocialLogin    bool           `json:"isSocialLogin" gorm:"column:is_social_login"`
	Hash             string         `gorm:"column:hash"`
}

func (User) TableName() string {
	return "Users"
}
