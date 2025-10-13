package social

import (
	"context"
	"fmt"

	lib_types "lib/types"
	"social/internal/app/models"
	"social/internal/app/models/types"
	social_models "social/internal/app/usecases/social/models"
)

func (uc *SocialService) AcceptFriendRequest(ctx context.Context, id types.RequestID) (*models.FriendRequest, error) {
	const api = "social.SocialService.AcceptFriendRequest"

	userID, err := uc.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	requests, err := uc.SocialRepository.GetFriendRequests(ctx, social_models.FriendRequestSelector{
		SelectorType: social_models.SelectorTypeAND,
		RequestID:    lib_types.WrapField(id),
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if len(requests) == 0 {
		return nil, fmt.Errorf("%s: request with id '%s': %w", api, id, models.ErrNotFound)
	}

	request := requests[0]

	if request.ToUser != types.UserID(userID) {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	friendRequest, err := uc.SocialRepository.UpdateFriendRequestStatus(ctx, id, models.FriendRequestStatusAccepted)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return friendRequest, nil
}
