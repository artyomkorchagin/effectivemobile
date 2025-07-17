package servicesubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (s *Service) UpdateSubscription(ctx context.Context, sub *types.Subscription) error {
	return s.repo.UpdateSubscription(ctx, sub)
}
