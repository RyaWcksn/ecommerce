package handlers

import (
	"net/http"

	"github.com/RyaWcksn/ecommerce/apis/v1/services"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

type IHandler interface {
	LoginHandler(w http.ResponseWriter, r *http.Request) error

	CreateProductHandler(w http.ResponseWriter, r *http.Request) error
}

type HandlerImpl struct {
	serviceImpl services.IService
	log         logger.ILogger
}

func NewHandlerImpl(s services.IService, l logger.ILogger) *HandlerImpl {
	return &HandlerImpl{
		serviceImpl: s,
		log:         l,
	}
}

var _ IHandler = (*HandlerImpl)(nil)
