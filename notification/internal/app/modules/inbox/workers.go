package inbox

import (
	"context"
	"errors"
	"log"
	"time"
)

type WorkerOption func(*Inbox)

func WithBatchSize(n int) WorkerOption {
	return func(w *Inbox) { w.batchSize = n }
}

func WithPollInterval(d time.Duration) WorkerOption {
	return func(w *Inbox) { w.pollInterval = d }
}

func WithMaxAttempts(n int) WorkerOption {
	return func(w *Inbox) { w.maxAttempts = n }
}

type Inbox struct {
	batchSize    int
	maxAttempts  int
	pollInterval time.Duration
}

func NewInbox(opts ...WorkerOption) Inbox {
	w := Inbox{
		batchSize:    100,
		maxAttempts:  10,
		pollInterval: 10 * time.Second,
	}
	for _, opt := range opts {
		opt(&w)
	}
	return w
}

type (
	// InboxMessagesHandler - обработчик событий по заявкам в друзья
	InboxMessagesHandler interface {
		// HandleBatch Возвращает списки успешных и проваленных id с их ошибками; err — для фатальных ошибок батча.
		HandleBatch(ctx context.Context, messages []*Message) (succeeded []string, failed map[string]error, err error)
	}
)

// InboxWorker — обработка inbox-событий именно по заявкам в друзья.
type InboxWorker struct {
	Inbox

	repo    Repository
	tm      TransactionManager
	handler InboxMessagesHandler
}

// NewInboxWorker конструктор с дефолтами.
func NewInboxWorker(
	repo Repository,
	tm TransactionManager,
	h InboxMessagesHandler,
	opts ...WorkerOption,
) *InboxWorker {
	w := &InboxWorker{
		Inbox:   NewInbox(opts...),
		repo:    repo,
		tm:      tm,
		handler: h,
	}

	return w
}

// Run — запускает бесконечный цикл обработки до отмены ctx.
// Селектит batch с FOR UPDATE SKIP LOCKED, обрабатывает, коммитит.
func (w *InboxWorker) Run(ctx context.Context) error {
	log.Println("InboxWorker started")

	t := time.NewTicker(w.pollInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			log.Println("InboxWorker tick")

			// Один "тик" — одна транзакция
			if err := w.tm.RunRepeatableRead(ctx, w.Fetch); err != nil {
				log.Printf("outbox: error: %v\n", err)
			}
		}
	}
}

// Fetch обработка событий
func (w *InboxWorker) Fetch(ctx context.Context) error {
	log.Println("InboxWorker.Fetch start")
	defer log.Println("InboxWorker.Fetch end")

	messages := w.repo.SearchMessages(
		ctx,
		WithStatus(StatusReceived, StatusFailed),
		WithMaxRetryCount(w.maxAttempts),
		WithLimit(w.batchSize),
		WithLock(), // FOR UPDATE
	)
	if len(messages) == 0 {
		log.Println("inbox no messages")
		return nil
	}

	if err := w.repo.UpdateMessages(
		ctx,
		WithUpdateIDs(extractIDsFromMessages(messages)...),
		SetStatus(StatusProcessing),
		IncAttempts(1),
	); err != nil {
		log.Printf("inbox UpdateMessages error: %v", err)
		return err
	}

	succeeded, failed, err := w.handler.HandleBatch(ctx, messages)
	if err != nil {
		log.Printf("inbox batch handle error: %v", err)
		return err
	}

	if len(succeeded) > 0 {
		e := w.repo.UpdateMessages(
			ctx,
			WithUpdateIDs(succeeded...),
			SetStatus(StatusProcessed),
			SetProcessedAt(time.Now().UTC()),
		)
		if e != nil {
			err = errors.Join(err, e)
		}
	}

	if len(failed) > 0 {
		e := w.repo.UpdateMessages(
			ctx,
			WithUpdateIDs(extractIDsFromFailed(failed)...),
			SetStatus(StatusFailed),
			SetLastErrorsByID(failed),
		)
		if e != nil {
			err = errors.Join(err, e)
		}
	}

	return err
}

func extractIDsFromMessages(messages []*Message) []string {
	if len(messages) == 0 {
		return nil
	}

	ids := make([]string, 0, len(messages))
	for _, msg := range messages {
		if msg != nil {
			ids = append(ids, msg.ID)
		}
	}
	return ids
}

func extractIDsFromFailed(failed map[string]error) []string {
	if len(failed) == 0 {
		return nil
	}

	ids := make([]string, 0, len(failed))
	for id := range failed {
		ids = append(ids, id)
	}
	return ids
}
