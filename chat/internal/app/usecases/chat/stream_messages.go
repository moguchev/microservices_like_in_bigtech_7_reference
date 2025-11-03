package chat

import (
	"context"
	"fmt"
	"log"
	"time"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	chat_models "chat/internal/app/usecases/chat/models"
)

func (s *ChatService) StreamMessages(ctx context.Context, chatID types.ChatID, sinceMessageTime time.Time) (<-chan *models.Message, error) {
	const api = "chat.ChatService.StreamMessages"

	_, err := s.GetChat(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	const (
		chanBufferSize      = 10
		messageHistoryLimit = 50
	)

	msgCh := make(chan *models.Message, chanBufferSize)

	go func() {
		defer close(msgCh)

		// Если пришла метка времени сообщения, то получаем историю сообщений
		if sinceMessageTime.Unix() != 0 {
			res, errGetMessages := s.ChatRepository.GetMessages(
				ctx,
				chatID,
				chat_models.WithGetMessagesLimit(messageHistoryLimit),
				chat_models.WithGetMessagesSinceMessageTime(sinceMessageTime),
			)
			if errGetMessages != nil {
				log.Printf("%s: GetMessages error: %s", api, errGetMessages)
				return
			}
			for _, m := range res.Messages {
				select {
				case msgCh <- m:
				case <-ctx.Done():
					return
				}
			}
		}

		// Подписка на новые сообщения
		streamCh, errStreamMessages := s.ChatRepository.StreamMessages(ctx, chatID)
		if errStreamMessages != nil {
			log.Printf("%s: StreamMessages error: %s", api, errStreamMessages)
			return
		}

		for {
			select {
			case msg, ok := <-streamCh:
				if !ok {
					return
				}
				select {
				case msgCh <- msg:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return msgCh, nil
}
