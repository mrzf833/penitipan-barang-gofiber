package router

import (
	"gofiber-penitipan-barang/application/controller"
	"gofiber-penitipan-barang/application/database"
	"gofiber-penitipan-barang/application/middleware"
	"gofiber-penitipan-barang/application/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthRouterApi(api fiber.Router, validate *validator.Validate) {
	authService := service.NewAuthService(database.DB, validate)
	authController := controller.NewAuthControllerImpl(authService)

	api.Post("/login", authController.Login)
	api.Post("/user-login", middleware.AuthMiddleware(), authController.User)
}
