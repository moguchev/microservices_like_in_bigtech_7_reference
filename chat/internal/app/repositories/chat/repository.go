package chat

import (
	"chat/internal/app/usecases/chat"
)

type Repository struct{}

var (
	_ chat.ChatRepository = (*Repository)(nil)
)

func NewRepository() *Repository {
	return &Repository{}
}
