// Code generated by MockGen. DO NOT EDIT.
// Source: order.go

// Package repositories is a generated GoMock package.
package repositories

import (
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