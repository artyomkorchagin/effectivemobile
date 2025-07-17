package psqlsubscription

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (r *Repository) GetSubscription(ctx context.Context, subscriptionID uint64) (*types.Subscription, error) {
	var (
		subscription  types.Subscription
		endDateNull   sql.NullTime
		startDateTime time.Time
	)
	query := `
        SELECT service_name, price, user_uuid, start_date, end_date
        FROM subscriptions
        WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, subscriptionID).Scan(
		&subscription.ServiceName,
		&subscription.Price,
		&subscription.UserUUID,
		&startDateTime,
		&endDateNull,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("subscription not found: %v", err)
		}
		return nil, fmt.Errorf("failed to get subscription: %v", err)
	}

	if endDateNull.Valid {
		subscription.EndDate = endDateNull.Time.Format("01-2006")
	} else {
		subscription.EndDate = ""
	}

	subscription.StartDate = startDateTime.Format("01-2006")
	subscription.ID = subscriptionID
	return &subscription, nil
}
