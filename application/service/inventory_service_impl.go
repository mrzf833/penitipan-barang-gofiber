package service

import (
	"database/sql"
	"gofiber-penitipan-barang/application/middleware"
	"gofiber-penitipan-barang/application/model"
	"gofiber-penitipan-barang/application/request"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InventoryServiceImpl struct {
	DB       *gorm.DB
	validate *validator.Validate
}

func NewInventoryService(DB *gorm.DB, validate *validator.Validate) InventoryService {
	return &InventoryServiceImpl{
		DB:       DB,
		validate: validate,
	}
}

func (service *InventoryServiceImpl) FindAll(c *fiber.Ctx) []model.Inventory {
	var inventories []model.Inventory
	service.DB.Preload("Category").Preload("DepositStudent").Preload("TakeStudent").Find(&inventories)
	return inventories
}

func (service *InventoryServiceImpl) FindById(c *fiber.Ctx, inventoryId int) model.Inventory {
	var inventory model.Inventory
	service.DB.Preload("Category").Preload("DepositStudent").Preload("TakeStudent").Find(&inventory, inventoryId)
	if inventory.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	return inventory
}

func (service *InventoryServiceImpl) Create(c *fiber.Ctx) model.Inventory {
	inventoryCreateRequest := request.InventoryCreateRequest{}

	err := c.BodyParser(&inventoryCreateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	err = service.validate.Struct(inventoryCreateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	inventory := model.Inventory{
		ItemName: inventoryCreateRequest.ItemName,
		Description: sql.NullString{
			String: inventoryCreateRequest.Description,
			Valid:  true,
		},
		DepositTime:  time.Now(),
		Status:       "deposit",
		DepositAdmin: middleware.AuthenticationGetUser.Id,
	}

	category := model.Category{}
	service.DB.Find(&category, inventoryCreateRequest.CategoryId)
	if category.Id == 0 {
		panic(fiber.NewError(404, "category not found"))
	}
	inventory.CategoryId = sql.NullInt64{
		Int64: int64(inventoryCreateRequest.CategoryId),
		Valid: true,
	}

	if inventoryCreateRequest.DepositStudentId != 0 {
		depositStudent := model.Student{}
		service.DB.Find(&depositStudent, inventoryCreateRequest.DepositStudentId)
		if depositStudent.Id == 0 {
			panic(fiber.NewError(404, "deposit student not found"))
		}
		inventory.DepositStudentId.Int64 = int64(depositStudent.Id)
		inventory.DepositStudentId.Valid = true
		inventory.DepositName.String = depositStudent.Fullname
		inventory.DepositName.Valid = true
	} else {
		inventory.DepositName.String = inventoryCreateRequest.DepositName
		inventory.DepositName.Valid = true
	}

	service.DB.Create(&inventory)
	service.DB.Preload("Category").Find(&inventory)
	return inventory
}

func (service *InventoryServiceImpl) Update(c *fiber.Ctx, inventoryId int) model.Inventory {
	inventoryUpdateRequest := request.InventoryUpdateRequest{}

	err := c.BodyParser(&inventoryUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	err = service.validate.Struct(inventoryUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	inventory := model.Inventory{
		Id:       inventoryId,
		ItemName: inventoryUpdateRequest.ItemName,
		Description: sql.NullString{
			String: inventoryUpdateRequest.Description,
			Valid:  true,
		},
		Status: inventoryUpdateRequest.Status,
	}

	// pengecekan inventory id
	cekInventory := model.Inventory{}
	service.DB.First(&cekInventory, inventoryId)
	if cekInventory.Id == 0 {
		panic(fiber.NewError(404, "inventory not found"))
	}

	category := model.Category{}
	service.DB.Find(&category, inventoryUpdateRequest.CategoryId)
	if category.Id == 0 {
		panic(fiber.NewError(404, "category not found"))
	}
	inventory.CategoryId = sql.NullInt64{
		Int64: int64(inventoryUpdateRequest.CategoryId),
		Valid: true,
	}

	// untuk deposit student id dan deposit name
	if inventoryUpdateRequest.DepositStudentId != 0 {
		depositStudent := model.Student{}
		service.DB.Find(&depositStudent, inventoryUpdateRequest.DepositStudentId)
		if depositStudent.Id == 0 {
			panic(fiber.NewError(404, "deposit student not found"))
		}
		inventory.DepositStudentId.Int64 = int64(depositStudent.Id)
		inventory.DepositStudentId.Valid = true
		inventory.DepositName.String = depositStudent.Fullname
		inventory.DepositName.Valid = true
	} else {
		inventory.DepositName.String = inventoryUpdateRequest.DepositName
		inventory.DepositName.Valid = true
	}

	inventory.DepositTime, err = time.Parse("2006-01-02 15:04:05", inventoryUpdateRequest.DepositTime)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	// untuk take student id dan take name
	if inventoryUpdateRequest.TakeStudentId != 0 {
		takeStudent := model.Student{}
		service.DB.Find(&takeStudent, inventoryUpdateRequest.TakeStudentId)
		if takeStudent.Id == 0 {
			panic(fiber.NewError(404, "take student not found"))
		}
		inventory.TakeStudentId.Int64 = int64(takeStudent.Id)
		inventory.TakeStudentId.Valid = true
		inventory.TakeName.String = takeStudent.Fullname
		inventory.TakeName.Valid = true
	} else {
		inventory.TakeName.String = inventoryUpdateRequest.TakeName
		inventory.TakeName.Valid = true
	}

	if inventoryUpdateRequest.TakeTime != "" {
		takeTime, err := time.Parse("2006-01-02 15:04:05", inventoryUpdateRequest.TakeTime)
		if err != nil {
			panic(fiber.NewError(500, err.Error()))
		}
		inventory.TakeTime = sql.NullTime{
			Time:  takeTime,
			Valid: true,
		}
	}

	tx := service.DB.First(&model.Inventory{}, inventoryId)
	if inventory.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	tx.Updates(&inventory)
	tx.Preload("Category").Find(&inventory)
	return inventory
}

func (service *InventoryServiceImpl) Delete(c *fiber.Ctx, inventoryId int) {
	inventory := model.Inventory{}

	tx := service.DB.First(&inventory, inventoryId)
	if inventory.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	tx.Delete(&inventory)
}

func (service *InventoryServiceImpl) TakeItem(c *fiber.Ctx, inventoryId int) model.Inventory {
	inventoryUpdateRequest := request.InventoryUpdateRequest{}

	err := c.BodyParser(&inventoryUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	err = service.validate.Struct(inventoryUpdateRequest)
	if err != nil {
		panic(fiber.NewError(500, err.Error()))
	}

	inventory := model.Inventory{
		Id:     inventoryId,
		Status: "take",
		TakeTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	// pengecekan inventory id
	cekInventory := model.Inventory{}
	service.DB.First(&cekInventory, inventoryId)
	if cekInventory.Id == 0 {
		panic(fiber.NewError(404, "inventory not found"))
	}

	// untuk take student id dan take name
	if inventoryUpdateRequest.TakeStudentId != 0 {
		takeStudent := model.Student{}
		service.DB.Find(&takeStudent, inventoryUpdateRequest.TakeStudentId)
		inventory.TakeStudentId.Int64 = int64(takeStudent.Id)
		inventory.TakeStudentId.Valid = true
		inventory.TakeName.String = takeStudent.Fullname
		inventory.TakeName.Valid = true
	} else {
		inventory.TakeName.String = inventoryUpdateRequest.TakeName
		inventory.TakeName.Valid = true
	}

	tx := service.DB.First(&model.Inventory{}, inventoryId)
	if inventory.Id == 0 {
		panic(fiber.NewError(404, "data not found"))
	}
	tx.Updates(&inventory)
	tx.Preload("Category").Find(&inventory)
	return inventory
}
