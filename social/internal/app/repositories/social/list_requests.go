package social

import (
	"context"
	"fmt"

	"social/internal/app/models"
	"social/internal/app/models/types"
	"social/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) ListRequests(ctx context.Context, userID types.UserID) ([]*models.FriendRequest, error) {
	const api = "social.Repository.ListRequests"

	qb := r.qb.
		Select(tableFriendRequestsColumns...).
		From(tableFriendRequests).
		Where(
			squirrel.And{
				squirrel.Eq{columnToUser: userID},
				squirrel.Eq{columnStatus: models.FriendRequestStatusPending},
			},
		).
		OrderBy(columnCreatedAt + " DESC")

	var rows []friendRequestRow
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, qb); err != nil {
		return nil, fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	friendRequests := make([]*models.FriendRequest, len(rows))
	for i := range rows {
		friendRequests[i] = toModel(&rows[i])
	}

	return friendRequests, nil
}
