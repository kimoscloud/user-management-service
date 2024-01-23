package team

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllByOrgId(orgId string) ([]organization.Team, error)
	GetPageByOrgId(orgId string, page int, size int) (types.Page[organization.Team], error)
	GetByID(id string) (*organization.Team, error)
	Create(team *organization.Team, tx *gorm.DB) (*organization.Team, error)
	Update(team *organization.Team, tx *gorm.DB) (*organization.Team, error)
	Delete(id string, tx *gorm.DB) error
	BeginTransaction() *gorm.DB
}
