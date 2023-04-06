package services

import (
	"context"
	"fmt"

	"github.com/RyaWcksn/ecommerce/constants"
	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/errors"
	"github.com/RyaWcksn/ecommerce/pkgs/tokens"
	"golang.org/x/crypto/bcrypt"
)

// Login implements IService
func (s *ServiceImpl) Login(ctx context.Context, payload *dto.LoginRequest) (token string, err error) {

	var info = &entities.LoginEntity{}
	switch payload.Role {
	case constants.BUYER:
		info, err = s.buyerImpl.GetEmail(ctx, payload.Email)
	case constants.SELLER:
		return "", errors.GetError(errors.InvalidRequest, fmt.Errorf("[ERR] %v", "Not implemented"))
	}
	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(payload.Password)); err != nil {
		s.log.Errorf("[ERR] While compare password with hash", err)
		return "", errors.GetError(errors.InvalidRequest, err)
	}

	token, err = tokens.GenerateJWT(payload)
	if err != nil {
		s.log.Errorf("[ERR] While generating token", err)
		return "", errors.GetError(errors.InternalServer, err)
	}

	return token, nil
}
