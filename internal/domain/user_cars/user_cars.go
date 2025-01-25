package user_cars

import "time"

type UserCar struct {
	ID        int64
	UserID    int64
	CarID     int64
	CreatedAt time.Time `db:"created_at"`
}

type UserCars struct {
	Cars []UserCar
}
