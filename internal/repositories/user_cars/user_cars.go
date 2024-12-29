package user_cars

import (
	"github.com/andReyM228/lib/errs"
	"gorm.io/gorm"
	"user_service/internal/domain"

	"github.com/andReyM228/lib/log"
)

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

func (r Repository) Create(userCar domain.UserCar) error {
	if err := r.db.Create(&userCar).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Delete(userID, carID int) error {
	query := r.db.Where("user_id = ? AND car_id = ?", userID, carID)

	if err := query.Delete(&domain.UserCar{}).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}
