package usecase

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	"github.com/kimoscloud/user-management-service/internal/core/model/request"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
	"github.com/kimoscloud/value-types/is_valid"
	"time"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepository
	logger         *logging.Logger
}

func NewCreateUserUseCase(ur repository.UserRepository, logger *logging.Logger) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: ur, logger: logger}
}

func (p *CreateUserUseCase) Handler(req *request.SignUpRequest) (*entity.User, *errors.AppError) {
	appError := validateSignUpRequest(req)
	if appError != nil {
		return nil, appError
	}
	user, err := p.userRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.NewInternalServerError("Error getting user by email", "", errors.ErrorGettingUser).AppError
	}
	user = &entity.User{
		Email:                      req.Email,
		Hash:                       req.Password,
		BadLoginAttempts:           0,
		IsLocked:                   false,
		AcceptTermsAndConditions:   req.AcceptTermsAndConditions,
		AcceptTermsAndConditionsAt: time.Now(),
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}
	createUserResult, err := p.userRepository.Create(user)
	if err != nil {
		p.logger.Error("Error creating user", "error", err.Error())
		return nil, errors.NewInternalServerError("Error creating user", "", errors.ErrorCreatingUser).AppError
	}
	return createUserResult, nil
}

func validateSignUpRequest(signUpRequest *request.SignUpRequest) *errors.AppError {
	if !signUpRequest.AcceptTermsAndConditions {
		return errors.NewBadRequestError("User must accept terms and conditions", "", errors.ErrorUserNotAcceptTermsAndConditions).AppError
	}
	if !is_valid.IsValidEmail(signUpRequest.Email) {
		return errors.NewBadRequestError("Invalid email", "", errors.ErrorInvalidEmail).AppError
	}
	if !is_valid.IsValidPassword(signUpRequest.Password) {
		return errors.NewBadRequestError("Invalid password", "", errors.ErrorPasswordDoesntHaveTheRequestedFormat).AppError
	}
	if signUpRequest.Password != signUpRequest.ConfirmPassword {
		return errors.NewBadRequestError("Password and confirm password don't match", "", errors.ErrorPasswordDoesntMatch).AppError
	}
	return nil
}
