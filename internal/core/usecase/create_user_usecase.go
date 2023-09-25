package usecase

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(ur repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: ur}
}

func (p *CreateUserUseCase) Handler(user *entity.User) (*entity.User, error) {
	//TODO make the validations here
	return p.userRepository.Create(user)
}
