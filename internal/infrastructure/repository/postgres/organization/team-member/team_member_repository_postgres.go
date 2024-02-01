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

func (r *RepositoryPostgres) GetAllByUserId(userId string, orgId string) ([]organization.TeamMember, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryPostgres) GetByTeamIdAndUserId(teamId string, userId string) (*organization.TeamMember, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryPostgres) Create(teamMember *organization.TeamMember, tx *gorm.DB) (*organization.TeamMember, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryPostgres) Update(teamMember *organization.TeamMember, tx *gorm.DB) (*organization.TeamMember, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryPostgres) DeleteByTeamIdAndUserId(teamId string, userId string, tx *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryPostgres) Delete(id string, tx *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryPostgres) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
