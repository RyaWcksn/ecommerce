package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Token        string `json:"token"`
}
