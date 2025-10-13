package chat

import (
	"context"
	"errors"

	"chat/internal/app/models/types"
	chat_models "chat/internal/app/usecases/chat/models"
)

func (r *Repository) GetMessages(ctx context.Context, chatID types.ChatID, opts ...chat_models.GetMessagesOption) (*chat_models.GetMessagesResult, error) {
	return nil, errors.New("unimplemented")
}
