package servicesubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (s *Service) CreateSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error {
	return s.repo.CreateSubscription(ctx, scr)
}
