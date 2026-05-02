package impl

import (
	"context"
	"strings"
	"time"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/domain"
	"github.com/YagoSchramm/gopher-social/internal/infrastructure/datastore/repository"
	"github.com/YagoSchramm/gopher-social/internal/usecase"
)

func NewUserUseCase(userRepo repository.UserRepository) usecase.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) GetByID(ctx context.Context, userID int64) (*domain.User, error) {
	if userID <= 0 {
		return nil, derr.NewBadRequestError("user id is required")
	}

	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, derr.JoinError("failed to get user by id", err)
	}

	return user, nil
}

func (u *userUseCase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return nil, derr.NewBadRequestError("email is required")
	}

	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, derr.JoinError("failed to get user by email", err)
	}

	return user, nil
}

func (u *userUseCase) CreateAndInvite(
	ctx context.Context,
	user *domain.User,
	token string,
	invitationExp time.Duration,
) error {
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
	}

	if err := u.userRepo.CreateAndInvite(ctx, user, token, invitationExp); err != nil {
		return derr.JoinError("failed to create and invite user", err)
	}

	return nil
}

func (u *userUseCase) Activate(ctx context.Context, token string) error {
	token = strings.TrimSpace(token)
	if token == "" {
		return derr.NewBadRequestError("invitation token is required")
	}

	if err := u.userRepo.Activate(ctx, token); err != nil {
		return derr.JoinError("failed to activate user", err)
	}

	return nil
}

func (u *userUseCase) Delete(ctx context.Context, userID int64) error {
	if userID <= 0 {
		return derr.NewBadRequestError("user id is required")
	}

	if err := u.userRepo.Delete(ctx, userID); err != nil {
		return derr.JoinError("failed to delete user", err)
	}

	return nil
}
