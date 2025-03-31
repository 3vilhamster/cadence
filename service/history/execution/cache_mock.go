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
// Source: cache.go
//
// Generated by this command:
//
//	mockgen -package execution -source cache.go -destination cache_mock.go
//

// Package execution is a generated GoMock package.
package execution

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"

	cache "github.com/uber/cadence/common/cache"
	types "github.com/uber/cadence/common/types"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
	isgomock struct{}
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCache) Delete(key any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", key)
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheMockRecorder) Delete(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCache)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockCache) Get(key any) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(any)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), key)
}

// GetAndCreateWorkflowExecution mocks base method.
func (m *MockCache) GetAndCreateWorkflowExecution(ctx context.Context, domainID string, execution types.WorkflowExecution) (Context, Context, ReleaseFunc, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAndCreateWorkflowExecution", ctx, domainID, execution)
	ret0, _ := ret[0].(Context)
	ret1, _ := ret[1].(Context)
	ret2, _ := ret[2].(ReleaseFunc)
	ret3, _ := ret[3].(bool)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// GetAndCreateWorkflowExecution indicates an expected call of GetAndCreateWorkflowExecution.
func (mr *MockCacheMockRecorder) GetAndCreateWorkflowExecution(ctx, domainID, execution any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAndCreateWorkflowExecution", reflect.TypeOf((*MockCache)(nil).GetAndCreateWorkflowExecution), ctx, domainID, execution)
}

// GetOrCreateCurrentWorkflowExecution mocks base method.
func (m *MockCache) GetOrCreateCurrentWorkflowExecution(ctx context.Context, domainID, workflowID string) (Context, ReleaseFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateCurrentWorkflowExecution", ctx, domainID, workflowID)
	ret0, _ := ret[0].(Context)
	ret1, _ := ret[1].(ReleaseFunc)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrCreateCurrentWorkflowExecution indicates an expected call of GetOrCreateCurrentWorkflowExecution.
func (mr *MockCacheMockRecorder) GetOrCreateCurrentWorkflowExecution(ctx, domainID, workflowID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateCurrentWorkflowExecution", reflect.TypeOf((*MockCache)(nil).GetOrCreateCurrentWorkflowExecution), ctx, domainID, workflowID)
}

// GetOrCreateWorkflowExecution mocks base method.
func (m *MockCache) GetOrCreateWorkflowExecution(ctx context.Context, domainID string, execution types.WorkflowExecution) (Context, ReleaseFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateWorkflowExecution", ctx, domainID, execution)
	ret0, _ := ret[0].(Context)
	ret1, _ := ret[1].(ReleaseFunc)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrCreateWorkflowExecution indicates an expected call of GetOrCreateWorkflowExecution.
func (mr *MockCacheMockRecorder) GetOrCreateWorkflowExecution(ctx, domainID, execution any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateWorkflowExecution", reflect.TypeOf((*MockCache)(nil).GetOrCreateWorkflowExecution), ctx, domainID, execution)
}

// GetOrCreateWorkflowExecutionForBackground mocks base method.
func (m *MockCache) GetOrCreateWorkflowExecutionForBackground(domainID string, execution types.WorkflowExecution) (Context, ReleaseFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateWorkflowExecutionForBackground", domainID, execution)
	ret0, _ := ret[0].(Context)
	ret1, _ := ret[1].(ReleaseFunc)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrCreateWorkflowExecutionForBackground indicates an expected call of GetOrCreateWorkflowExecutionForBackground.
func (mr *MockCacheMockRecorder) GetOrCreateWorkflowExecutionForBackground(domainID, execution any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateWorkflowExecutionForBackground", reflect.TypeOf((*MockCache)(nil).GetOrCreateWorkflowExecutionForBackground), domainID, execution)
}

// GetOrCreateWorkflowExecutionWithTimeout mocks base method.
func (m *MockCache) GetOrCreateWorkflowExecutionWithTimeout(domainID string, execution types.WorkflowExecution, timeout time.Duration) (Context, ReleaseFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateWorkflowExecutionWithTimeout", domainID, execution, timeout)
	ret0, _ := ret[0].(Context)
	ret1, _ := ret[1].(ReleaseFunc)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrCreateWorkflowExecutionWithTimeout indicates an expected call of GetOrCreateWorkflowExecutionWithTimeout.
func (mr *MockCacheMockRecorder) GetOrCreateWorkflowExecutionWithTimeout(domainID, execution, timeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateWorkflowExecutionWithTimeout", reflect.TypeOf((*MockCache)(nil).GetOrCreateWorkflowExecutionWithTimeout), domainID, execution, timeout)
}

// Iterator mocks base method.
func (m *MockCache) Iterator() cache.Iterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator")
	ret0, _ := ret[0].(cache.Iterator)
	return ret0
}

// Iterator indicates an expected call of Iterator.
func (mr *MockCacheMockRecorder) Iterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockCache)(nil).Iterator))
}

// Put mocks base method.
func (m *MockCache) Put(key, value any) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(any)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockCacheMockRecorder) Put(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockCache)(nil).Put), key, value)
}

// PutIfNotExist mocks base method.
func (m *MockCache) PutIfNotExist(key, value any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutIfNotExist", key, value)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutIfNotExist indicates an expected call of PutIfNotExist.
func (mr *MockCacheMockRecorder) PutIfNotExist(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutIfNotExist", reflect.TypeOf((*MockCache)(nil).PutIfNotExist), key, value)
}

// Release mocks base method.
func (m *MockCache) Release(key any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Release", key)
}

// Release indicates an expected call of Release.
func (mr *MockCacheMockRecorder) Release(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockCache)(nil).Release), key)
}

// Size mocks base method.
func (m *MockCache) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockCacheMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockCache)(nil).Size))
}
