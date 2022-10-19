package controller

import (
	"gofiber-penitipan-barang/application/helper"
	"gofiber-penitipan-barang/application/response"
	"gofiber-penitipan-barang/application/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AdminUserControllerImpl struct {
	AdminUserService service.AdminUserService
}

func NewAdminUserController(adminUserService service.AdminUserService) AdminUserController {
	return &AdminUserControllerImpl{
		AdminUserService: adminUserService,
	}
}

func (controller *AdminUserControllerImpl) FindAll(c *fiber.Ctx) error {
	adminUsers := controller.AdminUserService.FindAll(c)
	adminUserResponse := helper.ToAdminUserResponses(adminUsers)
	return c.JSON(response.WebResponse{
		Message: "Success get all admin user",
		Data:    adminUserResponse,
	})
}

func (controller *AdminUserControllerImpl) FindById(c *fiber.Ctx) error {
	adminUserId, _ := strconv.Atoi(c.Params("adminUserId"))
	adminUser := controller.AdminUserService.FindById(c, adminUserId)
	adminUserResponse := helper.ToAdminUserResponse(adminUser)
	return c.JSON(response.WebResponse{
		Message: "Success get admin user",
		Data:    adminUserResponse,
	})
}

func (controller *AdminUserControllerImpl) Create(c *fiber.Ctx) error {
	adminUser := controller.AdminUserService.Create(c)
	adminUserResponse := helper.ToAdminUserResponse(adminUser)
	return c.JSON(response.WebResponse{
		Message: "Success create admin user",
		Data:    adminUserResponse,
	})
}

func (controller *AdminUserControllerImpl) Update(c *fiber.Ctx) error {
	adminUserId, _ := strconv.Atoi(c.Params("adminUserId"))
	adminUser := controller.AdminUserService.Update(c, adminUserId)
	adminUserResponse := helper.ToAdminUserResponse(adminUser)
	return c.JSON(response.WebResponse{
		Message: "Success update admin user",
		Data:    adminUserResponse,
	})
}

func (controller *AdminUserControllerImpl) Delete(c *fiber.Ctx) error {
	adminUserId, _ := strconv.Atoi(c.Params("adminUserId"))
	controller.AdminUserService.Delete(c, adminUserId)
	return c.JSON(response.WebResponse{
		Message: "Success delete admin user",
	})
}

func (controller *AdminUserControllerImpl) UpdatePassword(c *fiber.Ctx) error {
	adminUserId, _ := strconv.Atoi(c.Params("adminUserId"))
	controller.AdminUserService.UpdatePassword(c, adminUserId)
	return c.JSON(response.WebResponse{
		Message: "Success update password admin user",
	})
}
