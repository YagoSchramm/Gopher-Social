package derr

import "errors"

type RepositoryError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e RepositoryError) Error() string {
	return e.Message
}

type ClientError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e ClientError) Error() string {
	return e.Message
}

func NewRepositoryError(code string, message string) RepositoryError {
	return RepositoryError{
		Code:    code,
		Message: message,
	}
}

func NewClientError(code string, message string) ClientError {
	return ClientError{
		Code:    code,
		Message: message,
	}
}

func JoinError(message string, err error) error {
	return errors.Join(errors.New(message), err)
}

var (
	NotFound         = NewRepositoryError("NOT_FOUND", "resource not found")
	Conflict         = NewClientError("CONFLICT", "resource already exists")
	CommentNotFound  = NotFound
	FollowerNotFound = NotFound
	FollowerConflict = Conflict
	PostNotFound     = NotFound
	RoleNotFound     = NotFound
	InvalidUserName  = NewClientError("INVALID_USER_NAME", "a user with that username already exists")
	InvalidUserEmail = NewClientError("INVALID_USER_EMAIL", "a user with that email already exists")
)
