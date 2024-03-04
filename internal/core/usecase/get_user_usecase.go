package usecase

import (
	errors2 "github.com/kimoscloud/user-management-service/internal/core/errors"
	"github.com/kimoscloud/user-management-service/internal/core/model/response"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
)

type GetUserUseCase struct {
	userRepository repository.Repository
	logger         logging.Logger
}

func NewGetUserUseCase(
	ur repository.Repository,
	logger logging.Logger,
) *GetUserUseCase {
	return &GetUserUseCase{userRepository: ur, logger: logger}
}

func (p *GetUserUseCase) Handler(id string) (
	*response.UserLightDTO,
	*errors.AppError,
) {
	result, err := p.userRepository.GetByID(id)
	if err != nil {
		return nil, errors.NewNotFoundError(
			"Error getting user by id",
			"",
			errors2.ErrorUserAuthenticatedNotFound,
		).AppError
	}
	if result == nil {
		return nil, errors.NewNotFoundError(
			"Error getting user by id",
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
