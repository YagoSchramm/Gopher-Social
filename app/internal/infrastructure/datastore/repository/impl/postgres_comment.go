package impl

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/gopher-social/internal/domain"
	"github.com/YagoSchramm/gopher-social/internal/infrastructure/datastore/repository"
	"github.com/YagoSchramm/gopher-social/internal/infrastructure/datastore/util"
)

//go:embed _query/comments/create.sql
var commentCreateQuery string

//go:embed _query/comments/get_by_post_id.sql
var commentGetByPostIDQuery string

func NewCommentRepository(db *sql.DB) repository.CommentRepository {
	return &CommentStore{db: db}
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, util.QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		commentCreateQuery,
		comment.Content,
		comment.PostID,
		comment.UserID,
	).Scan(
		&comment.ID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, util.QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, commentGetByPostIDQuery, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]domain.Comment, 0)
	for rows.Next() {
		var comment domain.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.PostID,
			&comment.UserID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.User.ID,
			&comment.User.Username,
		); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
