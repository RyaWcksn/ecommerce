package dto

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type CreateProductResponse struct {
	Code         int                  `json:"code"`
	Message      string               `json:"message"`
	ResponseTime string               `json:"responseTime"`
	Product      CreateProductRequest `json:"product"`
}
