package psqlsubscription

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteSubscription(ctx context.Context, subscriptionID uint64) error {
	const query = `
        DELETE * FROM subscriptions
		WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query,
		subscriptionID,
	)

	if err != nil {
		return fmt.Errorf("failed to delete subscription: %v", err)
	}

	return nil
}
