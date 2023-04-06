package tokens

import (
	"time"

	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(payload *dto.LoginRequest) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = payload.Email
	claims["role"] = payload.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// generate encoded token and return
	return token.SignedString([]byte("your-secret-key"))
}
