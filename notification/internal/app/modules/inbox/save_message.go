package inbox

import (
	"context"
)

func (p *Processor) SaveMessage(ctx context.Context, msg *Message) error {
	return p.Repository.SaveMessage(ctx, msg)
}
