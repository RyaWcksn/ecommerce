package handlers

import (
	"net/http"
	"testing"

	"github.com/RyaWcksn/ecommerce/apis/v1/services"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

func TestHandlerImpl_GetProductsHandler(t *testing.T) {
	type fields struct {
		serviceImpl services.IService
		log         logger.ILogger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerImpl{
				serviceImpl: tt.fields.serviceImpl,
				log:         tt.fields.log,
			}
			if err := h.GetProductsHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.GetProductsHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
