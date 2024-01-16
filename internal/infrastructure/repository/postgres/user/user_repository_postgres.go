package user

import (
	"errors"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/auth"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
)

type UserRepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

func (repo *UserRepositoryPostgres) GetAll() ([]auth.User, error) {
	var users []auth.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepositoryPostgres) GetPage(
	pageNumber int,
	pageSize int,
) (types.Page[auth.User], error) {
	var totalRows int64
	repo.db.Model(&auth.User{}).Count(&totalRows)
	var users []auth.User
	if err := repo.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return types.EmptyPage[auth.User](), err
	}
	pageBuilder := new(types.PageBuilder[auth.User])
	return pageBuilder.SetItems(users).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		Build(), nil
}

func (repo *UserRepositoryPostgres) GetByID(id string) (*auth.User, error) {
	var user auth.User
	result := repo.db.Model(&auth.User{}).Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepositoryPostgres) GetByEmail(email string) (
	*auth.User,
	error,
) {
	var user auth.User
	result := repo.db.Model(&auth.User{}).Where(
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

// TODO add context here
func (repo *UserRepositoryPostgres) Create(user *auth.User) (
	*auth.User,
	error,
) {
	result := repo.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetByID(user.ID)
}

func (repo *UserRepositoryPostgres) Update(user *auth.User) (
	*auth.User,
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
