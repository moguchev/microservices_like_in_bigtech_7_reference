package chat

import (
	"context"
	"fmt"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (s *ChatService) CreateDirectChat(ctx context.Context, participantID types.UserID) (*models.Chat, error) {
	const api = "chat.ChatService.CreateDirectChat"

	userID, err := s.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	if types.UserID(userID) == participantID {
		return nil, fmt.Errorf("%s: %w: user can't create chat with himself", api, models.ErrInvalidArgument)
	}

	chat, err := s.ChatRepository.CreateDirectChat(ctx, types.UserID(userID), participantID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return chat, nil
}
