package social_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"social/internal/app/models/types"
	"social/internal/app/usecases/social"
	"social/internal/app/usecases/social/mocks"
	social_models "social/internal/app/usecases/social/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSocialService_ListFriends(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := types.UserID(uuid.NewString())
	dbErr := errors.New("db error")

	friendList := []types.UserID{
		types.UserID(uuid.NewString()),
		types.UserID(uuid.NewString()),
	}

	tests := []struct {
		name                  string
		mock                  func(t *testing.T) social.Deps
		limit                 uint32
		lastAcceptRequestTime time.Time
		want                  []types.UserID
		wantErr               error
	}{
		{
			name:                  "Positive: list friends with specific limit and time",
			limit:                 2,
			lastAcceptRequestTime: time.Date(2025, 10, 1, 12, 0, 0, 0, time.UTC),
			want:                  friendList,
			wantErr:               nil,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)

				repoMock.EXPECT().
					ListFriends(
						ctx,
						userID,
						mock.AnythingOfType("models.ListFriendsOption"),
						mock.AnythingOfType("models.ListFriendsOption"),
					).
					Return(&social_models.ListFriendsResult{
						UserIDs:                   friendList,
						NextLastAcceptRequestTime: time.Now().UTC(),
					}, nil).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
				}
			},
		},
		{
			name:                  "Positive: limit zero replaced with default, lastAcceptRequestTime zero replaced with now",
			limit:                 0,
			lastAcceptRequestTime: time.Time{},
			want:                  friendList,
			wantErr:               nil,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)

				repoMock.EXPECT().
					ListFriends(
						ctx,
						userID,
						mock.AnythingOfType("models.ListFriendsOption"),
						mock.AnythingOfType("models.ListFriendsOption"),
					).
					Return(&social_models.ListFriendsResult{
						UserIDs:                   friendList,
						NextLastAcceptRequestTime: time.Now().UTC(),
					}, nil).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
				}
			},
		},
		{
			name:                  "Negative: repo returns error",
			limit:                 5,
			lastAcceptRequestTime: time.Now(),
			want:                  nil,
			wantErr:               dbErr,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)

				repoMock.EXPECT().
					ListFriends(
						ctx,
						userID,
						mock.AnythingOfType("models.ListFriendsOption"),
						mock.AnythingOfType("models.ListFriendsOption"),
					).
					Return(nil, dbErr).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := social.NewUsecase(tt.mock(t))

			res, err := uc.ListFriends(ctx, &social_models.ListFriendsRequest{
				UserID:                userID,
				Limit:                 tt.limit,
				LastAcceptRequestTime: tt.lastAcceptRequestTime,
			})

			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr), "expected error %v, got %v", tt.wantErr, err)
				require.Nil(t, res)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, res.UserIDs)
				require.False(t, res.NextLastAcceptRequestTime.IsZero(), "gotTime should not be zero")
			}
		})
	}
}
