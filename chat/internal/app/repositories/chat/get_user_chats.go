package chat

import (
	"context"
	"fmt"

	"chat/internal/app/models"
	"chat/internal/app/models/types"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) GetUserChats(ctx context.Context, userID types.UserID) ([]*models.Chat, error) {
	const api = "chat.Repository.GetUserChats"

	chatIDs, err := r.getChatIDsByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if len(chatIDs) == 0 {
		return []*models.Chat{}, nil
	}

	chats, err := r.getChatsByIDs(ctx, chatIDs)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return chats, nil
}

// getChatIDsByUser возвращает id чатов, где участвует пользователь
func (r *Repository) getChatIDsByUser(ctx context.Context, userID types.UserID) ([]string, error) {
	query := r.qb.Select(columnChatMemberChatID).
		From(tableChatMembers).
		Where(squirrel.Eq{columnChatMemberUserID: userID.String()})

	var ids []string
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &ids, query); err != nil {
		return nil, err
	}

	return ids, nil
}

// getChatsByIDs возвращает чаты по списку их id
func (r *Repository) getChatsByIDs(ctx context.Context, ids []string) ([]*models.Chat, error) {
	query := r.qb.Select(tableChatsColumns...).
		From(tableChats).
		Where(squirrel.Eq{columnChatID: ids})

	var rows []chatRow
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, query); err != nil {
		return nil, err
	}

	chats := make([]*models.Chat, len(rows))
	for i, row := range rows {
		chats[i] = &models.Chat{
			ID:        types.ChatID(row.ID),
			CreatedAt: row.CreatedAt,
			Members:   nil,
		}
	}

	if err := r.fillChatMembers(ctx, chats); err != nil {
		return nil, err
	}

	return chats, nil
}

// fillChatMembers подставляет участников для каждого чата
func (r *Repository) fillChatMembers(ctx context.Context, chats []*models.Chat) error {
	if len(chats) == 0 {
		return nil
	}

	ids := make([]string, len(chats))
	for i, c := range chats {
		ids[i] = c.ID.String()
	}

	query := r.qb.Select(tableChatMembersColumns...).
		From(tableChatMembers).
		Where(squirrel.Eq{columnChatMemberChatID: ids})

	var rows []chatMemberRow
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, query); err != nil {
		return err
	}

	chatMap := make(map[string]*models.Chat, len(chats))
	for _, c := range chats {
		chatMap[c.ID.String()] = c
	}

	for _, m := range rows {
		if c, ok := chatMap[m.ChatID]; ok {
			c.Members = append(c.Members, types.UserID(m.UserID))
		}
	}

	return nil
}
