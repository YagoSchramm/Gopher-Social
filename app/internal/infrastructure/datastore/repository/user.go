package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

type UserRepository interface {
	GetByID(context.Context, int64) (*domain.User, error)
	GetByEmail(context.Context, string) (*domain.User, error)
	Create(context.Context, *sql.Tx, *domain.User) error
	CreateAndInvite(ctx context.Context, user *domain.User, token string, exp time.Duration) error
	Activate(context.Context, string) error
	Delete(context.Context, int64) error
}
