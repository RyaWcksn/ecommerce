package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Role     string `json:"role" validate:"required"`
}

type LoginResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Token        string `json:"token"`
}

type TokenGenerator struct {
	Id        int
	SecretKey string
	Email     string
	Role      string
}
