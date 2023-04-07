package dto

import "github.com/RyaWcksn/ecommerce/entities"

type CreateOrderRequest struct {
	ProductId []int `json:"productId"`
	SellerId  int   `json:"sellerId"`
}

type CreateOrderResponse struct {
	Code         int                  `json:"code"`
	Message      string               `json:"message"`
	ResponseTime string               `json:"responseTime"`
	Order        entities.OrderStatus `json:"order"`
}