package users

import (
	"encoding/json"
	"github.com/andReyM228/lib/bus"
	"github.com/andReyM228/lib/rabbit"
	"github.com/andReyM228/lib/responder"
	"github.com/gofiber/fiber/v2"
	"user_service/internal/domain"
	"user_service/internal/domain/errs"
	"user_service/internal/repositories"
	"user_service/internal/services"
)

type Handler struct {
	userRepo    repositories.Users
	userService services.User
	rabbit      rabbit.Rabbit
}

func NewHandler(userRepo repositories.Users, userService services.User, rabbit rabbit.Rabbit) Handler {
	return Handler{
		userRepo:    userRepo,
		userService: userService,
		rabbit:      rabbit,
	}
}

func (h Handler) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	user, err := h.userService.GetUser(domain.FieldID, int64(id))
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	payload, err := json.Marshal(user)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.Send(payload)
}

func (h Handler) Update(ctx *fiber.Ctx) error {
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.userService.UpdateUser(user); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Create(ctx *fiber.Ctx) error {
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.userService.Registration(user); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.userService.DeleteUser(int64(id)); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Login(ctx *fiber.Ctx) error {
	var request loginRequest
	if err := ctx.BodyParser(&request); err != nil {
		return responder.HandleError(ctx, err)
	}

	if request.ChatID == 0 || request.Password == "" {
		return responder.HandleError(ctx, errs.BadRequestError{Cause: "wrong body"})
	}

	userID, err := h.userService.Login(request.ChatID, request.Password)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	payload, err := json.Marshal(loginResponse{userID})
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.Send(payload)
}

// -----------------------------------------------------------------------

func (h Handler) BrokerCreate(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var user domain.User
	if err := json.Unmarshal(req.Payload, &user); err != nil {
		return err
	}

	if err := h.userService.Registration(user); err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, nil)
}

func (h Handler) BrokerLogin(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var loginRequest bus.LoginRequest
	if err := json.Unmarshal(req.Payload, &loginRequest); err != nil {
		return err
	}

	if loginRequest.ChatID == 0 || loginRequest.Password == "" {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	userID, err := h.userService.Login(loginRequest.ChatID, loginRequest.Password)
	if err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, loginResponse{userID})
}

func (h Handler) BrokerGetUserByID(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var userRequest bus.GetUserByIDRequest
	if err := json.Unmarshal(req.Payload, &userRequest); err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	user, err := h.userService.GetUser(domain.FieldID, userRequest.ID)
	if err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, user)
}
