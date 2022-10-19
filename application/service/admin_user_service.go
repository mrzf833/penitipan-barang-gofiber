package service

import (
	"gofiber-penitipan-barang/application/model"

	"github.com/gofiber/fiber/v2"
)

type AdminUserService interface {
	FindAll(c *fiber.Ctx) []model.User
	FindById(c *fiber.Ctx, userId int) model.User
	Create(c *fiber.Ctx) model.User
	Update(c *fiber.Ctx, userId int) model.User
	Delete(c *fiber.Ctx, userId int)
	UpdatePassword(c *fiber.Ctx, adminUserId int)
}
