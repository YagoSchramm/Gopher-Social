package impl

import (
	"context"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/domain"
	domainrules "github.com/YagoSchramm/gopher-social/internal/domain/rules"
	"github.com/YagoSchramm/gopher-social/internal/infrastructure/datastore/repository"
	"github.com/YagoSchramm/gopher-social/internal/usecase"
)

func NewPostUseCase(postRepo repository.PostRepository) usecase.PostUseCase {
	return &postUseCase{
		postRepo: postRepo,
	}
}

type postUseCase struct {
	postRepo repository.PostRepository
}

func (p *postUseCase) GetByID(ctx context.Context, postID int64) (*domain.Post, error) {
	if err := domainrules.ValidatePostID(postID); err != nil {
		return nil, err
	}

	post, err := p.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, derr.JoinError("failed to get post by id", err)
	}

	return post, nil
}

func (p *postUseCase) Create(ctx context.Context, post *domain.Post) error {
	if err := domainrules.ValidatePostForCreate(post); err != nil {
		return err
	}

	if err := p.postRepo.Create(ctx, post); err != nil {
		return derr.JoinError("failed to create post", err)
	}

	return nil
}

func (p *postUseCase) Update(ctx context.Context, post *domain.Post) error {
	if err := domainrules.ValidatePostForUpdate(post); err != nil {
		return err
	}

	if err := p.postRepo.Update(ctx, post); err != nil {
		return derr.JoinError("failed to update post", err)
	}

	return nil
}

func (p *postUseCase) Delete(ctx context.Context, postID int64) error {
	if err := domainrules.ValidatePostID(postID); err != nil {
		return err
	}

	if err := p.postRepo.Delete(ctx, postID); err != nil {
		return derr.JoinError("failed to delete post", err)
	}

	return nil
}

func (p *postUseCase) GetUserFeed(ctx context.Context, userID int64, fq domain.PaginatedFeedQuery) ([]domain.PostWithMetadata, error) {
	if err := domainrules.ValidateUserID(userID); err != nil {
		return nil, err
	}
	if err := domainrules.ValidatePaginatedFeedQuery(&fq); err != nil {
		return nil, err
	}

	feed, err := p.postRepo.GetUserFeed(ctx, userID, fq)
	if err != nil {
		return nil, derr.JoinError("failed to get user feed", err)
	}

	return feed, nil
}
