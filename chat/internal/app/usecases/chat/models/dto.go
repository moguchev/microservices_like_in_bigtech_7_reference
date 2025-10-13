package models

import (
	"time"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
)

type ListMessagesRequest struct {
	ChatID          types.ChatID
	Limit           uint32
	LastMessageTime time.Time
}

type ListMessagesResult struct {
	Messages            []*models.Message
	NextLastMessageTime time.Time
}

type GetMessagesResult struct {
	Messages            []*models.Message
	NextLastMessageTime time.Time
}

type getMessagesOptions struct {
	Limit           uint32
	LastMessageTime time.Time
}

type GetMessagesOption func(o *getMessagesOptions)

func WithGetMessagesLimit(n uint32) GetMessagesOption {
	return func(o *getMessagesOptions) { o.Limit = n }
}

func WithGetMessagesLastMessageTime(t time.Time) GetMessagesOption {
	return func(o *getMessagesOptions) { o.LastMessageTime = t }
}

const MessagesLimit = 100

func CollectListMessagesOptions(opts ...GetMessagesOption) getMessagesOptions {
	res := getMessagesOptions{
		Limit:           MessagesLimit,
		LastMessageTime: time.Now().UTC(),
	}

	for _, opt := range opts {
		opt(&res)
	}
	return res
}
