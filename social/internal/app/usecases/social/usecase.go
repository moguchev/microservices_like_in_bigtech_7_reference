package social

import (
	"context"

	"social/internal/app/models"
	"social/internal/app/models/types"
	social_models "social/internal/app/usecases/social/models"
)

type Usecase interface {
	// SendFriendRequest - отправить заявку в друзья
	SendFriendRequest(ctx context.Context, id types.UserID) (*models.FriendRequest, error)
	// AcceptFriendRequest - принять заявку в друзья
	AcceptFriendRequest(ctx context.Context, id types.RequestID) (*models.FriendRequest, error)
	// DeclineFriendRequest - отклонить заявку в друзья
	DeclineFriendRequest(ctx context.Context, id types.RequestID) (*models.FriendRequest, error)
	// ListRequests - список входящих заявок
	ListRequests(ctx context.Context, id types.UserID) ([]*models.FriendRequest, error)
	// RemoveFriend - удалить пользователя из друзей
	RemoveFriend(ctx context.Context, id types.UserID) error
	// ListFriends - список друзей пользователя
	ListFriends(ctx context.Context, req *social_models.ListFriendsRequest) (*social_models.ListFriendsResult, error)
}

type (
	SocialRepository interface {
		// CreateFriendRequest - создать заявку в друзья
		//
		// errors: models.ErrAlreadyExists
		CreateFriendRequest(ctx context.Context, request *models.FriendRequest) (*models.FriendRequest, error)
		// UpdateFriendRequestStatus - обновить статус заявки в друзья
		//
		// errors: models.ErrNotFound
		UpdateFriendRequestStatus(ctx context.Context, id types.RequestID, status models.FriendRequestStatus) (*models.FriendRequest, error)
		// ListRequests - список входящих заявок
		ListRequests(ctx context.Context, userID types.UserID) ([]*models.FriendRequest, error)
		// GetFriendRequests - получить заявки по селектору
		GetFriendRequests(ctx context.Context, selector social_models.FriendRequestSelector) ([]*models.FriendRequest, error)

		// RemoveFriend - удалить пользователя из друзей
		//
		// errors: models.ErrNotFound
		RemoveFriend(ctx context.Context, userID, friendID types.UserID) error
		// ListFriends - список друзей пользователя
		ListFriends(ctx context.Context, id types.UserID, opts ...social_models.ListFriendsOption) (*social_models.ListFriendsResult, error)
	}

	UserIDProvider interface {
		GetUserIDFromIncomingContext(ctx context.Context) (string, error)
	}
)

type Deps struct {
	SocialRepository SocialRepository
	UserIDProvider
}

var _ Usecase = (*SocialService)(nil)

type SocialService struct {
	Deps
}

func NewUsecase(d Deps) Usecase {
	return &SocialService{Deps: d}
}
