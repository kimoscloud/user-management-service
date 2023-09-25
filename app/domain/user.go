package domain

import "time"

type User struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	LastLogin        time.Time `json:"lastLogin"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	DeletedAt        time.Time `json:"deletedAt"`
	BadLoginAttempts int       `json:"badLoginAttempts"`
	IsLocked         bool      `json:"isLocked"`
	IsSocialLogin    bool      `json:"isSocialLogin"`
	Hash             string
}
