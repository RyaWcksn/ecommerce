package repositories

import (
	"context"

	"github.com/RyaWcksn/ecommerce/entities"
)

//go:generate mockgen -source product.go -destination product_mock.go -package repositories
type IProduct interface {
	CreateProduct(ctx context.Context, entity *entities.CreateProductEntity) error
}
