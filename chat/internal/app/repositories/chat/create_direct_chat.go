package chat

import (
	"context"
	"errors"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (r *Repository) CreateDirectChat(ctx context.Context, userID, participantID types.UserID) (*models.Chat, error) {
	return nil, errors.New("unimplemented")
}
