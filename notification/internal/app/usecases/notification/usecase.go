package notification

import (
	"context"

	"notification/internal/app/modules/inbox"
)

type Usecase interface {
	// HandleBatch Возвращает списки успешных и проваленных id с их ошибками; err — для фатальных ошибок батча.
	HandleBatch(ctx context.Context, messages []*inbox.Message) (succeeded []string, failed map[string]error, err error)
}

var _ Usecase = (*NotificationService)(nil)

type NotificationService struct{}

func NewUsecase() Usecase {
	return &NotificationService{}
}
