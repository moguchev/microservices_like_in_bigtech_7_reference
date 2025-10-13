package social

import (
	"context"
	"errors"

	"social/internal/app/models"
)

func (r *Repository) CreateFriendRequest(ctx context.Context, request *models.FriendRequest) (*models.FriendRequest, error) {
	return nil, errors.New("unimplemented")
}
