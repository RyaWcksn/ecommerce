package order

import (
	"context"
	"time"

	"github.com/RyaWcksn/ecommerce/constants"
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

// SellerViewOrderList implements repositories.IOrder
func (o *OrderImpl) SellerViewOrderList(ctx context.Context, id int) (order *[]entities.Order, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	msg := constants.PendingMessage
	stts := constants.Pending

	var orderList []entities.Order

	rows, err := o.DB.QueryContext(ctxDb, GetBySellerId, id)
	if err != nil {
		o.log.Errorf("[ERR] While query the data := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}
	defer rows.Close()

	for rows.Next() {
		var order entities.Order
		var status int64

		if err := rows.Scan(
			&order.Id,
			&order.Buyer,
			&order.Seller,
			&order.DeliverySource,
			&order.DeliveryDestination,
			&order.Items,
			&order.Quantity,
			&order.Price,
			&order.TotalPrice,
			&status,
		); err != nil {
			o.log.Errorf("[ERR] While query the data := %v", err)
			return nil, errors.GetError(errors.InternalServer, err)
		}
		order.Status.Message = msg
		order.Status.Status = stts
		if status > 0 {
			order.Status.Message = constants.AcceptedMessage
			order.Status.Status = constants.Accepted
		}

		orderList = append(orderList, order)

	}
	return &orderList, nil
}

// UpdateOrder implements repositories.IOrder
func (o *OrderImpl) UpdateOrder(ctx context.Context, id int) (order *entities.Order, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	msg := constants.PendingMessage
	stts := constants.Pending

	tx, err := o.DB.Begin()
	if err != nil {
		o.log.Errorf("[ERR] While starting transaction := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}

	stmt, err := tx.PrepareContext(ctxDb, UpdateStatusOrder)
	if err != nil {
		o.log.Errorf("[ERR] While prepare statement := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctxDb, 1, id)
	if err != nil {
		tx.Rollback()
		o.log.Errorf("[ERR] While executing query := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}

	err = tx.Commit()
	if err != nil {
		o.log.Errorf("[ERR] While commit transaction := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}

	var orderData entities.Order
	var status int
	err = o.DB.QueryRowContext(ctxDb, GetById, id).Scan(
		&orderData.Id,
		&orderData.Buyer,
		&orderData.Seller,
		&orderData.DeliverySource,
		&orderData.DeliveryDestination,
		&orderData.Items,
		&orderData.Quantity,
		&orderData.Price,
		&orderData.TotalPrice,
		&status,
	)
	if err != nil {
		o.log.Errorf("[ERR] While getting new data := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}

	orderData.Status.Message = msg
	orderData.Status.Status = stts
	if status > 0 {
		orderData.Status.Message = constants.AcceptedMessage
		orderData.Status.Status = constants.Accepted
	}

	return &orderData, nil
}

// BuyerViewOrderList implements repositories.IOrder
func (o *OrderImpl) BuyerViewOrderList(ctx context.Context, id int) (order *[]entities.Order, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	msg := constants.PendingMessage
	stts := constants.Pending

	var orderList []entities.Order

	rows, err := o.DB.QueryContext(ctxDb, GetByBuyerId, id)
	if err != nil {
		o.log.Errorf("[ERR] While query the data := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}
	defer rows.Close()

	for rows.Next() {
		var order entities.Order
		var status int64

		if err := rows.Scan(
			&order.Id,
			&order.Buyer,
			&order.Seller,
			&order.DeliverySource,
			&order.DeliveryDestination,
			&order.Items,
			&order.Quantity,
			&order.Price,
			&order.TotalPrice,
			&status,
		); err != nil {
			o.log.Errorf("[ERR] While query the data := %v", err)
			return nil, errors.GetError(errors.InternalServer, err)
		}
		order.Status.Message = msg
		order.Status.Status = stts
		if status > 0 {
			order.Status.Message = constants.AcceptedMessage
			order.Status.Status = constants.Accepted
		}

		orderList = append(orderList, order)

	}
	return &orderList, nil
}
