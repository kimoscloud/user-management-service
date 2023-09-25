package repository

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	types "github.com/kimoscloud/value-types/domain"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetPage(pageNumber int, pageSize int) (types.Page[entity.User], error)
	GetByID(id string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id string) error
}
