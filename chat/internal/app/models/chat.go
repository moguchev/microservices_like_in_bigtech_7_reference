package models

import (
	"time"

	"chat/internal/app/models/types"

	"github.com/google/uuid"
)

// Chat - чат
type Chat struct {
	ID        types.ChatID
	Members   []types.UserID
	CreatedAt time.Time
}

func NewChat(members ...types.UserID) *Chat {
	return &Chat{
		ID:        types.ChatID(uuid.New().String()),
		Members:   members,
		CreatedAt: time.Now().UTC(),
	}
}
