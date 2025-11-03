package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"notification/internal/app/models"
	"notification/internal/app/modules/inbox"
)

// HandleBatch - симуляция отправки уведомления (вывод в консоль)
func (n *NotificationService) HandleBatch(ctx context.Context, messages []*inbox.Message) (succeeded []string, failed map[string]error, err error) {
	if len(messages) == 0 {
		log.Println("KafkaFriendRequestBatchHandler", "nothing to send")
		return nil, nil, nil
	}

	defer func() {
		if err != nil {
			log.Println("HandleBatch", err)
		} else {
			log.Println("HandleBatch", "succeeded", succeeded, "failed", failed)
		}
	}()

	succeeded = make([]string, 0, len(messages))
	failed = make(map[string]error, len(messages))

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, msg := range messages {
		// Демо случайной ошибки с вероятностью 33%
		randomErrNumber := rng.Intn(100)
		if randomErrNumber < 33 {
			failed[msg.ID] = fmt.Errorf("simulated random error %d", randomErrNumber)
			continue
		}

		switch msg.Topic {
		case "social.friend.request", "social.friend.updated":
			var fr models.FriendRequest
			if e := json.Unmarshal(msg.Payload, &fr); e != nil {
				failed[msg.ID] = fmt.Errorf("failed to parse FriendRequest: %w", e)
				continue
			}

			fmt.Printf("[Topic=%s] FriendRequest: ID=%s, From=%s, To=%s, Status=%s\n", msg.Topic, fr.ID, fr.FromUser, fr.ToUser, fr.Status)
			succeeded = append(succeeded, msg.ID)

		case "chat.message.sent":
			var m models.ChatMessage
			if e := json.Unmarshal(msg.Payload, &m); e != nil {
				failed[msg.ID] = fmt.Errorf("failed to parse ChatMessage: %w", e)
				continue
			}

			fmt.Printf("[Topic=%s] ChatMessage: ID=%s, ChatID=%s, SenderID=%s, Text=%s\n", msg.Topic, m.ID, m.ChatID, m.SenderID, m.Text)
			succeeded = append(succeeded, msg.ID)

		default:
			failed[msg.ID] = fmt.Errorf("unsupported topic: %s", msg.Topic)
		}
	}

	return succeeded, failed, nil
}
