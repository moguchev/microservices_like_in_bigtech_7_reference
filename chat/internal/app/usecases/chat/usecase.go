package chat

import (
	"context"
	"time"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	chat_models "chat/internal/app/usecases/chat/models"
)

type Usecase interface {
	// CreateDirectChat - создать личный чат
	CreateDirectChat(ctx context.Context, participantID types.UserID) (*models.Chat, error)
	// GetChat - получить информацию о чате
	GetChat(ctx context.Context, chatID types.ChatID) (*models.Chat, error)
	// ListUserChats - получить список чатов пользователя
	ListUserChats(ctx context.Context, userID types.UserID) ([]*models.Chat, error)
	// ListChatMembers - получить участников чата
	ListChatMembers(ctx context.Context, chatID types.ChatID) ([]types.UserID, error)
	// SendMessage - отправить сообщение
	SendMessage(ctx context.Context, chatID types.ChatID, text string) (*models.Message, error)
	// ListMessages - получить историю сообщений
	ListMessages(ctx context.Context, req *chat_models.ListMessagesRequest) (*chat_models.ListMessagesResult, error)
	// StreamMessages - серверный стрим новых сообщений
	StreamMessages(ctx context.Context, chatID types.ChatID, sinceMessageTime time.Time) (<-chan *models.Message, error)
}

type (
	ChatRepository interface {
		// CreateDirectChat - создать чат
		CreateDirectChat(ctx context.Context, chat *models.Chat) error
		// GetChat - получить информацию о чате
		//
		// errors: models.ErrNotFound
		GetChat(ctx context.Context, chatID types.ChatID) (*models.Chat, error)
		// GetUserChats - список чатов пользователя
		GetUserChats(ctx context.Context, userID types.UserID) ([]*models.Chat, error)
		// CreateMessage - создать сообщение
		CreateMessage(ctx context.Context, msg *models.Message) (*models.Message, error)
		// GetMessages - получить список сообщений
		GetMessages(ctx context.Context, chatID types.ChatID, opts ...chat_models.GetMessagesOption) (*chat_models.GetMessagesResult, error)
		// StreamMessages - подписка на новые сообщения
		StreamMessages(ctx context.Context, chatID types.ChatID) (<-chan *models.Message, error)
	}

	OutboxRepository interface {
		// SaveChatMessageSent - запись в Outbox сообщения по отправке сообщения чата
		SaveChatMessageSent(ctx context.Context, chatID types.ChatID, msg *models.Message) error
	}

	TransactionManager interface {
		RunReadCommitted(ctx context.Context, f func(ctx context.Context) error) error
	}

	UserIDProvider interface {
		GetUserIDFromIncomingContext(ctx context.Context) (string, error)
	}
)

type Deps struct {
	ChatRepository     ChatRepository
	TransactionManager TransactionManager
	OutboxRepository   OutboxRepository
	UserIDProvider
}

var _ Usecase = (*ChatService)(nil)

type ChatService struct {
	Deps
}

func NewUsecase(d Deps) Usecase {
	return &ChatService{Deps: d}
}
