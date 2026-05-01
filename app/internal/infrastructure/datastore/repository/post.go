package repository

import (
	"context"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

type PostRepository interface {
	GetByID(context.Context, int64) (*domain.Post, error)
	Create(context.Context, *domain.Post) error
	Delete(context.Context, int64) error
	Update(context.Context, *domain.Post) error
	GetUserFeed(context.Context, int64, domain.PaginatedFeedQuery) ([]domain.PostWithMetadata, error)
}
