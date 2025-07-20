package servicesubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

type Service struct {
	repo ReadWriter
}

func NewService(repo ReadWriter) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error {
	return s.repo.CreateSubscription(ctx, scr)
}

func (s *Service) DeleteSubscription(ctx context.Context, subscriptionID uint64) error {
	return s.repo.DeleteSubscription(ctx, subscriptionID)
}

func (s *Service) GetAllSubscriptions(ctx context.Context) ([]*types.Subscription, error) {
	return s.repo.GetAllSubscriptions(ctx)
}

func (s *Service) GetSubscription(ctx context.Context, subscriptionID uint64) (*types.Subscription, error) {
	return s.repo.GetSubscription(ctx, subscriptionID)
}

func (s *Service) GetSumOfSubscriptions(ctx context.Context, filter *types.Filter) (uint, error) {
	return s.repo.GetSumOfSubscriptions(ctx, filter)
}

func (s *Service) UpdateSubscription(ctx context.Context, sur *types.SubscriptionUpdateRequest) error {
	return s.repo.UpdateSubscription(ctx, sur)
}
