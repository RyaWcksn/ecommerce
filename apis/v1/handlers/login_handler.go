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

// LoginHandler implements IHandler
func (h *HandlerImpl) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	payload := dto.LoginRequest{}
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

	token, err := h.serviceImpl.Login(ctx, &payload)
	if err != nil {
		h.log.Errorf("[ERR] While getting token from service layer := %v", err)
		return err
	}
	resp := dto.LoginResponse{
		Code:         http.StatusCreated,
		Message:      "ok",
		ResponseTime: datetime.GetDateString(),
		Token:        token,
	}

	return ResponseJson(w, resp)
}
