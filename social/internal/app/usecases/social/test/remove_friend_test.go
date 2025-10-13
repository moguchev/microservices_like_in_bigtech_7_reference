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

func TestSocialService_RemoveFriend(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := types.UserID(uuid.NewString())
	friendID := types.UserID(uuid.NewString())
	dbErr := errors.New("db error")

	tests := []struct {
		name    string
		mock    func(t *testing.T) social.Deps
		argID   types.UserID
		wantErr error
	}{
		{
			name:    "Positive: remove friend successfully",
			argID:   friendID,
			wantErr: nil,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				repoMock.EXPECT().
					RemoveFriend(ctx, userID, friendID).
					Return(nil).
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
			name:    "Negative: repo returns error",
			argID:   friendID,
			wantErr: dbErr,
			mock: func(t *testing.T) social.Deps {
				repoMock := mocks.NewSocialRepository(t)
				providerMock := mocks.NewUserIDProvider(t)

				providerMock.EXPECT().
					GetUserIDFromIncomingContext(ctx).
					Return(userID.String(), nil).
					Once()

				repoMock.EXPECT().
					RemoveFriend(ctx, userID, friendID).
					Return(dbErr).
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
			err := uc.RemoveFriend(ctx, tt.argID)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr), "expected error %v, got %v", tt.wantErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
