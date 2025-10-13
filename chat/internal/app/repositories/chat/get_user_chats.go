package chat

import (
	"context"
	"errors"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func (r *Repository) GetUserChats(ctx context.Context, userID types.UserID) ([]*models.Chat, error) {
	return nil, errors.New("unimplemented")
}
