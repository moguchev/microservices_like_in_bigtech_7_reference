package social

import (
	"context"
	"errors"

	"social/internal/app/models/types"
	social_models "social/internal/app/usecases/social/models"
)

func (r *Repository) ListFriends(ctx context.Context, id types.UserID, opts ...social_models.ListFriendsOption) (*social_models.ListFriendsResult, error) {
	return nil, errors.New("unimplemented")
}
