package social

import (
	"context"
	"errors"

	"social/internal/app/models"
	"social/internal/app/models/types"
)

func (r *Repository) ListRequests(ctx context.Context, userID types.UserID) ([]*models.FriendRequest, error) {
	return nil, errors.New("unimplemented")
}
