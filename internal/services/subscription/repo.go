package subscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

type Reader interface {
	GetAllSubscriptions(ctx context.Context) ([]types.Subscription, error)
}

type Writer interface {
	AddSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error
}

type ReadWriter interface {
	Reader
	Writer
}
