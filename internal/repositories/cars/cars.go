package cars

import (
	"errors"
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/log"
	"gorm.io/gorm"
	"user_service/internal/domain/cars"
	"user_service/internal/repositories"
)

var _ repositories.Cars = Repository{}

type Repository struct {
	db  *gorm.DB
	log log.Logger
}

func NewRepository(database *gorm.DB, log log.Logger) Repository {
	return Repository{
		db:  database,
		log: log,
	}
}

func (r Repository) Get(id int64) (cars.Car, error) {
	var car carDB

	query := r.db.Where("id = ?", id)
	if err := query.First(&car).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Info(err.Error())
			return cars.Car{}, errs.NotFoundError{What: "car"}
		}

		r.log.Error(err.Error())
		return cars.Car{}, errs.InternalError{Cause: err.Error()}
	}

	return car.toDomain(), nil
}

func (r Repository) GetAll() (cars.Cars, error) {
	var carsFromDB carsDB

	if err := r.db.Find(&carsFromDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Info(err.Error())
			return nil, errs.NotFoundError{What: "cars"}
		}

		r.log.Error(err.Error())
		return nil, errs.InternalError{Cause: err.Error()}
	}

	return toDomainList(carsFromDB), nil
}

func (r Repository) Update(car cars.Car) error {
	query := r.db.Model(&carDB{}).Where("id = ?", car.ID)
	if err := query.Updates(fromDomain(car)).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Create(car cars.Car) error {
	if err := r.db.Create(fromDomain(car)).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	query := r.db.Where("id = ?", id)
	if err := query.Delete(&carDB{}).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}
