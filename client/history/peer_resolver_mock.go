// Code generated by MockGen. DO NOT EDIT.
// Source: peer_resolver.go
//
// Generated by this command:
//
//	mockgen -package history -source peer_resolver.go -destination peer_resolver_mock.go -package history github.com/uber/cadence/client/history PeerResolver
//

// Package history is a generated GoMock package.
package history

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPeerResolver is a mock of PeerResolver interface.
type MockPeerResolver struct {
	ctrl     *gomock.Controller
	recorder *MockPeerResolverMockRecorder
	isgomock struct{}
}

// MockPeerResolverMockRecorder is the mock recorder for MockPeerResolver.
type MockPeerResolverMockRecorder struct {
	mock *MockPeerResolver
}

// NewMockPeerResolver creates a new mock instance.
func NewMockPeerResolver(ctrl *gomock.Controller) *MockPeerResolver {
	mock := &MockPeerResolver{ctrl: ctrl}
	mock.recorder = &MockPeerResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerResolver) EXPECT() *MockPeerResolverMockRecorder {
	return m.recorder
}

// FromDomainID mocks base method.
func (m *MockPeerResolver) FromDomainID(domainID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromDomainID", domainID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromDomainID indicates an expected call of FromDomainID.
func (mr *MockPeerResolverMockRecorder) FromDomainID(domainID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromDomainID", reflect.TypeOf((*MockPeerResolver)(nil).FromDomainID), domainID)
}

// FromHostAddress mocks base method.
func (m *MockPeerResolver) FromHostAddress(hostAddress string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromHostAddress", hostAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromHostAddress indicates an expected call of FromHostAddress.
func (mr *MockPeerResolverMockRecorder) FromHostAddress(hostAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromHostAddress", reflect.TypeOf((*MockPeerResolver)(nil).FromHostAddress), hostAddress)
}

// FromShardID mocks base method.
func (m *MockPeerResolver) FromShardID(shardID int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromShardID", shardID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromShardID indicates an expected call of FromShardID.
func (mr *MockPeerResolverMockRecorder) FromShardID(shardID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromShardID", reflect.TypeOf((*MockPeerResolver)(nil).FromShardID), shardID)
}

// FromWorkflowID mocks base method.
func (m *MockPeerResolver) FromWorkflowID(workflowID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromWorkflowID", workflowID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromWorkflowID indicates an expected call of FromWorkflowID.
func (mr *MockPeerResolverMockRecorder) FromWorkflowID(workflowID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromWorkflowID", reflect.TypeOf((*MockPeerResolver)(nil).FromWorkflowID), workflowID)
}

// GetAllPeers mocks base method.
func (m *MockPeerResolver) GetAllPeers() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPeers")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPeers indicates an expected call of GetAllPeers.
func (mr *MockPeerResolverMockRecorder) GetAllPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPeers", reflect.TypeOf((*MockPeerResolver)(nil).GetAllPeers))
}

// GlobalRatelimitPeers mocks base method.
func (m *MockPeerResolver) GlobalRatelimitPeers(ratelimits []string) (map[Peer][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalRatelimitPeers", ratelimits)
	ret0, _ := ret[0].(map[Peer][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GlobalRatelimitPeers indicates an expected call of GlobalRatelimitPeers.
func (mr *MockPeerResolverMockRecorder) GlobalRatelimitPeers(ratelimits any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalRatelimitPeers", reflect.TypeOf((*MockPeerResolver)(nil).GlobalRatelimitPeers), ratelimits)
}
