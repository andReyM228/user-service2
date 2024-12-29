package users

import (
	"errors"
	"fmt"
	"github.com/andReyM228/lib/errs"
	"gorm.io/gorm"

	"github.com/andReyM228/lib/log"
	"user_service/internal/domain"
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

func (r Repository) Get(field string, value any) (domain.User, error) {
	var user domain.User

	query := r.db.Where(fmt.Sprintf("%s = ?", field), value)

	if err := query.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Info(err.Error())
			return domain.User{}, errs.NotFoundError{What: "user"}
		}

		r.log.Error(err.Error())
		return domain.User{}, errs.InternalError{Cause: err.Error()}
	}

	var cars []domain.Car
	query = r.db.Where("user_id = ?", user.ID)

	if err := query.Find(&cars).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Info(err.Error())
			return domain.User{}, errs.NotFoundError{What: "cars"}
		}

		r.log.Error(err.Error())
		return domain.User{}, errs.InternalError{Cause: err.Error()}
	}

	user.Cars = cars

	return user, nil
}

func (r Repository) Update(user domain.User) error {
	query := r.db.Model(&domain.User{}).Where("id = ?", user.ID)

	if err := query.Updates(domain.User{
		Name:           user.Name,
		Surname:        user.Surname,
		Phone:          user.Phone,
		Email:          user.Email,
		Password:       user.Password,
		ChatID:         user.ChatID,
		AccountAddress: user.AccountAddress,
	}).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Create(user domain.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	query := r.db.Where("id = ?", id)

	if err := query.Delete(&domain.User{}).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}
