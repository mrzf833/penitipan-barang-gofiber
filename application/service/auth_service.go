package service

import (
	"gofiber-penitipan-barang/application/model"
	"gofiber-penitipan-barang/application/response"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	CheckPasswordHash(password string, hash string) bool
	getUserByUsername(username string) (*model.User, error)
	Login(c *fiber.Ctx) response.LoginResponse
	User(c *fiber.Ctx) response.LoginResponse
}
