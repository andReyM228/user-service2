-- +migrate Up

ALTER TABLE cars
    ADD COLUMN IF NOT EXISTS price NUMERIC DEFAULT 0 ;


CREATE TABLE IF NOT EXISTS user_cars (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id),
    car_id integer REFERENCES cars(id),
    created_at TIMESTAMP DEFAULT current_timestamp
);

-- +migrate Down

DROP TABLE user_cars;

ALTER TABLE cars
    DROP COLUMN price;



