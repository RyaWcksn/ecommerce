package handlers

import (
	"net/http"

	"github.com/RyaWcksn/ecommerce/apis/v1/services"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

type IHandler interface {
	LoginHandler(w http.ResponseWriter, r *http.Response) error
}

type HandlerImpl struct {
	serviceImpl services.IService
	log         logger.ILogger
}
