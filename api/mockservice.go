// Code generated by MockGen. DO NOT EDIT.
// Source: Kafka/api (interfaces: Service)

// Package mocks is a generated GoMock package.
package api

import (

	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockService) CreateProduct(arg0 *ProductRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockServiceMockRecorder) CreateProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockService)(nil).CreateProduct), arg0)
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(arg0 *UserRequest) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), arg0)
}

// DownloadAndCompressImage mocks base method.
func (m *MockService) DownloadAndCompressImage(arg0, arg1 string , arg2 bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadAndCompressImage", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadAndCompressImage indicates an expected call of DownloadAndCompressImage.
func (mr *MockServiceMockRecorder) DownloadAndCompressImage(arg0, arg1 , arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadAndCompressImage", reflect.TypeOf((*MockService)(nil).DownloadAndCompressImage), arg0, arg1)
}

// GetProductFromKafka mocks base method.
func (m *MockService) GetProductFromKafka(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetProductFromKafka")
}

// GetProductFromKafka indicates an expected call of GetProductFromKafka.
func (mr *MockServiceMockRecorder) GetProductFromKafka(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductFromKafka", reflect.TypeOf((*MockService)(nil).GetProductFromKafka))
}
