package car_trading

import (
	"context"
	"encoding/json"
	"github.com/andReyM228/lib/auth"
	"github.com/andReyM228/lib/bus"
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/rabbit"
	"github.com/andReyM228/lib/responder"
	"github.com/gofiber/fiber/v2"
	"user_service/internal/services"
)

type Handler struct {
	carTrading services.CarTrading
	rabbit     rabbit.Rabbit
}

func NewHandler(carTrading services.CarTrading, rabbit rabbit.Rabbit) Handler {
	return Handler{
		carTrading: carTrading,
		rabbit:     rabbit,
	}
}

// TODO: передавать chat_id не как параметр, а в jwt токене, или переделать на rabbit

func (h Handler) BuyCar(ctx *fiber.Ctx) error {
	carID, err := ctx.ParamsInt("car_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	txHash := ctx.Params("tx_hash")
	if txHash == "" {
		return responder.HandleError(ctx, errs.BadRequestError{Cause: "empty tx_hash"})
	}

	chatID, err := auth.GetChatIDFromHeader(ctx)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.carTrading.BuyCar(ctx.Context(), chatID, int64(carID), txHash); err != nil {
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

//---------------------------------------------------------------------

func (h Handler) BrokerBuyCar(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var buyCarRequest bus.BuyCarRequest
	if err := json.Unmarshal(req.Payload, &buyCarRequest); err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	if err := h.carTrading.BuyCar(context.Background(), buyCarRequest.ChatID, buyCarRequest.CarID, buyCarRequest.TxHash); err != nil {
		return err
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, nil)
}

func (h Handler) GetUserCars(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var buyCarRequest bus.BuyCarRequest
	if err := json.Unmarshal(req.Payload, &buyCarRequest); err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	if err := h.carTrading.BuyCar(context.Background(), buyCarRequest.ChatID, buyCarRequest.CarID, buyCarRequest.TxHash); err != nil {
		return err
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, nil)
}
