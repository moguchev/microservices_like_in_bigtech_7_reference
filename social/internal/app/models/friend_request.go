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

// FriendRequest - запрос в друзья
type FriendRequest struct {
	ID        types.RequestID
	FromUser  types.UserID
	ToUser    types.UserID
	Status    FriendRequestStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewFriendRequest() *FriendRequest {
	return &FriendRequest{
		ID:        types.RequestID(uuid.New().String()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
