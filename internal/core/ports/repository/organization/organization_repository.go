package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	types "github.com/kimoscloud/value-types/domain"
)

type Repository interface {
	GetAll() ([]organization.Organization, error)
	GetPage(pageNumber int, pageSize int) (types.Page[organization.Organization], error)
	GetByID(id string) (*organization.Organization, error)
	GetAllByUserId(userId string) (*organization.Organization, error)
	Create(organization *organization.Organization) (*organization.Organization, error)
	Update(organization *organization.Organization) (*organization.Organization, error)
	Delete(id string) error
}
