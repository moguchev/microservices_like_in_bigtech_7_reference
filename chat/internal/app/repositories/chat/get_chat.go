package chat

import (
	"context"
	"fmt"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	"chat/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) GetChat(ctx context.Context, chatID types.ChatID) (*models.Chat, error) {
	const api = "chat.Repository.GetChat"

	chatInfo, err := r.getChatInfo(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	members, err := r.getChatMembers(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return &models.Chat{
		ID:        chatID,
		Members:   members,
		CreatedAt: chatInfo.CreatedAt,
	}, nil
}

func (r *Repository) getChatInfo(ctx context.Context, chatID types.ChatID) (*chatRow, error) {
	const api = "chat.Repository.getChatInfo"

	query := r.qb.
		Select(tableChatsColumns...).
		From(tableChats).
		Where(squirrel.Eq{columnChatID: chatID.String()})

	var row chatRow
	if err := r.db.GetQueryEngine(ctx).Getx(ctx, &row, query); err != nil {
		return nil, fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return &row, nil
}

func (r *Repository) getChatMembers(ctx context.Context, chatID types.ChatID) ([]types.UserID, error) {
	const api = "chat.Repository.getChatMembers"

	query := r.qb.
		Select(tableChatMembersColumns...).
		From(tableChatMembers).
		Where(squirrel.Eq{columnChatMemberChatID: chatID.String()})

	var rows []chatMemberRow
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, query); err != nil {
		return nil, fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	members := make([]types.UserID, len(rows))
	for i, m := range rows {
		members[i] = types.UserID(m.UserID)
	}

	return members, nil
}
