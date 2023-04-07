package services

import (
	"context"
	"testing"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/configs"
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
	cfg := configs.Cfg
	log := logger.New("", "", "")

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
		name     string
		args     args
		WantMock func()
		wantErr  bool
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
				buyerMock.EXPECT().GetEmail(gomock.Any(), gomock.Any()).Return(
					&entities.LoginEntity{
						Email:    "user@mail.com",
						Password: "$2a$10$xCniCdjZENytpQ7/NTNQduSJaZ6pl3bFMbd7bfF4OLwwbCyhGH8rC"},
					nil)
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.WantMock()
			s := NewServiceImpl().
				WithBuyer(buyerMock).
				WithSeller(sellerMock).
				WithOrder(orderMock).
				WithProduct(productMock).
				WithLog(log).WithConfig(*cfg)
			gotToken, err := s.Login(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken == "" {
				t.Errorf("ServiceImpl.Login()")
			}
		})
	}
}

func TestServiceImpl_CreateProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := repositories.NewMockIProduct(ctrl)
	log := logger.New("", "", "")

	cfg := configs.Cfg

	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "1")

	type args struct {
		ctx     context.Context
		payload *dto.CreateProductRequest
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
				ctx: ctx,
				payload: &dto.CreateProductRequest{
					Name:        "Dynames Gundam",
					Description: "HG Dynames Gundam from Kidou Senshi Gundam 00",
					Price:       "180000",
				},
			},
			wantMock: func() {
				productMock.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.wantMock()
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceImpl().
				WithProduct(productMock).
				WithLog(log).WithConfig(*cfg)
			if err := s.CreateProduct(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
