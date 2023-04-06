package services

import (
	"context"
	"fmt"

	"github.com/RyaWcksn/ecommerce/constants"
	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/errors"
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

	return info.Password, nil
}
