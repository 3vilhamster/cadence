// Code generated by MockGen. DO NOT EDIT.
// Source: types.go
//
// Generated by this command:
//
//	mockgen -package rpc -source types.go -destination factory_mock.go -self_package github.com/uber/cadence/common/rpc
//

// Package rpc is a generated GoMock package.
package rpc

import (
	reflect "reflect"

	membership "github.com/uber/cadence/common/membership"
	gomock "go.uber.org/mock/gomock"
	yarpc "go.uber.org/yarpc"
	tchannel "go.uber.org/yarpc/transport/tchannel"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
	isgomock struct{}
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// GetDispatcher mocks base method.
func (m *MockFactory) GetDispatcher() *yarpc.Dispatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDispatcher")
	ret0, _ := ret[0].(*yarpc.Dispatcher)
	return ret0
}

// GetDispatcher indicates an expected call of GetDispatcher.
func (mr *MockFactoryMockRecorder) GetDispatcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDispatcher", reflect.TypeOf((*MockFactory)(nil).GetDispatcher))
}

// GetMaxMessageSize mocks base method.
func (m *MockFactory) GetMaxMessageSize() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxMessageSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMaxMessageSize indicates an expected call of GetMaxMessageSize.
func (mr *MockFactoryMockRecorder) GetMaxMessageSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxMessageSize", reflect.TypeOf((*MockFactory)(nil).GetMaxMessageSize))
}

// GetTChannel mocks base method.
func (m *MockFactory) GetTChannel() tchannel.Channel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTChannel")
	ret0, _ := ret[0].(tchannel.Channel)
	return ret0
}

// GetTChannel indicates an expected call of GetTChannel.
func (mr *MockFactoryMockRecorder) GetTChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTChannel", reflect.TypeOf((*MockFactory)(nil).GetTChannel))
}

// Start mocks base method.
func (m *MockFactory) Start(arg0 PeerLister) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockFactoryMockRecorder) Start(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockFactory)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockFactory) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockFactoryMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockFactory)(nil).Stop))
}

// MockPeerLister is a mock of PeerLister interface.
type MockPeerLister struct {
	ctrl     *gomock.Controller
	recorder *MockPeerListerMockRecorder
	isgomock struct{}
}

// MockPeerListerMockRecorder is the mock recorder for MockPeerLister.
type MockPeerListerMockRecorder struct {
	mock *MockPeerLister
}

// NewMockPeerLister creates a new mock instance.
func NewMockPeerLister(ctrl *gomock.Controller) *MockPeerLister {
	mock := &MockPeerLister{ctrl: ctrl}
	mock.recorder = &MockPeerListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerLister) EXPECT() *MockPeerListerMockRecorder {
	return m.recorder
}

// Members mocks base method.
func (m *MockPeerLister) Members(service string) ([]membership.HostInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Members", service)
	ret0, _ := ret[0].([]membership.HostInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Members indicates an expected call of Members.
func (mr *MockPeerListerMockRecorder) Members(service any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockPeerLister)(nil).Members), service)
}

// Subscribe mocks base method.
func (m *MockPeerLister) Subscribe(service, name string, notifyChannel chan<- *membership.ChangedEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", service, name, notifyChannel)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockPeerListerMockRecorder) Subscribe(service, name, notifyChannel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockPeerLister)(nil).Subscribe), service, name, notifyChannel)
}

// Unsubscribe mocks base method.
func (m *MockPeerLister) Unsubscribe(service, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", service, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockPeerListerMockRecorder) Unsubscribe(service, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockPeerLister)(nil).Unsubscribe), service, name)
}
