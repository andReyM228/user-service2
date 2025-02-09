package repositories

import (
	"context"
	"user_service/internal/domain/car_txs"
	"user_service/internal/domain/cars"
	"user_service/internal/domain/user_cars"
	"user_service/internal/domain/users"
)

type (
	Cars interface {
		Get(id int64) (cars.Car, error)
		GetAll() (cars.Cars, error)
		Update(car cars.Car) error
		Create(car cars.Car) error
		Delete(id int64) error
	}

	Users interface {
		Get(field string, value any) (users.User, error)
		Update(user users.User) error
		Create(user users.User) error
		Delete(id int64) error
	}

	Transfers interface {
		Issue(ToAddress, Memo string, Amount int64) (string, error)
		Withdraw(ToAddress, Memo string, Amount int64) (string, error)
	}

	UserCars interface {
		Create(userCar user_cars.UserCar) error
		Delete(userID, carID int) error
	}

	CarTx interface {
		Get(ctx context.Context, txHash string) (car_txs.CarTx, error)
		GetAll(ctx context.Context, kind string) (car_txs.CarTxs, error)
		Create(ctx context.Context, transaction car_txs.CarTx) (car_txs.CarTx, error)
		Update(ctx context.Context, transaction car_txs.CarTx) error
		Delete(ctx context.Context, id int64) error
	}
)
