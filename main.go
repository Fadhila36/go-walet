package main

import (
	"go-walet/internal/api"
	"go-walet/internal/component"
	"go-walet/internal/config"
	"go-walet/internal/middleware"
	"go-walet/internal/repository"
	"go-walet/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	cacheConnection := component.GetCacheConnection()

	userRepository := repository.NewUser(dbConnection)

	UserService := service.NewUser(userRepository, cacheConnection)

	authMid := middleware.Authenticate(UserService)
	app := fiber.New()
	api.NewAuth(app, UserService, authMid)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
