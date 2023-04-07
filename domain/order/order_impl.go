package order

import (
	"context"
	"time"

	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/errors"
)

// CreateOrder implements repositories.IOrder
func (o *OrderImpl) CreateOrder(ctx context.Context, entity *entities.CreateOrder) error {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := o.DB.Begin()
	if err != nil {
		o.log.Errorf("[ERR] While starting transaction := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	stmt, err := tx.PrepareContext(ctxDb, InsertOrder)
	if err != nil {
		o.log.Errorf("[ERR] While prepare statement := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	defer stmt.Close()

	row, err := stmt.Exec(entity.Buyer, entity.Seller, entity.DeliverySource, entity.DeliveryDestination, entity.Items, entity.Quantity, entity.Price, entity.TotalPrice, entity.Status)
	if err != nil {
		tx.Rollback()
		o.log.Errorf("[ERR] While executing query := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	id, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		o.log.Errorf("[ERR] While executing query := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}
	if id == 0 {
		tx.Rollback()
		o.log.Errorf("[ERR] While executing query := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	err = tx.Commit()
	if err != nil {
		o.log.Errorf("[ERR] While commit transaction := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	return nil
}
