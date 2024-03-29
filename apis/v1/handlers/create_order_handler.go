package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/pkgs/datetime"
	"github.com/RyaWcksn/ecommerce/pkgs/errors"
	"github.com/RyaWcksn/ecommerce/pkgs/validations"
)

// CreateOrderHandler implements IHandler
func (h *HandlerImpl) CreateOrderHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	payload := dto.CreateOrderRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("[ERR] While read body := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}
	if err = json.Unmarshal(body, &payload); err != nil {
		h.log.Errorf("[ERR] While unmarshal body := %v", err)
		return errors.GetError(errors.InternalServer, err)
	}

	if err = validations.Validate(payload); err != nil {
		h.log.Errorf("[ERR] While validating body := %v", err)
		return errors.GetError(errors.InvalidRequest, err)
	}

	status, err := h.serviceImpl.CreateOrder(ctx, &payload)
	if err != nil {
		h.log.Errorf("[ERR] From service layer := %v", err)
		return err
	}

	resp := dto.CreateOrderResponse{
		Code:         http.StatusCreated,
		Message:      "ok",
		ResponseTime: datetime.GetDateString(),
		Order:        *status,
	}

	return ResponseJson(w, resp)

}
