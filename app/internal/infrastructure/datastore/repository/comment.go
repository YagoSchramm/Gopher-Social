package repository

import (
	"context"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

type CommentRepository interface {
	Create(context.Context, *domain.Comment) error
	GetByPostID(context.Context, int64) ([]domain.Comment, error)
}
