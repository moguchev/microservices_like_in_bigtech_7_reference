package types

// UserID - id пользователя
type UserID string

func (id UserID) String() string {
	return string(id)
}

// RequestID - id заявки в друзья
type RequestID string

func (id RequestID) String() string {
	return string(id)
}
