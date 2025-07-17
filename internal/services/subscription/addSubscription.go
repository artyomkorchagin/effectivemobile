package subscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (s *Service) AddSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error {
	return s.repo.AddSubscription(ctx, scr)
}
