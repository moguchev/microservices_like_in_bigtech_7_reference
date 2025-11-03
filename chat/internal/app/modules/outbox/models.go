package outbox

import (
	"time"

	"github.com/google/uuid"
)

type AggregateType string

const (
	AggregateTypeChatMessage AggregateType = "ChatMessage"
)

type EventType string

const (
	EventTypeChatMessageSent EventType = "ChatMessageSent"
)

type Event struct {
	ID            uuid.UUID
	AggregateType AggregateType
	AggregateID   string
	EventType     EventType
	Payload       []byte
	CreatedAt     time.Time
	PublishedAt   *time.Time
	RetryCount    int
}
