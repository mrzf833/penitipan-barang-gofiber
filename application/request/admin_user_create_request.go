package request

type AdminUserCreateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Username string `json:"username" validate:"required,min=1,max=225"`
	Password string `json:"password" validate:"required,min=1,max=225"`
}
