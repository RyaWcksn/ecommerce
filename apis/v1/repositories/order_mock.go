// Code generated by MockGen. DO NOT EDIT.
// Source: order.go

// Package repositories is a generated GoMock package.
package repositories

import (
	context "context"
	reflect "reflect"

	entities "github.com/RyaWcksn/ecommerce/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockIOrder is a mock of IOrder interface.
type MockIOrder struct {
	ctrl     *gomock.Controller
	recorder *MockIOrderMockRecorder
}

// MockIOrderMockRecorder is the mock recorder for MockIOrder.
type MockIOrderMockRecorder struct {
	mock *MockIOrder
}

// NewMockIOrder creates a new mock instance.
func NewMockIOrder(ctrl *gomock.Controller) *MockIOrder {
	mock := &MockIOrder{ctrl: ctrl}
	mock.recorder = &MockIOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOrder) EXPECT() *MockIOrderMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockIOrder) CreateOrder(ctx context.Context, entity *entities.CreateOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockIOrderMockRecorder) CreateOrder(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockIOrder)(nil).CreateOrder), ctx, entity)
}