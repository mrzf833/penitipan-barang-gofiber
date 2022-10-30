package service

import (
	"gofiber-penitipan-barang/application/model"

	"github.com/gofiber/fiber/v2"
)

type InventoryService interface {
	FindAll(c *fiber.Ctx) []model.Inventory
	FindById(c *fiber.Ctx, inventoryId int) model.Inventory
	Create(c *fiber.Ctx) model.Inventory
	Update(c *fiber.Ctx, inventoryId int) model.Inventory
	Delete(c *fiber.Ctx, inventoryId int)
	TakeItem(c *fiber.Ctx, inventoryId int) model.Inventory
}
