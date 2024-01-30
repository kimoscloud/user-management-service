package user_organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]organization.UserOrganization, error)
	GetPage(pageNumber int, pageSize int) (types.Page[organization.UserOrganization], error)
	GetByID(id string) (*organization.UserOrganization, error)
	GetAllByUserId(userId string) ([]organization.UserOrganization, error)
	GetUserOrganizationByUserAndOrganizationWithRolesAndPermissions(userId, orgId string) (
		*organization.UserOrganization,
		error,
	)
	Create(
		userOrganization *organization.UserOrganization,
		tx *gorm.DB,
	) (*organization.UserOrganization, error)
	CreateUserOrganizations(userOrganization []organization.UserOrganization, tx *gorm.DB) error
	Update(userOrganization *organization.UserOrganization) (*organization.UserOrganization, error)
	Delete(id string) error
	DeleteByOrganizationIdAndUserId(organizationId string, userId string) error
	BeginTransaction() *gorm.DB
}
