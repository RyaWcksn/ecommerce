package seller

import (
	"database/sql"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

type SellerImpl struct {
	DB  *sql.DB
	log logger.ILogger
}

func NewSellerImpl(sql *sql.DB, log logger.ILogger) *SellerImpl {
	return &SellerImpl{
		DB:  sql,
		log: log,
	}
}

var _ repositories.ISeller = (*SellerImpl)(nil)
