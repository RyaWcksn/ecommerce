// Code generated by MockGen. DO NOT EDIT.
// Source: buyer.go

// Package repositories is a generated GoMock package.
package repositories

import (
	context "context"
	reflect "reflect"

	entities "github.com/RyaWcksn/ecommerce/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockIBuyer is a mock of IBuyer interface.
type MockIBuyer struct {
	ctrl     *gomock.Controller
	recorder *MockIBuyerMockRecorder
}

// MockIBuyerMockRecorder is the mock recorder for MockIBuyer.
type MockIBuyerMockRecorder struct {
	mock *MockIBuyer
}

// NewMockIBuyer creates a new mock instance.
func NewMockIBuyer(ctrl *gomock.Controller) *MockIBuyer {
	mock := &MockIBuyer{ctrl: ctrl}
	mock.recorder = &MockIBuyerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBuyer) EXPECT() *MockIBuyerMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockIBuyer) GetData(ctx context.Context, id int) (*entities.BuyerEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", ctx, id)
	ret0, _ := ret[0].(*entities.BuyerEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockIBuyerMockRecorder) GetData(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockIBuyer)(nil).GetData), ctx, id)
}

// GetEmail mocks base method.
func (m *MockIBuyer) GetEmail(ctx context.Context, email string) (*entities.LoginEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail", ctx, email)
	ret0, _ := ret[0].(*entities.LoginEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockIBuyerMockRecorder) GetEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockIBuyer)(nil).GetEmail), ctx, email)
}