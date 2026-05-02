package impl

import (
	"context"
	"strings"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/domain"
	"github.com/YagoSchramm/gopher-social/internal/infrastructure/datastore/repository"
	"github.com/YagoSchramm/gopher-social/internal/usecase"
)

func NewCommentUseCase(commentRepo repository.CommentRepository) usecase.CommentUseCase {
	return &commentUseCase{
		commentRepo: commentRepo,
	}
}

type commentUseCase struct {
	commentRepo repository.CommentRepository
}

func (c *commentUseCase) Create(ctx context.Context, comment *domain.Comment) error {
	if comment == nil {
		return derr.NewBadRequestError("comment is required")
	}

	comment.Content = strings.TrimSpace(comment.Content)
	if comment.Content == "" {
		return derr.NewBadRequestError("content is required")
	}
	if comment.PostID <= 0 {
		return derr.NewBadRequestError("post id is required")
	}
	if comment.UserID <= 0 {
		return derr.NewBadRequestError("user id is required")
	}

	if err := c.commentRepo.Create(ctx, comment); err != nil {
		return derr.JoinError("failed to create comment", err)
	}

	return nil
}

func (c *commentUseCase) GetByPostID(ctx context.Context, postID int64) ([]domain.Comment, error) {
	if postID <= 0 {
		return nil, derr.NewBadRequestError("post id is required")
	}

	comments, err := c.commentRepo.GetByPostID(ctx, postID)
	if err != nil {
		return nil, derr.JoinError("failed to get comments by post id", err)
	}

	return comments, nil
}
