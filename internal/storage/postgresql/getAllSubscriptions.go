package psqlsubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) GetAllSubscriptions(ctx context.Context) ([]*types.Subscription, error) {

	query, _, err := goqu.Select("id", "service_name", "price", "user_id", "start_date", "end_date").
		From("subscriptions").
		ToSQL()
	if err != nil {
		return nil, types.ErrInternalServerError(err)
	}
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, types.ErrInternalServerError(err)
	}
	defer rows.Close()

	var subscriptions []*types.Subscription

	for rows.Next() {
		var sub types.Subscription

		err := rows.Scan(
			&sub.ID,
			&sub.ServiceName,
			&sub.Price,
			&sub.UserUUID,
			&sub.StartDate,
			&sub.EndDate,
		)
		if err != nil {
			return nil, types.ErrInternalServerError(err)
		}

		subscriptions = append(subscriptions, &sub)
	}

	if err := rows.Err(); err != nil {
		return nil, types.ErrInternalServerError(err)
	}

	return subscriptions, nil
}
