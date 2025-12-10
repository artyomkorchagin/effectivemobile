package psqlsubscription

import (
	"context"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) GetSumOfSubscriptions(ctx context.Context, filter *types.Filter) (uint, error) {

	priceSeriesExpr := goqu.L(
		`"price" * (SELECT COUNT(*) FROM generate_series(
			date_trunc('month', GREATEST("start_date", ?)),
			LEAST(COALESCE("end_date", CURRENT_DATE), ?),
			'1 month'
		))`,
		filter.StartDate,
		filter.EndDate,
	)

	sumExpr := goqu.COALESCE(goqu.SUM(priceSeriesExpr), goqu.L("0")).As("total_sum")

	ds := goqu.Dialect("postgres").
		Select(sumExpr).
		From("subscriptions").
		Where(
			goqu.C("start_date").Lte(goqu.COALESCE(goqu.C("end_date"), goqu.L("CURRENT_DATE"))),
		)

	if filter.UserUUID != nil {
		ds = ds.Where(goqu.C("user_id").Eq(filter.UserUUID))
	}
	if filter.ServiceName != nil {
		ds = ds.Where(goqu.C("service_name").Eq(filter.ServiceName))
	}

	query, args, err := ds.ToSQL()
	if err != nil {
		return 0, err
	}

	var sum uint
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&sum)
	return sum, err
}
