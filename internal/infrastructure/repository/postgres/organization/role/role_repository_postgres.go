package role

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}

func (repo *RepositoryPostgres) GetAll() ([]organization.Role, error) {
	var roles []organization.Role
	if err := repo.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (repo *RepositoryPostgres) GetByID(id string) (*organization.Role, error) {
	var role organization.Role
	if err := repo.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (repo *RepositoryPostgres) Create(role *organization.Role) (
	*organization.Role,
	error,
) {
	if err := repo.db.Create(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (repo *RepositoryPostgres) Update(role *organization.Role) (
	*organization.Role,
	error,
) {
	if err := repo.db.Save(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (repo *RepositoryPostgres) Delete(id string) error {
	if err := repo.db.Where("id = ?", id).Delete(&organization.Role{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryPostgres) BeginTransaction() *gorm.DB {
	return repo.db.Begin()
}
