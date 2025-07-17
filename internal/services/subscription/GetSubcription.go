package servicesubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (s *Service) GetSubscription(ctx context.Context, subscriptionID uint64) (*types.Subscription, error) {
	return s.repo.GetSubscription(ctx, subscriptionID)
}
