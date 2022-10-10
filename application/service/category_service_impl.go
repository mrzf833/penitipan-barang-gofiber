package service

import (
	"gofiber-penitipan-barang/application/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	DB       *gorm.DB
	validate *validator.Validate
}

func NewCategoryService(DB *gorm.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		DB:       DB,
		validate: validate,
	}
}

func (service *CategoryServiceImpl) FindAll(c *fiber.Ctx) []model.Category {
	var categories []model.Category
	service.DB.Find(&categories)
	return categories
}

func (service *CategoryServiceImpl) FindById(c *fiber.Ctx, categoryId int) model.Category {
	var category model.Category
	service.DB.Find(&category, categoryId)
	if category.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	return category
}

func (service *CategoryServiceImpl) Create(c *fiber.Ctx, categoryId int) model.Category {
	var category model.Category
	service.DB.Find(&category, categoryId)
	if category.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	return category
}
