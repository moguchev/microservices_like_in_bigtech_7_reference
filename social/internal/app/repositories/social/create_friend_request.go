package social

import (
	"context"
	"fmt"
	"strings"

	"social/internal/app/models"
	"social/internal/pkg/postgres"
)

func (r *Repository) CreateFriendRequest(ctx context.Context, request *models.FriendRequest) (*models.FriendRequest, error) {
	const api = "social.Repository.CreateFriendRequest"

	row := fromModel(request)

	query := r.qb.Insert(tableFriendRequests).
		Columns(tableFriendRequestsColumns...).
		Values(row.Values(tableFriendRequestsColumns...)...).
		Suffix("RETURNING " + strings.Join(tableFriendRequestsColumns, ","))

	var outRow friendRequestRow
	if err := r.db.GetQueryEngine(ctx).Getx(ctx, &outRow, query); err != nil {
		return nil, fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return toModel(&outRow), nil
}
