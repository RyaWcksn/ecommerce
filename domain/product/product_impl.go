package product

import (
	"context"
	"time"

	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/errors"
)

// CreateProduct implements repositories.IProduct
func (p *ProductImpl) CreateProduct(ctx context.Context, entity *entities.CreateProductEntity) error {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := p.DB.Begin()
	if err != nil {
		p.log.Errorf("[ERR] While starting transaction := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	stmt, err := tx.PrepareContext(ctxDb, InsertProduct)
	if err != nil {
		p.log.Errorf("[ERR] While prepare statement := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	defer stmt.Close()

	row, err := stmt.Exec(entity.Name, entity.Description, entity.Price, entity.Seller)
	if err != nil {
		tx.Rollback()
		p.log.Errorf("[ERR] While executing query := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	id, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		p.log.Errorf("[ERR] While executing query := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}
	if id == 0 {
		tx.Rollback()
		p.log.Errorf("[ERR] While executing query := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	err = tx.Commit()
	if err != nil {
		p.log.Errorf("[ERR] While commit transaction := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	return nil

}
