package user

import (
	"errors"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/auth"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository/user"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &RepositoryPostgres{db: db}
}

func (repo *RepositoryPostgres) FindUsersByEmails(emails []string) ([]auth.User, error) {
	var users []auth.User
	if err := repo.db.Where("email in ?", emails).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (repo *RepositoryPostgres) GetAll() ([]auth.User, error) {
	var users []auth.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *RepositoryPostgres) GetUserByEmailLike(search string, limit int) ([]auth.User, error) {
	var users []auth.User
	query := repo.db.Where("email ILIKE ? AND is_deleted ", search)
	if limit <= 10 {
		query.Limit(limit)
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *RepositoryPostgres) GetPage(
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

func (repo *RepositoryPostgres) GetByID(id string) (*auth.User, error) {
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

func (repo *RepositoryPostgres) GetByEmail(email string) (
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
func (repo *RepositoryPostgres) Create(user *auth.User) (
	*auth.User,
	error,
) {
	result := repo.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetByID(user.ID)
}

func (repo *RepositoryPostgres) Update(user *auth.User) (
	*auth.User,
	error,
) {
	result := repo.db.Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetByID(user.ID)
}
func (repo *RepositoryPostgres) Delete(id string) error {
	return errors.New("unimplemented")
}
