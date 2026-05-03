package impl

import (
	"context"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/domain"
	domainrules "github.com/YagoSchramm/gopher-social/internal/domain/rules"
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
	if err := domainrules.ValidateCommentForCreate(comment); err != nil {
		return err
	}

	if err := c.commentRepo.Create(ctx, comment); err != nil {
		return derr.JoinError("failed to create comment", err)
	}

	return nil
}

func (c *commentUseCase) GetByPostID(ctx context.Context, postID int64) ([]domain.Comment, error) {
	if err := domainrules.ValidatePostID(postID); err != nil {
		return nil, err
	}

	comments, err := c.commentRepo.GetByPostID(ctx, postID)
	if err != nil {
		return nil, derr.JoinError("failed to get comments by post id", err)
	}

	return comments, nil
}
