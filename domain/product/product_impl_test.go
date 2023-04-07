package product

import (
	"testing"
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
