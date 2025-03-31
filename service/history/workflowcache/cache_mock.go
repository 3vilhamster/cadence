// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/cadence/service/history/workflowcache (interfaces: WFCache)
//
// Generated by this command:
//
//	mockgen -package=workflowcache -destination=cache_mock.go github.com/uber/cadence/service/history/workflowcache WFCache
//

// Package workflowcache is a generated GoMock package.
package workflowcache

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWFCache is a mock of WFCache interface.
type MockWFCache struct {
	ctrl     *gomock.Controller
	recorder *MockWFCacheMockRecorder
	isgomock struct{}
}

// MockWFCacheMockRecorder is the mock recorder for MockWFCache.
type MockWFCacheMockRecorder struct {
	mock *MockWFCache
}

// NewMockWFCache creates a new mock instance.
func NewMockWFCache(ctrl *gomock.Controller) *MockWFCache {
	mock := &MockWFCache{ctrl: ctrl}
	mock.recorder = &MockWFCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWFCache) EXPECT() *MockWFCacheMockRecorder {
	return m.recorder
}

// AllowExternal mocks base method.
func (m *MockWFCache) AllowExternal(domainID, workflowID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowExternal", domainID, workflowID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AllowExternal indicates an expected call of AllowExternal.
func (mr *MockWFCacheMockRecorder) AllowExternal(domainID, workflowID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowExternal", reflect.TypeOf((*MockWFCache)(nil).AllowExternal), domainID, workflowID)
}

// AllowInternal mocks base method.
func (m *MockWFCache) AllowInternal(domainID, workflowID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowInternal", domainID, workflowID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AllowInternal indicates an expected call of AllowInternal.
func (mr *MockWFCacheMockRecorder) AllowInternal(domainID, workflowID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowInternal", reflect.TypeOf((*MockWFCache)(nil).AllowInternal), domainID, workflowID)
}
