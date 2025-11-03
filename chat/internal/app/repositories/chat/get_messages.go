package chat

import (
	"context"
	"fmt"
	"time"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	chat_models "chat/internal/app/usecases/chat/models"
	"chat/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) GetMessages(ctx context.Context, chatID types.ChatID, opts ...chat_models.GetMessagesOption) (*chat_models.GetMessagesResult, error) {
	const api = "chat.Repository.GetMessages"

	o := chat_models.CollectListMessagesOptions(opts...)

	qb := r.qb.
		Select(tableMessagesColumns...).
		From(tableMessages).
		Where(squirrel.Eq{columnMessageChatID: chatID.String()}).
		Limit(uint64(o.Limit))

	if o.BeforeMessageTime != nil {
		qb = qb.Where(squirrel.Lt{columnCreatedAt: *o.BeforeMessageTime}).
			OrderBy(columnCreatedAt + " DESC")
	}

	if o.SinceMessageTime != nil {
		qb = qb.Where(squirrel.Gt{columnCreatedAt: *o.SinceMessageTime}).
			OrderBy(columnCreatedAt + " ASC")
	}

	var rows []messageRow
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, qb); err != nil {
		return nil, fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	messages := make([]*models.Message, 0, len(rows))
	var nextTime time.Time

	for i := range rows {
		messages = append(messages, toMessageModel(&rows[i]))
	}

	if len(rows) > 0 {
		nextTime = rows[len(rows)-1].CreatedAt
	}

	return &chat_models.GetMessagesResult{
		Messages:            messages,
		NextLastMessageTime: nextTime,
	}, nil
}
