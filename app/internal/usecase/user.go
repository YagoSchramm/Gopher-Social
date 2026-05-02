package usecase

import (
	"context"
	"time"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

type UserUseCase interface {
	GetByID(context.Context, int64) (*domain.User, error)
	GetByEmail(context.Context, string) (*domain.User, error)
	CreateAndInvite(context.Context, *domain.User, string, time.Duration) error
	Activate(context.Context, string) error
	Delete(context.Context, int64) error
}
