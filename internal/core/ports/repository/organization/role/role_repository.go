package role

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]organization.Role, error)
	GetByID(id string) (*organization.Role, error)
	Create(organization *organization.Role) (*organization.Role, error)
	Update(organization *organization.Role) (*organization.Role, error)
	Delete(id string) error
	BeginTransaction() *gorm.DB
}
