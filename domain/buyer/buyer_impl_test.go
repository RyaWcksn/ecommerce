package buyer

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

func TestBuyerImpl_GetEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	log := logger.New("", "", "")

	query := GetPasswordByEmailQuery
	expectedRow := sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(1, "user@mail.com", "laksjd12kljlaksjv")
	mock.ExpectQuery(query).WithArgs("user@mail.com").WillReturnRows(expectedRow)

	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name     string
		args     args
		wantResp *entities.LoginEntity
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				ctx:   context.Background(),
				email: "user@mail.com",
			},
			wantResp: &entities.LoginEntity{
				Id:       1,
				Email:    "user@mail.com",
				Password: "laksjd12kljlaksjv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuyerImpl(db, log)
			gotResp, err := b.GetEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuyerImpl.GetEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("BuyerImpl.GetEmail() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestBuyerImpl_GetData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	log := logger.New("", "", "")

	query := GetBuyerDataById
	expectedRow := sqlmock.NewRows([]string{"name", "email", "alamat_pengiriman"}).AddRow("Arya", "user@mail.com", "Bandung")
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(expectedRow)
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name     string
		args     args
		wantResp *entities.BuyerEntity
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantResp: &entities.BuyerEntity{
				Name:             "Arya",
				Email:            "user@mail.com",
				AlamatPengiriman: "Bandung",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuyerImpl(db, log)
			gotResp, err := b.GetData(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuyerImpl.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("BuyerImpl.GetData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
