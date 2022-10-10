package service

import (
	"gofiber-penitipan-barang/application/model"

	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	FindAll(c *fiber.Ctx) []model.Category
	FindById(c *fiber.Ctx, categoryId int) model.Category
	Create(c *fiber.Ctx, categoryId int) model.Category
}
