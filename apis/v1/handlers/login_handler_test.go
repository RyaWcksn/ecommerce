package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RyaWcksn/ecommerce/apis/v1/services"
	"github.com/RyaWcksn/ecommerce/constants"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
	"github.com/golang/mock/gomock"
)

func TestHandlerImpl_LoginHandler(t *testing.T) {
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
						http.MethodPost, constants.LoginEndpoint, strings.NewReader(
							`{
							  "email": "user@mail.com",
							  "password": "password123",
							  "role": "buyer"
							}`,
						),
					)
					ctx := context.Background()
					return req.WithContext(ctx)
				}(),
			},
			wantMock: func() {
				s.EXPECT().Login(gomock.Any(), gomock.Any()).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJleHAiOjE2ODA4NTIzOTQsInJvbGUiOiJidXllciJ9.XTKKaRHj1hVNoz1oPb2ws6lZcvksKKGEzFpKlIA3C3o", nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			h := NewHandlerImpl(s, l)
			if err := h.LoginHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.LoginHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
