package product

import (
	"context"
	"database/sql"
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

// ListProduct implements repositories.IProduct
func (p *ProductImpl) ListProduct(ctx context.Context, id int) (products *[]entities.ProductListEntity, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var productsList []entities.ProductListEntity

	rows, err := p.DB.QueryContext(ctxDb, GetProductsSeller, id)
	if err != nil {
		p.log.Errorf("[ERR] While query the data := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}
	defer rows.Close()

	for rows.Next() {
		var product entities.ProductListEntity

		if err := rows.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Seller); err != nil {
			p.log.Errorf("[ERR] While query the data := %v", err)
			return nil, errors.GetError(errors.InternalServer, err)
		}

		productsList = append(productsList, product)

	}
	return &productsList, nil
}

// GetProductById implements repositories.IProduct
func (p *ProductImpl) GetProductById(ctx context.Context, id int) (product *entities.ProductListEntity, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	payload := entities.ProductListEntity{}
	err = p.DB.QueryRowContext(ctxDb, GetProductById, id).Scan(&payload.Id, &payload.ProductName, &payload.Description, &payload.Price, &payload.Seller)
	if err != nil {
		p.log.Errorf("[ERR] While getting email and password := %v", err)
		if err == sql.ErrNoRows {
			return nil, errors.GetError(errors.InvalidRequest, err)
		}
		return nil, errors.GetError(errors.InternalServer, err)
	}

	return &payload, nil
}

// GetAllProducts implements repositories.IProduct
func (p *ProductImpl) GetAllProducts(ctx context.Context) (productList *[]entities.ProductListEntity, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var productsList []entities.ProductListEntity

	rows, err := p.DB.QueryContext(ctxDb, GetAllProducts)
	if err != nil {
		p.log.Errorf("[ERR] While query the data := %v", err)
		return nil, errors.GetError(errors.InternalServer, err)
	}
	defer rows.Close()

	for rows.Next() {
		var product entities.ProductListEntity

		if err := rows.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Seller); err != nil {
			p.log.Errorf("[ERR] While query the data := %v", err)
			return nil, errors.GetError(errors.InternalServer, err)
		}

		productsList = append(productsList, product)

	}
	return &productsList, nil
}
