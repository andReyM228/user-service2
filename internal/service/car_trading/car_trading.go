package car_trading

import (
	"context"
	"errors"
	"github.com/andReyM228/one/chain_client"
	"user_service/internal/domain"
	"user_service/internal/repository"
	"user_service/internal/repository/cars"
	"user_service/internal/repository/transfers"
	"user_service/internal/repository/user_cars"
	"user_service/internal/repository/users"

	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/log"
)

const systemUser = 0

type Service struct {
	users           users.Repository
	cars            cars.Repository
	userCars        user_cars.Repository
	transfers       transfers.Repository
	chain           chain_client.Client
	carSystemWallet string
	log             log.Logger
}

func NewService(users users.Repository, cars cars.Repository, userCars user_cars.Repository, transfers transfers.Repository, chain chain_client.Client, carSystemWallet string, log log.Logger) Service {
	return Service{
		users:           users,
		cars:            cars,
		userCars:        userCars,
		transfers:       transfers,
		chain:           chain,
		carSystemWallet: carSystemWallet,
		log:             log,
	}
}

// TODO: проверить чтоб везде передавался ctx

func (s Service) BuyCar(ctx context.Context, chatID, carID int64, txHash string) error {
	user, err := s.users.Get(domain.FieldChatID, chatID)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}

	car, err := s.cars.Get(carID)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}

	tx, err := s.getTx(ctx, txHash)
	if err != nil {
		return err
	}

	if tx.ToAddress != s.carSystemWallet {
		err = errs.ForbiddenError{Cause: "invalid account_address_to"}
		s.log.Error(err.Error())
		return err
	}

	if user.AccountAddress != tx.FromAddress {
		err = errs.BadRequestError{Cause: "wrong account address"}
		s.log.Error(err.Error())
		return err
	}

	if tx.Amount.AmountOf(chain_client.DenomOne).Int64() < car.Price {
		err = errs.BadRequestError{Cause: "not enough transaction amount"}
		s.log.Error(err.Error())
		return err
	}

	if err := s.userCars.Create(user.ID, car.ID); err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info("Car sent")

	return nil
}

func (s Service) SellCar(chatID, carID int64) error {
	//user, err := s.users.Get(domain.FieldChatID, chatID)
	//if err != nil {
	//	s.log.Error(err.Error())
	//	return err
	//}
	//
	////проверка есть ли у юзера машина
	//
	//var existCar bool
	//
	//for _, car := range user.Cars {
	//	if car.ID == int(carID) {
	//		existCar = true
	//		break
	//	}
	//
	//}
	//
	//if existCar == false {
	//	s.log.Error(err.Error())
	//	return err
	//}
	//
	//car, err := s.cars.Get(carID)
	//if err != nil {
	//	s.log.Error(err.Error())
	//	return err
	//}
	//
	//s.log.Info("Sending transfer")
	//if err := s.transfers.Transfer(systemUser, user.ID, int(car.Price)); err != nil {
	//	s.log.Error(err.Error())
	//	return err
	//}
	//s.log.Info("Transfer sent")
	//
	//if err := s.userCars.Delete(user.ID, car.ID); err != nil {
	//	s.log.Error(err.Error())
	//	return err
	//}
	//s.log.Info("Car sell")
	//
	//if err := s.userCars.Create(systemUser, car.ID); err != nil {
	//	s.log.Error(err.Error())
	//	return err
	//}
	//s.log.Info("Car sell")

	return nil
}

//TODO: обработка ошибок

func (s Service) GetCar(id int64) (domain.Car, error) {
	car, err := s.cars.Get(id)
	if err != nil {
		if errors.As(err, &repository.InternalServerError{}) {
			s.log.Error(err.Error())
			return domain.Car{}, errs.InternalError{}
		}
		s.log.Debug(err.Error())

		return domain.Car{}, errs.NotFoundError{What: err.Error()}
	}

	return car, nil
}

func (s Service) GetCars(label string) (domain.Cars, error) {
	cars, err := s.cars.GetAll(label)
	if err != nil {
		if errors.As(err, &repository.InternalServerError{}) {
			s.log.Error(err.Error())
			return domain.Cars{}, errs.InternalError{}
		}
		s.log.Debug(err.Error())

		return domain.Cars{}, errs.NotFoundError{What: err.Error()}
	}

	return cars, nil
}

// TODO: вынести field "chat_id" в domain

func (s Service) GetUserCars(chatID int64) (domain.Cars, error) {
	user, err := s.users.Get("chat_id", chatID)
	if err != nil {
		if errors.As(err, &repository.InternalServerError{}) {
			s.log.Error(err.Error())
			return domain.Cars{}, errs.InternalError{}
		}
		s.log.Debug(err.Error())

		return domain.Cars{}, errs.NotFoundError{What: err.Error()}
	}

	return domain.Cars{Cars: user.Cars}, nil
}
