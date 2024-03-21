package repository

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	types "github.com/kimoscloud/value-types/domain"
)

type Repository interface {
	GetAll() ([]entity.User, error)
	GetPage(pageNumber int, pageSize int) (types.Page[entity.User], error)
	GetUserByEmailLike(email string, limit int) ([]entity.User, error)
	GetByID(id string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id string) error
	FindUsersByEmails(emails []string) ([]entity.User, error)
	IncrementBadLoginAttempts(id string) error
	LockUser(id string) error
}
