package grpc

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/grpc.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/common/types"
	"github.com/uber/cadence/common/types/mapper/proto"
)

func (g historyClient) CloseShard(ctx context.Context, cp1 *types.CloseShardRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.CloseShard(ctx, proto.FromHistoryCloseShardRequest(cp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) CountDLQMessages(ctx context.Context, cp1 *types.CountDLQMessagesRequest, p1 ...yarpc.CallOption) (hp1 *types.HistoryCountDLQMessagesResponse, err error) {
	response, err := g.c.CountDLQMessages(ctx, proto.FromHistoryCountDLQMessagesRequest(cp1), p1...)
	return proto.ToHistoryCountDLQMessagesResponse(response), proto.ToError(err)
}

func (g historyClient) DescribeHistoryHost(ctx context.Context, dp1 *types.DescribeHistoryHostRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeHistoryHostResponse, err error) {
	response, err := g.c.DescribeHistoryHost(ctx, proto.FromHistoryDescribeHistoryHostRequest(dp1), p1...)
	return proto.ToHistoryDescribeHistoryHostResponse(response), proto.ToError(err)
}

func (g historyClient) DescribeMutableState(ctx context.Context, dp1 *types.DescribeMutableStateRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeMutableStateResponse, err error) {
	response, err := g.c.DescribeMutableState(ctx, proto.FromHistoryDescribeMutableStateRequest(dp1), p1...)
	return proto.ToHistoryDescribeMutableStateResponse(response), proto.ToError(err)
}

func (g historyClient) DescribeQueue(ctx context.Context, dp1 *types.DescribeQueueRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeQueueResponse, err error) {
	response, err := g.c.DescribeQueue(ctx, proto.FromHistoryDescribeQueueRequest(dp1), p1...)
	return proto.ToHistoryDescribeQueueResponse(response), proto.ToError(err)
}

func (g historyClient) DescribeWorkflowExecution(ctx context.Context, hp1 *types.HistoryDescribeWorkflowExecutionRequest, p1 ...yarpc.CallOption) (dp1 *types.DescribeWorkflowExecutionResponse, err error) {
	response, err := g.c.DescribeWorkflowExecution(ctx, proto.FromHistoryDescribeWorkflowExecutionRequest(hp1), p1...)
	return proto.ToHistoryDescribeWorkflowExecutionResponse(response), proto.ToError(err)
}

func (g historyClient) GetCrossClusterTasks(ctx context.Context, gp1 *types.GetCrossClusterTasksRequest, p1 ...yarpc.CallOption) (gp2 *types.GetCrossClusterTasksResponse, err error) {
	response, err := g.c.GetCrossClusterTasks(ctx, proto.FromHistoryGetCrossClusterTasksRequest(gp1), p1...)
	return proto.ToHistoryGetCrossClusterTasksResponse(response), proto.ToError(err)
}

func (g historyClient) GetDLQReplicationMessages(ctx context.Context, gp1 *types.GetDLQReplicationMessagesRequest, p1 ...yarpc.CallOption) (gp2 *types.GetDLQReplicationMessagesResponse, err error) {
	response, err := g.c.GetDLQReplicationMessages(ctx, proto.FromHistoryGetDLQReplicationMessagesRequest(gp1), p1...)
	return proto.ToHistoryGetDLQReplicationMessagesResponse(response), proto.ToError(err)
}

func (g historyClient) GetFailoverInfo(ctx context.Context, gp1 *types.GetFailoverInfoRequest, p1 ...yarpc.CallOption) (gp2 *types.GetFailoverInfoResponse, err error) {
	response, err := g.c.GetFailoverInfo(ctx, proto.FromHistoryGetFailoverInfoRequest(gp1), p1...)
	return proto.ToHistoryGetFailoverInfoResponse(response), proto.ToError(err)
}

func (g historyClient) GetMutableState(ctx context.Context, gp1 *types.GetMutableStateRequest, p1 ...yarpc.CallOption) (gp2 *types.GetMutableStateResponse, err error) {
	response, err := g.c.GetMutableState(ctx, proto.FromHistoryGetMutableStateRequest(gp1), p1...)
	return proto.ToHistoryGetMutableStateResponse(response), proto.ToError(err)
}

func (g historyClient) GetReplicationMessages(ctx context.Context, gp1 *types.GetReplicationMessagesRequest, p1 ...yarpc.CallOption) (gp2 *types.GetReplicationMessagesResponse, err error) {
	response, err := g.c.GetReplicationMessages(ctx, proto.FromHistoryGetReplicationMessagesRequest(gp1), p1...)
	return proto.ToHistoryGetReplicationMessagesResponse(response), proto.ToError(err)
}

func (g historyClient) MergeDLQMessages(ctx context.Context, mp1 *types.MergeDLQMessagesRequest, p1 ...yarpc.CallOption) (mp2 *types.MergeDLQMessagesResponse, err error) {
	response, err := g.c.MergeDLQMessages(ctx, proto.FromHistoryMergeDLQMessagesRequest(mp1), p1...)
	return proto.ToHistoryMergeDLQMessagesResponse(response), proto.ToError(err)
}

func (g historyClient) NotifyFailoverMarkers(ctx context.Context, np1 *types.NotifyFailoverMarkersRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.NotifyFailoverMarkers(ctx, proto.FromHistoryNotifyFailoverMarkersRequest(np1), p1...)
	return proto.ToError(err)
}

func (g historyClient) PollMutableState(ctx context.Context, pp1 *types.PollMutableStateRequest, p1 ...yarpc.CallOption) (pp2 *types.PollMutableStateResponse, err error) {
	response, err := g.c.PollMutableState(ctx, proto.FromHistoryPollMutableStateRequest(pp1), p1...)
	return proto.ToHistoryPollMutableStateResponse(response), proto.ToError(err)
}

func (g historyClient) PurgeDLQMessages(ctx context.Context, pp1 *types.PurgeDLQMessagesRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.PurgeDLQMessages(ctx, proto.FromHistoryPurgeDLQMessagesRequest(pp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) QueryWorkflow(ctx context.Context, hp1 *types.HistoryQueryWorkflowRequest, p1 ...yarpc.CallOption) (hp2 *types.HistoryQueryWorkflowResponse, err error) {
	response, err := g.c.QueryWorkflow(ctx, proto.FromHistoryQueryWorkflowRequest(hp1), p1...)
	return proto.ToHistoryQueryWorkflowResponse(response), proto.ToError(err)
}

func (g historyClient) RatelimitUpdate(ctx context.Context, request *types.RatelimitUpdateRequest, opts ...yarpc.CallOption) (rp1 *types.RatelimitUpdateResponse, err error) {
	response, err := g.c.RatelimitUpdate(ctx, proto.FromHistoryRatelimitUpdateRequest(request), opts...)
	return proto.ToHistoryRatelimitUpdateResponse(response), proto.ToError(err)
}

func (g historyClient) ReadDLQMessages(ctx context.Context, rp1 *types.ReadDLQMessagesRequest, p1 ...yarpc.CallOption) (rp2 *types.ReadDLQMessagesResponse, err error) {
	response, err := g.c.ReadDLQMessages(ctx, proto.FromHistoryReadDLQMessagesRequest(rp1), p1...)
	return proto.ToHistoryReadDLQMessagesResponse(response), proto.ToError(err)
}

func (g historyClient) ReapplyEvents(ctx context.Context, hp1 *types.HistoryReapplyEventsRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.ReapplyEvents(ctx, proto.FromHistoryReapplyEventsRequest(hp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) RecordActivityTaskHeartbeat(ctx context.Context, hp1 *types.HistoryRecordActivityTaskHeartbeatRequest, p1 ...yarpc.CallOption) (rp1 *types.RecordActivityTaskHeartbeatResponse, err error) {
	response, err := g.c.RecordActivityTaskHeartbeat(ctx, proto.FromHistoryRecordActivityTaskHeartbeatRequest(hp1), p1...)
	return proto.ToHistoryRecordActivityTaskHeartbeatResponse(response), proto.ToError(err)
}

func (g historyClient) RecordActivityTaskStarted(ctx context.Context, rp1 *types.RecordActivityTaskStartedRequest, p1 ...yarpc.CallOption) (rp2 *types.RecordActivityTaskStartedResponse, err error) {
	response, err := g.c.RecordActivityTaskStarted(ctx, proto.FromHistoryRecordActivityTaskStartedRequest(rp1), p1...)
	return proto.ToHistoryRecordActivityTaskStartedResponse(response), proto.ToError(err)
}

func (g historyClient) RecordChildExecutionCompleted(ctx context.Context, rp1 *types.RecordChildExecutionCompletedRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RecordChildExecutionCompleted(ctx, proto.FromHistoryRecordChildExecutionCompletedRequest(rp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) RecordDecisionTaskStarted(ctx context.Context, rp1 *types.RecordDecisionTaskStartedRequest, p1 ...yarpc.CallOption) (rp2 *types.RecordDecisionTaskStartedResponse, err error) {
	response, err := g.c.RecordDecisionTaskStarted(ctx, proto.FromHistoryRecordDecisionTaskStartedRequest(rp1), p1...)
	return proto.ToHistoryRecordDecisionTaskStartedResponse(response), proto.ToError(err)
}

func (g historyClient) RefreshWorkflowTasks(ctx context.Context, hp1 *types.HistoryRefreshWorkflowTasksRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RefreshWorkflowTasks(ctx, proto.FromHistoryRefreshWorkflowTasksRequest(hp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) RemoveSignalMutableState(ctx context.Context, rp1 *types.RemoveSignalMutableStateRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RemoveSignalMutableState(ctx, proto.FromHistoryRemoveSignalMutableStateRequest(rp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) RemoveTask(ctx context.Context, rp1 *types.RemoveTaskRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RemoveTask(ctx, proto.FromHistoryRemoveTaskRequest(rp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) ReplicateEventsV2(ctx context.Context, rp1 *types.ReplicateEventsV2Request, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.ReplicateEventsV2(ctx, proto.FromHistoryReplicateEventsV2Request(rp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) RequestCancelWorkflowExecution(ctx context.Context, hp1 *types.HistoryRequestCancelWorkflowExecutionRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RequestCancelWorkflowExecution(ctx, proto.FromHistoryRequestCancelWorkflowExecutionRequest(hp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) ResetQueue(ctx context.Context, rp1 *types.ResetQueueRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.ResetQueue(ctx, proto.FromHistoryResetQueueRequest(rp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) ResetStickyTaskList(ctx context.Context, hp1 *types.HistoryResetStickyTaskListRequest, p1 ...yarpc.CallOption) (hp2 *types.HistoryResetStickyTaskListResponse, err error) {
	response, err := g.c.ResetStickyTaskList(ctx, proto.FromHistoryResetStickyTaskListRequest(hp1), p1...)
	return proto.ToHistoryResetStickyTaskListResponse(response), proto.ToError(err)
}

func (g historyClient) ResetWorkflowExecution(ctx context.Context, hp1 *types.HistoryResetWorkflowExecutionRequest, p1 ...yarpc.CallOption) (rp1 *types.ResetWorkflowExecutionResponse, err error) {
	response, err := g.c.ResetWorkflowExecution(ctx, proto.FromHistoryResetWorkflowExecutionRequest(hp1), p1...)
	return proto.ToHistoryResetWorkflowExecutionResponse(response), proto.ToError(err)
}

func (g historyClient) RespondActivityTaskCanceled(ctx context.Context, hp1 *types.HistoryRespondActivityTaskCanceledRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RespondActivityTaskCanceled(ctx, proto.FromHistoryRespondActivityTaskCanceledRequest(hp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) RespondActivityTaskCompleted(ctx context.Context, hp1 *types.HistoryRespondActivityTaskCompletedRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RespondActivityTaskCompleted(ctx, proto.FromHistoryRespondActivityTaskCompletedRequest(hp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) RespondActivityTaskFailed(ctx context.Context, hp1 *types.HistoryRespondActivityTaskFailedRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RespondActivityTaskFailed(ctx, proto.FromHistoryRespondActivityTaskFailedRequest(hp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) RespondCrossClusterTasksCompleted(ctx context.Context, rp1 *types.RespondCrossClusterTasksCompletedRequest, p1 ...yarpc.CallOption) (rp2 *types.RespondCrossClusterTasksCompletedResponse, err error) {
	response, err := g.c.RespondCrossClusterTasksCompleted(ctx, proto.FromHistoryRespondCrossClusterTasksCompletedRequest(rp1), p1...)
	return proto.ToHistoryRespondCrossClusterTasksCompletedResponse(response), proto.ToError(err)
}

func (g historyClient) RespondDecisionTaskCompleted(ctx context.Context, hp1 *types.HistoryRespondDecisionTaskCompletedRequest, p1 ...yarpc.CallOption) (hp2 *types.HistoryRespondDecisionTaskCompletedResponse, err error) {
	response, err := g.c.RespondDecisionTaskCompleted(ctx, proto.FromHistoryRespondDecisionTaskCompletedRequest(hp1), p1...)
	return proto.ToHistoryRespondDecisionTaskCompletedResponse(response), proto.ToError(err)
}

func (g historyClient) RespondDecisionTaskFailed(ctx context.Context, hp1 *types.HistoryRespondDecisionTaskFailedRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.RespondDecisionTaskFailed(ctx, proto.FromHistoryRespondDecisionTaskFailedRequest(hp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) ScheduleDecisionTask(ctx context.Context, sp1 *types.ScheduleDecisionTaskRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.ScheduleDecisionTask(ctx, proto.FromHistoryScheduleDecisionTaskRequest(sp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) SignalWithStartWorkflowExecution(ctx context.Context, hp1 *types.HistorySignalWithStartWorkflowExecutionRequest, p1 ...yarpc.CallOption) (sp1 *types.StartWorkflowExecutionResponse, err error) {
	response, err := g.c.SignalWithStartWorkflowExecution(ctx, proto.FromHistorySignalWithStartWorkflowExecutionRequest(hp1), p1...)
	return proto.ToHistorySignalWithStartWorkflowExecutionResponse(response), proto.ToError(err)
}

func (g historyClient) SignalWorkflowExecution(ctx context.Context, hp1 *types.HistorySignalWorkflowExecutionRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.SignalWorkflowExecution(ctx, proto.FromHistorySignalWorkflowExecutionRequest(hp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) StartWorkflowExecution(ctx context.Context, hp1 *types.HistoryStartWorkflowExecutionRequest, p1 ...yarpc.CallOption) (sp1 *types.StartWorkflowExecutionResponse, err error) {
	response, err := g.c.StartWorkflowExecution(ctx, proto.FromHistoryStartWorkflowExecutionRequest(hp1), p1...)
	return proto.ToHistoryStartWorkflowExecutionResponse(response), proto.ToError(err)
}

func (g historyClient) SyncActivity(ctx context.Context, sp1 *types.SyncActivityRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.SyncActivity(ctx, proto.FromHistorySyncActivityRequest(sp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) SyncShardStatus(ctx context.Context, sp1 *types.SyncShardStatusRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.SyncShardStatus(ctx, proto.FromHistorySyncShardStatusRequest(sp1), p1...)
	return proto.ToError(err)
}

func (g historyClient) TerminateWorkflowExecution(ctx context.Context, hp1 *types.HistoryTerminateWorkflowExecutionRequest, p1 ...yarpc.CallOption) (err error) {
	_, err = g.c.TerminateWorkflowExecution(ctx, proto.FromHistoryTerminateWorkflowExecutionRequest(hp1), p1...)
	return proto.ToError(err)
}
