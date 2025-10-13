package chat

import (
	"context"
	"fmt"
	"slices"
	"time"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	chat_models "chat/internal/app/usecases/chat/models"
)

func (s *ChatService) ListMessages(ctx context.Context, req *chat_models.ListMessagesRequest) (*chat_models.ListMessagesResult, error) {
	const api = "chat.ChatService.ListMessages"

	userID, err := s.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	chat, err := s.ChatRepository.GetChat(ctx, req.ChatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if !slices.Contains(chat.Members, types.UserID(userID)) {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	if req.Limit == 0 {
		req.Limit = chat_models.MessagesLimit
	}

	// Если не пришло время последнего сообщения, то считаем от текущего времени
	if req.LastMessageTime.Unix() == 0 {
		req.LastMessageTime = time.Now().UTC()
	}

	res, err := s.ChatRepository.GetMessages(
		ctx,
		req.ChatID,
		chat_models.WithGetMessagesLastMessageTime(req.LastMessageTime),
		chat_models.WithGetMessagesLimit(req.Limit),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return &chat_models.ListMessagesResult{
		Messages:            res.Messages,
		NextLastMessageTime: res.NextLastMessageTime,
	}, nil
}
