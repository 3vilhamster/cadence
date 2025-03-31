package errorinjectors

// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/errorinjector.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/persistence"
)

// injectorExecutionManager implements persistence.ExecutionManager interface instrumented with error injection.
type injectorExecutionManager struct {
	wrapped   persistence.ExecutionManager
	errorRate float64
	logger    log.Logger
}

// NewExecutionManager creates a new instance of ExecutionManager with error injection.
func NewExecutionManager(
	wrapped persistence.ExecutionManager,
	errorRate float64,
	logger log.Logger,
) persistence.ExecutionManager {
	return &injectorExecutionManager{
		wrapped:   wrapped,
		errorRate: errorRate,
		logger:    logger,
	}
}

func (c *injectorExecutionManager) Close() {
	c.wrapped.Close()
	return
}

func (c *injectorExecutionManager) CompleteHistoryTask(ctx context.Context, request *persistence.CompleteHistoryTaskRequest) (err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		err = c.wrapped.CompleteHistoryTask(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.CompleteHistoryTask", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) ConflictResolveWorkflowExecution(ctx context.Context, request *persistence.ConflictResolveWorkflowExecutionRequest) (cp1 *persistence.ConflictResolveWorkflowExecutionResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		cp1, err = c.wrapped.ConflictResolveWorkflowExecution(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.ConflictResolveWorkflowExecution", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) CreateFailoverMarkerTasks(ctx context.Context, request *persistence.CreateFailoverMarkersRequest) (err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		err = c.wrapped.CreateFailoverMarkerTasks(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.CreateFailoverMarkerTasks", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) CreateWorkflowExecution(ctx context.Context, request *persistence.CreateWorkflowExecutionRequest) (cp1 *persistence.CreateWorkflowExecutionResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		cp1, err = c.wrapped.CreateWorkflowExecution(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.CreateWorkflowExecution", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) DeleteCurrentWorkflowExecution(ctx context.Context, request *persistence.DeleteCurrentWorkflowExecutionRequest) (err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		err = c.wrapped.DeleteCurrentWorkflowExecution(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.DeleteCurrentWorkflowExecution", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) DeleteReplicationTaskFromDLQ(ctx context.Context, request *persistence.DeleteReplicationTaskFromDLQRequest) (err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		err = c.wrapped.DeleteReplicationTaskFromDLQ(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.DeleteReplicationTaskFromDLQ", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) DeleteWorkflowExecution(ctx context.Context, request *persistence.DeleteWorkflowExecutionRequest) (err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		err = c.wrapped.DeleteWorkflowExecution(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.DeleteWorkflowExecution", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) GetCurrentExecution(ctx context.Context, request *persistence.GetCurrentExecutionRequest) (gp1 *persistence.GetCurrentExecutionResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		gp1, err = c.wrapped.GetCurrentExecution(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.GetCurrentExecution", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) GetHistoryTasks(ctx context.Context, request *persistence.GetHistoryTasksRequest) (gp1 *persistence.GetHistoryTasksResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		gp1, err = c.wrapped.GetHistoryTasks(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.GetHistoryTasks", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) GetName() (s1 string) {
	return c.wrapped.GetName()
}

func (c *injectorExecutionManager) GetReplicationDLQSize(ctx context.Context, request *persistence.GetReplicationDLQSizeRequest) (gp1 *persistence.GetReplicationDLQSizeResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		gp1, err = c.wrapped.GetReplicationDLQSize(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.GetReplicationDLQSize", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) GetReplicationTasksFromDLQ(ctx context.Context, request *persistence.GetReplicationTasksFromDLQRequest) (gp1 *persistence.GetHistoryTasksResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		gp1, err = c.wrapped.GetReplicationTasksFromDLQ(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.GetReplicationTasksFromDLQ", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) GetShardID() (i1 int) {
	return c.wrapped.GetShardID()
}

func (c *injectorExecutionManager) GetWorkflowExecution(ctx context.Context, request *persistence.GetWorkflowExecutionRequest) (gp1 *persistence.GetWorkflowExecutionResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		gp1, err = c.wrapped.GetWorkflowExecution(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.GetWorkflowExecution", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) IsWorkflowExecutionExists(ctx context.Context, request *persistence.IsWorkflowExecutionExistsRequest) (ip1 *persistence.IsWorkflowExecutionExistsResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		ip1, err = c.wrapped.IsWorkflowExecutionExists(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.IsWorkflowExecutionExists", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) ListConcreteExecutions(ctx context.Context, request *persistence.ListConcreteExecutionsRequest) (lp1 *persistence.ListConcreteExecutionsResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		lp1, err = c.wrapped.ListConcreteExecutions(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.ListConcreteExecutions", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) ListCurrentExecutions(ctx context.Context, request *persistence.ListCurrentExecutionsRequest) (lp1 *persistence.ListCurrentExecutionsResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		lp1, err = c.wrapped.ListCurrentExecutions(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.ListCurrentExecutions", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) PutReplicationTaskToDLQ(ctx context.Context, request *persistence.PutReplicationTaskToDLQRequest) (err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		err = c.wrapped.PutReplicationTaskToDLQ(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.PutReplicationTaskToDLQ", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) RangeCompleteHistoryTask(ctx context.Context, request *persistence.RangeCompleteHistoryTaskRequest) (rp1 *persistence.RangeCompleteHistoryTaskResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		rp1, err = c.wrapped.RangeCompleteHistoryTask(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.RangeCompleteHistoryTask", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) RangeDeleteReplicationTaskFromDLQ(ctx context.Context, request *persistence.RangeDeleteReplicationTaskFromDLQRequest) (rp1 *persistence.RangeDeleteReplicationTaskFromDLQResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		rp1, err = c.wrapped.RangeDeleteReplicationTaskFromDLQ(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.RangeDeleteReplicationTaskFromDLQ", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorExecutionManager) UpdateWorkflowExecution(ctx context.Context, request *persistence.UpdateWorkflowExecutionRequest) (up1 *persistence.UpdateWorkflowExecutionResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		up1, err = c.wrapped.UpdateWorkflowExecution(ctx, request)
	}

	if fakeErr != nil {
		logErr(c.logger, "ExecutionManager.UpdateWorkflowExecution", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}
