package team_member

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllByTeamId(teamId string) ([]organization.TeamMember, error)
	//Get only the scope of the organization.
	GetAllByUserId(userId string, orgId string) ([]organization.TeamMember, error)
	GetByTeamIdAndUserId(teamId string, userId string) (*organization.TeamMember, error)
	Create(teamMember *organization.TeamMember, tx *gorm.DB) (*organization.TeamMember, error)
	Update(teamMember *organization.TeamMember, tx *gorm.DB) (*organization.TeamMember, error)
	DeleteByTeamIdAndUserId(teamId string, userId string, tx *gorm.DB) error
	Delete(id string, tx *gorm.DB) error
	BeginTransaction() *gorm.DB
}
