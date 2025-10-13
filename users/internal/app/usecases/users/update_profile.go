package users

import (
	"context"
	"fmt"

	"users/internal/app/models"
	"users/internal/app/models/types"
	users_models "users/internal/app/usecases/users/models"
)

func (uc *usecase) UpdateProfile(ctx context.Context, id types.UserID, f users_models.ProfileUpdateFields) (*models.UserProfile, error) {
	const api = "users.usecase.SearchByNickname"

	userID, err := uc.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	if userID != id.String() {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	updatedProfile, err := uc.ProfilesRepository.UpdateProfile(ctx, id, f)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return updatedProfile, nil
}
