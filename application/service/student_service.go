package service

import (
	"gofiber-penitipan-barang/application/model"

	"github.com/gofiber/fiber/v2"
)

type StudentService interface {
	FindAll(c *fiber.Ctx) []model.Student
	FindById(c *fiber.Ctx, studentId int) model.Student
	Create(c *fiber.Ctx) model.Student
	Update(c *fiber.Ctx, studentId int) model.Student
	Delete(c *fiber.Ctx, studentId int)
}
