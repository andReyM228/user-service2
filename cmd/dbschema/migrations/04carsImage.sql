-- +migrate Up
ALTER TABLE cars ADD COLUMN image TEXT DEFAULT NULL;

-- +migrate Down

ALTER TABLE cars DROP COLUMN image;