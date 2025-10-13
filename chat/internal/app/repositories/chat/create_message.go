package chat

import (
	"context"
	"errors"

	"chat/internal/app/models"
)

func (r *Repository) CreateMessage(ctx context.Context, msg *models.Message) (*models.Message, error) {
	return nil, errors.New("unimplemented")
}
