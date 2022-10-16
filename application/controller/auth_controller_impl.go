package controller

import (
	"gofiber-penitipan-barang/application/response"
	"gofiber-penitipan-barang/application/service"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(c *fiber.Ctx) error {
	data := controller.AuthService.Login(c)
	return c.JSON(response.WebResponse{
		Message: "Success login",
		Data:    data,
	})
}

func (controller *AuthControllerImpl) User(c *fiber.Ctx) error {
	data := controller.AuthService.User(c)
	return c.JSON(response.WebResponse{
		Message: "Get user login",
		Data:    data,
	})
}
