package outbox

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"chat/internal/app/models"
	"chat/internal/app/models/types"

	"github.com/google/uuid"
)

func (p *Processor) SaveChatMessageSent(ctx context.Context, chatID types.ChatID, msg *models.Message) error {
	const api = "outbox.Processor.SaveChatMessageSent"

	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("%s: Marshal: %w", api, err)
	}

	e := Event{
		ID:            uuid.New(),
		AggregateType: AggregateTypeChatMessage,
		AggregateID:   chatID.String(),
		EventType:     EventTypeChatMessageSent,
		Payload:       payload,
		CreatedAt:     time.Now().UTC(),
	}

	return p.Repository.SaveEvent(ctx, &e)
}
