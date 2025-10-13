package chat

import (
	"context"
	"fmt"
	"slices"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (s *ChatService) ListChatMembers(ctx context.Context, chatID types.ChatID) ([]types.UserID, error) {
	const api = "chat.ChatService.ListChatMembers"

	userID, err := s.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	chat, err := s.ChatRepository.GetChat(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	members := chat.Members

	if !slices.Contains(members, types.UserID(userID)) {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	return members, nil
}
