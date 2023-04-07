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

// CreateProductHandler implements IHandler
func (h *HandlerImpl) CreateProductHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	payload := dto.CreateProductRequest{}
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

	err = h.serviceImpl.CreateProduct(ctx, &payload)
	if err != nil {
		return err
	}

	resp := dto.CreateProductResponse{
		Code:         http.StatusOK,
		Message:      "ok",
		ResponseTime: datetime.GetDateString(),
		Product:      payload,
	}

	return ResponseJson(w, resp)
}
