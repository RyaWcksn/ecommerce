package seller

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

func TestSellerImpl_GetEmail(t *testing.T) {
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
			s := NewSellerImpl(db, log)
			gotResp, err := s.GetEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("SellerImpl.GetEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SellerImpl.GetEmail() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestSellerImpl_GetData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	log := logger.New("", "", "")

	query := GetSellerDataById
	expectedRow := sqlmock.NewRows([]string{"name", "email", "alamat_pickup"}).AddRow("Arya", "user@mail.com", "Jakarta")
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(expectedRow)
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name     string
		args     args
		wantResp *entities.SellerEntity
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantResp: &entities.SellerEntity{
				Name:         "Arya",
				Email:        "user@mail.com",
				AlamatPickup: "Jakarta",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSellerImpl(db, log)
			gotResp, err := s.GetData(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SellerImpl.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SellerImpl.GetData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
