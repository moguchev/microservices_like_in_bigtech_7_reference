package chat

import (
	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

func fromChatModel(m *models.Chat) *chatRow {
	if m == nil {
		return nil
	}
	return &chatRow{
		ID:        m.ID.String(),
		CreatedAt: m.CreatedAt,
	}
}

func fromMessageModel(m *models.Message) *messageRow {
	if m == nil {
		return nil
	}
	return &messageRow{
		ID:        m.ID.String(),
		ChatID:    m.ChatID.String(),
		SenderID:  m.SenderID.String(),
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
	}
}

func toMessageModel(r *messageRow) *models.Message {
	if r == nil {
		return nil
	}
	return &models.Message{
		ID:        types.MessageID(r.ID),
		ChatID:    types.ChatID(r.ChatID),
		SenderID:  types.UserID(r.SenderID),
		Text:      r.Text,
		CreatedAt: r.CreatedAt,
	}
}
