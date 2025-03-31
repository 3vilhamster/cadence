// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: workflow.go
//
// Generated by this command:
//
//	mockgen -package execution -source workflow.go -destination workflow_mock.go -self_package github.com/uber/cadence/service/history/execution
//

// Package execution is a generated GoMock package.
package execution

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWorkflow is a mock of Workflow interface.
type MockWorkflow struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowMockRecorder
	isgomock struct{}
}

// MockWorkflowMockRecorder is the mock recorder for MockWorkflow.
type MockWorkflowMockRecorder struct {
	mock *MockWorkflow
}

// NewMockWorkflow creates a new mock instance.
func NewMockWorkflow(ctrl *gomock.Controller) *MockWorkflow {
	mock := &MockWorkflow{ctrl: ctrl}
	mock.recorder = &MockWorkflowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkflow) EXPECT() *MockWorkflowMockRecorder {
	return m.recorder
}

// FlushBufferedEvents mocks base method.
func (m *MockWorkflow) FlushBufferedEvents() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushBufferedEvents")
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushBufferedEvents indicates an expected call of FlushBufferedEvents.
func (mr *MockWorkflowMockRecorder) FlushBufferedEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushBufferedEvents", reflect.TypeOf((*MockWorkflow)(nil).FlushBufferedEvents))
}

// GetContext mocks base method.
func (m *MockWorkflow) GetContext() Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContext")
	ret0, _ := ret[0].(Context)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockWorkflowMockRecorder) GetContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockWorkflow)(nil).GetContext))
}

// GetMutableState mocks base method.
func (m *MockWorkflow) GetMutableState() MutableState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMutableState")
	ret0, _ := ret[0].(MutableState)
	return ret0
}

// GetMutableState indicates an expected call of GetMutableState.
func (mr *MockWorkflowMockRecorder) GetMutableState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMutableState", reflect.TypeOf((*MockWorkflow)(nil).GetMutableState))
}

// GetReleaseFn mocks base method.
func (m *MockWorkflow) GetReleaseFn() ReleaseFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReleaseFn")
	ret0, _ := ret[0].(ReleaseFunc)
	return ret0
}

// GetReleaseFn indicates an expected call of GetReleaseFn.
func (mr *MockWorkflowMockRecorder) GetReleaseFn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReleaseFn", reflect.TypeOf((*MockWorkflow)(nil).GetReleaseFn))
}

// GetVectorClock mocks base method.
func (m *MockWorkflow) GetVectorClock() (int64, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVectorClock")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetVectorClock indicates an expected call of GetVectorClock.
func (mr *MockWorkflowMockRecorder) GetVectorClock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVectorClock", reflect.TypeOf((*MockWorkflow)(nil).GetVectorClock))
}

// HappensAfter mocks base method.
func (m *MockWorkflow) HappensAfter(that Workflow) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HappensAfter", that)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HappensAfter indicates an expected call of HappensAfter.
func (mr *MockWorkflowMockRecorder) HappensAfter(that any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HappensAfter", reflect.TypeOf((*MockWorkflow)(nil).HappensAfter), that)
}

// Revive mocks base method.
func (m *MockWorkflow) Revive() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revive")
	ret0, _ := ret[0].(error)
	return ret0
}

// Revive indicates an expected call of Revive.
func (mr *MockWorkflowMockRecorder) Revive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revive", reflect.TypeOf((*MockWorkflow)(nil).Revive))
}

// SuppressBy mocks base method.
func (m *MockWorkflow) SuppressBy(incomingWorkflow Workflow) (TransactionPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuppressBy", incomingWorkflow)
	ret0, _ := ret[0].(TransactionPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuppressBy indicates an expected call of SuppressBy.
func (mr *MockWorkflowMockRecorder) SuppressBy(incomingWorkflow any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuppressBy", reflect.TypeOf((*MockWorkflow)(nil).SuppressBy), incomingWorkflow)
}
