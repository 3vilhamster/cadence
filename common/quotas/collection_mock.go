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
// Source: github.com/uber/cadence/common/quotas (interfaces: ICollection)
//
// Generated by this command:
//
//	mockgen -package=quotas -destination=collection_mock.go github.com/uber/cadence/common/quotas ICollection
//

// Package quotas is a generated GoMock package.
package quotas

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockICollection is a mock of ICollection interface.
type MockICollection struct {
	ctrl     *gomock.Controller
	recorder *MockICollectionMockRecorder
	isgomock struct{}
}

// MockICollectionMockRecorder is the mock recorder for MockICollection.
type MockICollectionMockRecorder struct {
	mock *MockICollection
}

// NewMockICollection creates a new mock instance.
func NewMockICollection(ctrl *gomock.Controller) *MockICollection {
	mock := &MockICollection{ctrl: ctrl}
	mock.recorder = &MockICollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICollection) EXPECT() *MockICollectionMockRecorder {
	return m.recorder
}

// For mocks base method.
func (m *MockICollection) For(key string) Limiter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "For", key)
	ret0, _ := ret[0].(Limiter)
	return ret0
}

// For indicates an expected call of For.
func (mr *MockICollectionMockRecorder) For(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "For", reflect.TypeOf((*MockICollection)(nil).For), key)
}
