package psqlsubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) CreateSubscription(ctx context.Context, scr *types.SubscriptionCreateRequest) error {
	sql, args, err := goqu.Insert("subscriptions").
		Cols("service_name", "price", "user_id", "start_date", "end_date").
		Vals(
			goqu.Vals{scr.ServiceName, scr.Price, scr.UserUUID, scr.StartDate, scr.EndDate},
		).ToSQL()

	if err != nil {
		return types.ErrInternalServerError(err)
	}

	res, err := r.db.ExecContext(ctx, sql, args...)

	if err != nil {
		return types.ErrInternalServerError(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return types.ErrInternalServerError(err)
	}

	if rows == 0 {
		return types.ErrConflict(err)
	}

	return nil
}
