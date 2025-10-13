package users

import (
	"context"
	"fmt"

	"users/internal/app/models"
)

func (uc *usecase) SearchByNickname(ctx context.Context, query string, limit uint32) ([]*models.UserProfile, error) {
	const api = "users.usecase.SearchByNickname"

	const searchDefaultLimit = 10
	if limit == 0 {
		limit = searchDefaultLimit
	}

	profiles, err := uc.ProfilesRepository.SearchByNickname(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if len(profiles) == 0 {
		return nil, fmt.Errorf("%s: profile with query name '%s': %w", api, query, models.ErrNotFound)
	}

	return profiles, nil
}
