package psqlsubscription

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) CreateSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error {
	startDate, err := helpers.ParseTime(scr.StartDate)
	if err != nil {
		return types.ErrBadRequest(err)
	}

	var endDate sql.NullTime

	if scr.EndDate != "" {
		parsedEndDate, err := helpers.ParseTime(scr.EndDate)
		if err != nil {
			return types.ErrBadRequest(err)
		}

		if parsedEndDate.Before(startDate) {
			return types.ErrBadRequest(fmt.Errorf("end date must be after start date"))
		}

		endDate.Valid = true
		endDate.Time = parsedEndDate
	}

	sql, args, err := goqu.Insert("subscriptions").
		Cols("service_name", "price", "user_id", "start_date", "end_date").
		Vals(
			goqu.Vals{scr.ServiceName, scr.Price, scr.UserUUID, startDate, endDate},
		).ToSQL()

	if err != nil {
		return types.ErrInternalServerError(err)
	}

	res, err := r.db.ExecContext(ctx, sql, args)

	if err != nil {
		return types.ErrInternalServerError(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return types.ErrInternalServerError(err)
	}

	if rows == 0 {
		return types.ErrBadRequest(err)
	}

	return nil
}
