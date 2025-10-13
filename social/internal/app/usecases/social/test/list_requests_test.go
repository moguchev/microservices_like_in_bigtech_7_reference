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
	"github.com/stretchr/testify/require"
)

func TestSocialService_ListRequests(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := types.UserID(uuid.NewString())
	otherUserID := types.UserID(uuid.NewString())
	dbErr := errors.New("db error")

	sampleRequests := []*models.FriendRequest{
		{ID: types.RequestID(uuid.NewString()), FromUser: otherUserID, ToUser: userID, Status: models.FriendRequestStatusPending},
	}

	tests := []struct {
		name    string
		mock    func(t *testing.T) social.Deps
		argID   types.UserID
		want    []*models.FriendRequest
		wantErr error
	}{
		{
			name:    "Positive: list requests successfully",
			argID:   userID,
			want:    sampleRequests,
			wantErr: nil,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				repoMock.EXPECT().
					ListRequests(ctx, userID).
					Return(sampleRequests, nil).
					Once()

				return social.Deps{
					SocialRepository: repoMock,
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: unauthenticated",
			argID:   userID,
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
			name:    "Negative: permission denied",
			argID:   otherUserID,
			want:    nil,
			wantErr: models.ErrPermissionDenied,
			mock: func(t *testing.T) social.Deps {
				providerMock := mocks.NewUserIDProvider(t)
				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()
				return social.Deps{
					SocialRepository: mocks.NewSocialRepository(t),
					UserIDProvider:   providerMock,
				}
			},
		},
		{
			name:    "Negative: repo returns error",
			argID:   userID,
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
					ListRequests(ctx, userID).
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
			got, err := uc.ListRequests(ctx, tt.argID)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr), "expected error %v, got %v", tt.wantErr, err)
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
