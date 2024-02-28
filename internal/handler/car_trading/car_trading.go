package car_trading

import (
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/responder"
	"github.com/gofiber/fiber/v2"
	"user_service/internal/services"
)

type Handler struct {
	carTrading services.CarTrading
}

func NewHandler(carTrading services.CarTrading) Handler {
	return Handler{
		carTrading: carTrading,
	}
}

// TODO: передавать chat_id не как параметр, а в jwt токене, или переделать на rabbit

func (h Handler) BuyCar(ctx *fiber.Ctx) error {
	chatID, err := ctx.ParamsInt("chat_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	carID, err := ctx.ParamsInt("car_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	txHash := ctx.Params("tx_hash")
	if txHash == "" {
		return responder.HandleError(ctx, errs.BadRequestError{Cause: "empty tx_hash"})
	}

	if err := h.carTrading.BuyCar(ctx.Context(), int64(chatID), int64(carID), txHash); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) SellCar(ctx *fiber.Ctx) error {
	chatID, err := ctx.ParamsInt("chat_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	carID, err := ctx.ParamsInt("car_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.carTrading.SellCar(int64(chatID), int64(carID)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
