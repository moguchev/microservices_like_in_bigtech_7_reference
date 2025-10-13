package social

import (
	"context"
	"errors"

	"social/internal/app/models/types"
)

func (r *Repository) RemoveFriend(ctx context.Context, userID, friendID types.UserID) error {
	return errors.New("unimplemented")
}
