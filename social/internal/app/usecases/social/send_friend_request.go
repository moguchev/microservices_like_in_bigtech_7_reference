package social

import (
	"context"
	"fmt"

	"social/internal/app/models"
	"social/internal/app/models/types"
)

func (uc *SocialService) SendFriendRequest(ctx context.Context, id types.UserID) (*models.FriendRequest, error) {
	const api = "social.SocialService.SendFriendRequest"

	userID, err := uc.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	if types.UserID(userID) == id {
		return nil, fmt.Errorf("%s: %w: user can't send request to himself", api, models.ErrInvalidArgument)
	}

	newFriendRequest := models.NewFriendRequest()
	newFriendRequest.FromUser = types.UserID(userID)
	newFriendRequest.ToUser = id
	newFriendRequest.Status = models.FriendRequestStatusPending

	err = uc.TransactionManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			if _, err := uc.SocialRepository.CreateFriendRequest(txCtx, newFriendRequest); err != nil {
				return err
			}

			if err := uc.OutboxRepository.SaveFriendRequestCreated(txCtx, newFriendRequest.ToUser, newFriendRequest); err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return newFriendRequest, nil
}
