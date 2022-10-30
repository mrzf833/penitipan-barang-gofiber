package request

type InventoryTakeItemRequest struct {
	TakeStudentId int    `json:"take_student_id"`
	TakeName      string `json:"take_name"`
}
