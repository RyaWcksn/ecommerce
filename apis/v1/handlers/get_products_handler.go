package handlers

import (
	"net/http"

	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/pkgs/datetime"
)

// GetProductsHandler implements IHandler
func (h *HandlerImpl) GetProductsHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	products, err := h.serviceImpl.GetProducts(ctx)
	if err != nil {
		h.log.Errorf("[ERR] Error from service layer := %v", err)
		return err
	}
	resp := dto.ProductListResponse{
		Code:         http.StatusOK,
		Message:      "ok",
		ResponseTime: datetime.GetDateString(),
		Products:     *products,
	}

	return ResponseJson(w, resp)
}
