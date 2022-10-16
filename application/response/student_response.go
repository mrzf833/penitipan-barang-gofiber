package response

type StudentResponse struct {
	Id          int    `json:"id"`
	Fullname    string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
