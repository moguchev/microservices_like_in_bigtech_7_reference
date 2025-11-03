package inbox

import (
	"time"
)

// Status - статус обработки inbox сообщения
type Status string

const (
	StatusUnspecified Status = "unspecified"
	StatusReceived    Status = "received"
	StatusProcessing  Status = "processing"
	StatusProcessed   Status = "processed"
	StatusFailed      Status = "failed"
)

func (s Status) String() string {
	return string(s)
}

// Message - inbox сообщение
type Message struct {
	ID          string
	Topic       string
	Partition   int32
	Offset      int64
	Payload     []byte
	Status      Status
	Attempts    int
	LastError   string
	ReceivedAt  time.Time
	ProcessedAt time.Time
}
