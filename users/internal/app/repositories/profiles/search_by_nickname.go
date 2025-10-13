package profiles

import (
	"context"
	"errors"

	"users/internal/app/models"
)

func (r *Repository) SearchByNickname(ctx context.Context, query string, limit uint32) ([]*models.UserProfile, error) {
	return nil, errors.New("unimplemented")
}
