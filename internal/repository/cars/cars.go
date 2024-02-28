package cars

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

//TODO: обработка ошибок

func (r Repository) Get(id int64) (domain.Car, error) {
	var car domain.Car

	if err := r.db.Get(&car, "SELECT * FROM cars WHERE id = $1", id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.Car{}, repository.NotFound{NotFound: "car"}
		}

		r.log.Error(err.Error())
		return domain.Car{}, repository.InternalServerError{}
	}

	return car, nil
}

func (r Repository) GetAll(label string) (domain.Cars, error) {
	var cars []domain.Car

	if err := r.db.Select(&cars, "SELECT * FROM cars WHERE name = $1", label); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.Cars{}, repository.NotFound{NotFound: "cars"}
		}

		r.log.Error(err.Error())
		return domain.Cars{}, repository.InternalServerError{}
	}

	return domain.Cars{Cars: cars}, nil
}

func (r Repository) Update(car domain.Car) error {
	_, err := r.db.Exec("UPDATE cars SET name = $1, model = $2 WHERE id = $3", car.Name, car.Model, car.ID)

	if err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Create(car domain.Car) error {
	if _, err := r.db.Exec("INSERT INTO cars (name, model) VALUES ($1, $2)", car.Name, car.Model); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}
