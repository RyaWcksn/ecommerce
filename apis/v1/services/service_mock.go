// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package services is a generated GoMock package.
package services

import (
	context "context"
	reflect "reflect"

	dto "github.com/RyaWcksn/ecommerce/dto"
	entities "github.com/RyaWcksn/ecommerce/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockIService) CreateProduct(ctx context.Context, payload *dto.CreateProductRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockIServiceMockRecorder) CreateProduct(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockIService)(nil).CreateProduct), ctx, payload)
}

// GetProductsList mocks base method.
func (m *MockIService) GetProductsList(ctx context.Context, id int) (*[]entities.ProductListEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsList", ctx, id)
	ret0, _ := ret[0].(*[]entities.ProductListEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsList indicates an expected call of GetProductsList.
func (mr *MockIServiceMockRecorder) GetProductsList(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsList", reflect.TypeOf((*MockIService)(nil).GetProductsList), ctx, id)
}

// Login mocks base method.
func (m *MockIService) Login(ctx context.Context, payload *dto.LoginRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, payload)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIServiceMockRecorder) Login(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIService)(nil).Login), ctx, payload)
}