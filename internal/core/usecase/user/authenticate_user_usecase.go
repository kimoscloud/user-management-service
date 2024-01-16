package user

import (
	"github.com/kimoscloud/user-management-service/internal/core/auth"
	auth2 "github.com/kimoscloud/user-management-service/internal/core/model/request/auth"
	"github.com/kimoscloud/user-management-service/internal/core/model/response"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository/user"
	"github.com/kimoscloud/user-management-service/internal/core/utils"
	"github.com/kimoscloud/value-types/errors"
	"time"
)

type AuthenticateUserUseCase struct {
	userRepository user.Repository
	logger         logging.Logger
}

func NewAuthenticateUserUseCase(
	ur user.Repository,
	logger logging.Logger,
) *AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{userRepository: ur, logger: logger}
}

func (p *AuthenticateUserUseCase) Handler(request auth2.LoginRequest) (
	*response.BaseAuthenticationResponse,
	*errors.AppError,
) {
	result, err := p.userRepository.GetByEmail(request.Email)
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting user by email",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	if result == nil {
		p.logger.Error("User doesn't exist", "email", request.Email)
		return nil, errors.NewUnauthorizedError(
			"Email or password not exists",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	if result.IsLocked || result.BadLoginAttempts >= 5 {
		p.logger.Error("User is locked", "email", request.Email)
		return nil, errors.NewUnauthorizedError(
			"Email or password not exists",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	if !utils.ComparePasswords(result.Hash, request.Password) {
		return nil, errors.NewUnauthorizedError(
			"Email or password not exists",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	expirationTime := time.Now().Add(60 * time.Minute)
	jwt, err := auth.GenerateJWT(result.ID, result.Email, expirationTime)
	if err != nil {
		p.logger.Error("Error generating JWT", "errors", err.Error())
		return nil, errors.NewInternalServerError(
			"Error generating JWT",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	return &response.BaseAuthenticationResponse{
		AccessToken: jwt,
		ExpiresIn:   int(expirationTime.Unix()),
		TokenType:   "Bearer",
	}, nil

}
