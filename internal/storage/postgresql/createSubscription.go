package psqlsubscription

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
)

func (r *Repository) CreateSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error {
	startDate, err := helpers.ParseTime(scr.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start date: %w", err)
	}

	var endDate sql.NullTime

	if scr.EndDate != "" {
		parsedEndDate, err := helpers.ParseTime(scr.EndDate)
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
        INSERT INTO subscriptions (service_name, price, user_uuid, start_date, end_date)
        VALUES ($1, $2, $3, $4, $5)`

	_, err = r.db.ExecContext(ctx, query,
		scr.ServiceName,
		scr.Price,
		scr.UserUUID,
		startDate,
		endDate,
	)

	if err != nil {
		return fmt.Errorf("failed to insert subscription: %w", err)
	}

	return nil
}
