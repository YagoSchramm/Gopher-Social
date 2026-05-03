package rules

import (
	"strings"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/domain"
)

func ValidateCommentForCreate(comment *domain.Comment) error {
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

	return nil
}

func ValidatePostID(postID int64) error {
	if postID <= 0 {
		return derr.NewBadRequestError("post id is required")
	}

	return nil
}
