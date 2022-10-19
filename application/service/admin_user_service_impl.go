package service

import (
	"gofiber-penitipan-barang/application/model"
	"gofiber-penitipan-barang/application/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminUserServiceImpl struct {
	DB       *gorm.DB
	validate *validator.Validate
}

func NewAdminUserService(DB *gorm.DB, validate *validator.Validate) AdminUserService {
	return &AdminUserServiceImpl{
		DB:       DB,
		validate: validate,
	}
}

func (service *AdminUserServiceImpl) FindAll(c *fiber.Ctx) []model.User {
	var users []model.User
	service.DB.Where("role = ?", "admin").Find(&users)
	return users
}

func (service *AdminUserServiceImpl) FindById(c *fiber.Ctx, userId int) model.User {
	var user model.User
	service.DB.Find(&user, userId)
	if user.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	return user
}

func (service *AdminUserServiceImpl) Create(c *fiber.Ctx) model.User {
	adminUserCreateRequest := request.AdminUserCreateRequest{}

	err := c.BodyParser(&adminUserCreateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	err = service.validate.Struct(adminUserCreateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	password := []byte(adminUserCreateRequest.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	adminUser := model.User{
		Name:     adminUserCreateRequest.Name,
		Username: adminUserCreateRequest.Username,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	tx := service.DB.Create(&adminUser)
	if tx.Error != nil {
		panic(fiber.NewError(500, tx.Error.Error()))
	}
	return adminUser
}

func (service *AdminUserServiceImpl) Update(c *fiber.Ctx, adminUserId int) model.User {
	adminUserUpdateRequest := request.AdminUserUpdateRequest{}

	err := c.BodyParser(&adminUserUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	err = service.validate.Struct(adminUserUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	adminUser := model.User{}

	tx := service.DB.First(&adminUser, adminUserId)
	adminUser.Name = adminUserUpdateRequest.Name
	adminUser.Username = adminUserUpdateRequest.Username

	if adminUser.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}

	tx = tx.Updates(&adminUser)
	if tx.Error != nil {
		panic(fiber.NewError(500, tx.Error.Error()))
	}
	return adminUser
}

func (service *AdminUserServiceImpl) Delete(c *fiber.Ctx, adminUserId int) {
	adminUser := model.User{}

	tx := service.DB.First(&adminUser, adminUserId)
	if adminUser.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	tx.Delete(&adminUser)
}

func (service *AdminUserServiceImpl) UpdatePassword(c *fiber.Ctx, adminUserId int) {
	adminUserPasswordUpdateRequest := request.AdminUserPasswordUpdateRequest{}

	err := c.BodyParser(&adminUserPasswordUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	err = service.validate.Struct(adminUserPasswordUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	// --------
	password := []byte(adminUserPasswordUpdateRequest.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	adminUser := model.User{}

	tx := service.DB.First(&adminUser, adminUserId)
	adminUser.Password = string(hashedPassword)

	if adminUser.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}

	tx.Updates(&adminUser)
}
