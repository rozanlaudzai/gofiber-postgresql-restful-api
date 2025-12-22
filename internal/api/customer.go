package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/domain"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/dto"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/internal/util"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := &customerApi{
		customerService: customerService,
	}
	app.Get("/customers", ca.Index)
	app.Post("/customers", ca.Create)
	app.Put("/customers/:id", ca.Update)
}

func (ca *customerApi) Index(fCtx *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(fCtx.Context(), 10*time.Second)
	defer cancel()

	customerData, err := ca.customerService.Index(ctx)
	if err != nil {
		return fCtx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return fCtx.JSON(dto.CreateResponseSuccess(customerData))
}

func (ca *customerApi) Create(fCtx *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(fCtx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := fCtx.BodyParser(&req); err != nil {
		return fCtx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return fCtx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}
	err := ca.customerService.Create(ctx, req)
	if err != nil {
		return fCtx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return fCtx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess(""))
}

func (ca *customerApi) Update(fCtx *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(fCtx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	if err := fCtx.BodyParser(&req); err != nil {
		return fCtx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return fCtx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation error", fails))
	}

	req.ID = fCtx.Params("id")
	err := ca.customerService.Update(ctx, req)
	if err != nil {
		return fCtx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return fCtx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(""))
}
