package models

import (
	"time"

	"chat/internal/app/models/types"
)

// Chat - чат
type Chat struct {
	ID        types.ChatID
	Members   []types.UserID
	CreatedAt time.Time
}
