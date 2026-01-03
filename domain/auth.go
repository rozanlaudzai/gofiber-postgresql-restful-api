package domain

import (
	"context"

	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
