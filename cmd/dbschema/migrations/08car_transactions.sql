-- +migrate Up
CREATE TABLE IF NOT EXISTS car_transactions(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    tx_hash TEXT NOT NULL ,
    kind TEXT NOT NULL ,
    created_at TIMESTAMP DEFAULT current_timestamp
);

-- +migrate Down

DROP TABLE IF EXISTS car_transactions;