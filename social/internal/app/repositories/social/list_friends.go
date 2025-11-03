package social

import (
	"context"
	"fmt"
	"time"

	"social/internal/app/models/types"
	social_models "social/internal/app/usecases/social/models"
	"social/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) ListFriends(ctx context.Context, id types.UserID, opts ...social_models.ListFriendsOption) (*social_models.ListFriendsResult, error) {
	const api = "social.Repository.ListFriends"

	o := social_models.CollectListFriendsOptions(opts...)

	qb := r.qb.
		Select(tableFriendsColumns...).
		From(tableFriends).
		Where(
			squirrel.And{
				squirrel.Eq{columnUserID: id},
				squirrel.Lt{columnCreatedAt: o.LastAcceptRequestTime},
			},
		).
		OrderBy(columnCreatedAt + " DESC").
		Limit(uint64(o.Limit))

	var rows []friendRow
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, qb); err != nil {
		return nil, fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	var (
		userIDs  = make([]types.UserID, 0, len(rows))
		nextTime time.Time
	)

	for i := range rows {
		userIDs = append(userIDs, types.UserID(rows[i].FriendUserID))
	}

	if len(rows) > 0 {
		nextTime = rows[len(rows)-1].CreatedAt
	}

	return &social_models.ListFriendsResult{
		UserIDs:                   userIDs,
		NextLastAcceptRequestTime: nextTime,
	}, nil
}
