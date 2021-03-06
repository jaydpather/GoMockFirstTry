// Code generated by MockGen. DO NOT EDIT.
// Source: employeeRepository.go

// Package mock_employeeRepository is a generated GoMock package.
package mock_employeeRepository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIEmployeeRepository is a mock of IEmployeeRepository interface.
type MockIEmployeeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEmployeeRepositoryMockRecorder
}

// MockIEmployeeRepositoryMockRecorder is the mock recorder for MockIEmployeeRepository.
type MockIEmployeeRepositoryMockRecorder struct {
	mock *MockIEmployeeRepository
}

// NewMockIEmployeeRepository creates a new mock instance.
func NewMockIEmployeeRepository(ctrl *gomock.Controller) *MockIEmployeeRepository {
	mock := &MockIEmployeeRepository{ctrl: ctrl}
	mock.recorder = &MockIEmployeeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEmployeeRepository) EXPECT() *MockIEmployeeRepositoryMockRecorder {
	return m.recorder
}

// InsertEmployee mocks base method.
func (m *MockIEmployeeRepository) InsertEmployee() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertEmployee")
	ret0, _ := ret[0].(int)
	return ret0
}

// InsertEmployee indicates an expected call of InsertEmployee.
func (mr *MockIEmployeeRepositoryMockRecorder) InsertEmployee() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEmployee", reflect.TypeOf((*MockIEmployeeRepository)(nil).InsertEmployee))
}
