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
	err = s.DB.QueryRowContext(ctxDb, GetPasswordByEmailQuery, email).Scan(&payload.Id, &payload.Email, &payload.Password)
	if err != nil {
		s.log.Errorf("[ERR] While getting email and password := %v", err)
		if err == sql.ErrNoRows {
			return nil, errors.GetError(errors.InvalidRequest, err)
		}
		return nil, errors.GetError(errors.InternalServer, err)
	}

	return &payload, nil
}

// GetData implements repositories.ISeller
func (s *SellerImpl) GetData(ctx context.Context, id int) (resp *entities.SellerEntity, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	payload := entities.SellerEntity{}
	err = s.DB.QueryRowContext(ctxDb, GetSellerDataById, id).Scan(&payload.Name, &payload.Email, &payload.AlamatPickup)
	if err != nil {
		s.log.Errorf("[ERR] While getting email and password := %v", err)
		if err == sql.ErrNoRows {
			return nil, errors.GetError(errors.InvalidRequest, err)
		}
		return nil, errors.GetError(errors.InternalServer, err)
	}

	return &payload, nil
}
