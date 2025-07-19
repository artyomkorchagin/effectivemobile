package servicesubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (s *Service) UpdateSubscription(ctx context.Context, sur *types.SubscriptionUpdateRequest) error {
	return s.repo.UpdateSubscription(ctx, sur)
}
