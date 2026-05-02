package derr

import "net/http"

type HTTPError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewError(statusCode int, code string, message string) HTTPError {
	return HTTPError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

func NewBadRequestError(message string) HTTPError {
	return NewError(http.StatusBadRequest, "BAD_REQUEST", message)
}

func NewUnauthorizedError(message string) HTTPError {
	return NewError(http.StatusUnauthorized, "UNAUTHORIZED", message)
}

func NewNotFoundError(message string) HTTPError {
	return NewError(http.StatusNotFound, "NOT_FOUND", message)
}

func NewConflictError(message string) HTTPError {
	return NewError(http.StatusConflict, "CONFLICT", message)
}

func NewInternalError(message string) HTTPError {
	return NewError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", message)
}

var (
	BadRequestError     = NewBadRequestError("Bad Request")
	UnauthorizedError   = NewUnauthorizedError("Unauthorized")
	NotFoundError       = NewNotFoundError("Not Found")
	ConflictError       = NewConflictError("Conflict")
	InternalServerError = NewInternalError("Internal Server Error")
)
