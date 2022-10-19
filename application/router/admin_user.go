package router

import (
	"gofiber-penitipan-barang/application/controller"
	"gofiber-penitipan-barang/application/database"
	"gofiber-penitipan-barang/application/middleware"
	"gofiber-penitipan-barang/application/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AdminUserRouterApi(api fiber.Router, validate *validator.Validate) {
	adminUserService := service.NewAdminUserService(database.DB, validate)
	adminUserController := controller.NewAdminUserController(adminUserService)

	adminUserApi := api.Group("user", middleware.AuthMiddleware())
	adminUserApi.Get("/", adminUserController.FindAll)
	adminUserApi.Post("/", adminUserController.Create)
	adminUserApi.Get("/:adminUserId", adminUserController.FindById)
	adminUserApi.Put("/:adminUserId", adminUserController.Update)
	adminUserApi.Delete("/:adminUserId", adminUserController.Delete)
	adminUserApi.Put("/:adminUserId/password", adminUserController.UpdatePassword)
}
