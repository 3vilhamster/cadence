package ratelimited

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/ratelimited.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/cache"
	"github.com/uber/cadence/common/quotas"
	"github.com/uber/cadence/common/types"
	"github.com/uber/cadence/service/frontend/api"
	"github.com/uber/cadence/service/frontend/validate"
)

// apiHandler implements api.Handler interface instrumented with rate limiter.
type apiHandler struct {
	wrapped               api.Handler
	tokenSerializer       common.TaskTokenSerializer
	domainCache           cache.DomainCache
	userRateLimiter       quotas.Policy
	workerRateLimiter     quotas.Policy
	visibilityRateLimiter quotas.Policy
	asyncRateLimiter      quotas.Policy
}

// NewAPIHandler creates a new instance of Handler with ratelimiter.
func NewAPIHandler(
	wrapped api.Handler,
	domainCache cache.DomainCache,
	userRateLimiter quotas.Policy,
	workerRateLimiter quotas.Policy,
	visibilityRateLimiter quotas.Policy,
	asyncRateLimiter quotas.Policy,
) api.Handler {
	return &apiHandler{
		wrapped:               wrapped,
		tokenSerializer:       common.NewJSONTaskTokenSerializer(),
		domainCache:           domainCache,
		userRateLimiter:       userRateLimiter,
		workerRateLimiter:     workerRateLimiter,
		visibilityRateLimiter: visibilityRateLimiter,
		asyncRateLimiter:      asyncRateLimiter,
	}
}

func (h *apiHandler) CountWorkflowExecutions(ctx context.Context, cp1 *types.CountWorkflowExecutionsRequest) (cp2 *types.CountWorkflowExecutionsResponse, err error) {
	if cp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if cp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeVisibility, cp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.CountWorkflowExecutions(ctx, cp1)
}

func (h *apiHandler) DeprecateDomain(ctx context.Context, dp1 *types.DeprecateDomainRequest) (err error) {
	return h.wrapped.DeprecateDomain(ctx, dp1)
}

func (h *apiHandler) DescribeDomain(ctx context.Context, dp1 *types.DescribeDomainRequest) (dp2 *types.DescribeDomainResponse, err error) {
	return h.wrapped.DescribeDomain(ctx, dp1)
}

func (h *apiHandler) DescribeTaskList(ctx context.Context, dp1 *types.DescribeTaskListRequest) (dp2 *types.DescribeTaskListResponse, err error) {
	if dp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if dp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, dp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.DescribeTaskList(ctx, dp1)
}

func (h *apiHandler) DescribeWorkflowExecution(ctx context.Context, dp1 *types.DescribeWorkflowExecutionRequest) (dp2 *types.DescribeWorkflowExecutionResponse, err error) {
	if dp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if dp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, dp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.DescribeWorkflowExecution(ctx, dp1)
}

func (h *apiHandler) DiagnoseWorkflowExecution(ctx context.Context, dp1 *types.DiagnoseWorkflowExecutionRequest) (dp2 *types.DiagnoseWorkflowExecutionResponse, err error) {
	if dp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if dp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, dp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.DiagnoseWorkflowExecution(ctx, dp1)
}

func (h *apiHandler) GetClusterInfo(ctx context.Context) (cp1 *types.ClusterInfo, err error) {
	return h.wrapped.GetClusterInfo(ctx)
}

func (h *apiHandler) GetSearchAttributes(ctx context.Context) (gp1 *types.GetSearchAttributesResponse, err error) {
	return h.wrapped.GetSearchAttributes(ctx)
}

func (h *apiHandler) GetTaskListsByDomain(ctx context.Context, gp1 *types.GetTaskListsByDomainRequest) (gp2 *types.GetTaskListsByDomainResponse, err error) {
	if gp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if gp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, gp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.GetTaskListsByDomain(ctx, gp1)
}

func (h *apiHandler) GetWorkflowExecutionHistory(ctx context.Context, gp1 *types.GetWorkflowExecutionHistoryRequest) (gp2 *types.GetWorkflowExecutionHistoryResponse, err error) {
	if gp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if gp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, gp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.GetWorkflowExecutionHistory(ctx, gp1)
}

func (h *apiHandler) Health(ctx context.Context) (hp1 *types.HealthStatus, err error) {
	return h.wrapped.Health(ctx)
}

func (h *apiHandler) ListArchivedWorkflowExecutions(ctx context.Context, lp1 *types.ListArchivedWorkflowExecutionsRequest) (lp2 *types.ListArchivedWorkflowExecutionsResponse, err error) {
	if lp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if lp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeVisibility, lp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.ListArchivedWorkflowExecutions(ctx, lp1)
}

func (h *apiHandler) ListClosedWorkflowExecutions(ctx context.Context, lp1 *types.ListClosedWorkflowExecutionsRequest) (lp2 *types.ListClosedWorkflowExecutionsResponse, err error) {
	if lp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if lp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeVisibility, lp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.ListClosedWorkflowExecutions(ctx, lp1)
}

func (h *apiHandler) ListDomains(ctx context.Context, lp1 *types.ListDomainsRequest) (lp2 *types.ListDomainsResponse, err error) {
	return h.wrapped.ListDomains(ctx, lp1)
}

func (h *apiHandler) ListOpenWorkflowExecutions(ctx context.Context, lp1 *types.ListOpenWorkflowExecutionsRequest) (lp2 *types.ListOpenWorkflowExecutionsResponse, err error) {
	if lp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if lp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeVisibility, lp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.ListOpenWorkflowExecutions(ctx, lp1)
}

func (h *apiHandler) ListTaskListPartitions(ctx context.Context, lp1 *types.ListTaskListPartitionsRequest) (lp2 *types.ListTaskListPartitionsResponse, err error) {
	if lp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if lp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, lp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.ListTaskListPartitions(ctx, lp1)
}

func (h *apiHandler) ListWorkflowExecutions(ctx context.Context, lp1 *types.ListWorkflowExecutionsRequest) (lp2 *types.ListWorkflowExecutionsResponse, err error) {
	if lp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if lp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeVisibility, lp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.ListWorkflowExecutions(ctx, lp1)
}

func (h *apiHandler) PollForActivityTask(ctx context.Context, pp1 *types.PollForActivityTaskRequest) (pp2 *types.PollForActivityTaskResponse, err error) {
	if pp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if pp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeWorker, pp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.PollForActivityTask(ctx, pp1)
}

func (h *apiHandler) PollForDecisionTask(ctx context.Context, pp1 *types.PollForDecisionTaskRequest) (pp2 *types.PollForDecisionTaskResponse, err error) {
	if pp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if pp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeWorker, pp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.PollForDecisionTask(ctx, pp1)
}

func (h *apiHandler) QueryWorkflow(ctx context.Context, qp1 *types.QueryWorkflowRequest) (qp2 *types.QueryWorkflowResponse, err error) {
	if qp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if qp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, qp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.QueryWorkflow(ctx, qp1)
}

func (h *apiHandler) RecordActivityTaskHeartbeat(ctx context.Context, rp1 *types.RecordActivityTaskHeartbeatRequest) (rp2 *types.RecordActivityTaskHeartbeatResponse, err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.TaskToken == nil {
		err = validate.ErrTaskTokenNotSet
		return
	}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	if token.DomainID == "" {
		err = validate.ErrDomainNotSet
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, domainName)
	return h.wrapped.RecordActivityTaskHeartbeat(ctx, rp1)
}

func (h *apiHandler) RecordActivityTaskHeartbeatByID(ctx context.Context, rp1 *types.RecordActivityTaskHeartbeatByIDRequest) (rp2 *types.RecordActivityTaskHeartbeatResponse, err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, rp1.GetDomain())
	return h.wrapped.RecordActivityTaskHeartbeatByID(ctx, rp1)
}

func (h *apiHandler) RefreshWorkflowTasks(ctx context.Context, rp1 *types.RefreshWorkflowTasksRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, rp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.RefreshWorkflowTasks(ctx, rp1)
}

func (h *apiHandler) RegisterDomain(ctx context.Context, rp1 *types.RegisterDomainRequest) (err error) {
	return h.wrapped.RegisterDomain(ctx, rp1)
}

func (h *apiHandler) RequestCancelWorkflowExecution(ctx context.Context, rp1 *types.RequestCancelWorkflowExecutionRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, rp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.RequestCancelWorkflowExecution(ctx, rp1)
}

func (h *apiHandler) ResetStickyTaskList(ctx context.Context, rp1 *types.ResetStickyTaskListRequest) (rp2 *types.ResetStickyTaskListResponse, err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, rp1.GetDomain())
	return h.wrapped.ResetStickyTaskList(ctx, rp1)
}

func (h *apiHandler) ResetWorkflowExecution(ctx context.Context, rp1 *types.ResetWorkflowExecutionRequest) (rp2 *types.ResetWorkflowExecutionResponse, err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, rp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.ResetWorkflowExecution(ctx, rp1)
}

func (h *apiHandler) RespondActivityTaskCanceled(ctx context.Context, rp1 *types.RespondActivityTaskCanceledRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.TaskToken == nil {
		err = validate.ErrTaskTokenNotSet
		return
	}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	if token.DomainID == "" {
		err = validate.ErrDomainNotSet
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, domainName)
	return h.wrapped.RespondActivityTaskCanceled(ctx, rp1)
}

func (h *apiHandler) RespondActivityTaskCanceledByID(ctx context.Context, rp1 *types.RespondActivityTaskCanceledByIDRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, rp1.GetDomain())
	return h.wrapped.RespondActivityTaskCanceledByID(ctx, rp1)
}

func (h *apiHandler) RespondActivityTaskCompleted(ctx context.Context, rp1 *types.RespondActivityTaskCompletedRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.TaskToken == nil {
		err = validate.ErrTaskTokenNotSet
		return
	}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	if token.DomainID == "" {
		err = validate.ErrDomainNotSet
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, domainName)
	return h.wrapped.RespondActivityTaskCompleted(ctx, rp1)
}

func (h *apiHandler) RespondActivityTaskCompletedByID(ctx context.Context, rp1 *types.RespondActivityTaskCompletedByIDRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, rp1.GetDomain())
	return h.wrapped.RespondActivityTaskCompletedByID(ctx, rp1)
}

func (h *apiHandler) RespondActivityTaskFailed(ctx context.Context, rp1 *types.RespondActivityTaskFailedRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.TaskToken == nil {
		err = validate.ErrTaskTokenNotSet
		return
	}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	if token.DomainID == "" {
		err = validate.ErrDomainNotSet
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, domainName)
	return h.wrapped.RespondActivityTaskFailed(ctx, rp1)
}

func (h *apiHandler) RespondActivityTaskFailedByID(ctx context.Context, rp1 *types.RespondActivityTaskFailedByIDRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, rp1.GetDomain())
	return h.wrapped.RespondActivityTaskFailedByID(ctx, rp1)
}

func (h *apiHandler) RespondDecisionTaskCompleted(ctx context.Context, rp1 *types.RespondDecisionTaskCompletedRequest) (rp2 *types.RespondDecisionTaskCompletedResponse, err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.TaskToken == nil {
		err = validate.ErrTaskTokenNotSet
		return
	}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	if token.DomainID == "" {
		err = validate.ErrDomainNotSet
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, domainName)
	return h.wrapped.RespondDecisionTaskCompleted(ctx, rp1)
}

func (h *apiHandler) RespondDecisionTaskFailed(ctx context.Context, rp1 *types.RespondDecisionTaskFailedRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.TaskToken == nil {
		err = validate.ErrTaskTokenNotSet
		return
	}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	if token.DomainID == "" {
		err = validate.ErrDomainNotSet
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, domainName)
	return h.wrapped.RespondDecisionTaskFailed(ctx, rp1)
}

func (h *apiHandler) RespondQueryTaskCompleted(ctx context.Context, rp1 *types.RespondQueryTaskCompletedRequest) (err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.TaskToken == nil {
		err = validate.ErrTaskTokenNotSet
		return
	}
	token, err := h.tokenSerializer.DeserializeQueryTaskToken(rp1.TaskToken)
	if err != nil {
		return
	}
	if token.DomainID == "" {
		err = validate.ErrDomainNotSet
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	// Count the request in the host RPS,
	// but we still accept it even if RPS is exceeded
	h.allowDomain(ratelimitTypeWorker, domainName)
	return h.wrapped.RespondQueryTaskCompleted(ctx, rp1)
}

func (h *apiHandler) RestartWorkflowExecution(ctx context.Context, rp1 *types.RestartWorkflowExecutionRequest) (rp2 *types.RestartWorkflowExecutionResponse, err error) {
	if rp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if rp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, rp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.RestartWorkflowExecution(ctx, rp1)
}

func (h *apiHandler) ScanWorkflowExecutions(ctx context.Context, lp1 *types.ListWorkflowExecutionsRequest) (lp2 *types.ListWorkflowExecutionsResponse, err error) {
	if lp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if lp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeVisibility, lp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.ScanWorkflowExecutions(ctx, lp1)
}

func (h *apiHandler) SignalWithStartWorkflowExecution(ctx context.Context, sp1 *types.SignalWithStartWorkflowExecutionRequest) (sp2 *types.StartWorkflowExecutionResponse, err error) {
	if sp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if sp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, sp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.SignalWithStartWorkflowExecution(ctx, sp1)
}

func (h *apiHandler) SignalWithStartWorkflowExecutionAsync(ctx context.Context, sp1 *types.SignalWithStartWorkflowExecutionAsyncRequest) (sp2 *types.SignalWithStartWorkflowExecutionAsyncResponse, err error) {
	if sp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if sp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeAsync, sp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.SignalWithStartWorkflowExecutionAsync(ctx, sp1)
}

func (h *apiHandler) SignalWorkflowExecution(ctx context.Context, sp1 *types.SignalWorkflowExecutionRequest) (err error) {
	if sp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if sp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, sp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.SignalWorkflowExecution(ctx, sp1)
}

func (h *apiHandler) StartWorkflowExecution(ctx context.Context, sp1 *types.StartWorkflowExecutionRequest) (sp2 *types.StartWorkflowExecutionResponse, err error) {
	if sp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if sp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, sp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.StartWorkflowExecution(ctx, sp1)
}

func (h *apiHandler) StartWorkflowExecutionAsync(ctx context.Context, sp1 *types.StartWorkflowExecutionAsyncRequest) (sp2 *types.StartWorkflowExecutionAsyncResponse, err error) {
	if sp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if sp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeAsync, sp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.StartWorkflowExecutionAsync(ctx, sp1)
}

func (h *apiHandler) TerminateWorkflowExecution(ctx context.Context, tp1 *types.TerminateWorkflowExecutionRequest) (err error) {
	if tp1 == nil {
		err = validate.ErrRequestNotSet
		return
	}
	if tp1.GetDomain() == "" {
		err = validate.ErrDomainNotSet
		return
	}
	if ok := h.allowDomain(ratelimitTypeUser, tp1.GetDomain()); !ok {
		err = &types.ServiceBusyError{Message: "Too many outstanding requests to the cadence service"}
		return
	}
	return h.wrapped.TerminateWorkflowExecution(ctx, tp1)
}

func (h *apiHandler) UpdateDomain(ctx context.Context, up1 *types.UpdateDomainRequest) (up2 *types.UpdateDomainResponse, err error) {
	return h.wrapped.UpdateDomain(ctx, up1)
}
