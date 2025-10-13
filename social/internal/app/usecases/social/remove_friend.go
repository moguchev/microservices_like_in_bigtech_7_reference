package social

import (
	"context"
	"fmt"

	"social/internal/app/models"
	"social/internal/app/models/types"
)

func (uc *SocialService) RemoveFriend(ctx context.Context, friendID types.UserID) error {
	const api = "social.SocialService.RemoveFriend"

	userID, err := uc.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	if err = uc.SocialRepository.RemoveFriend(ctx, types.UserID(userID), friendID); err != nil {
		return fmt.Errorf("%s: %w", api, err)
	}

	return nil
}
