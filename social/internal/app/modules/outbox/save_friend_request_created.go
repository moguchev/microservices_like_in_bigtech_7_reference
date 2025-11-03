package outbox

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"social/internal/app/models"
	"social/internal/app/models/types"

	"github.com/google/uuid"
)

func (p *Processor) SaveFriendRequestCreated(ctx context.Context, toUserID types.UserID, req *models.FriendRequest) error {
	const api = "outbox.Processor.SaveFriendRequestCreated"

	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: Marshal: %w", api, err)
	}

	e := Event{
		ID:            uuid.New(),
		AggregateType: AggregateTypeFriendRequest,
		AggregateID:   toUserID.String(),
		EventType:     EventTypeFriendRequestCreated,
		Payload:       payload,
		CreatedAt:     time.Now().UTC(),
	}

	return p.Repository.SaveEvent(ctx, &e)
}
