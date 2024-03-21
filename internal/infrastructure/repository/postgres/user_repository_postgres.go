package postgres

import (
	"errors"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	types "github.com/kimoscloud/value-types/domain"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.Repository {
	return &RepositoryPostgres{db: db}
}

func (repo *RepositoryPostgres) FindUsersByEmails(emails []string) ([]entity.User, error) {
	var users []entity.User
	if err := repo.db.Where("email in ?", emails).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *RepositoryPostgres) LockUser(id string) error {
	return repo.db.Model(&entity.User{}).Where("id = ?", id).Update("is_locked", true).Error
}

func (repo *RepositoryPostgres) IncrementBadLoginAttempts(id string) error {
	return repo.db.Model(&entity.User{}).Where("id = ?", id).Update(
		"bad_attempts",
		gorm.Expr("bad_attempts + ?", 1),
	).Error
}

func (repo *RepositoryPostgres) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *RepositoryPostgres) GetUserByEmailLike(search string, limit int) ([]entity.User, error) {
	var users []entity.User
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
) (types.Page[entity.User], error) {
	var totalRows int64
	repo.db.Model(&entity.User{}).Count(&totalRows)
	var users []entity.User
	if err := repo.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return types.EmptyPage[entity.User](), err
	}
	pageBuilder := new(types.PageBuilder[entity.User])
	return pageBuilder.SetItems(users).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		Build(), nil
}

func (repo *RepositoryPostgres) GetByID(id string) (*entity.User, error) {
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

func (repo *RepositoryPostgres) GetByEmail(email string) (
	*entity.User,
	error,
) {
	var userResult entity.User
	result := repo.db.Model(&entity.User{}).Where(
		"email = ?",
		email,
	).First(&userResult)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &userResult, nil
}

// TODO add context here
func (repo *RepositoryPostgres) Create(user *entity.User) (
	*entity.User,
	error,
) {
	result := repo.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetByID(user.ID)
}

func (repo *RepositoryPostgres) Update(user *entity.User) (
	*entity.User,
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
