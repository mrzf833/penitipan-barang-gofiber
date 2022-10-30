package router

import (
	"gofiber-penitipan-barang/application/controller"
	"gofiber-penitipan-barang/application/database"
	"gofiber-penitipan-barang/application/middleware"
	"gofiber-penitipan-barang/application/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InventoryRouterApi(api fiber.Router, validate *validator.Validate) {
	inventoryService := service.NewInventoryService(database.DB, validate)
	inventoryController := controller.NewInventoryController(inventoryService)

	inventoryApi := api.Group("inventory", middleware.AuthMiddleware())
	inventoryApi.Get("/", inventoryController.FindAll)
	inventoryApi.Post("/", inventoryController.Create)
	inventoryApi.Get("/:inventoryId", inventoryController.FindById)
	inventoryApi.Put("/:inventoryId", inventoryController.Update)
	inventoryApi.Delete("/:inventoryId", inventoryController.Delete)
	inventoryApi.Put("/:inventoryId/take", inventoryController.TakeItem)
}
