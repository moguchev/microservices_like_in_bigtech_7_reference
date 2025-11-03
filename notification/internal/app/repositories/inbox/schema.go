package inbox

import (
	"database/sql"
	"time"
)

const tableInboxMessages = "public.inbox_messages"

const (
	columnInboxID          = "id"
	columnInboxTopic       = "topic"
	columnInboxPartition   = "partition"
	columnInboxOffset      = "ofset"
	columnInboxPayload     = "payload"
	columnInboxStatus      = "status"
	columnInboxAttempts    = "attempts"
	columnInboxLastError   = "last_error"
	columnInboxReceivedAt  = "received_at"
	columnInboxProcessedAt = "processed_at"
)

type inboxMessage struct {
	ID          string         `db:"id"`
	Topic       string         `db:"topic"`
	Partition   int32          `db:"partition"`
	Offset      int64          `db:"ofset"`
	Payload     []byte         `db:"payload"`
	Status      string         `db:"status"`
	Attempts    int            `db:"attempts"`
	LastError   sql.NullString `db:"last_error"`
	ReceivedAt  time.Time      `db:"received_at"`
	ProcessedAt sql.NullTime   `db:"processed_at"`
}

var (
	tableInboxMessagesColumns = []string{
		columnInboxID,
		columnInboxTopic,
		columnInboxPartition,
		columnInboxOffset,
		columnInboxPayload,
		columnInboxStatus,
		columnInboxAttempts,
		columnInboxLastError,
		columnInboxReceivedAt,
		columnInboxProcessedAt,
	}
)

func (m *inboxMessage) mapFields() map[string]any {
	return map[string]any{
		columnInboxID:          m.ID,
		columnInboxTopic:       m.Topic,
		columnInboxPartition:   m.Partition,
		columnInboxOffset:      m.Offset,
		columnInboxPayload:     m.Payload,
		columnInboxStatus:      m.Status,
		columnInboxAttempts:    m.Attempts,
		columnInboxLastError:   m.LastError,
		columnInboxReceivedAt:  m.ReceivedAt,
		columnInboxProcessedAt: m.ProcessedAt,
	}
}

func (m *inboxMessage) Values(columns ...string) []any {
	f := m.mapFields()
	values := make([]any, 0, len(columns))
	for _, c := range columns {
		if v, ok := f[c]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}
