package usecase

import (
	"context"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

type CommentUseCase interface {
	Create(context.Context, *domain.Comment) error
	GetByPostID(context.Context, int64) ([]domain.Comment, error)
}
