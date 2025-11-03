package profiles

import (
	"context"
	"fmt"

	"users/internal/app/models"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) SearchByNickname(ctx context.Context, query string, limit uint32) ([]*models.UserProfile, error) {
	const api = "profiles.Repository.SearchByNickname"

	qb := r.qb.
		Select(tableProfilesColumns...).
		From(tableProfiles).
		Where(squirrel.ILike{columnNickname: "%" + query + "%"}).
		OrderBy(columnCreatedAt + " DESC").
		Limit(uint64(limit))

	var rows []profileRow
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, qb); err != nil {
		return nil, fmt.Errorf("%s: Selectx: %w", api, err)
	}

	profiles := make([]*models.UserProfile, len(rows))
	for i := range rows {
		profiles[i] = toModel(&rows[i])
	}

	return profiles, nil
}
