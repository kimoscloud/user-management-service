package postgres

import (
	"errors"
	"github.com/kimoscloud/user-management-service/app/domain"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
	"math"
)

type UserRepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

func (repo *UserRepositoryPostgres) GetAll() ([]domain.User, error) {
	var users []domain.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func calculateTotalPages(totalRows int, pageSize int) int {
	if pageSize == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalRows) / float64(pageSize)))
}

func (repo *UserRepositoryPostgres) GetPage(pageNumber int, pageSize int) (types.Page[domain.User], error) {
	var totalRows int64
	repo.db.Model(&domain.User{}).Count(&totalRows)
	totalPages := calculateTotalPages(int(totalRows), pageSize)
	var users []domain.User
	if err := repo.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return types.EmptyPage[domain.User](), err
	}
	pageBuilder := new(types.PageBuilder[domain.User])
	return pageBuilder.SetItems(users).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		SetTotalPages(totalPages).
		Build(), nil
}
func (repo *UserRepositoryPostgres) GetByID(id string) (*domain.User, error) {
	return nil, errors.New("unimplemented")
}
func (repo *UserRepositoryPostgres) Create(user *domain.User) (*domain.User, error) {
	return nil, errors.New("unimplemented")
}
func (repo *UserRepositoryPostgres) Update(user *domain.User) (*domain.User, error) {
	return nil, errors.New("unimplemented")
}
func (repo *UserRepositoryPostgres) Delete(id string) error {
	return errors.New("unimplemented")
}
