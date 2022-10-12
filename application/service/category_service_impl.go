package service

import (
	"gofiber-penitipan-barang/application/model"
	"gofiber-penitipan-barang/application/request"

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

func (service *CategoryServiceImpl) Create(c *fiber.Ctx) model.Category {
	categoryCreateRequest := request.CategoryCreateRequest{}

	err := c.BodyParser(&categoryCreateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	category := model.Category{
		Name: categoryCreateRequest.Name,
	}

	err = service.validate.Struct(category)

	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	service.DB.Create(&category)
	return category
}

func (service *CategoryServiceImpl) Update(c *fiber.Ctx, categoryId int) model.Category {
	categoryUpdateRequest := request.CategoryUpdateRequest{}

	err := c.BodyParser(&categoryUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	category := model.Category{
		Id:   categoryId,
		Name: categoryUpdateRequest.Name,
	}

	err = service.validate.Struct(category)

	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	tx := service.DB.First(&model.Category{}, categoryId)
	if category.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	tx.Updates(&category)
	return category
}

func (service *CategoryServiceImpl) Delete(c *fiber.Ctx, categoryId int) {
	category := model.Category{}

	tx := service.DB.First(&category, categoryId)
	if category.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	tx.Delete(&category)
}
