package chat

import (
	"context"
	"fmt"
	"time"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	chat_models "chat/internal/app/usecases/chat/models"
)

func (r *Repository) StreamMessages(ctx context.Context, chatID types.ChatID) (<-chan *models.Message, error) {
	const api = "chat.Repository.StreamMessages"

	const pollInterval = 1 * time.Second

	msgCh := make(chan *models.Message, 10)

	go func() {
		defer close(msgCh)

		sinceMessageTime := time.Now().UTC()
		ticker := time.NewTicker(pollInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				res, err := r.GetMessages(
					ctx,
					chatID,
					chat_models.WithGetMessagesSinceMessageTime(sinceMessageTime),
				)
				if err != nil {
					fmt.Printf("%s: error fetching messages: %v\n", api, err)
					continue
				}

				if len(res.Messages) == 0 {
					continue
				}

				for _, msg := range res.Messages {
					select {
					case msgCh <- msg:
					case <-ctx.Done():
						return
					}
				}

				sinceMessageTime = res.Messages[0].CreatedAt
			}
		}
	}()

	return msgCh, nil
}
