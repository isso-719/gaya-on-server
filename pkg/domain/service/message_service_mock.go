// Code generated by MockGen. DO NOT EDIT.
// Source: message_service.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/isso-719/gaya-on-server/pkg/domain/model"
)

// MockIFMessageService is a mock of IFMessageService interface.
type MockIFMessageService struct {
	ctrl     *gomock.Controller
	recorder *MockIFMessageServiceMockRecorder
}

// MockIFMessageServiceMockRecorder is the mock recorder for MockIFMessageService.
type MockIFMessageServiceMockRecorder struct {
	mock *MockIFMessageService
}

// NewMockIFMessageService creates a new mock instance.
func NewMockIFMessageService(ctrl *gomock.Controller) *MockIFMessageService {
	mock := &MockIFMessageService{ctrl: ctrl}
	mock.recorder = &MockIFMessageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFMessageService) EXPECT() *MockIFMessageServiceMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method.
func (m *MockIFMessageService) CreateMessage(ctx context.Context, roomID int64, messageType, messageBody string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", ctx, roomID, messageType, messageBody)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockIFMessageServiceMockRecorder) CreateMessage(ctx, roomID, messageType, messageBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockIFMessageService)(nil).CreateMessage), ctx, roomID, messageType, messageBody)
}

// GetAllMessages mocks base method.
func (m *MockIFMessageService) GetAllMessages(ctx context.Context, roomID int64) ([]*model.MessageTypeAndBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMessages", ctx, roomID)
	ret0, _ := ret[0].([]*model.MessageTypeAndBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMessages indicates an expected call of GetAllMessages.
func (mr *MockIFMessageServiceMockRecorder) GetAllMessages(ctx, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMessages", reflect.TypeOf((*MockIFMessageService)(nil).GetAllMessages), ctx, roomID)
}

// Migrate mocks base method.
func (m *MockIFMessageService) Migrate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockIFMessageServiceMockRecorder) Migrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockIFMessageService)(nil).Migrate))
}