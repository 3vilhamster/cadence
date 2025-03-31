// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go
//
// Generated by this command:
//
//	mockgen -package events -source cache.go -destination cache_mock.go
//

// Package events is a generated GoMock package.
package events

import (
	context "context"
	reflect "reflect"

	types "github.com/uber/cadence/common/types"
	gomock "go.uber.org/mock/gomock"
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

// GetEvent mocks base method.
func (m *MockCache) GetEvent(ctx context.Context, shardID int, domainID, workflowID, runID string, firstEventID, eventID int64, branchToken []byte) (*types.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", ctx, shardID, domainID, workflowID, runID, firstEventID, eventID, branchToken)
	ret0, _ := ret[0].(*types.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockCacheMockRecorder) GetEvent(ctx, shardID, domainID, workflowID, runID, firstEventID, eventID, branchToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockCache)(nil).GetEvent), ctx, shardID, domainID, workflowID, runID, firstEventID, eventID, branchToken)
}

// PutEvent mocks base method.
func (m *MockCache) PutEvent(domainID, workflowID, runID string, eventID int64, event *types.HistoryEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutEvent", domainID, workflowID, runID, eventID, event)
}

// PutEvent indicates an expected call of PutEvent.
func (mr *MockCacheMockRecorder) PutEvent(domainID, workflowID, runID, eventID, event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutEvent", reflect.TypeOf((*MockCache)(nil).PutEvent), domainID, workflowID, runID, eventID, event)
}
