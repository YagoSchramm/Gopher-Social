package rules

import (
	"strings"
	"time"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/domain"
)

func ValidateUserID(userID int64) error {
	if userID <= 0 {
		return derr.NewBadRequestError("user id is required")
	}

	return nil
}

func ValidateUserEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return derr.NewBadRequestError("email is required")
	}

	return nil
}

func ValidateUserPassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return derr.NewBadRequestError("password is required")
	}

	return nil
}

func ValidateUserForCreate(user *domain.User, token string, invitationExp time.Duration) error {
	if user == nil {
		return derr.NewBadRequestError("user is required")
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
	token = strings.TrimSpace(token)

	switch {
	case user.Username == "":
		return derr.NewBadRequestError("username is required")
	case user.Email == "":
		return derr.NewBadRequestError("email is required")
	case user.Password == "":
		return derr.NewBadRequestError("password is required")
	case token == "":
		return derr.NewBadRequestError("invitation token is required")
	case invitationExp <= 0:
		return derr.NewBadRequestError("invitation expiration must be greater than zero")
	default:
		return nil
	}
}

func ValidateInvitationToken(token string) error {
	if strings.TrimSpace(token) == "" {
		return derr.NewBadRequestError("invitation token is required")
	}

	return nil
}
