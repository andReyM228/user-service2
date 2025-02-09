package services

import (
	"context"
	"user_service/internal/domain/cars"
	"user_service/internal/domain/users"
)

type (
	User interface {
		Login(chatID int64, password string) (int64, error)
		Registration(user users.User) error
		GetUser(field string, id int64) (users.User, error)
		UpdateUser(user users.User) error
		DeleteUser(id int64) error
	}

	CarTrading interface {
		BuyCar(ctx context.Context, chatID, carID int64, txHash string) error
		GetCar(id int64) (cars.Car, error)
		GetCars() (cars.Cars, error)
		GetUserCars(chatID int64) (cars.Cars, error)
		SellCar(chatID, carID int64) error
		CreateCar(car cars.Car) error
		UpdateCar(car cars.Car) error
		DeleteCar(id int64) error
	}
)
