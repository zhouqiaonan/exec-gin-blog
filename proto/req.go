package proto

type (
	ReqEditUser struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     int    `json:"role"`
	}
)
