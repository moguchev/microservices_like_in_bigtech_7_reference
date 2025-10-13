package social_test

import (
	"context"
	"errors"
	"testing"

	"social/internal/app/models"
	"social/internal/app/models/types"
	"social/internal/app/usecases/social"
	"social/internal/app/usecases/social/mocks"
	social_models "social/internal/app/usecases/social/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSocialService_AcceptFriendRequest(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := types.UserID(uuid.NewString())
	requestID := types.RequestID(uuid.NewString())
	friendUser := types.UserID(uuid.NewString())
	dbErr := errors.New("db error")

	tests := []struct {
		name    string
		mock    func(t *testing.T) social.Deps
		argID   types.RequestID
		want    *models.FriendRequest
		wantErr error
	}{
		{
			name:  "Positive: accept request successfully",
			argID: requestID,
			want: &models.FriendRequest{
				ID:       requestID,
				FromUser: friendUser,
				ToUser:   userID,
				Status:   models.FriendRequestStatusAccepted,
			},
			wantErr: nil,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				repoMock.EXPECT().
					GetFriendRequests(ctx, mock.MatchedBy(func(sel social_models.FriendRequestSelector) bool {
						return sel.RequestID.IsSet() &&
							sel.RequestID.Value() == requestID &&
							sel.SelectorType == social_models.SelectorTypeAND
					})).
					Return([]*models.FriendRequest{
						{
							ID:       requestID,
							FromUser: friendUser,
							ToUser:   userID,
							Status:   models.FriendRequestStatusPending,
						},
					}, nil).
					Once()

				repoMock.EXPECT().
					UpdateFriendRequestStatus(ctx, requestID, models.FriendRequestStatusAccepted).
					Return(&models.FriendRequest{
						ID:       requestID,
						FromUser: friendUser,
						ToUser:   userID,
						Status:   models.FriendRequestStatusAccepted,
					}, nil).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: unauthenticated",
			argID:   requestID,
			want:    nil,
			wantErr: models.ErrUnauthenticated,
			mock: func(t *testing.T) social.Deps {
				providerMock := mocks.NewUserIDProvider(t)
				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return("", errors.New("no user")).
					Once()
				return social.Deps{
					SocialRepository: mocks.NewSocialRepository(t),
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: request not found",
			argID:   requestID,
			want:    nil,
			wantErr: models.ErrNotFound,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				repoMock.EXPECT().
					GetFriendRequests(ctx, mock.Anything).
					Return([]*models.FriendRequest{}, nil).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: permission denied",
			argID:   requestID,
			want:    nil,
			wantErr: models.ErrPermissionDenied,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)
				otherUser := types.UserID(uuid.NewString())

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				repoMock.EXPECT().
					GetFriendRequests(ctx, mock.Anything).
					Return([]*models.FriendRequest{
						{ID: requestID, FromUser: friendUser, ToUser: otherUser, Status: models.FriendRequestStatusPending},
					}, nil).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: GetFriendRequests returns error",
			argID:   requestID,
			want:    nil,
			wantErr: dbErr,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				repoMock.EXPECT().
					GetFriendRequests(ctx, mock.Anything).
					Return(nil, dbErr).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: UpdateFriendRequestStatus returns error",
			argID:   requestID,
			want:    nil,
			wantErr: dbErr,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				repoMock.EXPECT().
					GetFriendRequests(ctx, mock.Anything).
					Return([]*models.FriendRequest{
						{ID: requestID, FromUser: friendUser, ToUser: userID, Status: models.FriendRequestStatusPending},
					}, nil).
					Once()

				repoMock.EXPECT().
					UpdateFriendRequestStatus(ctx, requestID, models.FriendRequestStatusAccepted).
					Return(nil, dbErr).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := social.NewUsecase(tt.mock(t))
			got, err := uc.AcceptFriendRequest(ctx, tt.argID)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr), "expected error %v, got %v", tt.wantErr, err)
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				require.Equal(t, tt.want.ID, got.ID)
				require.Equal(t, tt.want.FromUser, got.FromUser)
				require.Equal(t, tt.want.ToUser, got.ToUser)
				require.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}
