package models

import (
	"time"
)

// FriendRequest - запрос в друзья
type FriendRequest struct {
	ID        string    `json:"id"`
	FromUser  string    `json:"from_user"`
	ToUser    string    `json:"to_user"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
