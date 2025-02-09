package users

import (
	"errors"
	"fmt"
	"github.com/andReyM228/lib/errs"
	"gorm.io/gorm"
	"user_service/internal/domain/users"
	"user_service/internal/repositories"

	"github.com/andReyM228/lib/log"
)

// TODO: db models

var _ repositories.Users = Repository{}

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

func (r Repository) Get(field string, value any) (users.User, error) {
	var user userDB
	if err := r.db.Where(fmt.Sprintf("%s = ?", field), value).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Info(err.Error())
			return users.User{}, errs.NotFoundError{What: "user"}
		}

		r.log.Error(err.Error())
		return users.User{}, errs.InternalError{Cause: err.Error()}
	}

	if err := r.db.Model(&user).Association("Cars").Find(&user.Cars); err != nil {
		r.log.Error(err.Error())
		return users.User{}, errs.InternalError{Cause: err.Error()}
	}

	return user.toDomain(), nil
}

func (r Repository) Update(user users.User) error {
	query := r.db.Model(&userDB{}).Where("id = ?", user.ID)
	if err := query.Updates(fromDomain(user)).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Create(user users.User) error {
	if err := r.db.Create(fromDomain(user)).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	query := r.db.Where("id = ?", id)
	if err := query.Delete(&userDB{}).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}
