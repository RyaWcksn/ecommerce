package order

import (
	"database/sql"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

type OrderImpl struct {
	DB  *sql.DB
	log logger.ILogger
}

func NewOrderImpl(db *sql.DB, log logger.ILogger) *OrderImpl {
	return &OrderImpl{
		DB:  db,
		log: log,
	}
}

var _ repositories.IOrder = (*OrderImpl)(nil)
