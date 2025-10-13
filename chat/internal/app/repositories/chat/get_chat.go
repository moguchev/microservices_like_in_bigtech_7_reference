package chat

import (
	"context"
	"errors"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (r *Repository) GetChat(ctx context.Context, chatID types.ChatID) (*models.Chat, error) {
	return nil, errors.New("unimplemented")
}
