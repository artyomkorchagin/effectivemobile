-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS subscriptions (
    id BIGSERIAL PRIMARY KEY,
    service_name TEXT NOT NULL,
    price INTEGER NOT NULL CHECK (price >= 0),
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    CONSTRAINT chk_end_date_after_start CHECK (end_date IS NULL OR end_date >= start_date)
);

CREATE INDEX IF NOT EXISTS idx_subscriptions_user_id ON subscriptions(user_id);

CREATE INDEX IF NOT EXISTS idx_subscriptions_start_date ON subscriptions(start_date);

CREATE INDEX IF NOT EXISTS idx_subscriptions_user_active ON subscriptions(user_id, start_date, end_date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS idx_subscriptions_user_active;

DROP INDEX IF EXISTS idx_subscriptions_start_date;

DROP INDEX IF EXISTS idx_subscriptions_user_id;

DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
