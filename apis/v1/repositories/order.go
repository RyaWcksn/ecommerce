package repositories

import (
	"context"

	"github.com/RyaWcksn/ecommerce/entities"
)

//go:generate mockgen -source order.go -destination order_mock.go -package repositories
type IOrder interface {
	CreateOrder(ctx context.Context, entity *entities.CreateOrder) error
	SellerViewOrderList(ctx context.Context, id int) (order *[]entities.Order, err error)
}
