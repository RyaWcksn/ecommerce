package product

import (
	"database/sql"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

type ProductImpl struct {
	DB  *sql.DB
	log logger.ILogger
}

func NewProductImpl(db *sql.DB, l logger.ILogger) *ProductImpl {
	return &ProductImpl{
		DB:  db,
		log: l,
	}
}

var _ repositories.IProduct = (*ProductImpl)(nil)
