package repository

import "context"

type FollowerRepository interface {
	Follow(ctx context.Context, userID, followerID int64) error
	Unfollow(ctx context.Context, followerID, userID int64) error
}
