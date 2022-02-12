package handlers

import (
	"ms-fetch/domain/constants/messages"
	"ms-fetch/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ResourceHandler struct {
	Handler
}

func NewResourceHandler(handler Handler) ResourceHandler {
	return ResourceHandler{handler}
}

func (handler ResourceHandler) GetResource(ctx *fiber.Ctx) error {
	// database processing
	uc := usecase.NewResourceUseCase(handler.Contract)
	data, err := uc.ReadResouce()
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, data.FilterResourcePresenter, http.StatusOK)
}

func (handler ResourceHandler) GetResourceWithUSD(ctx *fiber.Ctx) error {
	// database processing
	uc := usecase.NewResourceUseCase(handler.Contract)
	data, err := uc.ReadResouceUsingCurrUSD()
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, data.FilterResourceCurrUSDPresenter, http.StatusOK)
}

func (handler ResourceHandler) GetReportWeekResource(ctx *fiber.Ctx) error {
	// database processing
	uc := usecase.NewResourceUseCase(handler.Contract)
	data, err := uc.ReportWeekResouce()
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, data.ReportResourcePresenter, http.StatusOK)
}
