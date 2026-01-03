package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/domain"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/dto"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	config         *config.Config
	userRepository domain.UserRepository
}

func NewAuth(config *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &authService{
		config:         config,
		userRepository: userRepository,
	}
}

func (as *authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := as.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	claim := jwt.MapClaims{
		"id":      user.Id,
		"expired": time.Now().Add(time.Duration(as.config.Jwt.Expired) * time.Minute).Unix(),
	}
	tokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := tokenByte.SignedString([]byte(as.config.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}
	return dto.AuthResponse{Token: tokenString}, nil
}
