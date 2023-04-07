package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/configs"
	"github.com/RyaWcksn/ecommerce/constants"
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

func TestServiceImpl_GetProductsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := repositories.NewMockIProduct(ctrl)
	log := logger.New("", "", "")

	cfg := configs.Cfg

	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "1")

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name            string
		args            args
		wantMock        func()
		wantProductList *[]entities.ProductListEntity
		wantErr         bool
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantMock: func() {
				productMock.EXPECT().ListProduct(gomock.Any(), gomock.Any()).Return(
					&[]entities.ProductListEntity{
						{
							Id:          1,
							ProductName: "HG Dynames Gundam",
							Description: "HG Dynames Gundam from Kidou Senshi Gundam 00",
							Price:       "180000",
							Seller:      1,
						},
						{
							Id:          2,
							ProductName: "HG Kyrios Gundam",
							Description: "HG Kyrios Gundam from Kidou Senshi Gundam 00",
							Price:       "180000",
							Seller:      1,
						},
						{
							Id:          3,
							ProductName: "HG Exia Gundam",
							Description: "HG Exia Gundam from Kidou Senshi Gundam 00",
							Price:       "180000",
							Seller:      1,
						},
						{
							Id:          4,
							ProductName: "HG Virtue Gundam",
							Description: "HG Virtue Gundam from Kidou Senshi Gundam 00",
							Price:       "180000",
							Seller:      1,
						},
					},
					nil,
				)
			},
			wantProductList: &[]entities.ProductListEntity{
				{
					Id:          1,
					ProductName: "HG Dynames Gundam",
					Description: "HG Dynames Gundam from Kidou Senshi Gundam 00",
					Price:       "180000",
					Seller:      1,
				},
				{
					Id:          2,
					ProductName: "HG Kyrios Gundam",
					Description: "HG Kyrios Gundam from Kidou Senshi Gundam 00",
					Price:       "180000",
					Seller:      1,
				},
				{
					Id:          3,
					ProductName: "HG Exia Gundam",
					Description: "HG Exia Gundam from Kidou Senshi Gundam 00",
					Price:       "180000",
					Seller:      1,
				},
				{
					Id:          4,
					ProductName: "HG Virtue Gundam",
					Description: "HG Virtue Gundam from Kidou Senshi Gundam 00",
					Price:       "180000",
					Seller:      1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.wantMock()
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceImpl().WithConfig(*cfg).WithLog(log).WithProduct(productMock)
			gotProductList, err := s.GetProductsList(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetProductsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProductList, tt.wantProductList) {
				t.Errorf("ServiceImpl.GetProductsList() = %v, want %v", gotProductList, tt.wantProductList)
			}
		})
	}
}

func TestServiceImpl_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := repositories.NewMockIProduct(ctrl)
	buyerMock := repositories.NewMockIBuyer(ctrl)
	sellerMock := repositories.NewMockISeller(ctrl)
	orderMock := repositories.NewMockIOrder(ctrl)
	log := logger.New("", "", "")
	cfg := configs.Cfg

	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "1")
	type args struct {
		ctx     context.Context
		payload *dto.CreateOrderRequest
	}
	tests := []struct {
		name     string
		args     args
		wantMock func()
		wantResp *entities.OrderStatus
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				payload: &dto.CreateOrderRequest{
					ProductId: []int{1},
					SellerId:  1,
				},
			},
			wantMock: func() {
				productMock.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(
					&entities.ProductListEntity{
						Id:          1,
						ProductName: "HG Dynames Gundam",
						Description: "HG Dynames Gundam from Kidou Senshi Gundam 00",
						Price:       "180000",
						Seller:      1,
					},
					nil,
				)
				buyerMock.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(
					&entities.BuyerEntity{
						Name:             "User",
						Email:            "user@mail.com",
						AlamatPengiriman: "Bandung",
					}, nil,
				)
				sellerMock.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(
					&entities.SellerEntity{
						Name:         "Arya",
						Email:        "arya@mail.com",
						AlamatPickup: "Jakarta",
					}, nil,
				)
				orderMock.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantResp: &entities.OrderStatus{
				Message: constants.PendingMessage,
				Status:  constants.Pending,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.wantMock()
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceImpl().
				WithBuyer(buyerMock).
				WithSeller(sellerMock).
				WithProduct(productMock).
				WithOrder(orderMock).
				WithLog(log).
				WithConfig(*cfg)
			gotResp, err := s.CreateOrder(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ServiceImpl.CreateOrder() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestServiceImpl_GetSellerOrderList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := repositories.NewMockIProduct(ctrl)
	buyerMock := repositories.NewMockIBuyer(ctrl)
	sellerMock := repositories.NewMockISeller(ctrl)
	orderMock := repositories.NewMockIOrder(ctrl)
	log := logger.New("", "", "")
	cfg := configs.Cfg

	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "1")

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name          string
		args          args
		wantMock      func()
		wantOrderList *[]entities.Order
		wantErr       bool
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
			},
			wantMock: func() {
				orderMock.EXPECT().SellerViewOrderList(gomock.Any(), gomock.Any()).
					Return(
						&[]entities.Order{
							{
								Id:                  1,
								Buyer:               2,
								Seller:              1,
								DeliverySource:      "Source",
								DeliveryDestination: "Destination",
								Items:               "Items",
								Quantity:            4,
								Price:               "180000",
								TotalPrice:          "180000",
								Status: entities.OrderStatus{
									Message: constants.PendingMessage,
									Status:  constants.Pending,
								},
							},
						}, nil)
			},
			wantOrderList: &[]entities.Order{
				{
					Id:                  1,
					Buyer:               2,
					Seller:              1,
					DeliverySource:      "Source",
					DeliveryDestination: "Destination",
					Items:               "Items",
					Quantity:            4,
					Price:               "180000",
					TotalPrice:          "180000",
					Status: entities.OrderStatus{
						Message: constants.PendingMessage,
						Status:  constants.Pending,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.wantMock()
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceImpl().
				WithBuyer(buyerMock).
				WithSeller(sellerMock).
				WithProduct(productMock).
				WithOrder(orderMock).
				WithLog(log).
				WithConfig(*cfg)
			gotOrderList, err := s.GetSellerOrderList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetSellerOrderList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrderList, tt.wantOrderList) {
				t.Errorf("ServiceImpl.GetSellerOrderList() = %v, want %v", gotOrderList, tt.wantOrderList)
			}
		})
	}
}

func TestServiceImpl_UpdateOrderStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := repositories.NewMockIProduct(ctrl)
	buyerMock := repositories.NewMockIBuyer(ctrl)
	sellerMock := repositories.NewMockISeller(ctrl)
	orderMock := repositories.NewMockIOrder(ctrl)
	log := logger.New("", "", "")
	cfg := configs.Cfg

	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "1")

	type args struct {
		ctx     context.Context
		payload *dto.AcceptOrderRequest
	}
	tests := []struct {
		name     string
		args     args
		wantMock func()
		wantResp *entities.Order
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				payload: &dto.AcceptOrderRequest{
					OrderId: 1,
				},
			},
			wantMock: func() {
				orderMock.EXPECT().UpdateOrder(gomock.Any(), gomock.Any()).Return(
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
			wantResp: &entities.Order{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.wantMock()
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceImpl().
				WithBuyer(buyerMock).
				WithSeller(sellerMock).
				WithProduct(productMock).
				WithOrder(orderMock).
				WithLog(log).
				WithConfig(*cfg)
			gotResp, err := s.UpdateOrderStatus(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.UpdateOrderStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ServiceImpl.UpdateOrderStatus() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestServiceImpl_GetBuyerOrderList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := repositories.NewMockIProduct(ctrl)
	buyerMock := repositories.NewMockIBuyer(ctrl)
	sellerMock := repositories.NewMockISeller(ctrl)
	orderMock := repositories.NewMockIOrder(ctrl)
	log := logger.New("", "", "")
	cfg := configs.Cfg

	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "1")
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name          string
		args          args
		wantMock      func()
		wantOrderList *[]entities.Order
		wantErr       bool
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
			},
			wantMock: func() {
				orderMock.EXPECT().BuyerViewOrderList(gomock.Any(), gomock.Any()).
					Return(
						&[]entities.Order{
							{
								Id:                  1,
								Buyer:               2,
								Seller:              1,
								DeliverySource:      "Source",
								DeliveryDestination: "Destination",
								Items:               "Items",
								Quantity:            4,
								Price:               "180000",
								TotalPrice:          "180000",
								Status: entities.OrderStatus{
									Message: constants.PendingMessage,
									Status:  constants.Pending,
								},
							},
						}, nil)
			},
			wantOrderList: &[]entities.Order{
				{
					Id:                  1,
					Buyer:               2,
					Seller:              1,
					DeliverySource:      "Source",
					DeliveryDestination: "Destination",
					Items:               "Items",
					Quantity:            4,
					Price:               "180000",
					TotalPrice:          "180000",
					Status: entities.OrderStatus{
						Message: constants.PendingMessage,
						Status:  constants.Pending,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.wantMock()
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceImpl().
				WithBuyer(buyerMock).
				WithSeller(sellerMock).
				WithProduct(productMock).
				WithOrder(orderMock).
				WithLog(log).
				WithConfig(*cfg)
			gotOrderList, err := s.GetBuyerOrderList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetBuyerOrderList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrderList, tt.wantOrderList) {
				t.Errorf("ServiceImpl.GetBuyerOrderList() = %v, want %v", gotOrderList, tt.wantOrderList)
			}
		})
	}
}
