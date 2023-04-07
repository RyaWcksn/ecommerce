package repositories

import (
	"context"

	"github.com/RyaWcksn/ecommerce/entities"
)

//go:generate mockgen -source buyer.go -destination buyer_mock.go -package repositories
type IBuyer interface {
	GetEmail(ctx context.Context, email string) (resp *entities.LoginEntity, err error)
	GetData(ctx context.Context, id int) (resp *entities.BuyerEntity, err error)
}
