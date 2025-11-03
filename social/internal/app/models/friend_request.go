package models

import (
	"time"

	"social/internal/app/models/types"

	"github.com/google/uuid"
)

// FriendRequestStatus - статус запроса в друзья
type FriendRequestStatus string

const (
	FriendRequestStatusUnspecified FriendRequestStatus = "UNSPECIFIED"
	FriendRequestStatusPending     FriendRequestStatus = "PENDING"
	FriendRequestStatusAccepted    FriendRequestStatus = "ACCEPTED"
	FriendRequestStatusDeclined    FriendRequestStatus = "DECLINED"
)

func (s FriendRequestStatus) String() string {
	return string(s)
}

// FriendRequest - запрос в друзья
type FriendRequest struct {
	ID        types.RequestID     `json:"id"`
	FromUser  types.UserID        `json:"from_user"`
	ToUser    types.UserID        `json:"to_user"`
	Status    FriendRequestStatus `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

func NewFriendRequest() *FriendRequest {
	return &FriendRequest{
		ID:        types.RequestID(uuid.New().String()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
