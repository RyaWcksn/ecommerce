package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RyaWcksn/ecommerce/apis/v1/services"
	"github.com/RyaWcksn/ecommerce/constants"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
	"github.com/golang/mock/gomock"
)

func TestHandlerImpl_AcceptOrderHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := services.NewMockIService(ctrl)
	l := logger.New("", "", "")
	r := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantMock func()
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				w: r,
				r: func() *http.Request {
					req := httptest.NewRequest(
						http.MethodPost, constants.CreateOrderEndpoint, strings.NewReader(
							`{
							  "orderId": 1
							}`,
						),
					)
					ctx := context.Background()
					req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJleHAiOjE2ODA4NDA2MTAsImlkIjoxLCJyb2xlIjoic2VsbGVyIn0.BqlpOmIFBGlWZYKvoRqPhD8_q4Smwob0QV47vIVz0QU")
					ctx = context.WithValue(ctx, "id", "1")
					return req.WithContext(ctx)
				}(),
			},
			wantMock: func() {
				s.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any()).Return(
					&entities.Order{
						Id:                  1,
						Buyer:               2,
						Seller:              1,
						DeliverySource:      "Jakarta",
						DeliveryDestination: "Bandung",
						Items:               "HG Gundam Dynames",
						Quantity:            4,
						Price:               "180000",
						TotalPrice:          "180000",
						Status: entities.OrderStatus{
							Message: constants.AcceptedMessage,
							Status:  constants.Accepted,
						},
					}, nil,
				)
			},
		},
	}
	for _, tt := range tests {
		tt.wantMock()
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandlerImpl(s, l)
			if err := h.AcceptOrderHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.AcceptOrderHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
