package inbox

import (
	"database/sql"

	"notification/internal/app/modules/inbox"
)

func fromModel(m *inbox.Message) *inboxMessage {
	if m == nil {
		return nil
	}

	return &inboxMessage{ID: m.ID,
		Topic:       m.Topic,
		Partition:   m.Partition,
		Offset:      m.Offset,
		Payload:     notnullJSON(m.Payload),
		Status:      m.Status.String(),
		Attempts:    m.Attempts,
		LastError:   sql.NullString{String: m.LastError, Valid: m.LastError != ""},
		ReceivedAt:  m.ReceivedAt,
		ProcessedAt: sql.NullTime{Time: m.ProcessedAt, Valid: !m.ProcessedAt.IsZero()},
	}
}

func notnullJSON(data []byte) []byte {
	if data == nil {
		return []byte("[]")
	}
	return data
}

func toModel(m *inboxMessage) *inbox.Message {
	if m == nil {
		return nil
	}

	return &inbox.Message{
		ID:          m.ID,
		Topic:       m.Topic,
		Partition:   m.Partition,
		Offset:      m.Offset,
		Payload:     m.Payload,
		Status:      inbox.Status(m.Status),
		Attempts:    m.Attempts,
		LastError:   m.LastError.String,
		ReceivedAt:  m.ReceivedAt,
		ProcessedAt: m.ProcessedAt.Time,
	}
}
