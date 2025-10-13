package profiles

import (
	"context"
	"errors"

	"users/internal/app/models"
	users_models "users/internal/app/usecases/users/models"
)

func (r *Repository) GetProfiles(ctx context.Context, selector users_models.ProfileSelector) ([]*models.UserProfile, error) {
	return nil, errors.New("unimplemented")
}
