package buyer

import (
	"context"
	"time"

	"github.com/RyaWcksn/ecommerce/entities"
)

// GetEmail implements repositories.IBuyer
func (b *BuyerImpl) GetEmail(ctx context.Context, email string) (resp *entities.LoginEntity, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	payload := entities.LoginEntity{}
	err = b.DB.QueryRowContext(ctxDb, GetPasswordByEmailQuery, email).Scan(&payload.Email, &payload.Password)
	if err != nil {
		b.log.Errorf("[ERR] While getting email and password := %v", err)
		return nil, err
	}

	return &payload, nil
}
