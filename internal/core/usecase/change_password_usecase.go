package usecase

import (
	"github.com/kimoscloud/user-management-service/internal/core/auth"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	"github.com/kimoscloud/user-management-service/internal/core/model/request"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	userRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/user-management-service/internal/core/utils"
	"github.com/kimoscloud/value-types/errors"
)

type ChangePasswordUseCase struct {
	userRepo userRepository.Repository
	logger   logging.Logger
}

func NewChangePasswordUseCase(
	ur userRepository.Repository,
	logger logging.Logger,
) *ChangePasswordUseCase {
	return &ChangePasswordUseCase{userRepo: ur, logger: logger}
}

func (p *ChangePasswordUseCase) Handle(
	userId string, request *request.ChangePasswordRequest,
) *errors.AppError {
	user, appError := p.checkIfUserExists(userId)
	if appError != nil {
		return appError
	}
	appError = p.comparePasswordsStep(user.Hash, request.OldPassword)
	if appError != nil {
		return appError
	}
	hashedPassword, appError := p.hashAndSaltStep(request.NewPassword)
	if appError != nil {
		return appError
	}
	user.Hash = hashedPassword
	user.BadLoginAttempts = 0
	user.IsLocked = false
	appError = p.updateUserStep(user)
	return nil
}

func (p *ChangePasswordUseCase) checkIfUserExists(userId string) (*entity.User, *errors.AppError) {
	user, err := p.userRepo.GetByID(userId)
	if err != nil {
		p.logger.Error("Error getting user by id", err)
		return nil, errors.NewInternalServerError(
			"Error getting user by email",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	if user == nil {
		p.logger.Error("User not found", err)
		return nil, errors.NewNotFoundError(
			"Error getting user by id",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	return user, nil

}

func (p *ChangePasswordUseCase) comparePasswordsStep(
	hash string,
	password string,
) *errors.AppError {
	if !utils.ComparePasswords(hash, password) {
		return errors.NewUnauthorizedError(
			"Old password is not correct",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	return nil
}

func (p *ChangePasswordUseCase) hashAndSaltStep(password string) (string, *errors.AppError) {
	hashedPassword, err := auth.HashAndSalt(password)
	if err != nil {
		p.logger.Error("Error hashing password", "errors", err.Error())
		return "", errors.NewInternalServerError(
			"Error updating the password",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	return hashedPassword, nil
}

func (p *ChangePasswordUseCase) updateUserStep(user *entity.User) *errors.AppError {
	_, err := p.userRepo.Update(user)
	if err != nil {
		p.logger.Error("Error updating user", err)
		return errors.NewInternalServerError(
			"Error updating the password",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	return nil
}
