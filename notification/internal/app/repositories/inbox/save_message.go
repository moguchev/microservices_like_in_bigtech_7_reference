package inbox

import (
	"context"
	"fmt"

	"notification/internal/app/modules/inbox"
	"notification/internal/pkg/postgres"
)

func (r *Repository) SaveMessage(ctx context.Context, msg *inbox.Message) error {
	const api = "inbox.Repository.SaveMessage"

	row := fromModel(msg)

	qb := r.qb.Insert(tableInboxMessages).
		Columns(tableInboxMessagesColumns...).
		Values(row.Values(tableInboxMessagesColumns...)...)

	if _, err := r.db.GetQueryEngine(ctx).Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return nil
}
