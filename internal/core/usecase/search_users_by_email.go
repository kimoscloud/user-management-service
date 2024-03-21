package usecase

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	userRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
)

type SearchUserByEmailUseCase struct {
	userRepo userRepository.Repository
	logger   logging.Logger
}

func NewSearchUserByEmailUseCase(userRepo userRepository.Repository, logger logging.Logger) *SearchUserByEmailUseCase {
	return &SearchUserByEmailUseCase{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (uc SearchUserByEmailUseCase) Handler(search string) ([]entity.User, *errors.AppError) {
	users, err := uc.userRepo.GetUserByEmailLike(search, 4)
	if err != nil {
		//TODO replace error code
		return nil, errors.NewInternalServerError("Error trying to get the users", "Error trying to get the users by email like", "000010").AppError
	}
	return users, nil
}
