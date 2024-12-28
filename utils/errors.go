package utils

import "net/http"

type AppError struct {
	HttpStatus int         `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Details    interface{} `json:"details,omitempty"`
}

func (e AppError) Error() string {
	return e.Message
}

func NewAppError(httpStatus int, code, message string, details interface{}) error {
	return AppError{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    message,
		Details:    details,
	}
}

func NewValidationError(details interface{}) error {
	return NewAppError(http.StatusBadRequest, "VALIDATION_ERROR", "Validation failed", details)
}

func NewBadRequestError(message string) error {
	return NewAppError(http.StatusBadRequest, "BAD_REQUEST", message, nil)
}

func NewNotFoundError(message string) error {
	return NewAppError(http.StatusNotFound, "NOT_FOUND", message, nil)
}

func NewConflictError(message string) error {
	return NewAppError(http.StatusConflict, "CONFLICT", message, nil)
}

func NewInternalError(message string) error {
	return NewAppError(http.StatusInternalServerError, "INTERNAL_ERROR", message, nil)
}

func NewUnauthorizedError(message string) error {
	return NewAppError(http.StatusUnauthorized, "UNAUTHORIZED", message, nil)
}

func NewForbiddenError(message string) error {
	return NewAppError(http.StatusForbidden, "FORBIDDEN", message, nil)
}

func GetStatusCode(err error) int {
	if appErr, ok := err.(AppError); ok {
		return appErr.HttpStatus
	}
	return http.StatusInternalServerError
}
