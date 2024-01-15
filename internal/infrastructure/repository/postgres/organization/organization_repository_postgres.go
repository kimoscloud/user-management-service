package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
	"math"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}

func (repo *RepositoryPostgres) GetAll() ([]organization.Organization, error) {
	var organizations []organization.Organization
	if err := repo.db.Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil
}

// Move to another side this function
func calculateTotalPages(totalRows int, pageSize int) int {
	if pageSize == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalRows) / float64(pageSize)))
}

func (repo *RepositoryPostgres) GetPage(
	pageNumber int,
	pageSize int,
) (types.Page[organization.Organization], error) {
	var totalRows int64
	repo.db.Model(&organization.Organization{}).Count(&totalRows)
	totalPages := calculateTotalPages(int(totalRows), pageSize)
	var organizations []organization.Organization
	if err := repo.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&organizations).Error; err != nil {
		return types.EmptyPage[organization.Organization](), err
	}
	pageBuilder := new(types.PageBuilder[organization.Organization])
	return pageBuilder.SetItems(organizations).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		SetTotalPages(totalPages).
		Build(), nil
}

func (repo *RepositoryPostgres) GetByID(id string) (*organization.Organization, error) {
	var organization organization.Organization
	if err := repo.db.Where("id = ?", id).First(&organization).Error; err != nil {
		return nil, err
	}
	return &organization, nil
}
func (repo *RepositoryPostgres) GetAllByUserId(userId string) ([]organization.Organization, error) {
	var organizations []organization.Organization
	if err := repo.db.Joins("inner join UserOrganization on organization.id = UserOrganization.organization_id").Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil

}
func (repo *RepositoryPostgres) Create(organization *organization.Organization) (*organization.Organization, error) {
	if err := repo.db.Create(&organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}
func (repo *RepositoryPostgres) Update(organization *organization.Organization) (*organization.Organization, error) {
	if err := repo.db.Save(&organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}
func (repo *RepositoryPostgres) Delete(id string) error {
	if err := repo.db.Where("id = ?", id).Delete(&organization.Organization{}).Error; err != nil {
		return err
	}
	return nil
}
