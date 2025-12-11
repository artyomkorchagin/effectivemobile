package servicesubscription

import (
	"context"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Service struct {
	repo     ReadWriter
	logger   *zap.Logger // по хорошему нужно сделать кастомный интерфейс
	validate *validator.Validate
}

func NewService(repo ReadWriter, logger *zap.Logger, validate *validator.Validate) *Service {
	return &Service{
		repo:     repo,
		logger:   logger,
		validate: validate,
	}
}

func (s *Service) CreateSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error {

	if err := s.validate.Struct(scr); err != nil {
		return types.ErrBadRequest(err)
	}

	if scr.EndDate != nil {
		if scr.EndDate.Before(scr.StartDate) {
			return types.ErrBadRequest(fmt.Errorf("end date must be after start date"))
		}
	}

	return s.repo.CreateSubscription(ctx, scr)
}

func (s *Service) DeleteSubscription(ctx context.Context, subscriptionID uint64) error {
	if err := s.repo.DeleteSubscription(ctx, subscriptionID); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAllSubscriptions(ctx context.Context) ([]*types.Subscription, error) {
	return s.repo.GetAllSubscriptions(ctx)
}

func (s *Service) GetSubscription(ctx context.Context, subscriptionID uint64) (*types.Subscription, error) {
	return s.repo.GetSubscription(ctx, subscriptionID)
}

func (s *Service) GetSumOfSubscriptions(ctx context.Context, filter *types.Filter) (uint, error) {

	if err := s.validate.Struct(filter); err != nil {
		return 0, types.ErrBadRequest(err)
	}

	return s.repo.GetSumOfSubscriptions(ctx, filter)
}

func (s *Service) UpdateSubscription(ctx context.Context, sur *types.SubscriptionUpdateRequest) error {

	if err := s.validate.Struct(sur); err != nil {
		return types.ErrBadRequest(err)
	}

	if sur.EndDate != nil && sur.StartDate != nil {
		if sur.EndDate.Before(*sur.StartDate) {
			return types.ErrBadRequest(fmt.Errorf("end date must be after start date"))
		}
	}
	return s.repo.UpdateSubscription(ctx, sur)
}
