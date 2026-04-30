package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Datastore interface {
	Posts() PostRepository
	Users() UserRepository
	Comments() CommentRepository
	Followers() FollowerRepository
	Roles() RoleRepository
}

type PostRepository interface {
	GetByID(context.Context, int64) (*domain.Post, error)
	Create(context.Context, *domain.Post) error
	Delete(context.Context, int64) error
	Update(context.Context, *domain.Post) error
	GetUserFeed(context.Context, int64, domain.PaginatedFeedQuery) ([]domain.PostWithMetadata, error)
}

type UserRepository interface {
	GetByID(context.Context, int64) (*domain.User, error)
	GetByEmail(context.Context, string) (*domain.User, error)
	Create(context.Context, *sql.Tx, *domain.User) error
	CreateAndInvite(ctx context.Context, user *domain.User, token string, exp time.Duration) error
	Activate(context.Context, string) error
	Delete(context.Context, int64) error
}

type CommentRepository interface {
	Create(context.Context, *domain.Comment) error
	GetByPostID(context.Context, int64) ([]domain.Comment, error)
}

type FollowerRepository interface {
	Follow(ctx context.Context, userID, followerID int64) error
	Unfollow(ctx context.Context, followerID, userID int64) error
}

type RoleRepository interface {
	GetByName(context.Context, string) (*domain.Role, error)
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
