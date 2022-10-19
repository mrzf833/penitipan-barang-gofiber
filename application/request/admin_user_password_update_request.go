package request

type AdminUserPasswordUpdateRequest struct {
	Password string `json:"password" validate:"required,min=1,max=225"`
}
