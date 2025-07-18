package servicesubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

type Reader interface {
	GetAllSubscriptions(ctx context.Context) ([]*types.Subscription, error)
	GetSubscription(ctx context.Context, subscriptionID uint64) (*types.Subscription, error)
	GetSumOfSubscriptions(ctx context.Context, filter *types.Filter) (uint, error)
}

type Writer interface {
	CreateSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error
	DeleteSubscription(ctx context.Context, subscriptionID uint64) error
	UpdateSubscription(ctx context.Context, sub *types.Subscription) error
}

type ReadWriter interface {
	Reader
	Writer
}
