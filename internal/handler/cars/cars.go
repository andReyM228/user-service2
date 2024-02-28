package cars

import (
	"encoding/json"
	"github.com/andReyM228/lib/bus"
	"github.com/andReyM228/lib/rabbit"

	"user_service/internal/domain"
	"user_service/internal/repository/cars"
	"user_service/internal/service/car_trading"

	"github.com/andReyM228/lib/auth"
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/responder"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	carRepo    cars.Repository
	carService car_trading.Service
	rabbit     rabbit.Rabbit
}

func NewHandler(repo cars.Repository, service car_trading.Service, rabbit rabbit.Rabbit) Handler {
	return Handler{
		carRepo:    repo,
		carService: service,
		rabbit:     rabbit,
	}
}

// TODO: обработка ошибок
// TODO: сервисный уровень

func (h Handler) Get(ctx *fiber.Ctx) error {
	token, err := responder.GetToken(ctx)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := auth.VerifyToken(token); err != nil {
		return responder.HandleError(ctx, errs.UnauthorizedError{Cause: err.Error()})
	}

	//TODO: сделать так везде
	//auth.GetChatID(token)

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	car, err := h.carService.GetCar(int64(id))
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	payload, err := json.Marshal(car)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.Send(payload)
}

func (h Handler) GetAll(ctx *fiber.Ctx) error {
	token, err := responder.GetToken(ctx)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := auth.VerifyToken(token); err != nil {
		return responder.HandleError(ctx, errs.UnauthorizedError{Cause: err.Error()})
	}

	label := ctx.Params("name", "bmw")

	cars, err := h.carService.GetCars(label)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	payload, err := json.Marshal(cars)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.Send(payload)
}

func (h Handler) GetUserCars(ctx *fiber.Ctx) error {
	token, err := responder.GetToken(ctx)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := auth.VerifyToken(token); err != nil {
		return responder.HandleError(ctx, errs.UnauthorizedError{Cause: err.Error()})
	}

	chatID, err := auth.GetChatID(token)
	if err != nil {
		return responder.HandleError(ctx, errs.UnauthorizedError{Cause: err.Error()})
	}

	cars, err := h.carService.GetUserCars(chatID)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	payload, err := json.Marshal(cars)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.Send(payload)
}

func (h Handler) Update(ctx *fiber.Ctx) error {
	var car domain.Car
	if err := ctx.BodyParser(&car); err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.carRepo.Update(car); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Create(ctx *fiber.Ctx) error {
	var car domain.Car
	if err := ctx.BodyParser(&car); err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.carRepo.Create(car); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.carRepo.Delete(int64(id)); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

//------------------------------------------------------------------

func (h Handler) BrokerGetCarByID(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var carRequest bus.GetCarByIDRequest
	if err := json.Unmarshal(req.Payload, &carRequest); err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	car, err := h.carRepo.Get(carRequest.ID)
	if err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, car)
}
