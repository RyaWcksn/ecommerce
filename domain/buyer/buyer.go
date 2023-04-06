package buyer

import (
	"database/sql"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

type BuyerImpl struct {
	DB  *sql.DB
	log logger.ILogger
}

func NewBuyerImpl(sql *sql.DB, log logger.ILogger) *BuyerImpl {
	return &BuyerImpl{
		DB:  sql,
		log: log,
	}
}

var _ repositories.IBuyer = (*BuyerImpl)(nil)
