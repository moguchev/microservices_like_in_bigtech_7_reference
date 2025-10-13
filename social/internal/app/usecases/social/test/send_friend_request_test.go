package social_test

import (
	"context"
	"errors"
	"testing"

	"social/internal/app/models"
	"social/internal/app/models/types"
	"social/internal/app/usecases/social"
	"social/internal/app/usecases/social/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSocialService_SendFriendRequest(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := types.UserID(uuid.NewString())
	friendID := types.UserID(uuid.NewString())
	anotherUser := types.UserID(uuid.NewString())
	dbErr := errors.New("db error")

	tests := []struct {
		name    string
		mock    func(t *testing.T) social.Deps
		argID   types.UserID
		want    *models.FriendRequest
		wantErr error
	}{
		{
			name:  "Positive: friend request created",
			argID: friendID,
			want: &models.FriendRequest{
				FromUser: userID,
				ToUser:   friendID,
				Status:   models.FriendRequestStatusPending,
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
					CreateFriendRequest(ctx, mock.MatchedBy(func(fr *models.FriendRequest) bool {
						return fr != nil &&
							fr.FromUser == userID &&
							fr.ToUser == friendID &&
							fr.Status == models.FriendRequestStatusPending
					})).
					Return(&models.FriendRequest{}, nil).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: unauthenticated",
			argID:   friendID,
			want:    nil,
			wantErr: models.ErrUnauthenticated,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return("", errors.New("no user")).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: repo error",
			argID:   anotherUser,
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
					CreateFriendRequest(ctx, mock.Anything).
					Return(nil, dbErr).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Self-request: user sends request to self",
			argID:   userID,
			want:    nil,
			wantErr: models.ErrInvalidArgument,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				// Repo не вызывается
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

			got, err := uc.SendFriendRequest(ctx, tt.argID)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr), "expected error %v, got %v", tt.wantErr, err)
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				require.Equal(t, tt.want.FromUser, got.FromUser)
				require.Equal(t, tt.want.ToUser, got.ToUser)
				require.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}
