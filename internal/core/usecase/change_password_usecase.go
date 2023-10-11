package usecase

import (
	"github.com/kimoscloud/user-management-service/internal/core/auth"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	"github.com/kimoscloud/user-management-service/internal/core/model/request"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
)

type ChangePasswordUsecase struct {
	userRepository repository.UserRepository
	logger         logging.Logger
}

func NewChangePasswordUseCase(
	ur repository.UserRepository,
	logger logging.Logger,
) *ChangePasswordUsecase {
	return &ChangePasswordUsecase{userRepository: ur, logger: logger}
}

func (p *ChangePasswordUsecase) Handle(
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

func (p *ChangePasswordUsecase) checkIfUserExists(userId string) (*entity.User, *errors.AppError) {
	user, err := p.userRepository.GetByID(userId)
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

func (p *ChangePasswordUsecase) comparePasswordsStep(
	hash string,
	password string,
) *errors.AppError {
	err := comparePasswords(hash, password)
	if err != nil {
		return errors.NewUnauthorizedError(
			"Old password is not correct",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	return nil
}

func (p *ChangePasswordUsecase) hashAndSaltStep(password string) (string, *errors.AppError) {
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

func (p *ChangePasswordUsecase) updateUserStep(user *entity.User) *errors.AppError {
	_, err := p.userRepository.Update(user)
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
