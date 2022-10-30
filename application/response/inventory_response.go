package response

type InventoryResponse struct {
	Id               int    `json:"id"`
	CategoryId       int    `json:"category_id"`
	CategoryName     string `json:"category_name"`
	DepositAdmin     int    `json:"deposit_admin"`
	DepositStudentId int    `json:"deposit_student_id"`
	DepositName      string `json:"deposit_name"`
	DepositTime      string `josn:"deposit_time"`
	ItemName         string `json:"item_name"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	TakeAdmin        int    `json:"take_admin"`
	TakeStudentId    int    `json:"take_student_id"`
	TakeName         string `json:"take_name"`
	TakeTime         string `json:"take_time"`
}
