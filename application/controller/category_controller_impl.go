package controller

import (
	"gofiber-penitipan-barang/application/helper"
	"gofiber-penitipan-barang/application/response"
	"gofiber-penitipan-barang/application/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) FindAll(c *fiber.Ctx) error {
	categories := controller.CategoryService.FindAll(c)
	categoryResponse := helper.ToCategoryResponses(categories)
	return c.JSON(response.WebResponse{
		Message: "Success get all category",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) FindById(c *fiber.Ctx) error {
	categoryId, _ := strconv.Atoi(c.Params("categoryId"))
	category := controller.CategoryService.FindById(c, categoryId)
	categoryResponse := helper.ToCategoryResponse(category)
	return c.JSON(response.WebResponse{
		Message: "Success get all category",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Create(c *fiber.Ctx) error {
	category := controller.CategoryService.Create(c)
	categoryResponse := helper.ToCategoryResponse(category)
	return c.JSON(response.WebResponse{
		Message: "Success create category",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Update(c *fiber.Ctx) error {
	categoryId, _ := strconv.Atoi(c.Params("categoryId"))
	category := controller.CategoryService.Update(c, categoryId)
	categoryResponse := helper.ToCategoryResponse(category)
	return c.JSON(response.WebResponse{
		Message: "Success update category",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Delete(c *fiber.Ctx) error {
	categoryId, _ := strconv.Atoi(c.Params("categoryId"))
	controller.CategoryService.Delete(c, categoryId)
	return c.JSON(response.WebResponse{
		Message: "Success delete category",
	})
}
