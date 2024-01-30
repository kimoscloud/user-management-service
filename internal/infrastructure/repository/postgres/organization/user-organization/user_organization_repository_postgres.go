package user_organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewUserOrganizationRepository(db *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}

func (repo *RepositoryPostgres) GetAll() ([]organization.UserOrganization, error) {
	var userOrganizations []organization.UserOrganization
	if err := repo.db.Find(&userOrganizations).Error; err != nil {
		return nil, err
	}
	return userOrganizations, nil
}

func (repo *RepositoryPostgres) GetPage(
	pageNumber int,
	pageSize int,
) (types.Page[organization.UserOrganization], error) {
	var totalRows int64
	repo.db.Model(&organization.UserOrganization{}).Count(&totalRows)
	var userOrganizations []organization.UserOrganization
	if err := repo.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&userOrganizations).Error; err != nil {
		return types.EmptyPage[organization.UserOrganization](), err
	}
	pageBuilder := new(types.PageBuilder[organization.UserOrganization])
	return pageBuilder.SetItems(userOrganizations).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		Build(), nil
}

func (repo *RepositoryPostgres) GetUserOrganizationByUserAndOrganizationWithRolesAndPermissions(userId, orgId string) (
	*organization.UserOrganization,
	error,
) {
	var userOrganization organization.UserOrganization
	if err := repo.db.Preload("Role.Permissions").
		Where("user_id = ?", userId).
		Where("organization_id = ?", orgId).
		Where("is_active = ? AND deleted_at IS NULL", true).
		First(&userOrganization).Error; err != nil {
		return nil, err
	}
	return &userOrganization, nil
}

func (repo *RepositoryPostgres) GetByID(id string) (*organization.UserOrganization, error) {
	var userOrganization organization.UserOrganization
	if err := repo.db.Where("id = ?", id).First(&userOrganization).Error; err != nil {
		return nil, err
	}
	return &userOrganization, nil
}

func (repo *RepositoryPostgres) GetAllByUserId(userId string) (
	[]organization.UserOrganization,
	error,
) {
	var userOrganizations []organization.UserOrganization
	if err := repo.db.Where("user_id = ?").Find(&userOrganizations).Error; err != nil {
		return nil, err
	}
	return userOrganizations, nil

}
func (repo *RepositoryPostgres) CreateUserOrganizations(
	userOrganization []organization.UserOrganization,
	tx *gorm.DB,
) error {
	if tx == nil {
		tx = repo.db
	}
	if err := tx.Create(&userOrganization).Error; err != nil {
		return err
	}
	return nil
}
func (repo *RepositoryPostgres) Create(
	userOrganization *organization.UserOrganization,
	tx *gorm.DB,
) (
	*organization.UserOrganization,
	error,
) {
	if tx == nil {
		tx = repo.db
	}
	if err := tx.Create(&userOrganization).Error; err != nil {
		return nil, err
	}
	return userOrganization, nil
}
func (repo *RepositoryPostgres) Update(userOrganization *organization.UserOrganization) (
	*organization.UserOrganization,
	error,
) {
	if err := repo.db.Save(&userOrganization).Error; err != nil {
		return nil, err
	}
	return userOrganization, nil
}
func (repo *RepositoryPostgres) Delete(id string) error {
	if err := repo.db.Where(
		"id = ?",
		id,
	).Delete(&organization.UserOrganization{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryPostgres) DeleteByUserId(userId string) error {
	if err := repo.db.Where(
		"user_id = ?",
		userId,
	).Delete(&organization.UserOrganization{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryPostgres) DeleteByOrganizationId(organizationId string) error {
	if err := repo.db.Where(
		"organization_id = ?",
		organizationId,
	).Delete(&organization.UserOrganization{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryPostgres) DeleteByOrganizationIdAndUserId(
	organizationId string,
	userId string,
) error {
	if err := repo.db.Where(
		"user_id = ? AND organization_id = ?",
		userId,
		organizationId,
	).Delete(&organization.UserOrganization{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepositoryPostgres) BeginTransaction() *gorm.DB {
	return repo.db.Begin()
}
