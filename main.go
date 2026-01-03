package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/internal/api"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/internal/config"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/internal/connection"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/internal/repository"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/internal/service"
)

func main() {
	config, err := config.Get()
	if err != nil {
		log.Fatalln("failed to get the config:", err.Error())
	}

	dbConn, err := connection.GetDatabase(config.Database)
	if err != nil {
		log.Fatalln("failed to connect to the database:", err.Error())
	}

	customerRepository := repository.NewCustomer(dbConn)
	userRepository := repository.NewUser(dbConn)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(config, userRepository)

	app := fiber.New()

	api.NewCustomer(app, customerService)
	api.NewAuth(app, authService)

	addr := fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port)
	if err = app.Listen(addr); err != nil {
		log.Fatalln("failed to listen:", err.Error())
	}
}
