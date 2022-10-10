package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, validate *validator.Validate) {
	// Middleware
	api := app.Group("/api")
	CategoryRouterApi(api, validate)
}
