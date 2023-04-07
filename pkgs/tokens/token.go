package tokens

import (
	"time"

	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(payload *dto.TokenGenerator) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = payload.Id
	claims["authorized"] = true
	claims["email"] = payload.Email
	claims["role"] = payload.Role
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	// generate encoded token and return
	return token.SignedString([]byte(payload.SecretKey))
}
