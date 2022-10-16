package router

import (
	"gofiber-penitipan-barang/application/controller"
	"gofiber-penitipan-barang/application/database"
	"gofiber-penitipan-barang/application/middleware"
	"gofiber-penitipan-barang/application/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CategoryRouterApi(api fiber.Router, validate *validator.Validate) {
	categoryService := service.NewCategoryService(database.DB, validate)
	categoryController := controller.NewCategoryController(categoryService)

	categoryApi := api.Group("category", middleware.AuthMiddleware())
	categoryApi.Get("/", categoryController.FindAll)
	categoryApi.Post("/", categoryController.Create)
	categoryApi.Get("/:categoryId", categoryController.FindById)
	categoryApi.Put("/:categoryId", categoryController.Update)
	categoryApi.Delete("/:categoryId", categoryController.Delete)
}
