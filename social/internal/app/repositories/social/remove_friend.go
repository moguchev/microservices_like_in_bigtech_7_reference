package social

import (
	"context"
	"fmt"

	"social/internal/app/models/types"
	"social/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) RemoveFriend(ctx context.Context, userID, friendID types.UserID) error {
	const api = "social.Repository.RemoveFriend"

	qb := r.qb.
		Delete(tableFriends).
		Where(
			squirrel.Or{
				squirrel.And{
					squirrel.Eq{columnUserID: userID},
					squirrel.Eq{columnFriendUserID: friendID},
				},
				squirrel.And{
					squirrel.Eq{columnUserID: friendID},
					squirrel.Eq{columnFriendUserID: userID},
				},
			},
		)

	if _, err := r.db.GetQueryEngine(ctx).Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return nil
}
