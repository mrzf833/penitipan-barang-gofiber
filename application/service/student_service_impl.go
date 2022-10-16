package service

import (
	"gofiber-penitipan-barang/application/model"
	"gofiber-penitipan-barang/application/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type StudentServiceImpl struct {
	DB       *gorm.DB
	validate *validator.Validate
}

func NewStudentService(DB *gorm.DB, validate *validator.Validate) StudentService {
	return &StudentServiceImpl{
		DB:       DB,
		validate: validate,
	}
}

func (service *StudentServiceImpl) FindAll(c *fiber.Ctx) []model.Student {
	var students []model.Student
	service.DB.Find(&students)
	return students
}

func (service *StudentServiceImpl) FindById(c *fiber.Ctx, studentId int) model.Student {
	var student model.Student
	service.DB.Find(&student, studentId)
	if student.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	return student
}

func (service *StudentServiceImpl) Create(c *fiber.Ctx) model.Student {
	studentCreateRequest := request.StudentCreateRequest{}

	err := c.BodyParser(&studentCreateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	student := model.Student{
		Fullname:    studentCreateRequest.Fullname,
		Email:       studentCreateRequest.Email,
		PhoneNumber: studentCreateRequest.PhoneNumber,
		Address:     studentCreateRequest.Address,
	}

	err = service.validate.Struct(student)

	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	service.DB.Create(&student)
	return student
}

func (service *StudentServiceImpl) Update(c *fiber.Ctx, studentId int) model.Student {
	studentUpdateRequest := request.StudentUpdateRequest{}

	err := c.BodyParser(&studentUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	student := model.Student{
		Id:          studentId,
		Fullname:    studentUpdateRequest.Fullname,
		Email:       studentUpdateRequest.Email,
		PhoneNumber: studentUpdateRequest.PhoneNumber,
		Address:     studentUpdateRequest.Address,
	}

	err = service.validate.Struct(student)

	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	tx := service.DB.First(&model.Student{}, studentId)
	if student.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	tx.Updates(&student)
	return student
}

func (service *StudentServiceImpl) Delete(c *fiber.Ctx, studentId int) {
	student := model.Student{}

	tx := service.DB.First(&student, studentId)
	if student.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	tx.Delete(&student)
}
