package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/RyaWcksn/ecommerce/configs"
	"github.com/RyaWcksn/ecommerce/constants"
	"github.com/RyaWcksn/ecommerce/pkgs/errors"
	"github.com/golang-jwt/jwt/v4"
)

type SellerNext func(http.ResponseWriter, *http.Request) error

func SellerMiddleware(cfg configs.Config, handler SellerNext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			xerr := errors.ErrorForm{
				Code:     http.StatusUnauthorized,
				Message:  "Unauthorized",
				Response: "Token is mandatory",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			// Return the secret key used to sign the token
			return []byte(cfg.App.SECRET), nil
		})
		if err != nil {
			xerr := errors.ErrorForm{
				Code:     http.StatusUnauthorized,
				Message:  "Unauthorized",
				Response: "Token is invalid",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			xerr := errors.ErrorForm{
				Code:     http.StatusUnauthorized,
				Message:  "Unauthorized",
				Response: "Token is not valid",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
		userRole, ok := claims["role"].(string)
		if !ok || userRole != constants.SELLER {
			xerr := errors.ErrorForm{
				Code:     http.StatusForbidden,
				Message:  "Forbidden",
				Response: "Role is not allowed to access this resource",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}

		id, ok := claims["id"].(float64)
		s := strconv.FormatFloat(id, 'f', -1, 64)

		ctx = context.WithValue(ctx, "id", s)
		if err := handler(w, r.WithContext(ctx)); err != nil {
			xerr := err.(*errors.ErrorForm)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(xerr.Code)
			json.NewEncoder(w).Encode(xerr)
		}

	})
}

func (fn SellerNext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			xerr := errors.ErrorForm{
				Code:     500,
				Message:  "Panic",
				Response: "Error",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(xerr.Code)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
	}()
	if err := fn(w, r); err != nil {
		xerr := err.(*errors.ErrorForm)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(xerr.Code)
		json.NewEncoder(w).Encode(xerr)
	}
}
