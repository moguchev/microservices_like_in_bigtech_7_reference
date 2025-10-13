package profiles

import (
	"context"
	"errors"

	"users/internal/app/models"
	"users/internal/app/models/types"
	users_models "users/internal/app/usecases/users/models"
)

func (r *Repository) UpdateProfile(ctx context.Context, id types.UserID, f users_models.ProfileUpdateFields) (*models.UserProfile, error) {
	return nil, errors.New("unimplemented")
}
