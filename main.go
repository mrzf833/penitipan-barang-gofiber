package main

import (
	"gofiber-penitipan-barang/application/database"
	"gofiber-penitipan-barang/application/router"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(), recover.New())
	validate := validator.New()

	database.ConnectDB()

	router.SetupRoutes(app, validate)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
