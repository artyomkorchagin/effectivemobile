package psqlsubscription

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
)

// NOTE: if endDate is null then it is considered as current date
func (r *Repository) GetSumOfSubscriptions(ctx context.Context, filter *types.Filter) (uint, error) {
	query, args := buildSumQuery(filter)

	var sum uint
	err := r.db.QueryRow(query, args...).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func buildSumQuery(filter *types.Filter) (string, []interface{}) {
	var whereClauses []string
	var args []interface{}
	paramIndex := 1

	if filter.UserUUID != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("user_id = $%d", paramIndex))
		args = append(args, filter.UserUUID)
		paramIndex++
	} else {
		args = append(args, nil)
		paramIndex++
	}

	startDate := filter.StartDate
	if startDate == "" {
		startDate = "01-01-2015" // some default date if startDate is null
	} else {
		startDate = fmt.Sprintf("01-%s", startDate)
	}
	args = append(args, startDate)
	paramIndex++

	endDate := filter.EndDate
	if endDate == "" {
		endDate = time.Now().Format("02-01-2006")
	} else {
		endDate = fmt.Sprintf("01-%s", endDate)
	}
	args = append(args, endDate)
	paramIndex++

	if filter.ServiceName != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("service_name = $%d", paramIndex))
		args = append(args, filter.ServiceName)
		paramIndex++
	} else {
		args = append(args, nil)
		paramIndex++
	}

	query := `
        SELECT COALESCE(SUM(
            price * (
                SELECT COUNT(*) FROM generate_series(
                    date_trunc('month', GREATEST(start_date, $2::date)),
                    LEAST(COALESCE(end_date, CURRENT_DATE), $3::date),
                    '1 month'
                )
            )
        ), 0) AS total_sum
        FROM subscriptions
        WHERE
            ($1::text IS NULL OR user_id = $1) AND
            ($4::text IS NULL OR service_name = $4) AND
            start_date <= COALESCE(end_date, CURRENT_DATE)
    `

	if len(whereClauses) > 0 {
		query += " AND " + strings.Join(whereClauses, " AND ")
	}

	return query, args
}
