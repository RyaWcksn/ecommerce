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
