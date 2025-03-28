package models

type LoginRequest struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Cookie  string `json:"cookie,omitempty"`
}
