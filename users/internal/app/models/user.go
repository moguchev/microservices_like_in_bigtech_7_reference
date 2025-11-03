package models

import (
	"users/internal/app/models/types"
)

// UserProfile - профиль пользователя
type UserProfile struct {
	ID        types.UserID
	Email     string
	Name      string
	Bio       string
	AvatarURL string
}
