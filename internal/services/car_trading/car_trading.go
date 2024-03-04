package car_trading

import (
	"context"
	"errors"
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/log"
	"github.com/andReyM228/one/chain_client"
	"user_service/internal/domain"
	"user_service/internal/repositories"
)

type Service struct {
	usersRepo       repositories.Users
	carsRepo        repositories.Cars
	userCarsRepo    repositories.UserCars
	transfersRepo   repositories.Transfers
	chain           chain_client.Client
	carSystemWallet string
	log             log.Logger
}

func NewService(usersRepo repositories.Users, carsRepo repositories.Cars, userCarsRepo repositories.UserCars, transfersRepo repositories.Transfers, chain chain_client.Client, carSystemWallet string, log log.Logger) Service {
	return Service{
		usersRepo:       usersRepo,
		carsRepo:        carsRepo,
		userCarsRepo:    userCarsRepo,
		transfersRepo:   transfersRepo,
		chain:           chain,
		carSystemWallet: carSystemWallet,
		log:             log,
	}
}

// TODO: проверить чтоб везде передавался ctx

func (s Service) BuyCar(ctx context.Context, chatID, carID int64, txHash string) error {
	user, err := s.usersRepo.Get(domain.FieldChatID, chatID)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}

	car, err := s.carsRepo.Get(carID)
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

	if err := s.userCarsRepo.Create(user.ID, car.ID); err != nil {
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

func (s Service) GetCar(id int64) (domain.Car, error) {
	car, err := s.carsRepo.Get(id)
	if err != nil {
		if errors.As(err, &repositories.InternalServerError{}) {
			s.log.Error(err.Error())
			return domain.Car{}, err
		}
		s.log.Debug(err.Error())

		return domain.Car{}, err
	}

	return car, nil
}

func (s Service) GetCars(label string) (domain.Cars, error) {
	cars, err := s.carsRepo.GetAll(label)
	if err != nil {
		if errors.As(err, &errs.InternalError{}) {
			s.log.Error(err.Error())
			return domain.Cars{}, err
		}
		s.log.Debug(err.Error())

		return domain.Cars{}, err
	}

	return cars, nil
}

func (s Service) GetUserCars(chatID int64) (domain.Cars, error) {
	user, err := s.usersRepo.Get(domain.FieldChatID, chatID)
	if err != nil {
		if errors.As(err, &errs.InternalError{}) {
			s.log.Error(err.Error())
			return domain.Cars{}, err
		}
		s.log.Debug(err.Error())

		return domain.Cars{}, err
	}

	return domain.Cars{Cars: user.Cars}, nil
}

func (s Service) CreateCar(car domain.Car) error {
	err := s.carsRepo.Create(car)
	if err != nil {
		if errors.As(err, &errs.InternalError{}) {
			s.log.Error(err.Error())
			return err
		}
		s.log.Debug(err.Error())

		return err
	}

	return nil
}

func (s Service) UpdateCar(car domain.Car) error {
	err := s.carsRepo.Update(car)
	if err != nil {
		if errors.As(err, &errs.InternalError{}) {
			s.log.Error(err.Error())
			return err
		}
		s.log.Debug(err.Error())

		return err
	}

	return nil
}

func (s Service) DeleteCar(id int64) error {
	err := s.carsRepo.Delete(id)
	if err != nil {
		if errors.As(err, &errs.InternalError{}) {
			s.log.Error(err.Error())
			return err
		}
		s.log.Debug(err.Error())

		return err
	}

	return nil
}
