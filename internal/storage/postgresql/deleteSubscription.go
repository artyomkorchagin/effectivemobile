package psqlsubscription

import (
	"context"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) DeleteSubscription(ctx context.Context, subscriptionID int64) error {

	sql, args, err := goqu.Delete("subscriptions").
		Where(goqu.C("id").Eq(subscriptionID)).
		ToSQL()

	if err != nil {
		return types.ErrInternalServerError(err)
	}
	res, err := r.db.ExecContext(ctx, sql, args...)

	if err != nil {
		return types.ErrInternalServerError(err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return types.ErrNotFound(fmt.Errorf("subscription not found"))
	}

	return nil
}
