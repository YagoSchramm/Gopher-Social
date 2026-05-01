package impl

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/infrastructure/datastore/repository"
	"github.com/YagoSchramm/gopher-social/internal/infrastructure/datastore/util"
	"github.com/lib/pq"
)

//go:embed _query/followers/follow.sql
var followerFollowQuery string

//go:embed _query/followers/unfollow.sql
var followerUnfollowQuery string

func NewFollowerRepository(db *sql.DB) repository.FollowerRepository {
	return &FollowerStore{db: db}
}

type FollowerStore struct {
	db *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, userID, followerID int64) error {
	ctx, cancel := context.WithTimeout(ctx, util.QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, followerFollowQuery, userID, followerID)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return derr.FollowerConflict
		}
		return err
	}

	return nil
}

func (s *FollowerStore) Unfollow(ctx context.Context, followerID, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, util.QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, followerUnfollowQuery, followerID, userID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return derr.FollowerNotFound
	}

	return nil
}
