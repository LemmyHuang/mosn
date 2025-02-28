// Code generated by MockGen. DO NOT EDIT.
// Source: ../stagemanager/stage_manager.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "mosn.io/mosn/pkg/config/v2"
)

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockApplication) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockApplicationMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockApplication)(nil).Close))
}

// InheritConnections mocks base method.
func (m *MockApplication) InheritConnections() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InheritConnections")
	ret0, _ := ret[0].(error)
	return ret0
}

// InheritConnections indicates an expected call of InheritConnections.
func (mr *MockApplicationMockRecorder) InheritConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InheritConnections", reflect.TypeOf((*MockApplication)(nil).InheritConnections))
}

// Init mocks base method.
func (m *MockApplication) Init(arg0 *v2.MOSNConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockApplicationMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockApplication)(nil).Init), arg0)
}

// Start mocks base method.
func (m *MockApplication) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockApplicationMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockApplication)(nil).Start))
}
