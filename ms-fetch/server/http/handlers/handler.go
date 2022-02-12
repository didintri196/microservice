package handlers

import (
	"ms-fetch/domain/presenters"
	"ms-fetch/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Contract *usecase.Contract
}

func NewHandler(ucContract *usecase.Contract) Handler {
	return Handler{Contract: ucContract}
}

func (handler Handler) SendResponseWithoutMeta(ctx *fiber.Ctx, message string, data interface{}, httpStatus int) error {
	return ctx.Status(httpStatus).JSON(presenters.ResponsePresenter{
		Message: message,
		Data:    data,
	})
}

func (handler Handler) SendResponseWithMeta(ctx *fiber.Ctx, data interface{}, message string, meta interface{}, httpStatus int) error {
	return ctx.Status(httpStatus).JSON(presenters.ResponsePresenter{
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}
