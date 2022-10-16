package request

type StudentUpdateRequest struct {
	Fullname    string `json:"fullname" validate:"required,min:1,max:225"`
	Email       string `json:"email" validate:"required,min:1,max:225"`
	PhoneNumber string `json:"phone_number" validate:"required,min:1,max:15"`
	Address     string `json:"address" validate:"omitempty"`
}
