package users

import (
	"context"
	"fmt"

	lib_types "lib/types"
	"users/internal/app/models"
	"users/internal/app/models/types"
	users_models "users/internal/app/usecases/users/models"
)

func (uc *usecase) GetProfileByID(ctx context.Context, id types.UserID) (*models.UserProfile, error) {
	const api = "users.usecase.GetProfileByID"

	profiles, err := uc.ProfilesRepository.GetProfiles(ctx, users_models.ProfileSelector{
		SelectorType: users_models.SelectorTypeAND,
		ID:           lib_types.WrapField(id),
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if len(profiles) == 0 {
		return nil, fmt.Errorf("%s: profile id '%s': %w", api, id, models.ErrNotFound)
	}

	return profiles[0], nil
}
