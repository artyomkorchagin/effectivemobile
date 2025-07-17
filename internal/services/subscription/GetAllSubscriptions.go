package servicesubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (s *Service) GetAllSubscriptions(ctx context.Context) ([]*types.Subscription, error) {
	return s.repo.GetAllSubscriptions(ctx)
}
