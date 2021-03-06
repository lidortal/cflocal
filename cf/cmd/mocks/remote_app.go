// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sclevine/cflocal/cf/cmd (interfaces: RemoteApp)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	remote "github.com/sclevine/cflocal/remote"
	service "github.com/sclevine/forge/service"
	io "io"
	reflect "reflect"
)

// MockRemoteApp is a mock of RemoteApp interface
type MockRemoteApp struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteAppMockRecorder
}

// MockRemoteAppMockRecorder is the mock recorder for MockRemoteApp
type MockRemoteAppMockRecorder struct {
	mock *MockRemoteApp
}

// NewMockRemoteApp creates a new mock instance
func NewMockRemoteApp(ctrl *gomock.Controller) *MockRemoteApp {
	mock := &MockRemoteApp{ctrl: ctrl}
	mock.recorder = &MockRemoteAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteApp) EXPECT() *MockRemoteAppMockRecorder {
	return m.recorder
}

// Command mocks base method
func (m *MockRemoteApp) Command(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "Command", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Command indicates an expected call of Command
func (mr *MockRemoteAppMockRecorder) Command(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockRemoteApp)(nil).Command), arg0)
}

// Droplet mocks base method
func (m *MockRemoteApp) Droplet(arg0 string) (io.ReadCloser, int64, error) {
	ret := m.ctrl.Call(m, "Droplet", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Droplet indicates an expected call of Droplet
func (mr *MockRemoteAppMockRecorder) Droplet(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Droplet", reflect.TypeOf((*MockRemoteApp)(nil).Droplet), arg0)
}

// Env mocks base method
func (m *MockRemoteApp) Env(arg0 string) (*remote.AppEnv, error) {
	ret := m.ctrl.Call(m, "Env", arg0)
	ret0, _ := ret[0].(*remote.AppEnv)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Env indicates an expected call of Env
func (mr *MockRemoteAppMockRecorder) Env(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Env", reflect.TypeOf((*MockRemoteApp)(nil).Env), arg0)
}

// Forward mocks base method
func (m *MockRemoteApp) Forward(arg0 string, arg1 service.Services) (service.Services, *service.ForwardConfig, error) {
	ret := m.ctrl.Call(m, "Forward", arg0, arg1)
	ret0, _ := ret[0].(service.Services)
	ret1, _ := ret[1].(*service.ForwardConfig)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Forward indicates an expected call of Forward
func (mr *MockRemoteAppMockRecorder) Forward(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forward", reflect.TypeOf((*MockRemoteApp)(nil).Forward), arg0, arg1)
}

// Restart mocks base method
func (m *MockRemoteApp) Restart(arg0 string) error {
	ret := m.ctrl.Call(m, "Restart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restart indicates an expected call of Restart
func (mr *MockRemoteAppMockRecorder) Restart(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restart", reflect.TypeOf((*MockRemoteApp)(nil).Restart), arg0)
}

// Services mocks base method
func (m *MockRemoteApp) Services(arg0 string) (service.Services, error) {
	ret := m.ctrl.Call(m, "Services", arg0)
	ret0, _ := ret[0].(service.Services)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Services indicates an expected call of Services
func (mr *MockRemoteAppMockRecorder) Services(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*MockRemoteApp)(nil).Services), arg0)
}

// SetDroplet mocks base method
func (m *MockRemoteApp) SetDroplet(arg0 string, arg1 io.Reader, arg2 int64) error {
	ret := m.ctrl.Call(m, "SetDroplet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDroplet indicates an expected call of SetDroplet
func (mr *MockRemoteAppMockRecorder) SetDroplet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDroplet", reflect.TypeOf((*MockRemoteApp)(nil).SetDroplet), arg0, arg1, arg2)
}

// SetEnv mocks base method
func (m *MockRemoteApp) SetEnv(arg0 string, arg1 map[string]string) error {
	ret := m.ctrl.Call(m, "SetEnv", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEnv indicates an expected call of SetEnv
func (mr *MockRemoteAppMockRecorder) SetEnv(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnv", reflect.TypeOf((*MockRemoteApp)(nil).SetEnv), arg0, arg1)
}
