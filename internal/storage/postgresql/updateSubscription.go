package psqlsubscription

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
)

func (r *Repository) UpdateSubscription(ctx context.Context, sur *types.SubscriptionUpdateRequest) error {
	var (
		serviceName sql.NullString
		price       sql.NullInt64
		userUUID    sql.NullString
		startDate   sql.NullTime
		endDate     sql.NullTime
	)

	if sur.ServiceName != "" {
		serviceName.Valid = true
		serviceName.String = sur.ServiceName
	}

	if sur.Price != 0 {
		price.Valid = true
		price.Int64 = int64(sur.Price)
	}

	if sur.UserUUID != "" {
		userUUID.Valid = true
		userUUID.String = sur.UserUUID
	}

	var start *time.Time
	if sur.StartDate != "" {
		t, err := helpers.ParseTime(sur.StartDate)
		if err != nil {
			return fmt.Errorf("invalid start date: %v", err)
		}
		start = &t
		startDate.Valid = true
		startDate.Time = t
	}

	if sur.EndDate != "" {
		t, err := helpers.ParseTime(sur.EndDate)
		if err != nil {
			return fmt.Errorf("invalid end date: %v", err)
		}

		if start != nil && t.Before(*start) {
			return fmt.Errorf("end date must be after start date")
		}

		endDate.Valid = true
		endDate.Time = t
	}

	query := `
        UPDATE subscriptions
        SET
            service_name = COALESCE(NULLIF($1, ''), service_name),
            price = COALESCE($2, price),
            user_id = COALESCE(NULLIF($3, ''), user_id),
            start_date = COALESCE($4, start_date),
            end_date = COALESCE($5, end_date)
        WHERE id = $6`

	_, err := r.db.ExecContext(ctx, query,
		serviceName,
		price,
		userUUID,
		startDate,
		endDate,
		sur.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update subscription: %v", err)
	}

	return nil
}
