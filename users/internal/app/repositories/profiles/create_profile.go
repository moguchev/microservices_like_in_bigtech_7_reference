package profiles

import (
	"context"
	"fmt"

	"users/internal/app/models"
	"users/internal/pkg/postgres"
)

func (r *Repository) CreateProfile(ctx context.Context, profile *models.UserProfile) error {
	const api = "profiles.Repository.CreateProfile"

	row := fromModel(profile)

	qb := r.qb.
		Insert(tableProfiles).
		Columns(tableProfilesColumns...).
		Values(row.Values(tableProfilesColumns...)...)

	if _, err := r.db.GetQueryEngine(ctx).Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: Execx: %w", api, postgres.ConvertPGError(err))
	}

	return nil
}
