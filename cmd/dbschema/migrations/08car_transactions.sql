-- +migrate Up
CREATE TYPE transaction_kind AS ENUM ('buy', 'sell');

CREATE TYPE car_transaction_status AS ENUM ('new', 'done', 'failed');

CREATE TABLE IF NOT EXISTS car_transactions(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    tx_hash TEXT NOT NULL UNIQUE,
    kind transaction_kind NOT NULL,
    status car_transaction_status NOT NULL DEFAULT 'new',
    error TEXT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- +migrate StatementBegin
create trigger car_transactions_set_updated_at
    before update
    on car_transactions
    for each row
    execute procedure trigger_set_updated_at();
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE IF EXISTS car_transactions;

DROP TYPE IF EXISTS transaction_kind;

DROP TYPE IF EXISTS car_transaction_status;