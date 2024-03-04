package services

import (
	"context"
	"user_service/internal/domain"
)

type (
	User interface {
		Login(chatID int64, password string) (int64, error)
		Registration(user domain.User) error
		GetUser(field string, id int64) (domain.User, error)
		UpdateUser(user domain.User) error
		DeleteUser(id int64) error
	}

	CarTrading interface {
		BuyCar(ctx context.Context, chatID, carID int64, txHash string) error
		GetCar(id int64) (domain.Car, error)
		GetCars(label string) (domain.Cars, error)
		GetUserCars(chatID int64) (domain.Cars, error)
		SellCar(chatID, carID int64) error
		CreateCar(car domain.Car) error
		UpdateCar(car domain.Car) error
		DeleteCar(id int64) error
	}
)
