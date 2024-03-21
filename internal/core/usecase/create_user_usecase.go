package usecase

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	auth2 "github.com/kimoscloud/user-management-service/internal/core/model/request/auth"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/user-management-service/internal/core/utils"
	"github.com/kimoscloud/value-types/errors"
	"github.com/kimoscloud/value-types/is_valid"
	"time"
)

type CreateUserUseCase struct {
	userRepository repository.Repository
	logger         logging.Logger
}

func NewCreateUserUseCase(
	ur repository.Repository,
	logger logging.Logger,
) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: ur, logger: logger}
}

func (p *CreateUserUseCase) Handler(req *auth2.SignUpRequest) (
	*entity.User,
	*errors.AppError,
) {
	appError := validateSignUpRequest(req)
	if appError != nil {
		return nil, appError
	}
	user, err := p.userRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting user by email",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	if user != nil {
		p.logger.Error("User already exists", "email", req.Email)
		return nil, errors.NewBadRequestError(
			"User already exists",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}

	hashedPassword, err := utils.GeneratePassword(req.Password)
	if err != nil {
		p.logger.Error("Error hashing password", "errors", err.Error())
		return nil, errors.NewInternalServerError(
			"Error creating the user",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	user = &entity.User{
		Email:                      req.Email,
		Hash:                       hashedPassword,
		BadLoginAttempts:           0,
		IsLocked:                   false,
		AcceptTermsAndConditions:   req.AcceptTermsAndConditions,
		AcceptTermsAndConditionsAt: time.Now(),
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}
	createUserResult, err := p.userRepository.Create(user)
	if err != nil {
		p.logger.Error("Error creating user", "errors", err.Error())
		return nil, errors.NewInternalServerError(
			"Error creating user",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	return createUserResult, nil
}

func validateSignUpRequest(signUpRequest *auth2.SignUpRequest) *errors.AppError {
	if !signUpRequest.AcceptTermsAndConditions {
		return errors.NewBadRequestError(
			"User must accept terms and conditions",
			"",
			errors.ErrorUserNotAcceptTermsAndConditions,
		).AppError
	}
	if !is_valid.IsValidEmail(signUpRequest.Email) {
		return errors.NewBadRequestError(
			"Invalid email",
			"",
			errors.ErrorInvalidEmail,
		).AppError
	}
	if !is_valid.IsValidPassword(signUpRequest.Password) {
		return errors.NewBadRequestError(
			"Invalid password",
			"",
			errors.ErrorPasswordDoesntHaveTheRequestedFormat,
		).AppError
	}
	if signUpRequest.Password != signUpRequest.ConfirmPassword {
		return errors.NewBadRequestError(
			"Password and confirm password don't match",
			"",
			errors.ErrorPasswordDoesntMatch,
		).AppError
	}
	return nil
}
