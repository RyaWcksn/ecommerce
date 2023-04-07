package buyer

import (
	"context"
	"database/sql"
	"time"

	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/errors"
)

// GetEmail implements repositories.IBuyer
func (b *BuyerImpl) GetEmail(ctx context.Context, email string) (resp *entities.LoginEntity, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	payload := entities.LoginEntity{}
	err = b.DB.QueryRowContext(ctxDb, GetPasswordByEmailQuery, email).Scan(&payload.Id, &payload.Email, &payload.Password)
	if err != nil {
		b.log.Errorf("[ERR] While getting email and password := %v", err)
		if err == sql.ErrNoRows {
			return nil, errors.GetError(errors.InvalidRequest, err)
		}
		return nil, errors.GetError(errors.InternalServer, err)
	}

	return &payload, nil
}
