package handlers

import (
	"ms-fetch/domain/constants/messages"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ResourceHandler struct {
	Handler
}

func NewResourceHandler(handler Handler) ResourceHandler {
	return ResourceHandler{handler}
}

func (handler ResourceHandler) GetCurrentResource(ctx *fiber.Ctx) error {

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, handler.Contract.Auth, http.StatusOK)
}
