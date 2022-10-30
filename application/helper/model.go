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

func ToAdminUserResponse(adminUser model.User) response.AdminUserResponse {
	return response.AdminUserResponse{
		Id:       adminUser.Id,
		Name:     adminUser.Name,
		Username: adminUser.Username,
	}
}

func ToAdminUserResponses(adminUsers []model.User) []response.AdminUserResponse {
	var adminUsersResponse []response.AdminUserResponse
	for _, adminUser := range adminUsers {
		adminUsersResponse = append(adminUsersResponse, ToAdminUserResponse(adminUser))
	}
	return adminUsersResponse
}

func ToInventoryResponse(inventory model.Inventory) response.InventoryResponse {
	inventoryResponse := response.InventoryResponse{
		Id:               inventory.Id,
		CategoryId:       int(inventory.CategoryId.Int64),
		CategoryName:     inventory.Category.Name,
		DepositAdmin:     inventory.DepositAdmin,
		DepositStudentId: int(inventory.DepositStudentId.Int64),
		DepositName:      inventory.DepositName.String,
		DepositTime:      inventory.DepositTime.Format("2006-01-02 15:04:05"),
		ItemName:         inventory.ItemName,
		Description:      inventory.Description.String,
		Status:           inventory.Status,
		TakeAdmin:        int(inventory.TakeAdmin.Int64),
		TakeStudentId:    inventory.TakeStudent.Id,
		TakeName:         inventory.TakeName.String,
	}
	if inventory.TakeTime.Valid {
		inventoryResponse.TakeTime = inventory.TakeTime.Time.Format("2006-01-02 15:04:05")
	}

	return inventoryResponse
}

func ToInventoryResponses(inventories []model.Inventory) []response.InventoryResponse {
	var inventoriesResponse []response.InventoryResponse
	for _, inventory := range inventories {
		inventoriesResponse = append(inventoriesResponse, ToInventoryResponse(inventory))
	}
	return inventoriesResponse
}
