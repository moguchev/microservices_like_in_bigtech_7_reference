package chat

import "time"

const (
	tableChats       = "public.chats"
	tableChatMembers = "public.chat_members"
	tableMessages    = "public.messages"
)

const (
	// chats
	columnChatID    = "id"
	columnCreatedAt = "created_at"

	// chat_members
	columnChatMemberChatID = "chat_id"
	columnChatMemberUserID = "user_id"

	// messages
	columnMessageID     = "id"
	columnMessageChatID = "chat_id"
	columnSenderID      = "sender_id"
	columnText          = "text"
)

var tableChatsColumns = []string{
	columnChatID,
	columnCreatedAt,
}

var tableChatMembersColumns = []string{
	columnChatMemberChatID,
	columnChatMemberUserID,
}

var tableMessagesColumns = []string{
	columnMessageID,
	columnMessageChatID,
	columnSenderID,
	columnText,
	columnCreatedAt,
}

type chatRow struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

type chatMemberRow struct {
	ChatID string `db:"chat_id"`
	UserID string `db:"user_id"`
}

type messageRow struct {
	ID        string    `db:"id"`
	ChatID    string    `db:"chat_id"`
	SenderID  string    `db:"sender_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *chatRow) mapFields() map[string]any {
	return map[string]any{
		columnChatID:    r.ID,
		columnCreatedAt: r.CreatedAt,
	}
}

func (r *chatRow) Values(columns ...string) []any {
	fields := r.mapFields()
	values := make([]any, 0, len(columns))
	for _, c := range columns {
		if v, ok := fields[c]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}

func (r *messageRow) mapFields() map[string]any {
	return map[string]any{
		columnMessageID:     r.ID,
		columnMessageChatID: r.ChatID,
		columnSenderID:      r.SenderID,
		columnText:          r.Text,
		columnCreatedAt:     r.CreatedAt,
	}
}

func (r *messageRow) Values(columns ...string) []any {
	fields := r.mapFields()
	values := make([]any, 0, len(columns))
	for _, c := range columns {
		if v, ok := fields[c]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}
