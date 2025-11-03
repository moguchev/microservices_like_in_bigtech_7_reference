package outbox

import (
	"time"

	"github.com/google/uuid"
)

type AggregateType string

const (
	AggregateTypeFriendRequest AggregateType = "FriendRequest"
)

type EventType string

const (
	EventTypeFriendRequestCreated EventType = "FriendRequestCreated"
	EventTypeFriendRequestUpdated EventType = "FriendRequestUpdated"
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
