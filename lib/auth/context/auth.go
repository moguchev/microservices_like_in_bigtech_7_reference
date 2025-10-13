package context

import (
	"context"
)

type MyUserIDProvider struct{}

func (p MyUserIDProvider) GetUserIDFromIncomingContext(ctx context.Context) (string, error) {
	return "", nil
}
