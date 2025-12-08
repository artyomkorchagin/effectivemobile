package psqlsubscription

import (
	"context"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) DeleteSubscription(ctx context.Context, subscriptionID uint64) (int64, error) {

	sql, args, err := goqu.From("subscriptions").Where(goqu.C("id").Eq(subscriptionID)).ToSQL()
	if err != nil {
		return 0, types.ErrInternalServerError(err)
	}
	res, err := r.db.ExecContext(ctx, sql, args...)

	if err != nil {
		return 0, types.ErrInternalServerError(fmt.Errorf("failed to delete subscription: %v", err))
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, types.ErrInternalServerError(fmt.Errorf("failed to get affected rows: %v", err))
	}
	return rows, nil
}
