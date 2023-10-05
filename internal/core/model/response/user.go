package response

import (
	"time"
)

type UserLightDTO struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	LastLogin time.Time
	CreatedAt time.Time
}
