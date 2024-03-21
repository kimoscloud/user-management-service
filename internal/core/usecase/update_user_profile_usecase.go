package usecase

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/request/auth"
	"github.com/kimoscloud/user-management-service/internal/core/model/response"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
)

type UpdateUserProfileUseCase struct {
	userRepository repository.Repository
	logger         logging.Logger
}

func NewUpdateUserProfileUseCase(
	ur repository.Repository,
	logger logging.Logger,
) *UpdateUserProfileUseCase {
	return &UpdateUserProfileUseCase{userRepository: ur, logger: logger}
}

func (p *UpdateUserProfileUseCase) Handler(
	id string,
	newInformation *auth.UpdateProfileRequest,
) (*response.UserLightDTO, *errors.AppError) {
	result, err := p.userRepository.GetByID(id)
	if err != nil {
		p.logger.Error("Error getting user by id", err)
		return nil, errors.NewInternalServerError(
			"Error getting user by id",
			"",
			errors.ErrorUserAuthenticatedNotFound,
		).AppError
	}
	if result == nil {
		p.logger.Error("User not found", err)
		return nil, errors.NewNotFoundError(
			"Error getting user by id",
			"",
			errors.ErrorUserAuthenticatedNotFound,
		).AppError
	}

	if resultGetByEmail, err := p.userRepository.GetByEmail(newInformation.Email); err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting user by email",
			"",
			errors.ErrorUserAuthenticatedNotFound,
		).AppError
	} else if resultGetByEmail != nil && resultGetByEmail.ID != result.ID {
		return nil, errors.NewBadRequestError(
			"Email already exists",
			"The email "+newInformation.Email+"exists in our database",
			errors.ErrorUserEmailAlreadyExists,
		).AppError
	}
	result.Email = newInformation.Email
	result.FirstName = newInformation.FirstName
	result.LastName = newInformation.LastName

	result, err = p.userRepository.Update(result)
	if err != nil {
		p.logger.Error("Error updating user", err)
		return nil, errors.NewInternalServerError(
			"Error updating user",
			"",
			errors.ErrorUserAuthenticatedNotFound,
		).AppError
	}
	return &response.UserLightDTO{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		LastLogin: result.LastLogin,
		CreatedAt: result.CreatedAt,
	}, nil
}
