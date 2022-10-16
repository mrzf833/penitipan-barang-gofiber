package controller

import (
	"gofiber-penitipan-barang/application/helper"
	"gofiber-penitipan-barang/application/response"
	"gofiber-penitipan-barang/application/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &StudentControllerImpl{
		StudentService: studentService,
	}
}

func (controller *StudentControllerImpl) FindAll(c *fiber.Ctx) error {
	students := controller.StudentService.FindAll(c)
	studentsResponse := helper.ToStudentResponses(students)
	return c.JSON(response.WebResponse{
		Message: "Success get all student",
		Data:    studentsResponse,
	})
}

func (controller *StudentControllerImpl) FindById(c *fiber.Ctx) error {
	studentId, _ := strconv.Atoi(c.Params("studentId"))
	student := controller.StudentService.FindById(c, studentId)
	studentResponse := helper.ToStudentResponse(student)
	return c.JSON(response.WebResponse{
		Message: "Success get student",
		Data:    studentResponse,
	})
}

func (controller *StudentControllerImpl) Create(c *fiber.Ctx) error {
	student := controller.StudentService.Create(c)
	studentResponse := helper.ToStudentResponse(student)
	return c.JSON(response.WebResponse{
		Message: "Success create student",
		Data:    studentResponse,
	})
}

func (controller *StudentControllerImpl) Update(c *fiber.Ctx) error {
	studentId, _ := strconv.Atoi(c.Params("studentId"))
	student := controller.StudentService.Update(c, studentId)
	studentResponse := helper.ToStudentResponse(student)
	return c.JSON(response.WebResponse{
		Message: "Success update student",
		Data:    studentResponse,
	})
}

func (controller *StudentControllerImpl) Delete(c *fiber.Ctx) error {
	studentId, _ := strconv.Atoi(c.Params("studentId"))
	controller.StudentService.Delete(c, studentId)
	return c.JSON(response.WebResponse{
		Message: "Success delete student",
	})
}
