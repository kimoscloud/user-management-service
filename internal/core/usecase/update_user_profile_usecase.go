package usecase

import (
	errors2 "github.com/kimoscloud/user-management-service/internal/core/errors"
	"github.com/kimoscloud/user-management-service/internal/core/model/request"
	"github.com/kimoscloud/user-management-service/internal/core/model/response"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
)

type UpdateUserProfileUseCase struct {
	userRepository repository.UserRepository
	logger         logging.Logger
}

func NewUpdateUserProfileUseCase(
	ur repository.UserRepository,
	logger logging.Logger,
) *UpdateUserProfileUseCase {
	return &UpdateUserProfileUseCase{userRepository: ur, logger: logger}
}

func (p *UpdateUserProfileUseCase) Handler(
	id string,
	newInformation *request.UpdateProfileRequest,
) (*response.UserLightDTO, *errors.AppError) {
	result, err := p.userRepository.GetByID(id)
	if err != nil {
		p.logger.Error("Error getting user by id", err)
		return nil, errors.NewInternalServerError(
			"Error getting user by id",
			"",
			errors2.ErrorUserAuthenticatedNotFound,
		).AppError
	}
	if result == nil {
		p.logger.Error("User not found", err)
		return nil, errors.NewNotFoundError(
			"Error getting user by id",
			"",
			errors2.ErrorUserAuthenticatedNotFound,
		).AppError
	}

	if resultGetByEmail, err := p.userRepository.GetByEmail(newInformation.Email); err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting user by email",
			"",
			errors2.ErrorUserAuthenticatedNotFound,
		).AppError
	} else if resultGetByEmail != nil && resultGetByEmail.ID != result.ID {
		return nil, errors.NewBadRequestError(
			"Email already exists",
			"The email "+newInformation.Email+"exists in our database",
			errors2.ErrorUserEmailAlreadyExists,
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
			errors2.ErrorUserAuthenticatedNotFound,
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
