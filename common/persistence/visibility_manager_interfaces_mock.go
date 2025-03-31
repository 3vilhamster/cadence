// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/cadence/common/persistence (interfaces: VisibilityManager)
//
// Generated by this command:
//
//	mockgen -package persistence -destination visibility_manager_interfaces_mock.go -self_package github.com/uber/cadence/common/persistence github.com/uber/cadence/common/persistence VisibilityManager
//

// Package persistence is a generated GoMock package.
package persistence

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockVisibilityManager is a mock of VisibilityManager interface.
type MockVisibilityManager struct {
	ctrl     *gomock.Controller
	recorder *MockVisibilityManagerMockRecorder
	isgomock struct{}
}

// MockVisibilityManagerMockRecorder is the mock recorder for MockVisibilityManager.
type MockVisibilityManagerMockRecorder struct {
	mock *MockVisibilityManager
}

// NewMockVisibilityManager creates a new mock instance.
func NewMockVisibilityManager(ctrl *gomock.Controller) *MockVisibilityManager {
	mock := &MockVisibilityManager{ctrl: ctrl}
	mock.recorder = &MockVisibilityManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVisibilityManager) EXPECT() *MockVisibilityManagerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockVisibilityManager) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockVisibilityManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockVisibilityManager)(nil).Close))
}

// CountWorkflowExecutions mocks base method.
func (m *MockVisibilityManager) CountWorkflowExecutions(ctx context.Context, request *CountWorkflowExecutionsRequest) (*CountWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountWorkflowExecutions", ctx, request)
	ret0, _ := ret[0].(*CountWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountWorkflowExecutions indicates an expected call of CountWorkflowExecutions.
func (mr *MockVisibilityManagerMockRecorder) CountWorkflowExecutions(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountWorkflowExecutions", reflect.TypeOf((*MockVisibilityManager)(nil).CountWorkflowExecutions), ctx, request)
}

// DeleteUninitializedWorkflowExecution mocks base method.
func (m *MockVisibilityManager) DeleteUninitializedWorkflowExecution(ctx context.Context, request *VisibilityDeleteWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUninitializedWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUninitializedWorkflowExecution indicates an expected call of DeleteUninitializedWorkflowExecution.
func (mr *MockVisibilityManagerMockRecorder) DeleteUninitializedWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUninitializedWorkflowExecution", reflect.TypeOf((*MockVisibilityManager)(nil).DeleteUninitializedWorkflowExecution), ctx, request)
}

// DeleteWorkflowExecution mocks base method.
func (m *MockVisibilityManager) DeleteWorkflowExecution(ctx context.Context, request *VisibilityDeleteWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflowExecution indicates an expected call of DeleteWorkflowExecution.
func (mr *MockVisibilityManagerMockRecorder) DeleteWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflowExecution", reflect.TypeOf((*MockVisibilityManager)(nil).DeleteWorkflowExecution), ctx, request)
}

// GetClosedWorkflowExecution mocks base method.
func (m *MockVisibilityManager) GetClosedWorkflowExecution(ctx context.Context, request *GetClosedWorkflowExecutionRequest) (*GetClosedWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClosedWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*GetClosedWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClosedWorkflowExecution indicates an expected call of GetClosedWorkflowExecution.
func (mr *MockVisibilityManagerMockRecorder) GetClosedWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClosedWorkflowExecution", reflect.TypeOf((*MockVisibilityManager)(nil).GetClosedWorkflowExecution), ctx, request)
}

// GetName mocks base method.
func (m *MockVisibilityManager) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockVisibilityManagerMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockVisibilityManager)(nil).GetName))
}

// ListClosedWorkflowExecutions mocks base method.
func (m *MockVisibilityManager) ListClosedWorkflowExecutions(ctx context.Context, request *ListWorkflowExecutionsRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutions", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutions indicates an expected call of ListClosedWorkflowExecutions.
func (mr *MockVisibilityManagerMockRecorder) ListClosedWorkflowExecutions(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutions", reflect.TypeOf((*MockVisibilityManager)(nil).ListClosedWorkflowExecutions), ctx, request)
}

// ListClosedWorkflowExecutionsByStatus mocks base method.
func (m *MockVisibilityManager) ListClosedWorkflowExecutionsByStatus(ctx context.Context, request *ListClosedWorkflowExecutionsByStatusRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutionsByStatus", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutionsByStatus indicates an expected call of ListClosedWorkflowExecutionsByStatus.
func (mr *MockVisibilityManagerMockRecorder) ListClosedWorkflowExecutionsByStatus(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutionsByStatus", reflect.TypeOf((*MockVisibilityManager)(nil).ListClosedWorkflowExecutionsByStatus), ctx, request)
}

// ListClosedWorkflowExecutionsByType mocks base method.
func (m *MockVisibilityManager) ListClosedWorkflowExecutionsByType(ctx context.Context, request *ListWorkflowExecutionsByTypeRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutionsByType", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutionsByType indicates an expected call of ListClosedWorkflowExecutionsByType.
func (mr *MockVisibilityManagerMockRecorder) ListClosedWorkflowExecutionsByType(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutionsByType", reflect.TypeOf((*MockVisibilityManager)(nil).ListClosedWorkflowExecutionsByType), ctx, request)
}

// ListClosedWorkflowExecutionsByWorkflowID mocks base method.
func (m *MockVisibilityManager) ListClosedWorkflowExecutionsByWorkflowID(ctx context.Context, request *ListWorkflowExecutionsByWorkflowIDRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutionsByWorkflowID", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutionsByWorkflowID indicates an expected call of ListClosedWorkflowExecutionsByWorkflowID.
func (mr *MockVisibilityManagerMockRecorder) ListClosedWorkflowExecutionsByWorkflowID(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutionsByWorkflowID", reflect.TypeOf((*MockVisibilityManager)(nil).ListClosedWorkflowExecutionsByWorkflowID), ctx, request)
}

// ListOpenWorkflowExecutions mocks base method.
func (m *MockVisibilityManager) ListOpenWorkflowExecutions(ctx context.Context, request *ListWorkflowExecutionsRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOpenWorkflowExecutions", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenWorkflowExecutions indicates an expected call of ListOpenWorkflowExecutions.
func (mr *MockVisibilityManagerMockRecorder) ListOpenWorkflowExecutions(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenWorkflowExecutions", reflect.TypeOf((*MockVisibilityManager)(nil).ListOpenWorkflowExecutions), ctx, request)
}

// ListOpenWorkflowExecutionsByType mocks base method.
func (m *MockVisibilityManager) ListOpenWorkflowExecutionsByType(ctx context.Context, request *ListWorkflowExecutionsByTypeRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOpenWorkflowExecutionsByType", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenWorkflowExecutionsByType indicates an expected call of ListOpenWorkflowExecutionsByType.
func (mr *MockVisibilityManagerMockRecorder) ListOpenWorkflowExecutionsByType(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenWorkflowExecutionsByType", reflect.TypeOf((*MockVisibilityManager)(nil).ListOpenWorkflowExecutionsByType), ctx, request)
}

// ListOpenWorkflowExecutionsByWorkflowID mocks base method.
func (m *MockVisibilityManager) ListOpenWorkflowExecutionsByWorkflowID(ctx context.Context, request *ListWorkflowExecutionsByWorkflowIDRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOpenWorkflowExecutionsByWorkflowID", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenWorkflowExecutionsByWorkflowID indicates an expected call of ListOpenWorkflowExecutionsByWorkflowID.
func (mr *MockVisibilityManagerMockRecorder) ListOpenWorkflowExecutionsByWorkflowID(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenWorkflowExecutionsByWorkflowID", reflect.TypeOf((*MockVisibilityManager)(nil).ListOpenWorkflowExecutionsByWorkflowID), ctx, request)
}

// ListWorkflowExecutions mocks base method.
func (m *MockVisibilityManager) ListWorkflowExecutions(ctx context.Context, request *ListWorkflowExecutionsByQueryRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkflowExecutions", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkflowExecutions indicates an expected call of ListWorkflowExecutions.
func (mr *MockVisibilityManagerMockRecorder) ListWorkflowExecutions(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflowExecutions", reflect.TypeOf((*MockVisibilityManager)(nil).ListWorkflowExecutions), ctx, request)
}

// RecordWorkflowExecutionClosed mocks base method.
func (m *MockVisibilityManager) RecordWorkflowExecutionClosed(ctx context.Context, request *RecordWorkflowExecutionClosedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordWorkflowExecutionClosed", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordWorkflowExecutionClosed indicates an expected call of RecordWorkflowExecutionClosed.
func (mr *MockVisibilityManagerMockRecorder) RecordWorkflowExecutionClosed(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordWorkflowExecutionClosed", reflect.TypeOf((*MockVisibilityManager)(nil).RecordWorkflowExecutionClosed), ctx, request)
}

// RecordWorkflowExecutionStarted mocks base method.
func (m *MockVisibilityManager) RecordWorkflowExecutionStarted(ctx context.Context, request *RecordWorkflowExecutionStartedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordWorkflowExecutionStarted", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordWorkflowExecutionStarted indicates an expected call of RecordWorkflowExecutionStarted.
func (mr *MockVisibilityManagerMockRecorder) RecordWorkflowExecutionStarted(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordWorkflowExecutionStarted", reflect.TypeOf((*MockVisibilityManager)(nil).RecordWorkflowExecutionStarted), ctx, request)
}

// RecordWorkflowExecutionUninitialized mocks base method.
func (m *MockVisibilityManager) RecordWorkflowExecutionUninitialized(ctx context.Context, request *RecordWorkflowExecutionUninitializedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordWorkflowExecutionUninitialized", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordWorkflowExecutionUninitialized indicates an expected call of RecordWorkflowExecutionUninitialized.
func (mr *MockVisibilityManagerMockRecorder) RecordWorkflowExecutionUninitialized(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordWorkflowExecutionUninitialized", reflect.TypeOf((*MockVisibilityManager)(nil).RecordWorkflowExecutionUninitialized), ctx, request)
}

// ScanWorkflowExecutions mocks base method.
func (m *MockVisibilityManager) ScanWorkflowExecutions(ctx context.Context, request *ListWorkflowExecutionsByQueryRequest) (*ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanWorkflowExecutions", ctx, request)
	ret0, _ := ret[0].(*ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanWorkflowExecutions indicates an expected call of ScanWorkflowExecutions.
func (mr *MockVisibilityManagerMockRecorder) ScanWorkflowExecutions(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanWorkflowExecutions", reflect.TypeOf((*MockVisibilityManager)(nil).ScanWorkflowExecutions), ctx, request)
}

// UpsertWorkflowExecution mocks base method.
func (m *MockVisibilityManager) UpsertWorkflowExecution(ctx context.Context, request *UpsertWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkflowExecution indicates an expected call of UpsertWorkflowExecution.
func (mr *MockVisibilityManagerMockRecorder) UpsertWorkflowExecution(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkflowExecution", reflect.TypeOf((*MockVisibilityManager)(nil).UpsertWorkflowExecution), ctx, request)
}
