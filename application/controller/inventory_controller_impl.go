package controller

import (
	"gofiber-penitipan-barang/application/helper"
	"gofiber-penitipan-barang/application/response"
	"gofiber-penitipan-barang/application/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type InventoryControllerImpl struct {
	InventoryService service.InventoryService
}

func NewInventoryController(inventoryService service.InventoryService) InventoryController {
	return &InventoryControllerImpl{
		InventoryService: inventoryService,
	}
}

func (controller *InventoryControllerImpl) FindAll(c *fiber.Ctx) error {
	inventories := controller.InventoryService.FindAll(c)
	inventoryResponse := helper.ToInventoryResponses(inventories)
	return c.JSON(response.WebResponse{
		Message: "Success get all inventory",
		Data:    inventoryResponse,
	})
}

func (controller *InventoryControllerImpl) FindById(c *fiber.Ctx) error {
	inventoryId, _ := strconv.Atoi(c.Params("inventoryId"))
	inventory := controller.InventoryService.FindById(c, inventoryId)
	inventoryResponse := helper.ToInventoryResponse(inventory)
	return c.JSON(response.WebResponse{
		Message: "Success get inventory",
		Data:    inventoryResponse,
	})
}

func (controller *InventoryControllerImpl) Create(c *fiber.Ctx) error {
	inventory := controller.InventoryService.Create(c)
	inventoryResponse := helper.ToInventoryResponse(inventory)
	return c.JSON(response.WebResponse{
		Message: "Success create inventory",
		Data:    inventoryResponse,
	})
}

func (controller *InventoryControllerImpl) Update(c *fiber.Ctx) error {
	inventoryId, _ := strconv.Atoi(c.Params("inventoryId"))
	inventory := controller.InventoryService.Update(c, inventoryId)
	inventoryResponse := helper.ToInventoryResponse(inventory)
	return c.JSON(response.WebResponse{
		Message: "Success update inventory",
		Data:    inventoryResponse,
	})
}

func (controller *InventoryControllerImpl) Delete(c *fiber.Ctx) error {
	inventoryId, _ := strconv.Atoi(c.Params("inventoryId"))
	controller.InventoryService.Delete(c, inventoryId)
	return c.JSON(response.WebResponse{
		Message: "Success delete inventory",
	})
}

func (controller *InventoryControllerImpl) TakeItem(c *fiber.Ctx) error {
	inventoryId, _ := strconv.Atoi(c.Params("inventoryId"))
	inventory := controller.InventoryService.TakeItem(c, inventoryId)
	inventoryResponse := helper.ToInventoryResponse(inventory)
	return c.JSON(response.WebResponse{
		Message: "Success take item inventory",
		Data:    inventoryResponse,
	})
}
