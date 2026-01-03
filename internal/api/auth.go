package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/domain"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/dto"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuth(app *fiber.App, authService domain.AuthService) {
	aa := &authApi{
		authService: authService,
	}
	app.Post("/auth", aa.Login)
}

func (aa *authApi) Login(fCtx *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(fCtx.Context(), 10*time.Minute)
	defer cancel()

	var req dto.AuthRequest
	if err := fCtx.BodyParser(&req); err != nil {
		return fCtx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, err := aa.authService.Login(ctx, req)
	if err != nil {
		return fCtx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return fCtx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(res))
}
