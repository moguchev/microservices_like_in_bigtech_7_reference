package types

// UserID - id пользователя
type UserID string

func (id UserID) String() string {
	return string(id)
}

// ChatID - id чата
type ChatID string

func (id ChatID) String() string {
	return string(id)
}

// MessageID - id сообщения
type MessageID string

func (id MessageID) String() string {
	return string(id)
}
