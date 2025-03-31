// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -package api -source interface.go -destination interface_mock.go -self_package github.com/uber/cadence/service/frontend/api
//

// Package api is a generated GoMock package.
package api

import (
	context "context"
	reflect "reflect"

	types "github.com/uber/cadence/common/types"
	gomock "go.uber.org/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
	isgomock struct{}
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// CountWorkflowExecutions mocks base method.
func (m *MockHandler) CountWorkflowExecutions(arg0 context.Context, arg1 *types.CountWorkflowExecutionsRequest) (*types.CountWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*types.CountWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountWorkflowExecutions indicates an expected call of CountWorkflowExecutions.
func (mr *MockHandlerMockRecorder) CountWorkflowExecutions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountWorkflowExecutions", reflect.TypeOf((*MockHandler)(nil).CountWorkflowExecutions), arg0, arg1)
}

// DeprecateDomain mocks base method.
func (m *MockHandler) DeprecateDomain(arg0 context.Context, arg1 *types.DeprecateDomainRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeprecateDomain", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeprecateDomain indicates an expected call of DeprecateDomain.
func (mr *MockHandlerMockRecorder) DeprecateDomain(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeprecateDomain", reflect.TypeOf((*MockHandler)(nil).DeprecateDomain), arg0, arg1)
}

// DescribeDomain mocks base method.
func (m *MockHandler) DescribeDomain(arg0 context.Context, arg1 *types.DescribeDomainRequest) (*types.DescribeDomainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDomain", arg0, arg1)
	ret0, _ := ret[0].(*types.DescribeDomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDomain indicates an expected call of DescribeDomain.
func (mr *MockHandlerMockRecorder) DescribeDomain(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDomain", reflect.TypeOf((*MockHandler)(nil).DescribeDomain), arg0, arg1)
}

// DescribeTaskList mocks base method.
func (m *MockHandler) DescribeTaskList(arg0 context.Context, arg1 *types.DescribeTaskListRequest) (*types.DescribeTaskListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTaskList", arg0, arg1)
	ret0, _ := ret[0].(*types.DescribeTaskListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTaskList indicates an expected call of DescribeTaskList.
func (mr *MockHandlerMockRecorder) DescribeTaskList(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTaskList", reflect.TypeOf((*MockHandler)(nil).DescribeTaskList), arg0, arg1)
}

// DescribeWorkflowExecution mocks base method.
func (m *MockHandler) DescribeWorkflowExecution(arg0 context.Context, arg1 *types.DescribeWorkflowExecutionRequest) (*types.DescribeWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*types.DescribeWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkflowExecution indicates an expected call of DescribeWorkflowExecution.
func (mr *MockHandlerMockRecorder) DescribeWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).DescribeWorkflowExecution), arg0, arg1)
}

// DiagnoseWorkflowExecution mocks base method.
func (m *MockHandler) DiagnoseWorkflowExecution(arg0 context.Context, arg1 *types.DiagnoseWorkflowExecutionRequest) (*types.DiagnoseWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiagnoseWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*types.DiagnoseWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiagnoseWorkflowExecution indicates an expected call of DiagnoseWorkflowExecution.
func (mr *MockHandlerMockRecorder) DiagnoseWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiagnoseWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).DiagnoseWorkflowExecution), arg0, arg1)
}

// GetClusterInfo mocks base method.
func (m *MockHandler) GetClusterInfo(arg0 context.Context) (*types.ClusterInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInfo", arg0)
	ret0, _ := ret[0].(*types.ClusterInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInfo indicates an expected call of GetClusterInfo.
func (mr *MockHandlerMockRecorder) GetClusterInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInfo", reflect.TypeOf((*MockHandler)(nil).GetClusterInfo), arg0)
}

// GetSearchAttributes mocks base method.
func (m *MockHandler) GetSearchAttributes(arg0 context.Context) (*types.GetSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSearchAttributes", arg0)
	ret0, _ := ret[0].(*types.GetSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSearchAttributes indicates an expected call of GetSearchAttributes.
func (mr *MockHandlerMockRecorder) GetSearchAttributes(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSearchAttributes", reflect.TypeOf((*MockHandler)(nil).GetSearchAttributes), arg0)
}

// GetTaskListsByDomain mocks base method.
func (m *MockHandler) GetTaskListsByDomain(arg0 context.Context, arg1 *types.GetTaskListsByDomainRequest) (*types.GetTaskListsByDomainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskListsByDomain", arg0, arg1)
	ret0, _ := ret[0].(*types.GetTaskListsByDomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskListsByDomain indicates an expected call of GetTaskListsByDomain.
func (mr *MockHandlerMockRecorder) GetTaskListsByDomain(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskListsByDomain", reflect.TypeOf((*MockHandler)(nil).GetTaskListsByDomain), arg0, arg1)
}

// GetWorkflowExecutionHistory mocks base method.
func (m *MockHandler) GetWorkflowExecutionHistory(arg0 context.Context, arg1 *types.GetWorkflowExecutionHistoryRequest) (*types.GetWorkflowExecutionHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowExecutionHistory", arg0, arg1)
	ret0, _ := ret[0].(*types.GetWorkflowExecutionHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowExecutionHistory indicates an expected call of GetWorkflowExecutionHistory.
func (mr *MockHandlerMockRecorder) GetWorkflowExecutionHistory(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowExecutionHistory", reflect.TypeOf((*MockHandler)(nil).GetWorkflowExecutionHistory), arg0, arg1)
}

// Health mocks base method.
func (m *MockHandler) Health(arg0 context.Context) (*types.HealthStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(*types.HealthStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health.
func (mr *MockHandlerMockRecorder) Health(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockHandler)(nil).Health), arg0)
}

// ListArchivedWorkflowExecutions mocks base method.
func (m *MockHandler) ListArchivedWorkflowExecutions(arg0 context.Context, arg1 *types.ListArchivedWorkflowExecutionsRequest) (*types.ListArchivedWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListArchivedWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*types.ListArchivedWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArchivedWorkflowExecutions indicates an expected call of ListArchivedWorkflowExecutions.
func (mr *MockHandlerMockRecorder) ListArchivedWorkflowExecutions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArchivedWorkflowExecutions", reflect.TypeOf((*MockHandler)(nil).ListArchivedWorkflowExecutions), arg0, arg1)
}

// ListClosedWorkflowExecutions mocks base method.
func (m *MockHandler) ListClosedWorkflowExecutions(arg0 context.Context, arg1 *types.ListClosedWorkflowExecutionsRequest) (*types.ListClosedWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*types.ListClosedWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutions indicates an expected call of ListClosedWorkflowExecutions.
func (mr *MockHandlerMockRecorder) ListClosedWorkflowExecutions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutions", reflect.TypeOf((*MockHandler)(nil).ListClosedWorkflowExecutions), arg0, arg1)
}

// ListDomains mocks base method.
func (m *MockHandler) ListDomains(arg0 context.Context, arg1 *types.ListDomainsRequest) (*types.ListDomainsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDomains", arg0, arg1)
	ret0, _ := ret[0].(*types.ListDomainsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDomains indicates an expected call of ListDomains.
func (mr *MockHandlerMockRecorder) ListDomains(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomains", reflect.TypeOf((*MockHandler)(nil).ListDomains), arg0, arg1)
}

// ListOpenWorkflowExecutions mocks base method.
func (m *MockHandler) ListOpenWorkflowExecutions(arg0 context.Context, arg1 *types.ListOpenWorkflowExecutionsRequest) (*types.ListOpenWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOpenWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*types.ListOpenWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenWorkflowExecutions indicates an expected call of ListOpenWorkflowExecutions.
func (mr *MockHandlerMockRecorder) ListOpenWorkflowExecutions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenWorkflowExecutions", reflect.TypeOf((*MockHandler)(nil).ListOpenWorkflowExecutions), arg0, arg1)
}

// ListTaskListPartitions mocks base method.
func (m *MockHandler) ListTaskListPartitions(arg0 context.Context, arg1 *types.ListTaskListPartitionsRequest) (*types.ListTaskListPartitionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTaskListPartitions", arg0, arg1)
	ret0, _ := ret[0].(*types.ListTaskListPartitionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTaskListPartitions indicates an expected call of ListTaskListPartitions.
func (mr *MockHandlerMockRecorder) ListTaskListPartitions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTaskListPartitions", reflect.TypeOf((*MockHandler)(nil).ListTaskListPartitions), arg0, arg1)
}

// ListWorkflowExecutions mocks base method.
func (m *MockHandler) ListWorkflowExecutions(arg0 context.Context, arg1 *types.ListWorkflowExecutionsRequest) (*types.ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*types.ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkflowExecutions indicates an expected call of ListWorkflowExecutions.
func (mr *MockHandlerMockRecorder) ListWorkflowExecutions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflowExecutions", reflect.TypeOf((*MockHandler)(nil).ListWorkflowExecutions), arg0, arg1)
}

// PollForActivityTask mocks base method.
func (m *MockHandler) PollForActivityTask(arg0 context.Context, arg1 *types.PollForActivityTaskRequest) (*types.PollForActivityTaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollForActivityTask", arg0, arg1)
	ret0, _ := ret[0].(*types.PollForActivityTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollForActivityTask indicates an expected call of PollForActivityTask.
func (mr *MockHandlerMockRecorder) PollForActivityTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollForActivityTask", reflect.TypeOf((*MockHandler)(nil).PollForActivityTask), arg0, arg1)
}

// PollForDecisionTask mocks base method.
func (m *MockHandler) PollForDecisionTask(arg0 context.Context, arg1 *types.PollForDecisionTaskRequest) (*types.PollForDecisionTaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollForDecisionTask", arg0, arg1)
	ret0, _ := ret[0].(*types.PollForDecisionTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollForDecisionTask indicates an expected call of PollForDecisionTask.
func (mr *MockHandlerMockRecorder) PollForDecisionTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollForDecisionTask", reflect.TypeOf((*MockHandler)(nil).PollForDecisionTask), arg0, arg1)
}

// QueryWorkflow mocks base method.
func (m *MockHandler) QueryWorkflow(arg0 context.Context, arg1 *types.QueryWorkflowRequest) (*types.QueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryWorkflow", arg0, arg1)
	ret0, _ := ret[0].(*types.QueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryWorkflow indicates an expected call of QueryWorkflow.
func (mr *MockHandlerMockRecorder) QueryWorkflow(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryWorkflow", reflect.TypeOf((*MockHandler)(nil).QueryWorkflow), arg0, arg1)
}

// RecordActivityTaskHeartbeat mocks base method.
func (m *MockHandler) RecordActivityTaskHeartbeat(arg0 context.Context, arg1 *types.RecordActivityTaskHeartbeatRequest) (*types.RecordActivityTaskHeartbeatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordActivityTaskHeartbeat", arg0, arg1)
	ret0, _ := ret[0].(*types.RecordActivityTaskHeartbeatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskHeartbeat indicates an expected call of RecordActivityTaskHeartbeat.
func (mr *MockHandlerMockRecorder) RecordActivityTaskHeartbeat(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskHeartbeat", reflect.TypeOf((*MockHandler)(nil).RecordActivityTaskHeartbeat), arg0, arg1)
}

// RecordActivityTaskHeartbeatByID mocks base method.
func (m *MockHandler) RecordActivityTaskHeartbeatByID(arg0 context.Context, arg1 *types.RecordActivityTaskHeartbeatByIDRequest) (*types.RecordActivityTaskHeartbeatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordActivityTaskHeartbeatByID", arg0, arg1)
	ret0, _ := ret[0].(*types.RecordActivityTaskHeartbeatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskHeartbeatByID indicates an expected call of RecordActivityTaskHeartbeatByID.
func (mr *MockHandlerMockRecorder) RecordActivityTaskHeartbeatByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskHeartbeatByID", reflect.TypeOf((*MockHandler)(nil).RecordActivityTaskHeartbeatByID), arg0, arg1)
}

// RefreshWorkflowTasks mocks base method.
func (m *MockHandler) RefreshWorkflowTasks(arg0 context.Context, arg1 *types.RefreshWorkflowTasksRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWorkflowTasks", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshWorkflowTasks indicates an expected call of RefreshWorkflowTasks.
func (mr *MockHandlerMockRecorder) RefreshWorkflowTasks(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWorkflowTasks", reflect.TypeOf((*MockHandler)(nil).RefreshWorkflowTasks), arg0, arg1)
}

// RegisterDomain mocks base method.
func (m *MockHandler) RegisterDomain(arg0 context.Context, arg1 *types.RegisterDomainRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterDomain", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterDomain indicates an expected call of RegisterDomain.
func (mr *MockHandlerMockRecorder) RegisterDomain(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterDomain", reflect.TypeOf((*MockHandler)(nil).RegisterDomain), arg0, arg1)
}

// RequestCancelWorkflowExecution mocks base method.
func (m *MockHandler) RequestCancelWorkflowExecution(arg0 context.Context, arg1 *types.RequestCancelWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestCancelWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestCancelWorkflowExecution indicates an expected call of RequestCancelWorkflowExecution.
func (mr *MockHandlerMockRecorder) RequestCancelWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCancelWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).RequestCancelWorkflowExecution), arg0, arg1)
}

// ResetStickyTaskList mocks base method.
func (m *MockHandler) ResetStickyTaskList(arg0 context.Context, arg1 *types.ResetStickyTaskListRequest) (*types.ResetStickyTaskListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetStickyTaskList", arg0, arg1)
	ret0, _ := ret[0].(*types.ResetStickyTaskListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetStickyTaskList indicates an expected call of ResetStickyTaskList.
func (mr *MockHandlerMockRecorder) ResetStickyTaskList(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetStickyTaskList", reflect.TypeOf((*MockHandler)(nil).ResetStickyTaskList), arg0, arg1)
}

// ResetWorkflowExecution mocks base method.
func (m *MockHandler) ResetWorkflowExecution(arg0 context.Context, arg1 *types.ResetWorkflowExecutionRequest) (*types.ResetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*types.ResetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetWorkflowExecution indicates an expected call of ResetWorkflowExecution.
func (mr *MockHandlerMockRecorder) ResetWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).ResetWorkflowExecution), arg0, arg1)
}

// RespondActivityTaskCanceled mocks base method.
func (m *MockHandler) RespondActivityTaskCanceled(arg0 context.Context, arg1 *types.RespondActivityTaskCanceledRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCanceled", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCanceled indicates an expected call of RespondActivityTaskCanceled.
func (mr *MockHandlerMockRecorder) RespondActivityTaskCanceled(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCanceled", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskCanceled), arg0, arg1)
}

// RespondActivityTaskCanceledByID mocks base method.
func (m *MockHandler) RespondActivityTaskCanceledByID(arg0 context.Context, arg1 *types.RespondActivityTaskCanceledByIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCanceledByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCanceledByID indicates an expected call of RespondActivityTaskCanceledByID.
func (mr *MockHandlerMockRecorder) RespondActivityTaskCanceledByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCanceledByID", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskCanceledByID), arg0, arg1)
}

// RespondActivityTaskCompleted mocks base method.
func (m *MockHandler) RespondActivityTaskCompleted(arg0 context.Context, arg1 *types.RespondActivityTaskCompletedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCompleted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCompleted indicates an expected call of RespondActivityTaskCompleted.
func (mr *MockHandlerMockRecorder) RespondActivityTaskCompleted(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCompleted", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskCompleted), arg0, arg1)
}

// RespondActivityTaskCompletedByID mocks base method.
func (m *MockHandler) RespondActivityTaskCompletedByID(arg0 context.Context, arg1 *types.RespondActivityTaskCompletedByIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCompletedByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCompletedByID indicates an expected call of RespondActivityTaskCompletedByID.
func (mr *MockHandlerMockRecorder) RespondActivityTaskCompletedByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCompletedByID", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskCompletedByID), arg0, arg1)
}

// RespondActivityTaskFailed mocks base method.
func (m *MockHandler) RespondActivityTaskFailed(arg0 context.Context, arg1 *types.RespondActivityTaskFailedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskFailed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskFailed indicates an expected call of RespondActivityTaskFailed.
func (mr *MockHandlerMockRecorder) RespondActivityTaskFailed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskFailed", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskFailed), arg0, arg1)
}

// RespondActivityTaskFailedByID mocks base method.
func (m *MockHandler) RespondActivityTaskFailedByID(arg0 context.Context, arg1 *types.RespondActivityTaskFailedByIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskFailedByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskFailedByID indicates an expected call of RespondActivityTaskFailedByID.
func (mr *MockHandlerMockRecorder) RespondActivityTaskFailedByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskFailedByID", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskFailedByID), arg0, arg1)
}

// RespondDecisionTaskCompleted mocks base method.
func (m *MockHandler) RespondDecisionTaskCompleted(arg0 context.Context, arg1 *types.RespondDecisionTaskCompletedRequest) (*types.RespondDecisionTaskCompletedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondDecisionTaskCompleted", arg0, arg1)
	ret0, _ := ret[0].(*types.RespondDecisionTaskCompletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RespondDecisionTaskCompleted indicates an expected call of RespondDecisionTaskCompleted.
func (mr *MockHandlerMockRecorder) RespondDecisionTaskCompleted(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondDecisionTaskCompleted", reflect.TypeOf((*MockHandler)(nil).RespondDecisionTaskCompleted), arg0, arg1)
}

// RespondDecisionTaskFailed mocks base method.
func (m *MockHandler) RespondDecisionTaskFailed(arg0 context.Context, arg1 *types.RespondDecisionTaskFailedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondDecisionTaskFailed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondDecisionTaskFailed indicates an expected call of RespondDecisionTaskFailed.
func (mr *MockHandlerMockRecorder) RespondDecisionTaskFailed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondDecisionTaskFailed", reflect.TypeOf((*MockHandler)(nil).RespondDecisionTaskFailed), arg0, arg1)
}

// RespondQueryTaskCompleted mocks base method.
func (m *MockHandler) RespondQueryTaskCompleted(arg0 context.Context, arg1 *types.RespondQueryTaskCompletedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondQueryTaskCompleted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondQueryTaskCompleted indicates an expected call of RespondQueryTaskCompleted.
func (mr *MockHandlerMockRecorder) RespondQueryTaskCompleted(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondQueryTaskCompleted", reflect.TypeOf((*MockHandler)(nil).RespondQueryTaskCompleted), arg0, arg1)
}

// RestartWorkflowExecution mocks base method.
func (m *MockHandler) RestartWorkflowExecution(arg0 context.Context, arg1 *types.RestartWorkflowExecutionRequest) (*types.RestartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*types.RestartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestartWorkflowExecution indicates an expected call of RestartWorkflowExecution.
func (mr *MockHandlerMockRecorder) RestartWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).RestartWorkflowExecution), arg0, arg1)
}

// ScanWorkflowExecutions mocks base method.
func (m *MockHandler) ScanWorkflowExecutions(arg0 context.Context, arg1 *types.ListWorkflowExecutionsRequest) (*types.ListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*types.ListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanWorkflowExecutions indicates an expected call of ScanWorkflowExecutions.
func (mr *MockHandlerMockRecorder) ScanWorkflowExecutions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanWorkflowExecutions", reflect.TypeOf((*MockHandler)(nil).ScanWorkflowExecutions), arg0, arg1)
}

// SignalWithStartWorkflowExecution mocks base method.
func (m *MockHandler) SignalWithStartWorkflowExecution(arg0 context.Context, arg1 *types.SignalWithStartWorkflowExecutionRequest) (*types.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWithStartWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*types.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignalWithStartWorkflowExecution indicates an expected call of SignalWithStartWorkflowExecution.
func (mr *MockHandlerMockRecorder) SignalWithStartWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWithStartWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).SignalWithStartWorkflowExecution), arg0, arg1)
}

// SignalWithStartWorkflowExecutionAsync mocks base method.
func (m *MockHandler) SignalWithStartWorkflowExecutionAsync(arg0 context.Context, arg1 *types.SignalWithStartWorkflowExecutionAsyncRequest) (*types.SignalWithStartWorkflowExecutionAsyncResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWithStartWorkflowExecutionAsync", arg0, arg1)
	ret0, _ := ret[0].(*types.SignalWithStartWorkflowExecutionAsyncResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignalWithStartWorkflowExecutionAsync indicates an expected call of SignalWithStartWorkflowExecutionAsync.
func (mr *MockHandlerMockRecorder) SignalWithStartWorkflowExecutionAsync(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWithStartWorkflowExecutionAsync", reflect.TypeOf((*MockHandler)(nil).SignalWithStartWorkflowExecutionAsync), arg0, arg1)
}

// SignalWorkflowExecution mocks base method.
func (m *MockHandler) SignalWorkflowExecution(arg0 context.Context, arg1 *types.SignalWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignalWorkflowExecution indicates an expected call of SignalWorkflowExecution.
func (mr *MockHandlerMockRecorder) SignalWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).SignalWorkflowExecution), arg0, arg1)
}

// StartWorkflowExecution mocks base method.
func (m *MockHandler) StartWorkflowExecution(arg0 context.Context, arg1 *types.StartWorkflowExecutionRequest) (*types.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*types.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartWorkflowExecution indicates an expected call of StartWorkflowExecution.
func (mr *MockHandlerMockRecorder) StartWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).StartWorkflowExecution), arg0, arg1)
}

// StartWorkflowExecutionAsync mocks base method.
func (m *MockHandler) StartWorkflowExecutionAsync(arg0 context.Context, arg1 *types.StartWorkflowExecutionAsyncRequest) (*types.StartWorkflowExecutionAsyncResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorkflowExecutionAsync", arg0, arg1)
	ret0, _ := ret[0].(*types.StartWorkflowExecutionAsyncResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartWorkflowExecutionAsync indicates an expected call of StartWorkflowExecutionAsync.
func (mr *MockHandlerMockRecorder) StartWorkflowExecutionAsync(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflowExecutionAsync", reflect.TypeOf((*MockHandler)(nil).StartWorkflowExecutionAsync), arg0, arg1)
}

// TerminateWorkflowExecution mocks base method.
func (m *MockHandler) TerminateWorkflowExecution(arg0 context.Context, arg1 *types.TerminateWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateWorkflowExecution indicates an expected call of TerminateWorkflowExecution.
func (mr *MockHandlerMockRecorder) TerminateWorkflowExecution(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).TerminateWorkflowExecution), arg0, arg1)
}

// UpdateDomain mocks base method.
func (m *MockHandler) UpdateDomain(arg0 context.Context, arg1 *types.UpdateDomainRequest) (*types.UpdateDomainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDomain", arg0, arg1)
	ret0, _ := ret[0].(*types.UpdateDomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDomain indicates an expected call of UpdateDomain.
func (mr *MockHandlerMockRecorder) UpdateDomain(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDomain", reflect.TypeOf((*MockHandler)(nil).UpdateDomain), arg0, arg1)
}
