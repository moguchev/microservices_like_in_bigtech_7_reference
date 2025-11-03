package profiles

import (
	"context"
	"fmt"

	"users/internal/app/models"
	users_models "users/internal/app/usecases/users/models"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) GetProfiles(ctx context.Context, selector users_models.ProfileSelector) ([]*models.UserProfile, error) {
	const api = "profiles.Repository.GetProfiles"

	q := r.qb.
		Select(tableProfilesColumns...).
		From(tableProfiles)

	q = applyProfilesSelector(q, selector)

	var rows []profileRow
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, q); err != nil {
		return nil, fmt.Errorf("%s: Selectx: %w", api, err)
	}

	profiles := make([]*models.UserProfile, len(rows))
	for i := range rows {
		profiles[i] = toModel(&rows[i])
	}

	return profiles, nil
}

func applyProfilesSelector(sb squirrel.SelectBuilder, selector users_models.ProfileSelector) squirrel.SelectBuilder {
	var conditions []squirrel.Sqlizer

	switch selector.SelectorType {
	case users_models.SelectorTypeOR:
		conditions = applyProfilesSelectorConditions(conditions, selector)
		if len(conditions) > 0 {
			sb = sb.Where(squirrel.Or(conditions))
		}
	case users_models.SelectorTypeAND:
		conditions = applyProfilesSelectorConditions(conditions, selector)
		if len(conditions) > 0 {
			sb = sb.Where(squirrel.And(conditions))
		}
	default:
		return sb
	}

	return sb
}

func applyProfilesSelectorConditions(conds []squirrel.Sqlizer, selector users_models.ProfileSelector) []squirrel.Sqlizer {
	if selector.ID.IsSet() {
		conds = append(conds, squirrel.Eq{
			columnID: selector.ID.Value(),
		})
	}

	if selector.Name.IsSet() {
		conds = append(conds, squirrel.Eq{
			columnNickname: selector.Name.Value(),
		})
	}

	return conds
}
