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
// Source: peer_chooser.go
//
// Generated by this command:
//
//	mockgen -package rpc -source peer_chooser.go -destination peer_chooser_mock.go -self_package github.com/uber/cadence/common/rpc
//

// Package rpc is a generated GoMock package.
package rpc

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	peer "go.uber.org/yarpc/api/peer"
	transport "go.uber.org/yarpc/api/transport"

	membership "github.com/uber/cadence/common/membership"
)

// MockPeerChooserFactory is a mock of PeerChooserFactory interface.
type MockPeerChooserFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPeerChooserFactoryMockRecorder
	isgomock struct{}
}

// MockPeerChooserFactoryMockRecorder is the mock recorder for MockPeerChooserFactory.
type MockPeerChooserFactoryMockRecorder struct {
	mock *MockPeerChooserFactory
}

// NewMockPeerChooserFactory creates a new mock instance.
func NewMockPeerChooserFactory(ctrl *gomock.Controller) *MockPeerChooserFactory {
	mock := &MockPeerChooserFactory{ctrl: ctrl}
	mock.recorder = &MockPeerChooserFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerChooserFactory) EXPECT() *MockPeerChooserFactoryMockRecorder {
	return m.recorder
}

// CreatePeerChooser mocks base method.
func (m *MockPeerChooserFactory) CreatePeerChooser(transport peer.Transport, opts PeerChooserOptions) (PeerChooser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeerChooser", transport, opts)
	ret0, _ := ret[0].(PeerChooser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePeerChooser indicates an expected call of CreatePeerChooser.
func (mr *MockPeerChooserFactoryMockRecorder) CreatePeerChooser(transport, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeerChooser", reflect.TypeOf((*MockPeerChooserFactory)(nil).CreatePeerChooser), transport, opts)
}

// MockPeerChooser is a mock of PeerChooser interface.
type MockPeerChooser struct {
	ctrl     *gomock.Controller
	recorder *MockPeerChooserMockRecorder
	isgomock struct{}
}

// MockPeerChooserMockRecorder is the mock recorder for MockPeerChooser.
type MockPeerChooserMockRecorder struct {
	mock *MockPeerChooser
}

// NewMockPeerChooser creates a new mock instance.
func NewMockPeerChooser(ctrl *gomock.Controller) *MockPeerChooser {
	mock := &MockPeerChooser{ctrl: ctrl}
	mock.recorder = &MockPeerChooserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerChooser) EXPECT() *MockPeerChooserMockRecorder {
	return m.recorder
}

// Choose mocks base method.
func (m *MockPeerChooser) Choose(arg0 context.Context, arg1 *transport.Request) (peer.Peer, func(error), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(peer.Peer)
	ret1, _ := ret[1].(func(error))
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Choose indicates an expected call of Choose.
func (mr *MockPeerChooserMockRecorder) Choose(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockPeerChooser)(nil).Choose), arg0, arg1)
}

// IsRunning mocks base method.
func (m *MockPeerChooser) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning.
func (mr *MockPeerChooserMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockPeerChooser)(nil).IsRunning))
}

// Start mocks base method.
func (m *MockPeerChooser) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockPeerChooserMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPeerChooser)(nil).Start))
}

// Stop mocks base method.
func (m *MockPeerChooser) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockPeerChooserMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockPeerChooser)(nil).Stop))
}

// UpdatePeers mocks base method.
func (m *MockPeerChooser) UpdatePeers(serviceName string, members []membership.HostInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePeers", serviceName, members)
}

// UpdatePeers indicates an expected call of UpdatePeers.
func (mr *MockPeerChooserMockRecorder) UpdatePeers(serviceName, members any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePeers", reflect.TypeOf((*MockPeerChooser)(nil).UpdatePeers), serviceName, members)
}
