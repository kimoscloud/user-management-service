package postgres

import (
	"errors"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
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

func (repo *UserRepositoryPostgres) GetAll() ([]entity.User, error) {
	var users []entity.User
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

func (repo *UserRepositoryPostgres) GetPage(
	pageNumber int,
	pageSize int,
) (types.Page[entity.User], error) {
	var totalRows int64
	repo.db.Model(&entity.User{}).Count(&totalRows)
	totalPages := calculateTotalPages(int(totalRows), pageSize)
	var users []entity.User
	if err := repo.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return types.EmptyPage[entity.User](), err
	}
	pageBuilder := new(types.PageBuilder[entity.User])
	return pageBuilder.SetItems(users).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		SetTotalPages(totalPages).
		Build(), nil
}

func (repo *UserRepositoryPostgres) GetByID(id string) (*entity.User, error) {
	var user entity.User
	result := repo.db.Model(&entity.User{}).Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepositoryPostgres) GetByEmail(email string) (
	*entity.User,
	error,
) {
	var user entity.User
	result := repo.db.Model(&entity.User{}).Where(
		"email = ?",
		email,
	).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepositoryPostgres) Create(user *entity.User) (
	*entity.User,
	error,
) {
	result := repo.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetByID(user.ID)
}

func (repo *UserRepositoryPostgres) Update(user *entity.User) (
	*entity.User,
	error,
) {
	result := repo.db.Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetByID(user.ID)
}
func (repo *UserRepositoryPostgres) Delete(id string) error {
	return errors.New("unimplemented")
}
