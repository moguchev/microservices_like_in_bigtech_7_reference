package profiles

import (
	"context"
	"fmt"
	"strings"

	"users/internal/app/models"
	"users/internal/app/models/types"
	users_models "users/internal/app/usecases/users/models"
	"users/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) UpdateProfile(ctx context.Context, id types.UserID, f users_models.ProfileUpdateFields) (*models.UserProfile, error) {
	const api = "profiles.Repository.UpdateProfile"

	qb := r.qb.
		Update(tableProfiles).
		Where(squirrel.Eq{columnID: id}).
		Suffix("RETURNING " + strings.Join(tableProfilesColumns, ","))

	qb = applyProfileUpdateFields(qb, f)

	var row profileRow
	if err := r.db.GetQueryEngine(ctx).Getx(ctx, &row, qb); err != nil {
		return nil, fmt.Errorf("%s: Getx: %w", api, postgres.ConvertPGError(err))
	}

	return toModel(&row), nil
}

func applyProfileUpdateFields(ub squirrel.UpdateBuilder, f users_models.ProfileUpdateFields) squirrel.UpdateBuilder {
	m := make(map[string]interface{})

	if f.Name.IsSet() {
		m[columnNickname] = f.Name.Value()
	}
	if f.Bio.IsSet() {
		m[columnBio] = f.Bio.Value()
	}
	if f.AvatarURL.IsSet() {
		m[columnAvatarURL] = f.AvatarURL.Value()
	}

	if len(m) > 0 {
		ub = ub.SetMap(m)
	}

	return ub
}
