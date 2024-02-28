-- +migrate Up
CREATE TABLE IF NOT EXISTS cars(
                      id BIGSERIAL PRIMARY KEY NOT NULL ,
                      name TEXT NOT NULL ,
                      model TEXT NOT NULL,
                      created_at TIMESTAMP DEFAULT current_timestamp
);

-- +migrate Down

DROP TABLE IF EXISTS cars;