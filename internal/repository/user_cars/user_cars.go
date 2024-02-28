package user_cars

import (
	"database/sql"
	"errors"

	"user_service/internal/domain"
	"user_service/internal/repository"

	"github.com/andReyM228/lib/log"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewRepository(database *sqlx.DB, log log.Logger) Repository {
	return Repository{
		db:  database,
		log: log,
	}
}

func (r Repository) Create(userID, carID int) error {
	if _, err := r.db.Exec("INSERT INTO user_cars (user_id, car_id) VALUES ($1, $2)", userID, carID); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Delete(userID, carID int) error {
	if _, err := r.db.Exec("DELETE FROM user_cars WHERE user_id = $1 AND car_id = $2", userID, carID); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) GetUserCars(userID int64) (domain.UserCars, error) {
	var userCars []domain.UserCar

	if err := r.db.Select(&userCars, "SELECT * FROM user_cars WHERE user_id = $1", userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.UserCars{}, repository.NotFound{NotFound: "user_cars"}
		}

		r.log.Error(err.Error())
		return domain.UserCars{}, repository.InternalServerError{}
	}

	return domain.UserCars{Cars: userCars}, nil
}
