// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -package rpc -source client.go -destination client_mock.go -self_package github.com/uber/cadence/common/quotas/global/rpc Client
//

// Package rpc is a generated GoMock package.
package rpc

import (
	context "context"
	reflect "reflect"
	time "time"

	shared "github.com/uber/cadence/common/quotas/global/shared"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockClient) Update(ctx context.Context, period time.Duration, load map[shared.GlobalKey]Calls) UpdateResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, period, load)
	ret0, _ := ret[0].(UpdateResult)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockClientMockRecorder) Update(ctx, period, load any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClient)(nil).Update), ctx, period, load)
}
