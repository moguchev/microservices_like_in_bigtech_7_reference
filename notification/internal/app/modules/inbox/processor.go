package inbox

import (
	"context"
)

type (
	// Repository - репозиторий inbox
	Repository interface {
		SaveMessage(ctx context.Context, msg *Message) error
		SearchMessages(ctx context.Context, opts ...SearchMessageOption) []*Message
		UpdateMessages(ctx context.Context, opts ...UpdateMessageOption) error
	}

	TransactionManager interface {
		RunRepeatableRead(ctx context.Context, f func(tctx context.Context) error) error
	}
)

type Deps struct {
	Repository Repository
}

type Processor struct {
	Deps
}

func NewProcessor(d Deps) *Processor {
	return &Processor{
		Deps: d,
	}
}
