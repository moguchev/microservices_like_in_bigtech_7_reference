package chat

import (
	"context"
	"fmt"

	"chat/internal/app/models"
	"chat/internal/pkg/postgres"
)

func (r *Repository) CreateDirectChat(ctx context.Context, chat *models.Chat) error {
	const api = "chat.Repository.CreateDirectChat"

	if chat == nil || len(chat.Members) != 2 {
		return fmt.Errorf("%s: %w", api, models.ErrInvalidArgument)
	}

	err := r.tm.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			if err := r.createChat(txCtx, chat); err != nil {
				return err
			}

			if err := r.createChatMembers(txCtx, chat); err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("%s: %w", api, err)
	}

	return nil
}

func (r *Repository) createChat(ctx context.Context, chat *models.Chat) error {
	const api = "chat.Repository.createChat"

	row := fromChatModel(chat)

	query := r.qb.Insert(tableChats).
		Columns(tableChatsColumns...).
		Values(row.Values(tableChatsColumns...)...)

	if _, err := r.db.GetQueryEngine(ctx).Execx(ctx, query); err != nil {
		return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return nil
}

func (r *Repository) createChatMembers(ctx context.Context, chat *models.Chat) error {
	const api = "chat.Repository.createChatMembers"

	query := r.qb.Insert(tableChatMembers).
		Columns(tableChatMembersColumns...)

	for _, id := range chat.Members {
		query = query.Values(chat.ID, id.String())
	}

	if _, err := r.db.GetQueryEngine(ctx).Execx(ctx, query); err != nil {
		return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return nil
}
