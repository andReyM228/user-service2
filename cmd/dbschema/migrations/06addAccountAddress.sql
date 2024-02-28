-- +migrate Up
ALTER TABLE users ADD COLUMN account_address TEXT;

-- +migrate Down
ALTER TABLE users DROP COLUMN account_address;