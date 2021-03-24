// Code generated by MockGen. DO NOT EDIT.
// Source: router/router.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// GetExternalData mocks base method.
func (m *MockController) GetExternalData(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetExternalData", w, r)
}

// GetExternalData indicates an expected call of GetExternalData.
func (mr *MockControllerMockRecorder) GetExternalData(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalData", reflect.TypeOf((*MockController)(nil).GetExternalData), w, r)
}

// GetUserById mocks base method.
func (m *MockController) GetUserById(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUserById", w, r)
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockControllerMockRecorder) GetUserById(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockController)(nil).GetUserById), w, r)
}

// GetUsers mocks base method.
func (m *MockController) GetUsers(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUsers", w, r)
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockControllerMockRecorder) GetUsers(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockController)(nil).GetUsers), w, r)
}
