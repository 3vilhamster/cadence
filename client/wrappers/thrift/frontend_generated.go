package thrift

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/thrift.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/common/types"
	"github.com/uber/cadence/common/types/mapper/thrift"
)

func (g frontendClient) CountWorkflowExecutions(ctx context.Context, cp1 *types.CountWorkflowExecutionsRequest, p1 ...yarpc.CallOption) (cp2 *types.CountWorkflowExecutionsResponse, err error) {
	response, err := g.c.CountWorkflowExecutions(ctx, thrift.FromCountWorkflowExecutionsRequest(cp1), p1...)
	return thrift.ToCountWorkflowExecutionsResponse(response), thrift.ToError(err)
}

func (g frontendClient) DeprecateDomain(ctx context.Context, dp1 *types.DeprecateDomainRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.DeprecateDomain(ctx, thrift.FromDeprecateDomainRequest(dp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) DescribeDomain(ctx context.Context, dp1 *types.DescribeDomainRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeDomainResponse, err error) {
	response, err := g.c.DescribeDomain(ctx, thrift.FromDescribeDomainRequest(dp1), p1...)
	return thrift.ToDescribeDomainResponse(response), thrift.ToError(err)
}

func (g frontendClient) DescribeTaskList(ctx context.Context, dp1 *types.DescribeTaskListRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeTaskListResponse, err error) {
	response, err := g.c.DescribeTaskList(ctx, thrift.FromDescribeTaskListRequest(dp1), p1...)
	return thrift.ToDescribeTaskListResponse(response), thrift.ToError(err)
}

func (g frontendClient) DescribeWorkflowExecution(ctx context.Context, dp1 *types.DescribeWorkflowExecutionRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeWorkflowExecutionResponse, err error) {
	response, err := g.c.DescribeWorkflowExecution(ctx, thrift.FromDescribeWorkflowExecutionRequest(dp1), p1...)
	return thrift.ToDescribeWorkflowExecutionResponse(response), thrift.ToError(err)
}

func (g frontendClient) DiagnoseWorkflowExecution(ctx context.Context, dp1 *types.DiagnoseWorkflowExecutionRequest, p1 ...yarpc.CallOption) (dp2 *types.DiagnoseWorkflowExecutionResponse, err error) {
	response, err := g.c.DiagnoseWorkflowExecution(ctx, thrift.FromDiagnoseWorkflowExecutionRequest(dp1), p1...)
	return thrift.ToDiagnoseWorkflowExecutionResponse(response), thrift.ToError(err)
}

func (g frontendClient) GetClusterInfo(ctx context.Context, p1 ...yarpc.CallOption) (cp1 *types.ClusterInfo, err error) {
	response, err := g.c.GetClusterInfo(ctx, p1...)
	return thrift.ToGetClusterInfoResponse(response), thrift.ToError(err)
}

func (g frontendClient) GetSearchAttributes(ctx context.Context, p1 ...yarpc.CallOption) (gp1 *types.GetSearchAttributesResponse, err error) {
	response, err := g.c.GetSearchAttributes(ctx, p1...)
	return thrift.ToGetSearchAttributesResponse(response), thrift.ToError(err)
}

func (g frontendClient) GetTaskListsByDomain(ctx context.Context, gp1 *types.GetTaskListsByDomainRequest, p1 ...yarpc.CallOption) (gp2 *types.GetTaskListsByDomainResponse, err error) {
	response, err := g.c.GetTaskListsByDomain(ctx, thrift.FromGetTaskListsByDomainRequest(gp1), p1...)
	return thrift.ToGetTaskListsByDomainResponse(response), thrift.ToError(err)
}

func (g frontendClient) GetWorkflowExecutionHistory(ctx context.Context, gp1 *types.GetWorkflowExecutionHistoryRequest, p1 ...yarpc.CallOption) (gp2 *types.GetWorkflowExecutionHistoryResponse, err error) {
	response, err := g.c.GetWorkflowExecutionHistory(ctx, thrift.FromGetWorkflowExecutionHistoryRequest(gp1), p1...)
	return thrift.ToGetWorkflowExecutionHistoryResponse(response), thrift.ToError(err)
}

func (g frontendClient) ListArchivedWorkflowExecutions(ctx context.Context, lp1 *types.ListArchivedWorkflowExecutionsRequest, p1 ...yarpc.CallOption) (lp2 *types.ListArchivedWorkflowExecutionsResponse, err error) {
	response, err := g.c.ListArchivedWorkflowExecutions(ctx, thrift.FromListArchivedWorkflowExecutionsRequest(lp1), p1...)
	return thrift.ToListArchivedWorkflowExecutionsResponse(response), thrift.ToError(err)
}

func (g frontendClient) ListClosedWorkflowExecutions(ctx context.Context, lp1 *types.ListClosedWorkflowExecutionsRequest, p1 ...yarpc.CallOption) (lp2 *types.ListClosedWorkflowExecutionsResponse, err error) {
	response, err := g.c.ListClosedWorkflowExecutions(ctx, thrift.FromListClosedWorkflowExecutionsRequest(lp1), p1...)
	return thrift.ToListClosedWorkflowExecutionsResponse(response), thrift.ToError(err)
}

func (g frontendClient) ListDomains(ctx context.Context, lp1 *types.ListDomainsRequest, p1 ...yarpc.CallOption) (lp2 *types.ListDomainsResponse, err error) {
	response, err := g.c.ListDomains(ctx, thrift.FromListDomainsRequest(lp1), p1...)
	return thrift.ToListDomainsResponse(response), thrift.ToError(err)
}

func (g frontendClient) ListOpenWorkflowExecutions(ctx context.Context, lp1 *types.ListOpenWorkflowExecutionsRequest, p1 ...yarpc.CallOption) (lp2 *types.ListOpenWorkflowExecutionsResponse, err error) {
	response, err := g.c.ListOpenWorkflowExecutions(ctx, thrift.FromListOpenWorkflowExecutionsRequest(lp1), p1...)
	return thrift.ToListOpenWorkflowExecutionsResponse(response), thrift.ToError(err)
}

func (g frontendClient) ListTaskListPartitions(ctx context.Context, lp1 *types.ListTaskListPartitionsRequest, p1 ...yarpc.CallOption) (lp2 *types.ListTaskListPartitionsResponse, err error) {
	response, err := g.c.ListTaskListPartitions(ctx, thrift.FromListTaskListPartitionsRequest(lp1), p1...)
	return thrift.ToListTaskListPartitionsResponse(response), thrift.ToError(err)
}

func (g frontendClient) ListWorkflowExecutions(ctx context.Context, lp1 *types.ListWorkflowExecutionsRequest, p1 ...yarpc.CallOption) (lp2 *types.ListWorkflowExecutionsResponse, err error) {
	response, err := g.c.ListWorkflowExecutions(ctx, thrift.FromListWorkflowExecutionsRequest(lp1), p1...)
	return thrift.ToListWorkflowExecutionsResponse(response), thrift.ToError(err)
}

func (g frontendClient) PollForActivityTask(ctx context.Context, pp1 *types.PollForActivityTaskRequest, p1 ...yarpc.CallOption) (pp2 *types.PollForActivityTaskResponse, err error) {
	response, err := g.c.PollForActivityTask(ctx, thrift.FromPollForActivityTaskRequest(pp1), p1...)
	return thrift.ToPollForActivityTaskResponse(response), thrift.ToError(err)
}

func (g frontendClient) PollForDecisionTask(ctx context.Context, pp1 *types.PollForDecisionTaskRequest, p1 ...yarpc.CallOption) (pp2 *types.PollForDecisionTaskResponse, err error) {
	response, err := g.c.PollForDecisionTask(ctx, thrift.FromPollForDecisionTaskRequest(pp1), p1...)
	return thrift.ToPollForDecisionTaskResponse(response), thrift.ToError(err)
}

func (g frontendClient) QueryWorkflow(ctx context.Context, qp1 *types.QueryWorkflowRequest, p1 ...yarpc.CallOption) (qp2 *types.QueryWorkflowResponse, err error) {
	response, err := g.c.QueryWorkflow(ctx, thrift.FromQueryWorkflowRequest(qp1), p1...)
	return thrift.ToQueryWorkflowResponse(response), thrift.ToError(err)
}

func (g frontendClient) RecordActivityTaskHeartbeat(ctx context.Context, rp1 *types.RecordActivityTaskHeartbeatRequest, p1 ...yarpc.CallOption) (rp2 *types.RecordActivityTaskHeartbeatResponse, err error) {
	response, err := g.c.RecordActivityTaskHeartbeat(ctx, thrift.FromRecordActivityTaskHeartbeatRequest(rp1), p1...)
	return thrift.ToRecordActivityTaskHeartbeatResponse(response), thrift.ToError(err)
}

func (g frontendClient) RecordActivityTaskHeartbeatByID(ctx context.Context, rp1 *types.RecordActivityTaskHeartbeatByIDRequest, p1 ...yarpc.CallOption) (rp2 *types.RecordActivityTaskHeartbeatResponse, err error) {
	response, err := g.c.RecordActivityTaskHeartbeatByID(ctx, thrift.FromRecordActivityTaskHeartbeatByIDRequest(rp1), p1...)
	return thrift.ToRecordActivityTaskHeartbeatByIDResponse(response), thrift.ToError(err)
}

func (g frontendClient) RefreshWorkflowTasks(ctx context.Context, rp1 *types.RefreshWorkflowTasksRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RefreshWorkflowTasks(ctx, thrift.FromRefreshWorkflowTasksRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RegisterDomain(ctx context.Context, rp1 *types.RegisterDomainRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RegisterDomain(ctx, thrift.FromRegisterDomainRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RequestCancelWorkflowExecution(ctx context.Context, rp1 *types.RequestCancelWorkflowExecutionRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RequestCancelWorkflowExecution(ctx, thrift.FromRequestCancelWorkflowExecutionRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) ResetStickyTaskList(ctx context.Context, rp1 *types.ResetStickyTaskListRequest, p1 ...yarpc.CallOption) (rp2 *types.ResetStickyTaskListResponse, err error) {
	response, err := g.c.ResetStickyTaskList(ctx, thrift.FromResetStickyTaskListRequest(rp1), p1...)
	return thrift.ToResetStickyTaskListResponse(response), thrift.ToError(err)
}

func (g frontendClient) ResetWorkflowExecution(ctx context.Context, rp1 *types.ResetWorkflowExecutionRequest, p1 ...yarpc.CallOption) (rp2 *types.ResetWorkflowExecutionResponse, err error) {
	response, err := g.c.ResetWorkflowExecution(ctx, thrift.FromResetWorkflowExecutionRequest(rp1), p1...)
	return thrift.ToResetWorkflowExecutionResponse(response), thrift.ToError(err)
}

func (g frontendClient) RespondActivityTaskCanceled(ctx context.Context, rp1 *types.RespondActivityTaskCanceledRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondActivityTaskCanceled(ctx, thrift.FromRespondActivityTaskCanceledRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RespondActivityTaskCanceledByID(ctx context.Context, rp1 *types.RespondActivityTaskCanceledByIDRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondActivityTaskCanceledByID(ctx, thrift.FromRespondActivityTaskCanceledByIDRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RespondActivityTaskCompleted(ctx context.Context, rp1 *types.RespondActivityTaskCompletedRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondActivityTaskCompleted(ctx, thrift.FromRespondActivityTaskCompletedRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RespondActivityTaskCompletedByID(ctx context.Context, rp1 *types.RespondActivityTaskCompletedByIDRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondActivityTaskCompletedByID(ctx, thrift.FromRespondActivityTaskCompletedByIDRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RespondActivityTaskFailed(ctx context.Context, rp1 *types.RespondActivityTaskFailedRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondActivityTaskFailed(ctx, thrift.FromRespondActivityTaskFailedRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RespondActivityTaskFailedByID(ctx context.Context, rp1 *types.RespondActivityTaskFailedByIDRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondActivityTaskFailedByID(ctx, thrift.FromRespondActivityTaskFailedByIDRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RespondDecisionTaskCompleted(ctx context.Context, rp1 *types.RespondDecisionTaskCompletedRequest, p1 ...yarpc.CallOption) (rp2 *types.RespondDecisionTaskCompletedResponse, err error) {
	response, err := g.c.RespondDecisionTaskCompleted(ctx, thrift.FromRespondDecisionTaskCompletedRequest(rp1), p1...)
	return thrift.ToRespondDecisionTaskCompletedResponse(response), thrift.ToError(err)
}

func (g frontendClient) RespondDecisionTaskFailed(ctx context.Context, rp1 *types.RespondDecisionTaskFailedRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondDecisionTaskFailed(ctx, thrift.FromRespondDecisionTaskFailedRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RespondQueryTaskCompleted(ctx context.Context, rp1 *types.RespondQueryTaskCompletedRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.RespondQueryTaskCompleted(ctx, thrift.FromRespondQueryTaskCompletedRequest(rp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) RestartWorkflowExecution(ctx context.Context, rp1 *types.RestartWorkflowExecutionRequest, p1 ...yarpc.CallOption) (rp2 *types.RestartWorkflowExecutionResponse, err error) {
	response, err := g.c.RestartWorkflowExecution(ctx, thrift.FromRestartWorkflowExecutionRequest(rp1), p1...)
	return thrift.ToRestartWorkflowExecutionResponse(response), thrift.ToError(err)
}

func (g frontendClient) ScanWorkflowExecutions(ctx context.Context, lp1 *types.ListWorkflowExecutionsRequest, p1 ...yarpc.CallOption) (lp2 *types.ListWorkflowExecutionsResponse, err error) {
	response, err := g.c.ScanWorkflowExecutions(ctx, thrift.FromScanWorkflowExecutionsRequest(lp1), p1...)
	return thrift.ToScanWorkflowExecutionsResponse(response), thrift.ToError(err)
}

func (g frontendClient) SignalWithStartWorkflowExecution(ctx context.Context, sp1 *types.SignalWithStartWorkflowExecutionRequest, p1 ...yarpc.CallOption) (sp2 *types.StartWorkflowExecutionResponse, err error) {
	response, err := g.c.SignalWithStartWorkflowExecution(ctx, thrift.FromSignalWithStartWorkflowExecutionRequest(sp1), p1...)
	return thrift.ToSignalWithStartWorkflowExecutionResponse(response), thrift.ToError(err)
}

func (g frontendClient) SignalWithStartWorkflowExecutionAsync(ctx context.Context, sp1 *types.SignalWithStartWorkflowExecutionAsyncRequest, p1 ...yarpc.CallOption) (sp2 *types.SignalWithStartWorkflowExecutionAsyncResponse, err error) {
	response, err := g.c.SignalWithStartWorkflowExecutionAsync(ctx, thrift.FromSignalWithStartWorkflowExecutionAsyncRequest(sp1), p1...)
	return thrift.ToSignalWithStartWorkflowExecutionAsyncResponse(response), thrift.ToError(err)
}

func (g frontendClient) SignalWorkflowExecution(ctx context.Context, sp1 *types.SignalWorkflowExecutionRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.SignalWorkflowExecution(ctx, thrift.FromSignalWorkflowExecutionRequest(sp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) StartWorkflowExecution(ctx context.Context, sp1 *types.StartWorkflowExecutionRequest, p1 ...yarpc.CallOption) (sp2 *types.StartWorkflowExecutionResponse, err error) {
	response, err := g.c.StartWorkflowExecution(ctx, thrift.FromStartWorkflowExecutionRequest(sp1), p1...)
	return thrift.ToStartWorkflowExecutionResponse(response), thrift.ToError(err)
}

func (g frontendClient) StartWorkflowExecutionAsync(ctx context.Context, sp1 *types.StartWorkflowExecutionAsyncRequest, p1 ...yarpc.CallOption) (sp2 *types.StartWorkflowExecutionAsyncResponse, err error) {
	response, err := g.c.StartWorkflowExecutionAsync(ctx, thrift.FromStartWorkflowExecutionAsyncRequest(sp1), p1...)
	return thrift.ToStartWorkflowExecutionAsyncResponse(response), thrift.ToError(err)
}

func (g frontendClient) TerminateWorkflowExecution(ctx context.Context, tp1 *types.TerminateWorkflowExecutionRequest, p1 ...yarpc.CallOption) (err error) {
	err = g.c.TerminateWorkflowExecution(ctx, thrift.FromTerminateWorkflowExecutionRequest(tp1), p1...)
	return thrift.ToError(err)
}

func (g frontendClient) UpdateDomain(ctx context.Context, up1 *types.UpdateDomainRequest, p1 ...yarpc.CallOption) (up2 *types.UpdateDomainResponse, err error) {
	response, err := g.c.UpdateDomain(ctx, thrift.FromUpdateDomainRequest(up1), p1...)
	return thrift.ToUpdateDomainResponse(response), thrift.ToError(err)
}
