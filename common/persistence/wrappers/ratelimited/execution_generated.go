package ratelimited

// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/ratelimited.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"github.com/uber/cadence/common/persistence"
	"github.com/uber/cadence/common/quotas"
)

// ratelimitedExecutionManager implements persistence.ExecutionManager interface instrumented with rate limiter.
type ratelimitedExecutionManager struct {
	wrapped     persistence.ExecutionManager
	rateLimiter quotas.Limiter
}

// NewExecutionManager creates a new instance of ExecutionManager with ratelimiter.
func NewExecutionManager(
	wrapped persistence.ExecutionManager,
	rateLimiter quotas.Limiter,
) persistence.ExecutionManager {
	return &ratelimitedExecutionManager{
		wrapped:     wrapped,
		rateLimiter: rateLimiter,
	}
}

func (c *ratelimitedExecutionManager) Close() {
	c.wrapped.Close()
	return
}

func (c *ratelimitedExecutionManager) CompleteHistoryTask(ctx context.Context, request *persistence.CompleteHistoryTaskRequest) (err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.CompleteHistoryTask(ctx, request)
}

func (c *ratelimitedExecutionManager) ConflictResolveWorkflowExecution(ctx context.Context, request *persistence.ConflictResolveWorkflowExecutionRequest) (cp1 *persistence.ConflictResolveWorkflowExecutionResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.ConflictResolveWorkflowExecution(ctx, request)
}

func (c *ratelimitedExecutionManager) CreateFailoverMarkerTasks(ctx context.Context, request *persistence.CreateFailoverMarkersRequest) (err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.CreateFailoverMarkerTasks(ctx, request)
}

func (c *ratelimitedExecutionManager) CreateWorkflowExecution(ctx context.Context, request *persistence.CreateWorkflowExecutionRequest) (cp1 *persistence.CreateWorkflowExecutionResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.CreateWorkflowExecution(ctx, request)
}

func (c *ratelimitedExecutionManager) DeleteCurrentWorkflowExecution(ctx context.Context, request *persistence.DeleteCurrentWorkflowExecutionRequest) (err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.DeleteCurrentWorkflowExecution(ctx, request)
}

func (c *ratelimitedExecutionManager) DeleteReplicationTaskFromDLQ(ctx context.Context, request *persistence.DeleteReplicationTaskFromDLQRequest) (err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.DeleteReplicationTaskFromDLQ(ctx, request)
}

func (c *ratelimitedExecutionManager) DeleteWorkflowExecution(ctx context.Context, request *persistence.DeleteWorkflowExecutionRequest) (err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.DeleteWorkflowExecution(ctx, request)
}

func (c *ratelimitedExecutionManager) GetCurrentExecution(ctx context.Context, request *persistence.GetCurrentExecutionRequest) (gp1 *persistence.GetCurrentExecutionResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.GetCurrentExecution(ctx, request)
}

func (c *ratelimitedExecutionManager) GetHistoryTasks(ctx context.Context, request *persistence.GetHistoryTasksRequest) (gp1 *persistence.GetHistoryTasksResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.GetHistoryTasks(ctx, request)
}

func (c *ratelimitedExecutionManager) GetName() (s1 string) {
	return c.wrapped.GetName()
}

func (c *ratelimitedExecutionManager) GetReplicationDLQSize(ctx context.Context, request *persistence.GetReplicationDLQSizeRequest) (gp1 *persistence.GetReplicationDLQSizeResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.GetReplicationDLQSize(ctx, request)
}

func (c *ratelimitedExecutionManager) GetReplicationTasksFromDLQ(ctx context.Context, request *persistence.GetReplicationTasksFromDLQRequest) (gp1 *persistence.GetHistoryTasksResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.GetReplicationTasksFromDLQ(ctx, request)
}

func (c *ratelimitedExecutionManager) GetShardID() (i1 int) {
	return c.wrapped.GetShardID()
}

func (c *ratelimitedExecutionManager) GetWorkflowExecution(ctx context.Context, request *persistence.GetWorkflowExecutionRequest) (gp1 *persistence.GetWorkflowExecutionResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.GetWorkflowExecution(ctx, request)
}

func (c *ratelimitedExecutionManager) IsWorkflowExecutionExists(ctx context.Context, request *persistence.IsWorkflowExecutionExistsRequest) (ip1 *persistence.IsWorkflowExecutionExistsResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.IsWorkflowExecutionExists(ctx, request)
}

func (c *ratelimitedExecutionManager) ListConcreteExecutions(ctx context.Context, request *persistence.ListConcreteExecutionsRequest) (lp1 *persistence.ListConcreteExecutionsResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.ListConcreteExecutions(ctx, request)
}

func (c *ratelimitedExecutionManager) ListCurrentExecutions(ctx context.Context, request *persistence.ListCurrentExecutionsRequest) (lp1 *persistence.ListCurrentExecutionsResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.ListCurrentExecutions(ctx, request)
}

func (c *ratelimitedExecutionManager) PutReplicationTaskToDLQ(ctx context.Context, request *persistence.PutReplicationTaskToDLQRequest) (err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.PutReplicationTaskToDLQ(ctx, request)
}

func (c *ratelimitedExecutionManager) RangeCompleteHistoryTask(ctx context.Context, request *persistence.RangeCompleteHistoryTaskRequest) (rp1 *persistence.RangeCompleteHistoryTaskResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.RangeCompleteHistoryTask(ctx, request)
}

func (c *ratelimitedExecutionManager) RangeDeleteReplicationTaskFromDLQ(ctx context.Context, request *persistence.RangeDeleteReplicationTaskFromDLQRequest) (rp1 *persistence.RangeDeleteReplicationTaskFromDLQResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.RangeDeleteReplicationTaskFromDLQ(ctx, request)
}

func (c *ratelimitedExecutionManager) UpdateWorkflowExecution(ctx context.Context, request *persistence.UpdateWorkflowExecutionRequest) (up1 *persistence.UpdateWorkflowExecutionResponse, err error) {
	if ok := c.rateLimiter.Allow(); !ok {
		err = ErrPersistenceLimitExceeded
		return
	}
	return c.wrapped.UpdateWorkflowExecution(ctx, request)
}
