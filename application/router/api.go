package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, validate *validator.Validate) {
	// Middleware
	api := app.Group("/api")
	AuthRouterApi(api, validate)
	CategoryRouterApi(api, validate)
	StudentRouterApi(api, validate)
	AdminUserRouterApi(api, validate)
}
