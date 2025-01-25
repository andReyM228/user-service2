package user_cars

import (
	"time"
	"user_service/internal/domain/user_cars"
)

type UserCarDB struct {
	ID        int64
	UserID    int64
	CarID     int64
	CreatedAt time.Time `db:"created_at"`
}

type UserCarsDB struct {
	Cars []UserCarDB
}

// TODO: подумать над entity
func fromDomain(userCar user_cars.UserCar) UserCarDB {
	return UserCarDB{
		ID:        userCar.ID,
		UserID:    userCar.UserID,
		CarID:     userCar.CarID,
		CreatedAt: userCar.CreatedAt,
	}
}

func (u UserCarDB) toDomain() user_cars.UserCar {
	return user_cars.UserCar{
		ID:        u.ID,
		UserID:    u.UserID,
		CarID:     u.CarID,
		CreatedAt: u.CreatedAt,
	}
}

func toDomainList(userCarsDB UserCarsDB) user_cars.UserCars {
	result := make([]user_cars.UserCar, 0, len(userCarsDB.Cars))

	for _, car := range userCarsDB.Cars {
		result = append(result, car.toDomain())
	}

	return user_cars.UserCars{
		Cars: result,
	}
}
