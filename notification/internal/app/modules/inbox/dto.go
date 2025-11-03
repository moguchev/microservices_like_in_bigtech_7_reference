package inbox

import (
	"time"
)

// ----- Search options DTO -----

type searchMessageOptions struct {
	Status []Status

	MaxRetryCount int
	Limit         int
	WithLock      bool
}

type SearchMessageOption func(o *searchMessageOptions)

// ----- Option builders -----

func WithLimit(n int) SearchMessageOption {
	return func(o *searchMessageOptions) { o.Limit = n }
}

func WithStatus(status ...Status) SearchMessageOption {
	return func(o *searchMessageOptions) { o.Status = status }
}

func WithMaxRetryCount(n int) SearchMessageOption {
	return func(o *searchMessageOptions) { o.MaxRetryCount = n }
}

func WithLock() SearchMessageOption {
	return func(o *searchMessageOptions) { o.WithLock = true }
}

func CollectSearchMessageOptions(opts ...SearchMessageOption) searchMessageOptions {
	res := searchMessageOptions{
		Limit:         10,
		MaxRetryCount: 3,
	}

	for _, opt := range opts {
		opt(&res)
	}
	return res
}

// ----- Update options DTO -----

type updateMessageOptions struct {
	IDs []string

	// что обновляем
	SetStatus         Status
	SetLastErrorsByID map[string]error
	IncAttemptsBy     int
	SetProcessedAt    time.Time
}

type UpdateMessageOption func(*updateMessageOptions)

// ----- Option builders -----

func WithUpdateIDs(ids ...string) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.IDs = append(o.IDs, ids...) }
}

func SetStatus(status Status) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.SetStatus = status }
}

func SetLastErrorsByID(lastErrorsByID map[string]error) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.SetLastErrorsByID = lastErrorsByID }
}

func IncAttempts(by int) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.IncAttemptsBy = by }
}

func SetProcessedAt(processedAt time.Time) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.SetProcessedAt = processedAt }
}

func CollectUpdateMessageOptions(opts ...UpdateMessageOption) updateMessageOptions {
	o := updateMessageOptions{
		SetStatus: StatusUnspecified,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
