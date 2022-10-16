package router

import (
	"gofiber-penitipan-barang/application/controller"
	"gofiber-penitipan-barang/application/database"
	"gofiber-penitipan-barang/application/middleware"
	"gofiber-penitipan-barang/application/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func StudentRouterApi(api fiber.Router, validate *validator.Validate) {
	studentService := service.NewStudentService(database.DB, validate)
	studentController := controller.NewStudentController(studentService)

	studentApi := api.Group("student", middleware.AuthMiddleware())
	studentApi.Get("/", studentController.FindAll)
	studentApi.Post("/", studentController.Create)
	studentApi.Get("/:categoryId", studentController.FindById)
	studentApi.Put("/:categoryId", studentController.Update)
	studentApi.Delete("/:categoryId", studentController.Delete)
}
