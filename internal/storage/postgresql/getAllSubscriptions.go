package psqlsubscription

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

func (r *Repository) GetAllSubscriptions(ctx context.Context) ([]*types.Subscription, error) {

	var (
		subscriptions []*types.Subscription
		endDateNull   sql.NullTime
		startDateTime time.Time
	)

	query := `
        SELECT id, service_name, price, user_uuid, start_date, end_date
        FROM subscriptions`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query subscriptions: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sub types.Subscription

		err := rows.Scan(
			&sub.ID,
			&sub.ServiceName,
			&sub.Price,
			&sub.UserUUID,
			&startDateTime,
			&endDateNull,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan subscription row: %w", err)
		}

		if endDateNull.Valid {
			sub.EndDate = endDateNull.Time.Format("01-2006")
		} else {
			sub.EndDate = ""
		}

		sub.StartDate = startDateTime.Format("01-2006")

		subscriptions = append(subscriptions, &sub)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return subscriptions, nil
}
