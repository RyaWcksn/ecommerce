package services

import (
	"context"

	"github.com/RyaWcksn/ecommerce/apis/v1/repositories"
	"github.com/RyaWcksn/ecommerce/configs"
	"github.com/RyaWcksn/ecommerce/dto"
	"github.com/RyaWcksn/ecommerce/entities"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
)

//go:generate mockgen -source service.go -destination service_mock.go -package services
type IService interface {
	Login(ctx context.Context, payload *dto.LoginRequest) (token string, err error)

	// Seller
	CreateProduct(ctx context.Context, payload *dto.CreateProductRequest) error
	GetProductsList(ctx context.Context, id int) (productList *[]entities.ProductListEntity, err error)

	// Buyer
	CreateOrder(ctx context.Context, payload *dto.CreateOrderRequest) (resp *entities.OrderStatus, err error)
}

type ServiceImpl struct {
	cfg         configs.Config
	buyerImpl   repositories.IBuyer
	sellerImpl  repositories.ISeller
	orderImpl   repositories.IOrder
	productImpl repositories.IProduct
	log         logger.ILogger
}

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) WithConfig(cfg configs.Config) *ServiceImpl {
	s.cfg = cfg
	return s
}

func (s *ServiceImpl) WithBuyer(b repositories.IBuyer) *ServiceImpl {
	s.buyerImpl = b
	return s
}

func (s *ServiceImpl) WithSeller(sl repositories.ISeller) *ServiceImpl {
	s.sellerImpl = sl
	return s
}

func (s *ServiceImpl) WithOrder(o repositories.IOrder) *ServiceImpl {
	s.orderImpl = o
	return s
}

func (s *ServiceImpl) WithProduct(p repositories.IProduct) *ServiceImpl {
	s.productImpl = p
	return s
}

func (s *ServiceImpl) WithLog(l logger.ILogger) *ServiceImpl {
	s.log = l
	return s
}

var _ IService = (*ServiceImpl)(nil)
