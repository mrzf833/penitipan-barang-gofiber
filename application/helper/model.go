package helper

import (
	"gofiber-penitipan-barang/application/model"
	"gofiber-penitipan-barang/application/response"
)

func ToCategoryResponse(category model.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []model.Category) []response.CategoryResponse {
	var categoriesResponse []response.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}

func ToStudentResponse(student model.Student) response.StudentResponse {
	return response.StudentResponse{
		Id:          student.Id,
		Fullname:    student.Fullname,
		Email:       student.Email,
		PhoneNumber: student.PhoneNumber,
		Address:     student.Address,
	}
}

func ToStudentResponses(students []model.Student) []response.StudentResponse {
	var studentsResponse []response.StudentResponse
	for _, student := range students {
		studentsResponse = append(studentsResponse, ToStudentResponse(student))
	}
	return studentsResponse
}
