package psqlsubscription

import (
	"context"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

func (r *Repository) UpdateSubscription(ctx context.Context, sur *types.SubscriptionUpdateRequest) error {
	update := goqu.Update("subscriptions").Where(goqu.C("id").Eq(sur.ID))

	setValues := make(goqu.Record)

	if sur.ServiceName != "" {
		setValues["service_name"] = sur.ServiceName
	}

	if sur.Price != 0 {
		setValues["price"] = sur.Price
	}

	if sur.UserUUID != uuid.Nil {
		setValues["user_id"] = sur.UserUUID
	}

	if sur.StartDate != nil {
		setValues["start_date"] = sur.StartDate
	}

	if sur.EndDate != nil {
		setValues["end_date"] = sur.EndDate
	}

	if len(setValues) == 0 {
		return types.ErrBadRequest(fmt.Errorf("nothing to update"))
	}

	update = update.Set(setValues)

	sqlStr, args, err := update.ToSQL()
	if err != nil {
		return types.ErrInternalServerError(fmt.Errorf("failed to build update query: %w", err))
	}

	result, err := r.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		return types.ErrInternalServerError(fmt.Errorf("failed to update subscription: %w", err))
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return types.ErrNotFound(fmt.Errorf("subscription not found with id %d", sur.ID))
	}

	return nil
}
