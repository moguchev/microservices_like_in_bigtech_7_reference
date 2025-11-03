package chat

import (
	"context"
	"fmt"
	"slices"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (s *ChatService) CreateDirectChat(ctx context.Context, participantID types.UserID) (*models.Chat, error) {
	const api = "chat.ChatService.CreateDirectChat"

	rawUserID, err := s.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	userID := types.UserID(rawUserID)

	if userID == participantID {
		return nil, fmt.Errorf("%s: %w: user can't create chat with himself", api, models.ErrInvalidArgument)
	}

	chats, err := s.ChatRepository.GetUserChats(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	for _, chat := range chats {
		if slices.Contains(chat.Members, userID) && slices.Contains(chat.Members, participantID) {
			return nil, fmt.Errorf("%s: %w", api, models.ErrAlreadyExists)
		}
	}

	chat := models.NewChat(userID, participantID)

	if err := s.ChatRepository.CreateDirectChat(ctx, chat); err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return chat, nil
}
