package controller

import "github.com/gofiber/fiber/v2"

type CategoryController interface {
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}
