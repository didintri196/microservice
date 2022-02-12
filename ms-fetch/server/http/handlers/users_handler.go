package handlers

import (
	"github.com/gofiber/fiber/v2"
	"ms-fetch/domain/constants/messages"
	"net/http"
)

type UserHandler struct {
	Handler
}

func NewUserHandler(handler Handler) UserHandler {
	return UserHandler{handler}
}

func (handler UserHandler) Status(ctx *fiber.Ctx) error {
	// database processing
	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, handler.Handler.Contract.Auth, http.StatusOK)
}