// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/trustwallet/go-libs/database (interfaces: DBContextGetter,TrxContextGetter)

// Package database is a generated GoMock package.
package database

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDBContextGetter is a mock of DBContextGetter interface.
type MockDBContextGetter struct {
	ctrl     *gomock.Controller
	recorder *MockDBContextGetterMockRecorder
}

// MockDBContextGetterMockRecorder is the mock recorder for MockDBContextGetter.
type MockDBContextGetterMockRecorder struct {
	mock *MockDBContextGetter
}

// NewMockDBContextGetter creates a new mock instance.
func NewMockDBContextGetter(ctrl *gomock.Controller) *MockDBContextGetter {
	mock := &MockDBContextGetter{ctrl: ctrl}
	mock.recorder = &MockDBContextGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBContextGetter) EXPECT() *MockDBContextGetterMockRecorder {
	return m.recorder
}

// DBFrom mocks base method.
func (m *MockDBContextGetter) DBFrom(arg0 context.Context) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBFrom", arg0)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// DBFrom indicates an expected call of DBFrom.
func (mr *MockDBContextGetterMockRecorder) DBFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBFrom", reflect.TypeOf((*MockDBContextGetter)(nil).DBFrom), arg0)
}

// ReadOnlyDB mocks base method.
func (m *MockDBContextGetter) ReadOnlyDB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOnlyDB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// ReadOnlyDB indicates an expected call of ReadOnlyDB.
func (mr *MockDBContextGetterMockRecorder) ReadOnlyDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOnlyDB", reflect.TypeOf((*MockDBContextGetter)(nil).ReadOnlyDB))
}

// MockTrxContextGetter is a mock of TrxContextGetter interface.
type MockTrxContextGetter struct {
	ctrl     *gomock.Controller
	recorder *MockTrxContextGetterMockRecorder
}

// MockTrxContextGetterMockRecorder is the mock recorder for MockTrxContextGetter.
type MockTrxContextGetterMockRecorder struct {
	mock *MockTrxContextGetter
}

// NewMockTrxContextGetter creates a new mock instance.
func NewMockTrxContextGetter(ctrl *gomock.Controller) *MockTrxContextGetter {
	mock := &MockTrxContextGetter{ctrl: ctrl}
	mock.recorder = &MockTrxContextGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrxContextGetter) EXPECT() *MockTrxContextGetterMockRecorder {
	return m.recorder
}

// Transaction mocks base method.
func (m *MockTrxContextGetter) Transaction(arg0 context.Context, arg1 func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockTrxContextGetterMockRecorder) Transaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockTrxContextGetter)(nil).Transaction), arg0, arg1)
}
