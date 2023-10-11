package usecase

import (
	"github.com/kimoscloud/user-management-service/internal/core/auth"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	"github.com/kimoscloud/user-management-service/internal/core/model/request"
	"github.com/kimoscloud/user-management-service/internal/core/model/response"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthenticateUserUseCase struct {
	userRepository repository.UserRepository
	logger         logging.Logger
}

func NewAuthenticateUserUseCase(
	ur repository.UserRepository,
	logger logging.Logger,
) *AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{userRepository: ur, logger: logger}
}

func (p *AuthenticateUserUseCase) Handler(request request.LoginRequest) (
	*response.BaseAuthenticationResponse,
	*errors.AppError,
) {
	result, appError := p.checkIfUserExists(request)
	if appError != nil {
		return nil, appError
	}
	appError = checkIfUserIsLocked(result, p.logger)
	if appError != nil {
		return nil, appError
	}
	appError = comparePasswords(result.Hash, request.Password)
	if appError != nil {
		return nil, appError
	}
	jwt, expirationTime, appError := generateJWT(result, p.logger)
	if appError != nil {
		return nil, appError
	}
	return &response.BaseAuthenticationResponse{
		AccessToken: jwt,
		ExpiresIn:   int(expirationTime.Unix()),
		TokenType:   "Bearer",
	}, nil

}

func (p *AuthenticateUserUseCase) checkIfUserExists(request request.LoginRequest) (
	*entity.User,
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
	return result, nil
}

func generateJWT(result *entity.User, logger logging.Logger) (
	string, time.Time,
	*errors.AppError,
) {
	expirationTime := time.Now().Add(60 * time.Minute)
	jwt, err := auth.GenerateJWT(result.ID, result.Email, expirationTime)
	if err != nil {
		logger.Error("Error generating JWT", "errors", err.Error())
		return "", expirationTime, errors.NewInternalServerError(
			"Error generating JWT",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	return jwt, expirationTime, nil
}

func checkIfUserIsLocked(result *entity.User, logger logging.Logger) *errors.AppError {
	if result.IsLocked || result.BadLoginAttempts >= 5 {
		logger.Error("User is locked", "email", result.Email)
		return errors.NewUnauthorizedError(
			"Email or password not exists",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	return nil
}

func comparePasswords(hash string, password string) *errors.AppError {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.NewUnauthorizedError(
			"Email or password not exists",
			"",
			errors.ErrorAuthenticatingUser,
		).AppError
	}
	return nil
}
