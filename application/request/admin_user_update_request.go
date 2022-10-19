package request

type AdminUserUpdateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=225"`
	Username string `json:"username" validate:"required,min=1,max=225"`
}
