package dto

type CreateOrderRequest struct {
	ProductId []int `json:"productId"`
	SellerId  int   `json:"sellerId"`
}

type CreateOrderResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Order        struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"order"`
}
