package social

import (
	"context"
	"fmt"
	"time"

	"social/internal/app/models/types"
	"social/internal/pkg/postgres"
)

func (r *Repository) CreateFriendPair(ctx context.Context, userID, friendID types.UserID) error {
	const api = "social.Repository.CreateFriendPair"

	var (
		pairs = [][2]types.UserID{
			{userID, friendID},
			{friendID, userID},
		}

		now = time.Now().UTC()
	)

	for _, pair := range pairs {
		qb := r.qb.
			Insert(tableFriends).
			Columns(tableFriendsColumns...).
			Values(pair[0], pair[1], now).
			Suffix("ON CONFLICT DO NOTHING")

		if _, err := r.db.GetQueryEngine(ctx).Execx(ctx, qb); err != nil {
			return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
		}
	}

	return nil
}
