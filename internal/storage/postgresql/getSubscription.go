package psqlsubscription

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) GetSubscription(ctx context.Context, subscriptionID uint64) (*types.Subscription, error) {

	var sub types.Subscription

	query, args, err := goqu.Select("service_name", "price", "user_id", "start_date", "end_date").
		From("subscriptions").
		Where(goqu.C("id").Eq(subscriptionID)).
		ToSQL()

	if err != nil {
		return nil, types.ErrInternalServerError(err)
	}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(
		&sub.ServiceName,
		&sub.Price,
		&sub.UserUUID,
		&sub.StartDate,
		&sub.EndDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.ErrNotFound(fmt.Errorf("subscription not found: %w", err))
		}
		return nil, types.ErrInternalServerError(err)
	}

	return &sub, nil
}
