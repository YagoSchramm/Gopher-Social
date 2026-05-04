package usecase

import (
	"context"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

type PostUseCase interface {
	GetByID(context.Context, int64) (*domain.Post, error)
	Create(context.Context, *domain.Post) error
	Update(context.Context, *domain.Post) error
	Delete(context.Context, int64) error
	GetUserFeed(context.Context, int64, domain.PaginatedFeedQuery) ([]domain.PostWithMetadata, error)
}
