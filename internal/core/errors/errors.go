package errors

import (
	"github.com/kimoscloud/value-types/errors"
	"net/http"
)

type ForbiddenError struct {
	*errors.AppError
}

type ConflictError struct {
	*errors.AppError
}

func NewForbiddenError(message, description string, code errors.ErrorCode) *ForbiddenError {
	return &ForbiddenError{
		AppError: &errors.AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusForbidden,
		},
	}
}

func NewConflictError(message, description string, code errors.ErrorCode) *ConflictError {
	return &ConflictError{
		AppError: &errors.AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusConflict,
		},
	}
}
