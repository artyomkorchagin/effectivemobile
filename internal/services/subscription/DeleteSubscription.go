package servicesubscription

import (
	"context"
)

func (s *Service) DeleteSubscription(ctx context.Context, subscriptionID uint64) error {
	return s.repo.DeleteSubscription(ctx, subscriptionID)
}
