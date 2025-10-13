package social

import (
	"context"
	"errors"

	"social/internal/app/models"
	"social/internal/app/models/types"
)

func (r *Repository) UpdateFriendRequestStatus(ctx context.Context, id types.RequestID, status models.FriendRequestStatus) (*models.FriendRequest, error) {
	return nil, errors.New("unimplemented")
}
