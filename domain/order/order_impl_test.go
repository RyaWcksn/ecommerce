package order

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/ecommerce/constants"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

func TestOrderImpl_CreateOrder(t *testing.T) {
	type fields struct {
		DB  *sql.DB
		log logger.ILogger
	}
	type args struct {
		ctx    context.Context
		entity *entities.CreateOrder
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
			o := &OrderImpl{
				DB:  tt.fields.DB,
				log: tt.fields.log,
			}
			if err := o.CreateOrder(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("OrderImpl.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrderImpl_SellerViewOrderList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	log := logger.New("", "", "")

	query := regexp.QuoteMeta(GetBySellerId)
	rows := sqlmock.NewRows([]string{"id", "buyer", "seller", "delivery_source", "delivery_destination", "items", "quantity", "price", "total_price", "status"}).
		AddRow(1, 2, 1, "Source", "Destination", "Items", 4, "180000", "180000", 0)

	// Set up the mock expectations
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name      string
		args      args
		wantOrder *[]entities.Order
		wantErr   bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantOrder: &[]entities.Order{
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
		t.Run(tt.name, func(t *testing.T) {
			o := NewOrderImpl(db, log)
			gotOrder, err := o.SellerViewOrderList(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderImpl.SellerViewOrderList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("OrderImpl.SellerViewOrderList() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
