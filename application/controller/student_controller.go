package controller

import "github.com/gofiber/fiber/v2"

type StudentController interface {
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
