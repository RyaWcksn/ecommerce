package handlers

import (
	"net/http"

	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/pkgs/datetime"
)

// GetSellerOrdersHandler implements IHandler
func (h *HandlerImpl) GetSellerOrdersHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	orders, err := h.serviceImpl.GetSellerOrderList(ctx)
	if err != nil {
		h.log.Errorf("[ERR] Error from service layer := %v", err)
		return err
	}

	resp := dto.OrdersResponse{
		Code:         http.StatusOK,
		Message:      "ok",
		ResponseTime: datetime.GetDateString(),
		Orders:       *orders,
	}

	return ResponseJson(w, resp)

}
