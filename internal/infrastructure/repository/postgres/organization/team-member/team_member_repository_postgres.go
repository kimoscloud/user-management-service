package team_member

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewTeamMemberRepository(db *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{
		db: db,
	}
}

func (r *RepositoryPostgres) GetAllByTeamId(teamId string) ([]organization.TeamMember, error) {
	var teamMembers []organization.TeamMember
	r.db.Model(&organization.TeamMember{}).Where("team_id = ?", teamId).Find(&teamMembers)
	return teamMembers, nil
}

func (r *RepositoryPostgres) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
