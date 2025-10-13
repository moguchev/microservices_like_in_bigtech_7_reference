package social

import (
	"context"
	"fmt"
	"time"

	social_models "social/internal/app/usecases/social/models"
)

func (uc *SocialService) ListFriends(ctx context.Context, req *social_models.ListFriendsRequest) (*social_models.ListFriendsResult, error) {
	const api = "social.SocialService.ListFriends"

	if req.Limit == 0 {
		req.Limit = social_models.FriendsLimit
	}

	// Если не пришло время последней принятой заявки, то считаем от текущего времени
	if req.LastAcceptRequestTime.Unix() == 0 {
		req.LastAcceptRequestTime = time.Now().UTC()
	}

	res, err := uc.SocialRepository.ListFriends(
		ctx,
		req.UserID,
		social_models.WithGetMessagesLimit(req.Limit),
		social_models.WithGetMessagesLastMessageTime(req.LastAcceptRequestTime),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return res, nil
}
