package errors

import (
	"github.com/kimoscloud/value-types/errors"
	"net/http"
)

type ForbiddenError struct {
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
