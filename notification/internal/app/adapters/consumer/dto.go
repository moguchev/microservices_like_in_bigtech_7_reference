package consumer

import (
	"time"

	"notification/internal/app/modules/inbox"

	"github.com/IBM/sarama"
)

func toInboxMessage(eventID string, msg *sarama.ConsumerMessage) *inbox.Message {
	return &inbox.Message{
		ID:         eventID,
		Topic:      msg.Topic,
		Partition:  msg.Partition,
		Offset:     msg.Offset,
		Payload:    msg.Value,
		Status:     inbox.StatusReceived,
		ReceivedAt: time.Now(),
	}
}
