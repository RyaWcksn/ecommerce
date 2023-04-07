package dto

import "github.com/RyaWcksn/ecommerce/entities"

type ProductListResponse struct {
	Code         int                          `json:"code"`
	Message      string                       `json:"message"`
	ResponseTime string                       `json:"responseTime"`
	Products     []entities.ProductListEntity `json:"products"`
}
