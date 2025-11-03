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

	if request.Status == models.FriendRequestStatusAccepted {
		return nil, fmt.Errorf("%s: request is already accepted %w", api, models.ErrInvalidArgument)
	}

	var friendRequest *models.FriendRequest
	err = uc.TransactionManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			// Обновляем статус заявки
			var updateErr error
			friendRequest, updateErr = uc.SocialRepository.UpdateFriendRequestStatus(txCtx, id, models.FriendRequestStatusAccepted)
			if updateErr != nil {
				return updateErr
			}

			// Создаем пару друзей
			if err := uc.SocialRepository.CreateFriendPair(txCtx, request.FromUser, request.ToUser); err != nil {
				return err
			}

			// Добавляем событие в outbox
			if err := uc.OutboxRepository.SaveFriendRequestUpdated(txCtx, request.ToUser, friendRequest); err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return friendRequest, nil
}
