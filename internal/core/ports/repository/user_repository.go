package repository

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/auth"
	types "github.com/kimoscloud/value-types/domain"
)

type Repository interface {
	GetAll() ([]auth.User, error)
	GetPage(pageNumber int, pageSize int) (types.Page[auth.User], error)
	GetUserByEmailLike(email string, limit int) ([]auth.User, error)
	GetByID(id string) (*auth.User, error)
	GetByEmail(email string) (*auth.User, error)
	Create(user *auth.User) (*auth.User, error)
	Update(user *auth.User) (*auth.User, error)
	Delete(id string) error
	FindUsersByEmails(emails []string) ([]auth.User, error)
	IncrementBadLoginAttempts(id string) error
	LockUser(id string) error
}
