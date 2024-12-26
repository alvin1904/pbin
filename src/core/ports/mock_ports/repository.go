// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/core/ports/repository.go
//
// Generated by this command:
//
//	mockgen --source=./src/core/ports/repository.go --destination=./src/core/ports/mocks/repository.go
//

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	context "context"
	reflect "reflect"

	domain "github.com/Abhishekkarunakaran/pbin/src/core/domain"
	uuid "github.com/gofrs/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddData mocks base method.
func (m *MockRepository) AddData(ctx context.Context, id uuid.UUID, data domain.Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddData", ctx, id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddData indicates an expected call of AddData.
func (mr *MockRepositoryMockRecorder) AddData(ctx, id, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddData", reflect.TypeOf((*MockRepository)(nil).AddData), ctx, id, data)
}

// GetData mocks base method.
func (m *MockRepository) GetData(ctx context.Context, id uuid.UUID) (*domain.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", ctx, id)
	ret0, _ := ret[0].(*domain.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockRepositoryMockRecorder) GetData(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockRepository)(nil).GetData), ctx, id)
}

// IsContentPresent mocks base method.
func (m *MockRepository) IsContentPresent(ctx context.Context, id uuid.UUID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsContentPresent", ctx, id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsContentPresent indicates an expected call of IsContentPresent.
func (mr *MockRepositoryMockRecorder) IsContentPresent(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsContentPresent", reflect.TypeOf((*MockRepository)(nil).IsContentPresent), ctx, id)
}

// RemoveData mocks base method.
func (m *MockRepository) RemoveData(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveData", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveData indicates an expected call of RemoveData.
func (mr *MockRepositoryMockRecorder) RemoveData(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveData", reflect.TypeOf((*MockRepository)(nil).RemoveData), ctx, id)
}
