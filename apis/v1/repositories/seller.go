package repositories

import (
	"context"

	"github.com/RyaWcksn/ecommerce/entities"
)

//go:generate mockgen -source seller.go -destination seller_mock.go -package repositories
type ISeller interface {
	GetEmail(ctx context.Context, email string) (resp *entities.LoginEntity, err error)
	GetData(ctx context.Context, id int) (resp *entities.SellerEntity, err error)
}
