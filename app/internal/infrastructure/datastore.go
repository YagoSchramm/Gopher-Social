package infrastructure

import "context"

type PostsRepository interface {
	GetAll(ctx context.Context, filter string) ([]*model.Post, error)
}
