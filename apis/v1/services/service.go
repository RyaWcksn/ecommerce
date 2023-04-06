package services

import (
	"context"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

type IService interface {
	Login(ctx context.Context, payload *dto.LoginRequest) (token string, err error)
}

type ServiceImpl struct {
	buyerImpl   repositories.IBuyer
	sellerImpl  repositories.ISeller
	orderImpl   repositories.IOrder
	productImpl repositories.IProduct
	log         logger.ILogger
}

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) WithBuyer(b repositories.IBuyer) *ServiceImpl {
	return &ServiceImpl{
		buyerImpl: b,
	}
}

func (s *ServiceImpl) WithSeller(sl repositories.ISeller) *ServiceImpl {
	return &ServiceImpl{
		sellerImpl: sl,
	}
}

func (s *ServiceImpl) WithOrder(o repositories.IOrder) *ServiceImpl {
	return &ServiceImpl{
		orderImpl: o,
	}
}

func (s *ServiceImpl) WithProduct(p repositories.IProduct) *ServiceImpl {
	return &ServiceImpl{
		productImpl: p,
	}
}

var _ IService = (*ServiceImpl)(nil)
