package users

import (
	"errors"
	"user_service/internal/domain"
	"user_service/internal/domain/errs"
	"user_service/internal/repository"
	"user_service/internal/repository/cars"
	"user_service/internal/repository/transfers"
	"user_service/internal/repository/user_cars"
	"user_service/internal/repository/users"

	"github.com/andReyM228/lib/log"
)

const systemUser = 0

type Service struct {
	users     users.Repository
	cars      cars.Repository
	userCars  user_cars.Repository
	transfers transfers.Repository
	log       log.Logger
}

func NewService(users users.Repository, log log.Logger) Service {
	return Service{
		users: users,
		log:   log,
	}
}

func (s Service) Login(chatID int64, password string) (int64, error) {
	user, err := s.users.Get(domain.FieldChatID, chatID)
	if err != nil {
		if errors.As(err, &repository.NotFound{}) {
			return 0, errs.NotFoundError{What: "user"}
		}

		s.log.Error(err.Error())

		return 0, errs.InternalError{}
	}

	if password != user.Password {
		return 0, errs.Unauthorized{Cause: "wrong password"}
	}

	return int64(user.ID), nil
}

func (s Service) Registration(user domain.User) error {
	_, err := s.users.Get(domain.FieldChatID, user.ChatID)
	if err == nil {
		return errors.New("this user already registered")
	}

	_, err = s.users.Get(domain.FieldPhone, user.Phone)
	if err == nil {
		return errors.New("this phone number already taken")
	}

	err = s.users.Create(user)
	if err != nil {
		return err
	}

	return nil
}
