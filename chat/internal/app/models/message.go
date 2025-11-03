package models

import (
	"time"

	"chat/internal/app/models/types"

	"github.com/google/uuid"
)

// Message - сообщение
type Message struct {
	ID        types.MessageID `json:"id"`
	ChatID    types.ChatID    `json:"chat_id"`
	SenderID  types.UserID    `json:"sender_id"`
	Text      string          `json:"text"`
	CreatedAt time.Time       `json:"created_at"`
}

func NewMessage() *Message {
	return &Message{
		ID:        types.MessageID(uuid.New().String()),
		CreatedAt: time.Now(),
	}
}
