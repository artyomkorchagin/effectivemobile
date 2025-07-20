package psqlsubscription

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func teardownTestDB(repo *Repository) {
	repo.db.Exec("DROP TABLE IF EXISTS subscriptions;")
}

func insertTestSubscription(t *testing.T, repo *Repository, sub *types.SubscriptionCreateRequest) uint64 {
	startDate, _ := helpers.ParseTime(sub.StartDate)
	var endDateSQL sql.NullTime
	if sub.EndDate != "" {
		endDate, _ := helpers.ParseTime(sub.EndDate)
		endDateSQL = sql.NullTime{Time: endDate, Valid: true}
	}

	var id uint64
	err := repo.db.QueryRow(`
        INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`,
		sub.ServiceName, sub.Price, sub.UserUUID, startDate, endDateSQL,
	).Scan(&id)

	require.NoError(t, err)
	return id
}

func TestCreateSubscription(t *testing.T) {
	repo := setupTest(t)
	defer teardownTestDB(repo)
	ctx := context.Background()
	sub := types.NewSubscriptionCreateRequest("Netflix", 10, "123e4567-e89b-42d3-a456-556642440000", "01-2025", "02-2025")

	err := repo.CreateSubscription(ctx, &sub)
	require.NoError(t, err)

	var count int
	err = repo.db.QueryRow("SELECT COUNT(*) FROM subscriptions WHERE service_name = 'Netflix'").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestGetSubscription(t *testing.T) {
	repo := setupTest(t)

	ctx := context.Background()
	sub := types.NewSubscriptionCreateRequest("Spotify", 15, "3a7b2c8d-5f9e-4a1b-8c7d-0e2f5a4b3c1d", "01-2025", "02-2025")
	id := insertTestSubscription(t, repo, &sub)

	result, err := repo.GetSubscription(ctx, id)
	require.NoError(t, err)
	assert.Equal(t, "Spotify", result.ServiceName)
	assert.Equal(t, "3a7b2c8d-5f9e-4a1b-8c7d-0e2f5a4b3c1d", result.UserUUID)
	assert.Equal(t, "01-2025", result.StartDate)
	assert.Equal(t, "02-2025", result.EndDate)
}

func TestUpdateSubscription(t *testing.T) {
	repo := setupTest(t)

	ctx := context.Background()
	sub := types.NewSubscriptionCreateRequest("YouTube", 20, "7c6d3a1b-9e8f-4d7c-8b5a-2e9c1d0e4f3a", "01-2025", "02-2025")
	id := insertTestSubscription(t, repo, &sub)

	updated := &types.SubscriptionUpdateRequest{
		ID:          id,
		ServiceName: "YouTube Premium",
		Price:       25,
		UserUUID:    "1a2b3c4d-5e6f-7a8b-9c0d-0e1f2a3b4c5d",
		StartDate:   "03-2025",
		EndDate:     "04-2025",
	}

	err := repo.UpdateSubscription(ctx, updated)
	require.NoError(t, err)

	result, err := repo.GetSubscription(ctx, id)
	require.NoError(t, err)
	assert.Equal(t, "YouTube Premium", result.ServiceName)
	assert.Equal(t, "1a2b3c4d-5e6f-7a8b-9c0d-0e1f2a3b4c5d", result.UserUUID)
	assert.Equal(t, "03-2025", result.StartDate)
	assert.Equal(t, "04-2025", result.EndDate)
}

func TestDeleteSubscription(t *testing.T) {
	repo := setupTest(t)

	ctx := context.Background()
	sub := types.NewSubscriptionCreateRequest("Prime", 12, "0a1b2c3d-4e5f-6a7b-8c9d-0e1f2a3b4c5d", "01-2025", "")
	id := insertTestSubscription(t, repo, &sub)

	err := repo.DeleteSubscription(ctx, id)
	require.NoError(t, err)

	_, err = repo.GetSubscription(ctx, id)
	assert.Error(t, err)
}

func TestGetAllSubscriptions(t *testing.T) {
	repo := setupTest(t)

	ctx := context.Background()
	sub1 := types.NewSubscriptionCreateRequest("Prime", 12, "6f5a4e3d-2c1b-9d8e-7f0a-1e2d3c4b5a6f", "01-2025", "02-2025")
	sub2 := types.NewSubscriptionCreateRequest("Disney", 9, "2c4e6d8f-0a1b-2c3d-4e5f-6a7b8c9d0e1f", "03-2025", "")

	insertTestSubscription(t, repo, &sub1)
	insertTestSubscription(t, repo, &sub2)

	subs, err := repo.GetAllSubscriptions(ctx)
	require.NoError(t, err)
	assert.Len(t, subs, 2)
}

func TestGetSumOfSubscriptions(t *testing.T) {
	repo := setupTest(t)

	ctx := context.Background()

	req := types.NewSubscriptionCreateRequest("Netflix", 10, "3d5f7e9c-1a2b-3d4e-5f6a-7b8c9d0e1f2a", "01-2024", "03-2024")
	id := insertTestSubscription(t, repo, &req)
	req2 := types.NewSubscriptionCreateRequest("Netflix", 10, "3d5f7e9c-1a2b-3d4e-5f6a-7b8c9d0e1f2a", "01-2024", "02-2024")
	insertTestSubscription(t, repo, &req2)

	filter := &types.Filter{
		UserUUID:    "3d5f7e9c-1a2b-3d4e-5f6a-7b8c9d0e1f2a",
		ServiceName: "Netflix",
		StartDate:   "01-2024",
		EndDate:     "03-2024",
	}

	sum, err := repo.GetSumOfSubscriptions(ctx, filter)
	require.NoError(t, err)

	assert.Equal(t, uint(50), sum)

	repo.db.Exec("DELETE FROM subscriptions WHERE id = $1", id)
}

func setupTest(t *testing.T) *Repository {
	db, err := sql.Open("pgx", helpers.GetEnv("TESTDB_DSN", ""))
	require.NoError(t, err)

	repo, err := NewRepository(db)
	require.NoError(t, err)

	return repo
}
