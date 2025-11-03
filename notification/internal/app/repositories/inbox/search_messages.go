package inbox

import (
	"context"
	"fmt"
	"log"

	"notification/internal/app/modules/inbox"

	"github.com/Masterminds/squirrel"
)

// SearchMessages возвращает сообщения по фильтрам.
// При ошибке возвращает пустой срез.
func (r *Repository) SearchMessages(ctx context.Context, opts ...inbox.SearchMessageOption) []*inbox.Message {
	const api = "inbox.Repository.SearchMessages"

	o := inbox.CollectSearchMessageOptions(opts...)

	qb := r.qb.
		Select(tableInboxMessagesColumns...).
		From(tableInboxMessages).
		OrderBy(columnInboxReceivedAt).
		Limit(uint64(o.Limit))

	if len(o.Status) > 0 {
		qb = qb.Where(squirrel.Eq{columnInboxStatus: o.Status})
	}

	if o.MaxRetryCount > 0 {
		qb = qb.Where(squirrel.Lt{columnInboxAttempts: o.MaxRetryCount})
	}

	if o.WithLock {
		qb = qb.Suffix("FOR UPDATE SKIP LOCKED")
	}

	var rows []inboxMessage
	if err := r.db.GetQueryEngine(ctx).Selectx(ctx, &rows, qb); err != nil {
		log.Println(fmt.Sprintf("%s: Selectx: %s", api, err))
		return nil
	}

	messages := make([]*inbox.Message, 0, len(rows))
	for i := range rows {
		messages = append(messages, toModel(&rows[i]))
	}

	return messages
}
