package servicesubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (s *Service) GetSumOfSubscriptions(ctx context.Context, filter *types.Filter) (uint, error) {
	return s.repo.GetSumOfSubscriptions(ctx, filter)
}
