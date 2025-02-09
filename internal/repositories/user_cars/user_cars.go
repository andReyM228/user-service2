package user_cars

import (
	"github.com/andReyM228/lib/errs"
	"gorm.io/gorm"
	"user_service/internal/domain/user_cars"
	"user_service/internal/repositories"

	"github.com/andReyM228/lib/log"
)

var _ repositories.UserCars = Repository{}

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

func (r Repository) Create(userCar user_cars.UserCar) error {
	if err := r.db.Create(fromDomain(userCar)).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Delete(userID, carID int) error {
	query := r.db.Where("user_id = ? AND car_id = ?", userID, carID)

	if err := query.Delete(&userCarDB{}).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}
