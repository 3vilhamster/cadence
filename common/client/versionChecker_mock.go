// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/cadence/common/client (interfaces: VersionChecker)
//
// Generated by this command:
//
//	mockgen -package client -destination versionChecker_mock.go -self_package github.com/uber/cadence/common/client github.com/uber/cadence/common/client VersionChecker
//

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	shared "github.com/uber/cadence/.gen/go/shared"
	gomock "go.uber.org/mock/gomock"
)

// MockVersionChecker is a mock of VersionChecker interface.
type MockVersionChecker struct {
	ctrl     *gomock.Controller
	recorder *MockVersionCheckerMockRecorder
	isgomock struct{}
}

// MockVersionCheckerMockRecorder is the mock recorder for MockVersionChecker.
type MockVersionCheckerMockRecorder struct {
	mock *MockVersionChecker
}

// NewMockVersionChecker creates a new mock instance.
func NewMockVersionChecker(ctrl *gomock.Controller) *MockVersionChecker {
	mock := &MockVersionChecker{ctrl: ctrl}
	mock.recorder = &MockVersionCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionChecker) EXPECT() *MockVersionCheckerMockRecorder {
	return m.recorder
}

// ClientSupported mocks base method.
func (m *MockVersionChecker) ClientSupported(ctx context.Context, enableClientVersionCheck bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSupported", ctx, enableClientVersionCheck)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientSupported indicates an expected call of ClientSupported.
func (mr *MockVersionCheckerMockRecorder) ClientSupported(ctx, enableClientVersionCheck any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSupported", reflect.TypeOf((*MockVersionChecker)(nil).ClientSupported), ctx, enableClientVersionCheck)
}

// SupportsConsistentQuery mocks base method.
func (m *MockVersionChecker) SupportsConsistentQuery(clientImpl, clientFeatureVersion string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsConsistentQuery", clientImpl, clientFeatureVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// SupportsConsistentQuery indicates an expected call of SupportsConsistentQuery.
func (mr *MockVersionCheckerMockRecorder) SupportsConsistentQuery(clientImpl, clientFeatureVersion any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsConsistentQuery", reflect.TypeOf((*MockVersionChecker)(nil).SupportsConsistentQuery), clientImpl, clientFeatureVersion)
}

// SupportsRawHistoryQuery mocks base method.
func (m *MockVersionChecker) SupportsRawHistoryQuery(clientImpl, clientFeatureVersion string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsRawHistoryQuery", clientImpl, clientFeatureVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// SupportsRawHistoryQuery indicates an expected call of SupportsRawHistoryQuery.
func (mr *MockVersionCheckerMockRecorder) SupportsRawHistoryQuery(clientImpl, clientFeatureVersion any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsRawHistoryQuery", reflect.TypeOf((*MockVersionChecker)(nil).SupportsRawHistoryQuery), clientImpl, clientFeatureVersion)
}

// SupportsStickyQuery mocks base method.
func (m *MockVersionChecker) SupportsStickyQuery(clientImpl, clientFeatureVersion string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsStickyQuery", clientImpl, clientFeatureVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// SupportsStickyQuery indicates an expected call of SupportsStickyQuery.
func (mr *MockVersionCheckerMockRecorder) SupportsStickyQuery(clientImpl, clientFeatureVersion any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsStickyQuery", reflect.TypeOf((*MockVersionChecker)(nil).SupportsStickyQuery), clientImpl, clientFeatureVersion)
}

// SupportsWorkflowAlreadyCompletedError mocks base method.
func (m *MockVersionChecker) SupportsWorkflowAlreadyCompletedError(clientImpl, clientFeatureVersion string, featureFlags shared.FeatureFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsWorkflowAlreadyCompletedError", clientImpl, clientFeatureVersion, featureFlags)
	ret0, _ := ret[0].(error)
	return ret0
}

// SupportsWorkflowAlreadyCompletedError indicates an expected call of SupportsWorkflowAlreadyCompletedError.
func (mr *MockVersionCheckerMockRecorder) SupportsWorkflowAlreadyCompletedError(clientImpl, clientFeatureVersion, featureFlags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsWorkflowAlreadyCompletedError", reflect.TypeOf((*MockVersionChecker)(nil).SupportsWorkflowAlreadyCompletedError), clientImpl, clientFeatureVersion, featureFlags)
}
