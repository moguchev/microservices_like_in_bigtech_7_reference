package users

import (
	"context"
	"fmt"

	lib_types "lib/types"
	"users/internal/app/models"
	users_models "users/internal/app/usecases/users/models"
)

func (uc *usecase) GetProfileByNickname(ctx context.Context, name string) (*models.UserProfile, error) {
	const api = "users.usecase.GetProfileByNickname"

	profiles, err := uc.ProfilesRepository.GetProfiles(ctx, users_models.ProfileSelector{
		SelectorType: users_models.SelectorTypeAND,
		Name:         lib_types.WrapField(name),
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if len(profiles) == 0 {
		return nil, fmt.Errorf("%s: profile with name '%s': %w", api, name, models.ErrNotFound)
	}

	return profiles[0], nil
}
