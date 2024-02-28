package repositories

import "user_service/internal/domain"

type (
	Cars interface {
		Get(id int64) (domain.Car, error)
		GetAll(label string) (domain.Cars, error)
		Update(car domain.Car) error
		Create(car domain.Car) error
		Delete(id int64) error
	}

	Users interface {
		Get(field string, value any) (domain.User, error)
		Update(user domain.User) error
		Create(user domain.User) error
		Delete(id int64) error
	}

	Transfers interface {
		Issue(ToAddress, Memo string, Amount int64) (string, error)
		Withdraw(ToAddress, Memo string, Amount int64) (string, error)
	}

	UserCars interface {
		Create(userID, carID int) error
		Delete(userID, carID int) error
		GetUserCars(userID int64) (domain.UserCars, error)
	}
)
