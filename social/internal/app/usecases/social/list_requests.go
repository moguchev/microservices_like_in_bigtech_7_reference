package social

import (
	"context"
	"fmt"

	"social/internal/app/models"
	"social/internal/app/models/types"
)

func (uc *SocialService) ListRequests(ctx context.Context, id types.UserID) ([]*models.FriendRequest, error) {
	const api = "social.SocialService.ListRequests"

	userID, err := uc.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	if id != types.UserID(userID) {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	requests, err := uc.SocialRepository.ListRequests(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return requests, nil
}
