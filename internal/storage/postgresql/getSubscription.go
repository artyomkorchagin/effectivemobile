package psqlsubscription

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

func (r *Repository) GetSubscription(ctx context.Context, subscriptionID uint64) (*types.Subscription, error) {
	var (
		serviceName string
		price       uint
		userUUID    uuid.UUID
		startDate   time.Time
		endDate     *time.Time
	)

	query, args, err := goqu.Select("service_name", "price", "user_id", "start_date", "end_date").
		From("subscriptions").
		Where(goqu.C("id").Eq(subscriptionID)).
		ToSQL()

	if err != nil {
		return nil, types.ErrInternalServerError(err)
	}

	err = r.db.QueryRowContext(ctx, query, args).Scan(
		&serviceName,
		&price,
		&userUUID,
		&startDate,
		&endDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.ErrNotFound(fmt.Errorf("subscription not found: %w", err))
		}
		return nil, types.ErrInternalServerError(err)
	}

	sub := &types.Subscription{
		ID:          subscriptionID,
		ServiceName: serviceName,
		Price:       price,
		UserUUID:    userUUID,
		StartDate:   startDate.Format("01-2006"),
		EndDate:     "",
	}

	if endDate != nil {
		sub.EndDate = endDate.Format("01-2006")
	}

	return sub, nil
}
