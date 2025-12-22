package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/domain"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/dto"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := &customerApi{
		customerService: customerService,
	}
	app.Get("/customers", ca.Index)
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
