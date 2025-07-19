-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
    id BIGSERIAL PRIMARY KEY,
    service_name TEXT NOT NULL,
    price INTEGER NOT NULL CHECK (price >= 0),
    user_id TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
