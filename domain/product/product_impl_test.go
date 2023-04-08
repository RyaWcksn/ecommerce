package product

import (
	"context"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

func TestProductImpl_CreateProduct(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("Failed to create mock database connection: %s", err)
	// }
	// defer db.Close()
	// log := logger.New("", "", "")

	// mock.Expect()
	// mock.ExpectPrepare(InsertProduct).ExpectExec().WithArgs("Dyames Gundam", "HG Dynames Gundam from Kidou Senshi Gundam 00", "180000", 1).WillReturnResult(sqlmock.NewResult(1, 1))

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatal(err)
	// }

	// // Verify that all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Unfulfilled expectations: %s", err)
	// }

	// type args struct {
	// 	ctx    context.Context
	// 	entity *entities.CreateProductEntity
	// }
	// tests := []struct {
	// 	name    string
	// 	args    args
	// 	wantErr bool
	// }{
	// 	{
	// 		name: "success",
	// 		args: args{
	// 			ctx: context.Background(),
	// 			entity: &entities.CreateProductEntity{
	// 				Name:        "Gundam Dynames",
	// 				Description: "HG Dynames Gundam from Kidou Senshi Gundam 00",
	// 				Price:       "180000",
	// 				Seller:      1,
	// 			},
	// 		},
	// 		wantErr: false,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		p := NewProductImpl(db, log)
	// 		if err := p.CreateProduct(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
	// 			t.Errorf("ProductImpl.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
	// 		}
	// 	})
	// }
}

func TestProductImpl_ListProduct(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	log := logger.New("", "", "")

	query := regexp.QuoteMeta(GetProductsSeller)
	expectedRow := sqlmock.NewRows([]string{"id", "product_name", "description", "price", "seller"}).
		AddRow(1, "HG Dynames Gundam", "HG Dynames Gundam from Kidou Senshi Gundam 00", "180000", 1).
		AddRow(2, "HG Kyrios Gundam", "HG Kyrios Gundam from Kidou Senshi Gundam 00", "180000", 1).
		AddRow(3, "HG Exia Gundam", "HG Exia Gundam from Kidou Senshi Gundam 00", "180000", 1).
		AddRow(4, "HG Virtue Gundam", "HG Virtue Gundam from Kidou Senshi Gundam 00", "180000", 1)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(expectedRow)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name         string
		args         args
		wantProducts *[]entities.ProductListEntity
		wantErr      bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantProducts: &[]entities.ProductListEntity{
				{
					Id:          1,
					ProductName: "HG Dynames Gundam",
					Description: "HG Dynames Gundam from Kidou Senshi Gundam 00",
					Price:       180000.00,
					Seller:      1,
				},
				{
					Id:          2,
					ProductName: "HG Kyrios Gundam",
					Description: "HG Kyrios Gundam from Kidou Senshi Gundam 00",
					Price:       180000.00,
					Seller:      1,
				},
				{
					Id:          3,
					ProductName: "HG Exia Gundam",
					Description: "HG Exia Gundam from Kidou Senshi Gundam 00",
					Price:       180000.00,
					Seller:      1,
				},
				{
					Id:          4,
					ProductName: "HG Virtue Gundam",
					Description: "HG Virtue Gundam from Kidou Senshi Gundam 00",
					Price:       180000.00,
					Seller:      1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProductImpl(db, log)
			gotProducts, err := p.ListProduct(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductImpl.ListProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProducts, tt.wantProducts) {
				t.Errorf("ProductImpl.ListProduct() = %v, want %v", gotProducts, tt.wantProducts)
			}
		})
	}
}

func TestProductImpl_GetProductById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	log := logger.New("", "", "")

	query := GetProductById
	expectedRow := sqlmock.NewRows([]string{"id", "product_name", "description", "price", "seller"}).AddRow(1, "HG Dynames Gundam", "HG Dynames Gundam from Kidou Senshi Gundam 00", "180000", 1)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(expectedRow)
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name        string
		args        args
		wantProduct *entities.ProductListEntity
		wantErr     bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantProduct: &entities.ProductListEntity{
				Id:          1,
				ProductName: "HG Dynames Gundam",
				Description: "HG Dynames Gundam from Kidou Senshi Gundam 00",
				Price:       180000.00,
				Seller:      1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProductImpl(db, log)
			gotProduct, err := p.GetProductById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductImpl.GetProductById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("ProductImpl.GetProductById() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func TestProductImpl_GetAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	log := logger.New("", "", "")

	query := regexp.QuoteMeta(GetAllProducts)
	expectedRow := sqlmock.NewRows([]string{"id", "product_name", "description", "price", "seller"}).
		AddRow(1, "HG Dynames Gundam", "HG Dynames Gundam from Kidou Senshi Gundam 00", "180000", 1).
		AddRow(2, "HG Kyrios Gundam", "HG Kyrios Gundam from Kidou Senshi Gundam 00", "180000", 1).
		AddRow(3, "HG Exia Gundam", "HG Exia Gundam from Kidou Senshi Gundam 00", "180000", 1).
		AddRow(4, "HG Virtue Gundam", "HG Virtue Gundam from Kidou Senshi Gundam 00", "180000", 1)
	mock.ExpectQuery(query).WillReturnRows(expectedRow)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name         string
		args         args
		wantProducts *[]entities.ProductListEntity
		wantErr      bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
			},
			wantProducts: &[]entities.ProductListEntity{
				{
					Id:          1,
					ProductName: "HG Dynames Gundam",
					Description: "HG Dynames Gundam from Kidou Senshi Gundam 00",
					Price:       180000.00,
					Seller:      1,
				},
				{
					Id:          2,
					ProductName: "HG Kyrios Gundam",
					Description: "HG Kyrios Gundam from Kidou Senshi Gundam 00",
					Price:       180000.00,
					Seller:      1,
				},
				{
					Id:          3,
					ProductName: "HG Exia Gundam",
					Description: "HG Exia Gundam from Kidou Senshi Gundam 00",
					Price:       180000.00,
					Seller:      1,
				},
				{
					Id:          4,
					ProductName: "HG Virtue Gundam",
					Description: "HG Virtue Gundam from Kidou Senshi Gundam 00",
					Price:       180000.00,
					Seller:      1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProductImpl(db, log)
			gotProducts, err := p.GetAllProducts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductImpl.ListProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProducts, tt.wantProducts) {
				t.Errorf("ProductImpl.ListProduct() = %v, want %v", gotProducts, tt.wantProducts)
			}
		})
	}
}
