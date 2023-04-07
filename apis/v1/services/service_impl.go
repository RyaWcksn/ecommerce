package services

import (
	"context"
	"fmt"
	"strconv"

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
		if err != nil {
			s.log.Errorf("[ERR] While getting password", err)
			return "", errors.GetError(errors.InternalServer, err)
		}
	case constants.SELLER:
		info, err = s.sellerImpl.GetEmail(ctx, payload.Email)
		if err != nil {
			s.log.Errorf("[ERR] While getting password", err)
			return "", errors.GetError(errors.InternalServer, err)
		}
	default:
		return "", errors.GetError(errors.InvalidRequest, fmt.Errorf("[ERR] Role not found"))
	}
	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(payload.Password)); err != nil {
		s.log.Errorf("[ERR] While compare password with hash", err)
		return "", errors.GetError(errors.InvalidRequest, err)
	}

	tokenPayload := dto.TokenGenerator{
		Id:        info.Id,
		SecretKey: s.cfg.App.SECRET,
		Email:     payload.Email,
		Role:      payload.Role,
	}
	token, err = tokens.GenerateJWT(&tokenPayload)
	if err != nil {
		s.log.Errorf("[ERR] While generating token", err)
		return "", errors.GetError(errors.InternalServer, err)
	}

	return token, nil
}

// CreateProduct implements IService
func (s *ServiceImpl) CreateProduct(ctx context.Context, payload *dto.CreateProductRequest) error {
	idStr := ctx.Value("id").(string)
	id, _ := strconv.Atoi(idStr)

	entity := entities.CreateProductEntity{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Seller:      id,
	}

	err := s.productImpl.CreateProduct(ctx, &entity)
	if err != nil {
		s.log.Errorf("[ERR] Err from domain := %v", err)
		return err
	}

	return nil
}

// GetProductsList implements IService
func (s *ServiceImpl) GetProductsList(ctx context.Context, id int) (productList *[]entities.ProductListEntity, err error) {
	idStr := ctx.Value("id").(string)
	id, _ = strconv.Atoi(idStr)

	products, err := s.productImpl.ListProduct(ctx, id)
	if err != nil {
		s.log.Errorf("[ERR] Err from domain := %v", err)
		return nil, err
	}

	return products, nil

}

// CreateOrder implements IService
func (s *ServiceImpl) CreateOrder(ctx context.Context, payload *dto.CreateOrderRequest) (resp *entities.OrderStatus, err error) {
	idStr := ctx.Value("id").(string)
	id, _ := strconv.Atoi(idStr)
	fmt.Println("ID", id)

	buyerData, err := s.buyerImpl.GetData(ctx, id)
	if err != nil {
		s.log.Errorf("[ERR] Error while getting buyer data := %v", err)
		return nil, err
	}

	sellerData, err := s.sellerImpl.GetData(ctx, payload.SellerId)
	if err != nil {
		s.log.Errorf("[ERR] Error while getting seller data := %v", err)
		return nil, err
	}

	var items string
	var total int
	var grandTotal string
	var quantity int
	for _, productId := range payload.ProductId {
		product, err := s.productImpl.GetProductById(ctx, productId)
		if err != nil {
			s.log.Errorf("[ERR] Error while getting product data := %v", err)
			return nil, err
		}
		price, _ := strconv.Atoi(product.Price)
		total += price
		items += ", " + product.ProductName
		grandTotal = strconv.Itoa(total)
		quantity++
	}

	createOrderEntity := entities.CreateOrder{
		Buyer:               id,
		Seller:              payload.SellerId,
		DeliverySource:      buyerData.AlamatPengiriman,
		DeliveryDestination: sellerData.AlamatPickup,
		Items:               items,
		Quantity:            quantity,
		Price:               strconv.Itoa(total),
		TotalPrice:          grandTotal,
		Status:              0,
	}

	err = s.orderImpl.CreateOrder(ctx, &createOrderEntity)
	if err != nil {
		s.log.Errorf("[ERR] Error while create order := %v", err)
		return nil, err
	}

	return &entities.OrderStatus{
		Message: constants.PendingMessage,
		Status:  constants.Pending,
	}, nil
}

// GetSellerOrderList implements IService
func (s *ServiceImpl) GetSellerOrderList(ctx context.Context) (orderList *[]entities.Order, err error) {
	idStr := ctx.Value("id").(string)
	id, _ := strconv.Atoi(idStr)

	orders, err := s.orderImpl.SellerViewOrderList(ctx, id)
	if err != nil {
		s.log.Errorf("[ERR] Error while getting orders data := %v", err)
		return nil, err
	}

	return orders, nil
}

// UpdateOrderStatus implements IService
func (s *ServiceImpl) UpdateOrderStatus(ctx context.Context, payload *dto.AcceptOrderRequest) (resp *entities.Order, err error) {
	order, err := s.orderImpl.UpdateOrder(ctx, payload.OrderId)
	if err != nil {
		s.log.Errorf("[ERR] Error from domain layer := %v", err)
		return nil, err
	}
	return order, nil
}
