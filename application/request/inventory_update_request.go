package request

type InventoryUpdateRequest struct {
	CategoryId       int    `json:"category_id" validate:"required,number"`
	DepositStudentId int    `json:"deposit_student_id"`
	DepositName      string `json:"deposit_name"`
	DepositTime      string `json:"deposit_time" validate:"required"`
	ItemName         string `json:"item_name" validate:"required"`
	Description      string `json:"description"`
	Status           string `json:"status" validate:"required"`
	TakeStudentId    int    `json:"take_student_id"`
	TakeName         string `json:"take_name"`
	TakeTime         string `json:"take_time"`
}
