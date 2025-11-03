package chat

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

const maxTextSize = 4096

func (s *ChatService) SendMessage(ctx context.Context, chatID types.ChatID, text string) (*models.Message, error) {
	const api = "chat.ChatService.SendMessage"

	userID, err := s.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	if len(text) > maxTextSize {
		return nil, fmt.Errorf("%s: %w: max messsage size is %d", api, models.ErrInvalidArgument, maxTextSize)
	}

	chat, err := s.ChatRepository.GetChat(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if !slices.Contains(chat.Members, types.UserID(userID)) {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	msg := models.NewMessage()
	msg.ChatID = chatID
	msg.SenderID = types.UserID(userID)
	msg.Text = strings.TrimSpace(text)

	err = s.TransactionManager.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			if _, err := s.ChatRepository.CreateMessage(txCtx, msg); err != nil {
				return err
			}

			if err := s.OutboxRepository.SaveChatMessageSent(txCtx, msg.ChatID, msg); err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return msg, nil
}
