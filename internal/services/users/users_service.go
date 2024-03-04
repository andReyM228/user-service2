package users

import (
	"errors"
	"github.com/andReyM228/lib/log"
	"user_service/internal/domain"
	"user_service/internal/domain/errs"
	"user_service/internal/repositories"
)

type Service struct {
	usersRepo repositories.Users
	log       log.Logger
}

func NewService(usersRepo repositories.Users, log log.Logger) Service {
	return Service{
		usersRepo: usersRepo,
		log:       log,
	}
}

func (s Service) Login(chatID int64, password string) (int64, error) {
	user, err := s.usersRepo.Get(domain.FieldChatID, chatID)
	if err != nil {
		if errors.As(err, &repositories.NotFound{}) {
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
	_, err := s.usersRepo.Get(domain.FieldChatID, user.ChatID)
	if err == nil {
		return errors.New("this user already registered")
	}

	_, err = s.usersRepo.Get(domain.FieldPhone, user.Phone)
	if err == nil {
		return errors.New("this phone number already taken")
	}

	err = s.usersRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetUser(field string, id int64) (domain.User, error) {
	user, err := s.usersRepo.Get(field, id)
	if err != nil {
		if errors.As(err, &repositories.NotFound{}) {
			return domain.User{}, errs.NotFoundError{What: "user"}
		}

		s.log.Error(err.Error())

		return domain.User{}, errs.InternalError{}
	}

	return user, nil
}

func (s Service) UpdateUser(user domain.User) error {
	err := s.usersRepo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) DeleteUser(id int64) error {
	err := s.usersRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
