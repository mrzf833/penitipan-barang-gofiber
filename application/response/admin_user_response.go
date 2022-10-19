package response

type AdminUserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
