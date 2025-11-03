package social

import (
	"context"
	"fmt"
	"strings"

	"social/internal/app/models"
	"social/internal/app/models/types"
	"social/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) UpdateFriendRequestStatus(ctx context.Context, id types.RequestID, status models.FriendRequestStatus) (*models.FriendRequest, error) {
	const api = "social.Repository.UpdateFriendRequestStatus"

	qb := r.qb.
		Update(tableFriendRequests).
		Set(columnStatus, status).
		Where(squirrel.Eq{columnID: id}).
		Suffix("RETURNING " + strings.Join(tableFriendRequestsColumns, ","))

	var row friendRequestRow
	if err := r.db.GetQueryEngine(ctx).Getx(ctx, &row, qb); err != nil {
		return nil, fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return toModel(&row), nil
}
