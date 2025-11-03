package chat

import (
	"context"
	"fmt"
	"strings"

	"chat/internal/app/models"
	"chat/internal/pkg/postgres"
)

func (r *Repository) CreateMessage(ctx context.Context, msg *models.Message) (*models.Message, error) {
	const api = "chat.Repository.CreateMessage"

	row := fromMessageModel(msg)

	query := r.qb.
		Insert(tableMessages).
		Columns(tableMessagesColumns...).
		Values(row.Values(tableMessagesColumns...)...).
		Suffix("RETURNING " + strings.Join(tableMessagesColumns, ","))

	var outRow messageRow
	if err := r.db.GetQueryEngine(ctx).Getx(ctx, &outRow, query); err != nil {
		return nil, fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return toMessageModel(&outRow), nil
}
