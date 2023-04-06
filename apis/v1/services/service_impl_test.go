package services

import (
	"context"
	"testing"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
	"github.com/golang/mock/gomock"
)

func TestServiceImpl_Login(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerMock := repositories.NewMockIBuyer(ctrl)
	sellerMock := repositories.NewMockISeller(ctrl)
	productMock := repositories.NewMockIProduct(ctrl)
	orderMock := repositories.NewMockIOrder(ctrl)

	type fields struct {
		buyerImpl   repositories.IBuyer
		sellerImpl  repositories.ISeller
		orderImpl   repositories.IOrder
		productImpl repositories.IProduct
		log         logger.ILogger
	}
	type args struct {
		ctx     context.Context
		payload *dto.LoginRequest
	}
	tests := []struct {
		name      string
		args      args
		WantMock  func()
		wantToken string
		wantErr   bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: &dto.LoginRequest{
					Email:    "user@mail.com",
					Password: "password123",
					Role:     "buyer",
				},
			},
			WantMock: func() {
				buyerMock.EXPECT().GetEmail(gomock.Any(), gomock.Any()).Return(&entities.LoginEntity{Email: "user@mail.com", Password: "password123"}, nil)
			},

			wantToken: "Token Dummy",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceImpl().
				WithBuyer(buyerMock).
				WithSeller(sellerMock).
				WithOrder(orderMock).
				WithProduct(productMock)
			gotToken, err := s.Login(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("ServiceImpl.Login() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}
