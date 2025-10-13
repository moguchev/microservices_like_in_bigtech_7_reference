package chat

import (
	"context"
	"errors"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (r *Repository) StreamMessages(ctx context.Context, chatID types.ChatID) (<-chan *models.Message, error) {
	return nil, errors.New("unimplemented")
}
