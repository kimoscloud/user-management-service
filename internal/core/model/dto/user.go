package dto

import (
	"gorm.io/gorm"
	"time"
)

type UserDTO struct {
	ID               string
	Name             string
	Email            string
	LastLogin        time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	BadLoginAttempts int
	IsLocked         bool
	IsSocialLogin    bool
	Hash             string
}
