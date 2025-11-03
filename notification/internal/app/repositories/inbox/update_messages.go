package inbox

import (
	"context"
	"fmt"

	"notification/internal/app/modules/inbox"
	"notification/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

// UpdateMessages обновляет сообщения в inbox по заданным опциям.
func (r *Repository) UpdateMessages(ctx context.Context, opts ...inbox.UpdateMessageOption) error {
	const api = "inbox.Repository.UpdateMessages"

	o := inbox.CollectUpdateMessageOptions(opts...)

	// защита от noop
	if o.IncAttemptsBy == 0 && len(o.SetLastErrorsByID) == 0 && o.SetProcessedAt.IsZero() && o.SetStatus == "" {
		return nil
	}

	qb := r.qb.Update(tableInboxMessages)

	if len(o.IDs) > 0 {
		qb = qb.Where(squirrel.Eq{columnInboxID: o.IDs})
	}

	if o.SetStatus != inbox.StatusUnspecified {
		qb = qb.Set(columnInboxStatus, o.SetStatus)
	}
	if o.IncAttemptsBy > 0 {
		qb = qb.Set(columnInboxAttempts, squirrel.Expr(columnInboxAttempts+" + ?", o.IncAttemptsBy))
	}
	if !o.SetProcessedAt.IsZero() {
		qb = qb.Set(columnInboxProcessedAt, o.SetProcessedAt)
	}
	if len(o.SetLastErrorsByID) > 0 {
		caseExpr := "CASE " + columnInboxID + " "
		for id, err := range o.SetLastErrorsByID {
			caseExpr += fmt.Sprintf("WHEN '%s' THEN '%s' ", id, err.Error())
		}
		caseExpr += "END"
		qb = qb.Set(columnInboxLastError, squirrel.Expr(caseExpr))
	}

	if _, err := r.db.GetQueryEngine(ctx).Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return nil
}
