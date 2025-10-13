package chat

import (
	"context"
	"fmt"
	"slices"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (s *ChatService) GetChat(ctx context.Context, chatID types.ChatID) (*models.Chat, error) {
	const api = "chat.ChatService.GetChat"

	userID, err := s.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	chat, err := s.ChatRepository.GetChat(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if !slices.Contains(chat.Members, types.UserID(userID)) {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	return chat, nil
}
