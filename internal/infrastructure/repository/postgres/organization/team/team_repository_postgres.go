package team

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}

func (repo *RepositoryPostgres) GetAllByOrgId(
	orgId string,
) ([]organization.Team, error) {
	var teams []organization.Team
	if err := repo.db.Where("organization_id = ?", orgId).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
func (repo *RepositoryPostgres) GetPageByOrgId(
	orgId string,
	pageNumber int,
	pageSize int,
) (types.Page[organization.Team], error) {
	var totalRows int64
	repo.db.Model(&organization.Team{}).Where("organization_id = ?", orgId).Count(&totalRows)
	var teams []organization.Team
	if err := repo.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&teams).Error; err != nil {
		return types.EmptyPage[organization.Team](), err
	}
	pageBuilder := new(types.PageBuilder[organization.Team])
	return pageBuilder.SetItems(teams).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		Build(), nil
}

func (repo *RepositoryPostgres) GetByID(id string) (*organization.Team, error) {
	var team organization.Team
	if err := repo.db.Where("id = ?", id).First(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
func (repo *RepositoryPostgres) Create(
	team *organization.Team,
	tx *gorm.DB,
) (*organization.Team, error) {
	if tx == nil {
		tx = repo.db
	}
	if err := tx.Create(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}
func (repo *RepositoryPostgres) Update(
	team *organization.Team,
	tx *gorm.DB,
) (*organization.Team, error) {
	if tx == nil {
		tx = repo.db
	}
	if err := tx.Save(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}
func (repo *RepositoryPostgres) Delete(id string, tx *gorm.DB) error {
	if tx == nil {
		tx = repo.db
	}
	return tx.Delete(&organization.Team{}, id).Error
}
func (repo *RepositoryPostgres) BeginTransaction() *gorm.DB {
	return repo.db.Begin()
}
