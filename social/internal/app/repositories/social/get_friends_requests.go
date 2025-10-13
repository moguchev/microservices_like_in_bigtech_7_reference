package social

import (
	"context"
	"errors"

	"social/internal/app/models"
	social_models "social/internal/app/usecases/social/models"
)

func (r *Repository) GetFriendRequests(ctx context.Context, selector social_models.FriendRequestSelector) ([]*models.FriendRequest, error) {
	return nil, errors.New("unimplemented")
}
