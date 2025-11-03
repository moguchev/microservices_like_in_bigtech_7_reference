package models

import (
	"time"

	lib_types "lib/types"
	"social/internal/app/models"
	"social/internal/app/models/types"
)

type ListFriendsRequest struct {
	UserID                types.UserID
	Limit                 uint32
	LastAcceptRequestTime time.Time
}

type ListFriendsResult struct {
	UserIDs                   []types.UserID
	NextLastAcceptRequestTime time.Time
}

type listFriendsOptions struct {
	Limit                 uint32
	LastAcceptRequestTime time.Time
}

type ListFriendsOption func(o *listFriendsOptions)

func WithGetMessagesLimit(n uint32) ListFriendsOption {
	return func(o *listFriendsOptions) { o.Limit = n }
}

func WithGetMessagesLastMessageTime(t time.Time) ListFriendsOption {
	return func(o *listFriendsOptions) { o.LastAcceptRequestTime = t }
}

const FriendsLimit = 100

func CollectListFriendsOptions(opts ...ListFriendsOption) listFriendsOptions {
	res := listFriendsOptions{
		Limit:                 FriendsLimit,
		LastAcceptRequestTime: time.Now().UTC(),
	}

	for _, opt := range opts {
		opt(&res)
	}
	return res
}

type (
	// FriendRequestSelector структура для подбора и поиска заявок в друзья
	FriendRequestSelector struct {
		// Тип поиска (OR/AND)
		SelectorType

		RequestID lib_types.Field[types.RequestID]
		FromUser  lib_types.Field[types.UserID]
		ToUser    lib_types.Field[types.UserID]
		Status    lib_types.Field[models.FriendRequestStatus]
	}
)

// SelectorType тип селектора
type SelectorType int

const (
	SelectorTypeOR SelectorType = iota
	SelectorTypeAND
)
