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

type Storage struct {
	Posts interface {
		GetByID(context.Context, int64) (*domain.Post, error)
		Create(context.Context, *domain.Post) error
		Delete(context.Context, int64) error
		Update(context.Context, *domain.Post) error
		GetUserFeed(context.Context, int64, domain.PaginatedFeedQuery) ([]domain.PostWithMetadata, error)
	}
	Users interface {
		GetByID(context.Context, int64) (*domain.User, error)
		GetByEmail(context.Context, string) (*domain.User, error)
		Create(context.Context, *sql.Tx, *domain.User) error
		CreateAndInvite(ctx context.Context, user *domain.User, token string, exp time.Duration) error
		Activate(context.Context, string) error
		Delete(context.Context, int64) error
	}
	Comments interface {
		Create(context.Context, *domain.Comment) error
		GetByPostID(context.Context, int64) ([]domain.Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, userID, followerID int64) error
		Unfollow(ctx context.Context, followerID, userID int64) error
	}
	Roles interface {
		GetByName(context.Context, string) (*domain.Role, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
		Roles:     &RoleStore{db},
	}
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
