package types

// UserID - id пользователя
type UserID string

func (id UserID) String() string {
	return string(id)
}
