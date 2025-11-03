package social

import (
	"context"
	"fmt"

	"social/internal/app/models"
	social_models "social/internal/app/usecases/social/models"
	"social/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) GetFriendRequests(ctx context.Context, selector social_models.FriendRequestSelector) ([]*models.FriendRequest, error) {
	const api = "social.Repository.GetFriendRequests"

	qb := r.qb.
		Select(tableFriendRequestsColumns...).
		From(tableFriendRequests)

	qb = applyFriendRequestSelector(qb, selector)

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

func applyFriendRequestSelector(sb squirrel.SelectBuilder, selector social_models.FriendRequestSelector) squirrel.SelectBuilder {
	var sqlizer []squirrel.Sqlizer

	switch selector.SelectorType {
	case social_models.SelectorTypeOR:
		sqlizer = squirrel.Or{}
	case social_models.SelectorTypeAND:
		sqlizer = squirrel.And{}
	default:
		return sb
	}

	if selector.RequestID.IsSet() {
		sqlizer = append(sqlizer, squirrel.Eq{columnID: selector.RequestID.Value()})
	}
	if selector.FromUser.IsSet() {
		sqlizer = append(sqlizer, squirrel.Eq{columnFromUser: selector.FromUser.Value()})
	}
	if selector.ToUser.IsSet() {
		sqlizer = append(sqlizer, squirrel.Eq{columnToUser: selector.ToUser.Value()})
	}
	if selector.Status.IsSet() {
		sqlizer = append(sqlizer, squirrel.Eq{columnStatus: selector.Status.Value()})
	}

	switch selector.SelectorType {
	case social_models.SelectorTypeOR:
		return sb.Where(squirrel.Or(sqlizer))
	case social_models.SelectorTypeAND:
		return sb.Where(squirrel.And(sqlizer))
	default:
		return sb
	}
}
