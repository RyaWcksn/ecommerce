package seller

import (
	"context"
	"database/sql"
	"time"

	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/errors"
)

// GetEmail implements repositories.ISeller
func (s *SellerImpl) GetEmail(ctx context.Context, email string) (resp *entities.LoginEntity, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	payload := entities.LoginEntity{}
	err = s.DB.QueryRowContext(ctxDb, GetPasswordByEmailQuery, email).Scan(&payload.Email, &payload.Password)
	if err != nil {
		s.log.Errorf("[ERR] While getting email and password := %v", err)
		if err == sql.ErrNoRows {
			return nil, errors.GetError(errors.InvalidRequest, err)
		}
		return nil, errors.GetError(errors.InternalServer, err)
	}

	return &payload, nil
}
