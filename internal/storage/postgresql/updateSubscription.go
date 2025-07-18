package psqlsubscription

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
)

func (r *Repository) UpdateSubscription(ctx context.Context, sub *types.Subscription) error {
	startDate, err := helpers.ParseTime(sub.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start date: %w", err)
	}

	var endDate sql.NullTime

	if sub.EndDate != "" {
		parsedEndDate, err := helpers.ParseTime(sub.EndDate)
		if err != nil {
			return fmt.Errorf("invalid end date: %v", err)
		}

		if parsedEndDate.Before(startDate) {
			return fmt.Errorf("end date must be after start date")
		}

		endDate.Valid = true
		endDate.Time = parsedEndDate
	}

	const query = `
        UPDATE subscriptions
        SET service_name = $1, price = $2, user_id = $3, start_date = $4, end_date = $5
        WHERE id = $6`

	_, err = r.db.ExecContext(ctx, query,
		sub.ServiceName,
		sub.Price,
		sub.UserUUID,
		startDate,
		endDate,
		sub.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update subscription: %v", err)
	}

	return nil
}
