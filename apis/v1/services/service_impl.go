package services

import (
	"context"

	"github.com/RyaWcksn/ecommerce/dto"
)

// Login implements IService
func (s *ServiceImpl) Login(ctx context.Context, payload *dto.LoginRequest) (token string, err error) {
	panic("unimplemented")
}
