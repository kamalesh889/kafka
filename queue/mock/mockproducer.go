// Code generated by MockGen. DO NOT EDIT.
// Source: Kafka/queue (interfaces: Producer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kafka "github.com/segmentio/kafka-go"
)

// MockProducer is a mock of Producer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// WriteMessage mocks base method.
func (m *MockProducer) WriteMessage(arg0 context.Context, arg1 kafka.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage.
func (mr *MockProducerMockRecorder) WriteMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockProducer)(nil).WriteMessage), arg0, arg1)
}