package models

import (
	"time"

	"chat/internal/app/models/types"

	"github.com/google/uuid"
)

// Message - сообщение
type Message struct {
	ID        types.MessageID
	ChatID    types.ChatID
	SenderID  types.UserID
	Text      string
	CreatedAt time.Time
}

func NewMessage() *Message {
	return &Message{
		ID:        types.MessageID(uuid.New().String()),
		CreatedAt: time.Now(),
	}
}
