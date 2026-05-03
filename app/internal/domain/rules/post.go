package rules

import (
	"strings"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/domain"
)

func ValidatePostForCreate(post *domain.Post) error {
	if post == nil {
		return derr.NewBadRequestError("post is required")
	}

	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
	if post.Title == "" {
		return derr.NewBadRequestError("title is required")
	}
	if post.Content == "" {
		return derr.NewBadRequestError("content is required")
	}
	if post.UserID <= 0 {
		return derr.NewBadRequestError("user id is required")
	}

	return nil
}

func ValidatePaginatedFeedQuery(q *domain.PaginatedFeedQuery) error {
	if q == nil {
		return derr.NewBadRequestError("feed query is required")
	}
	if q.Limit <= 0 {
		return derr.NewBadRequestError("limit must be greater than zero")
	}
	if q.Offset < 0 {
		return derr.NewBadRequestError("offset must be greater than or equal to zero")
	}

	q.Sort = strings.TrimSpace(q.Sort)
	if q.Sort == "" {
		q.Sort = "desc"
	}

	return nil
}
