package handlers

import (
	"net/http"
	"strconv"

	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/pkgs/datetime"
)

// GetProductListsHandler implements IHandler
func (h *HandlerImpl) GetProductListsHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	idStr := ctx.Value("id").(string)
	id, _ := strconv.Atoi(idStr)

	products, err := h.serviceImpl.GetProductsList(ctx, id)
	if err != nil {
		h.log.Errorf("[ERR] Error from service layer := %v", err)
		return err
	}
	resp := dto.ProductListResponse{
		Code:         http.StatusCreated,
		Message:      "ok",
		ResponseTime: datetime.GetDateString(),
		Products:     *products,
	}

	return ResponseJson(w, resp)
}
