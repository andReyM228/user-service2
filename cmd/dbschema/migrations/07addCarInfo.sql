-- +migrate Up
ALTER TABLE cars ADD COLUMN info TEXT DEFAULT NULL;

-- +migrate Down

ALTER TABLE cars DROP COLUMN info;