package usecase

import (
	"github.com/kimoscloud/user-management-service/app/domain"
	"github.com/kimoscloud/user-management-service/app/repository"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(ur repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: ur}
}

func (p *CreateUserUseCase) Handler(user *domain.User) (*domain.User, error) {
	//TODO make the validations here
	return p.userRepository.Create(user)
}
