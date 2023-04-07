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
