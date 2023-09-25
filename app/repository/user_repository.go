package repository

import (
	"github.com/kimoscloud/user-management-service/app/domain"
	types "github.com/kimoscloud/value-types/domain"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetPage(pageNumber int, pageSize int) (types.Page[domain.User], error)
	GetByID(id string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(id string) error
}
