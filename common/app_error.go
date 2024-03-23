package common

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr   error  `json:"_"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewFullErrorResponse(statusCode int, err error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    err,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(err error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    err,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorizedResponse(err error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    err,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("You have no permission: %s", err.Error()), "NO_PERMISSION_ERROR")
}

func ErrDB(err error) error {
	return errors.New("something went wrong with Database");
}
