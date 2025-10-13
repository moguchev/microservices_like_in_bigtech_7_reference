package chat

import (
	"context"
	"fmt"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (s *ChatService) ListUserChats(ctx context.Context, userID types.UserID) ([]*models.Chat, error) {
	const api = "chat.ChatService.ListUserChats"

	currentUserID, err := s.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	if types.UserID(currentUserID) != userID {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	chats, err := s.ChatRepository.GetUserChats(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return chats, nil
}
