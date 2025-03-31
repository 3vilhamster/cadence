package metered

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/metered.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/cache"
	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/log/tag"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/common/types"
	"github.com/uber/cadence/service/frontend/api"
	"github.com/uber/cadence/service/frontend/config"
)

// apiHandler frontend handler wrapper for authentication and authorization
type apiHandler struct {
	handler         api.Handler
	logger          log.Logger
	metricsClient   metrics.Client
	domainCache     cache.DomainCache
	cfg             *config.Config
	tokenSerializer common.TaskTokenSerializer
}

// NewAPIHandler creates frontend handler with metrics and logging
func NewAPIHandler(handler api.Handler, logger log.Logger, metricsClient metrics.Client, domainCache cache.DomainCache, cfg *config.Config) api.Handler {
	return &apiHandler{
		handler:         handler,
		logger:          logger,
		metricsClient:   metricsClient,
		domainCache:     domainCache,
		cfg:             cfg,
		tokenSerializer: common.NewJSONTaskTokenSerializer(),
	}
}

func (h *apiHandler) CountWorkflowExecutions(ctx context.Context, cp1 *types.CountWorkflowExecutionsRequest) (cp2 *types.CountWorkflowExecutionsResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("CountWorkflowExecutions")}
	tags = append(tags, toCountWorkflowExecutionsRequestTags(cp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendCountWorkflowExecutionsScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(cp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	cp2, err = h.handler.CountWorkflowExecutions(ctx, cp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return cp2, err
}
func (h *apiHandler) DeprecateDomain(ctx context.Context, dp1 *types.DeprecateDomainRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("DeprecateDomain")}
	scope := h.metricsClient.Scope(metrics.FrontendDeprecateDomainScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainUnknownTag())...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.DeprecateDomain(ctx, dp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) DescribeDomain(ctx context.Context, dp1 *types.DescribeDomainRequest) (dp2 *types.DescribeDomainResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("DescribeDomain")}
	scope := h.metricsClient.Scope(metrics.FrontendDescribeDomainScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainUnknownTag())...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	dp2, err = h.handler.DescribeDomain(ctx, dp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return dp2, err
}
func (h *apiHandler) DescribeTaskList(ctx context.Context, dp1 *types.DescribeTaskListRequest) (dp2 *types.DescribeTaskListResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("DescribeTaskList")}
	tags = append(tags, toDescribeTaskListRequestTags(dp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendDescribeTaskListScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(dp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	dp2, err = h.handler.DescribeTaskList(ctx, dp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return dp2, err
}
func (h *apiHandler) DescribeWorkflowExecution(ctx context.Context, dp1 *types.DescribeWorkflowExecutionRequest) (dp2 *types.DescribeWorkflowExecutionResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("DescribeWorkflowExecution")}
	tags = append(tags, toDescribeWorkflowExecutionRequestTags(dp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendDescribeWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(dp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	dp2, err = h.handler.DescribeWorkflowExecution(ctx, dp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return dp2, err
}
func (h *apiHandler) DiagnoseWorkflowExecution(ctx context.Context, dp1 *types.DiagnoseWorkflowExecutionRequest) (dp2 *types.DiagnoseWorkflowExecutionResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("DiagnoseWorkflowExecution")}
	tags = append(tags, toDiagnoseWorkflowExecutionRequestTags(dp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendDiagnoseWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(dp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	dp2, err = h.handler.DiagnoseWorkflowExecution(ctx, dp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return dp2, err
}
func (h *apiHandler) GetClusterInfo(ctx context.Context) (cp1 *types.ClusterInfo, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("GetClusterInfo")}
	scope := h.metricsClient.Scope(metrics.FrontendGetClusterInfoScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainUnknownTag())...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	cp1, err = h.handler.GetClusterInfo(ctx)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return cp1, err
}
func (h *apiHandler) GetSearchAttributes(ctx context.Context) (gp1 *types.GetSearchAttributesResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("GetSearchAttributes")}
	scope := h.metricsClient.Scope(metrics.FrontendGetSearchAttributesScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainUnknownTag())...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	gp1, err = h.handler.GetSearchAttributes(ctx)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return gp1, err
}
func (h *apiHandler) GetTaskListsByDomain(ctx context.Context, gp1 *types.GetTaskListsByDomainRequest) (gp2 *types.GetTaskListsByDomainResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("GetTaskListsByDomain")}
	tags = append(tags, toGetTaskListsByDomainRequestTags(gp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendGetTaskListsByDomainScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(gp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	gp2, err = h.handler.GetTaskListsByDomain(ctx, gp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return gp2, err
}
func (h *apiHandler) GetWorkflowExecutionHistory(ctx context.Context, gp1 *types.GetWorkflowExecutionHistoryRequest) (gp2 *types.GetWorkflowExecutionHistoryResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("GetWorkflowExecutionHistory")}
	tags = append(tags, toGetWorkflowExecutionHistoryRequestTags(gp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendGetWorkflowExecutionHistoryScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(gp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	gp2, err = h.handler.GetWorkflowExecutionHistory(ctx, gp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return gp2, err
}
func (h *apiHandler) Health(ctx context.Context) (hp1 *types.HealthStatus, err error) {
	return h.handler.Health(ctx)
}
func (h *apiHandler) ListArchivedWorkflowExecutions(ctx context.Context, lp1 *types.ListArchivedWorkflowExecutionsRequest) (lp2 *types.ListArchivedWorkflowExecutionsResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ListArchivedWorkflowExecutions")}
	tags = append(tags, toListArchivedWorkflowExecutionsRequestTags(lp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendListArchivedWorkflowExecutionsScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(lp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	lp2, err = h.handler.ListArchivedWorkflowExecutions(ctx, lp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return lp2, err
}
func (h *apiHandler) ListClosedWorkflowExecutions(ctx context.Context, lp1 *types.ListClosedWorkflowExecutionsRequest) (lp2 *types.ListClosedWorkflowExecutionsResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ListClosedWorkflowExecutions")}
	tags = append(tags, toListClosedWorkflowExecutionsRequestTags(lp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendListClosedWorkflowExecutionsScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(lp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	lp2, err = h.handler.ListClosedWorkflowExecutions(ctx, lp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return lp2, err
}
func (h *apiHandler) ListDomains(ctx context.Context, lp1 *types.ListDomainsRequest) (lp2 *types.ListDomainsResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ListDomains")}
	scope := h.metricsClient.Scope(metrics.FrontendListDomainsScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainUnknownTag())...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	lp2, err = h.handler.ListDomains(ctx, lp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return lp2, err
}
func (h *apiHandler) ListOpenWorkflowExecutions(ctx context.Context, lp1 *types.ListOpenWorkflowExecutionsRequest) (lp2 *types.ListOpenWorkflowExecutionsResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ListOpenWorkflowExecutions")}
	tags = append(tags, toListOpenWorkflowExecutionsRequestTags(lp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendListOpenWorkflowExecutionsScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(lp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	lp2, err = h.handler.ListOpenWorkflowExecutions(ctx, lp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return lp2, err
}
func (h *apiHandler) ListTaskListPartitions(ctx context.Context, lp1 *types.ListTaskListPartitionsRequest) (lp2 *types.ListTaskListPartitionsResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ListTaskListPartitions")}
	tags = append(tags, toListTaskListPartitionsRequestTags(lp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendListTaskListPartitionsScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(lp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	lp2, err = h.handler.ListTaskListPartitions(ctx, lp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return lp2, err
}
func (h *apiHandler) ListWorkflowExecutions(ctx context.Context, lp1 *types.ListWorkflowExecutionsRequest) (lp2 *types.ListWorkflowExecutionsResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ListWorkflowExecutions")}
	tags = append(tags, toListWorkflowExecutionsRequestTags(lp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendListWorkflowExecutionsScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(lp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	lp2, err = h.handler.ListWorkflowExecutions(ctx, lp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return lp2, err
}
func (h *apiHandler) PollForActivityTask(ctx context.Context, pp1 *types.PollForActivityTaskRequest) (pp2 *types.PollForActivityTaskResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("PollForActivityTask")}
	tags = append(tags, toPollForActivityTaskRequestTags(pp1)...)
	scope := common.NewPerTaskListScope(pp1.Domain, pp1.TaskList.GetName(), pp1.TaskList.GetKind(), h.metricsClient, metrics.FrontendPollForActivityTaskScope).Tagged(metrics.GetContextTags(ctx)...)
	scope.IncCounter(metrics.CadenceRequestsPerTaskList)
	sw := scope.StartTimer(metrics.CadenceLatencyPerTaskList)
	defer sw.Stop()
	swPerDomain := h.metricsClient.Scope(metrics.FrontendPollForActivityTaskScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(pp1.GetDomain()))...).StartTimer(metrics.CadenceLatency)
	defer swPerDomain.Stop()
	logger := h.logger.WithTags(tags...)

	pp2, err = h.handler.PollForActivityTask(ctx, pp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return pp2, err
}
func (h *apiHandler) PollForDecisionTask(ctx context.Context, pp1 *types.PollForDecisionTaskRequest) (pp2 *types.PollForDecisionTaskResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("PollForDecisionTask")}
	tags = append(tags, toPollForDecisionTaskRequestTags(pp1)...)
	scope := common.NewPerTaskListScope(pp1.Domain, pp1.TaskList.GetName(), pp1.TaskList.GetKind(), h.metricsClient, metrics.FrontendPollForDecisionTaskScope).Tagged(metrics.GetContextTags(ctx)...)
	scope.IncCounter(metrics.CadenceRequestsPerTaskList)
	sw := scope.StartTimer(metrics.CadenceLatencyPerTaskList)
	defer sw.Stop()
	swPerDomain := h.metricsClient.Scope(metrics.FrontendPollForDecisionTaskScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(pp1.GetDomain()))...).StartTimer(metrics.CadenceLatency)
	defer swPerDomain.Stop()
	logger := h.logger.WithTags(tags...)

	pp2, err = h.handler.PollForDecisionTask(ctx, pp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return pp2, err
}
func (h *apiHandler) QueryWorkflow(ctx context.Context, qp1 *types.QueryWorkflowRequest) (qp2 *types.QueryWorkflowResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("QueryWorkflow")}
	tags = append(tags, toQueryWorkflowRequestTags(qp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendQueryWorkflowScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(qp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	qp2, err = h.handler.QueryWorkflow(ctx, qp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return qp2, err
}
func (h *apiHandler) RecordActivityTaskHeartbeat(ctx context.Context, rp1 *types.RecordActivityTaskHeartbeatRequest) (rp2 *types.RecordActivityTaskHeartbeatResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RecordActivityTaskHeartbeat")}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	tags = append(tags, tag.WorkflowDomainName(domainName), tag.WorkflowID(token.WorkflowID), tag.WorkflowRunID(token.RunID))
	scope := h.metricsClient.Scope(metrics.FrontendRecordActivityTaskHeartbeatScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(domainName))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	rp2, err = h.handler.RecordActivityTaskHeartbeat(ctx, rp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return rp2, err
}
func (h *apiHandler) RecordActivityTaskHeartbeatByID(ctx context.Context, rp1 *types.RecordActivityTaskHeartbeatByIDRequest) (rp2 *types.RecordActivityTaskHeartbeatResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RecordActivityTaskHeartbeatByID")}
	tags = append(tags, toRecordActivityTaskHeartbeatByIDRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendRecordActivityTaskHeartbeatByIDScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	rp2, err = h.handler.RecordActivityTaskHeartbeatByID(ctx, rp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return rp2, err
}
func (h *apiHandler) RefreshWorkflowTasks(ctx context.Context, rp1 *types.RefreshWorkflowTasksRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RefreshWorkflowTasks")}
	tags = append(tags, toRefreshWorkflowTasksRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendRefreshWorkflowTasksScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RefreshWorkflowTasks(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RegisterDomain(ctx context.Context, rp1 *types.RegisterDomainRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RegisterDomain")}
	scope := h.metricsClient.Scope(metrics.FrontendRegisterDomainScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainUnknownTag())...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RegisterDomain(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RequestCancelWorkflowExecution(ctx context.Context, rp1 *types.RequestCancelWorkflowExecutionRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RequestCancelWorkflowExecution")}
	tags = append(tags, toRequestCancelWorkflowExecutionRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendRequestCancelWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RequestCancelWorkflowExecution(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) ResetStickyTaskList(ctx context.Context, rp1 *types.ResetStickyTaskListRequest) (rp2 *types.ResetStickyTaskListResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ResetStickyTaskList")}
	tags = append(tags, toResetStickyTaskListRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendResetStickyTaskListScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	rp2, err = h.handler.ResetStickyTaskList(ctx, rp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return rp2, err
}
func (h *apiHandler) ResetWorkflowExecution(ctx context.Context, rp1 *types.ResetWorkflowExecutionRequest) (rp2 *types.ResetWorkflowExecutionResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ResetWorkflowExecution")}
	tags = append(tags, toResetWorkflowExecutionRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendResetWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	rp2, err = h.handler.ResetWorkflowExecution(ctx, rp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return rp2, err
}
func (h *apiHandler) RespondActivityTaskCanceled(ctx context.Context, rp1 *types.RespondActivityTaskCanceledRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondActivityTaskCanceled")}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	tags = append(tags, tag.WorkflowDomainName(domainName), tag.WorkflowID(token.WorkflowID), tag.WorkflowRunID(token.RunID))
	scope := h.metricsClient.Scope(metrics.FrontendRespondActivityTaskCanceledScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(domainName))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RespondActivityTaskCanceled(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RespondActivityTaskCanceledByID(ctx context.Context, rp1 *types.RespondActivityTaskCanceledByIDRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondActivityTaskCanceledByID")}
	tags = append(tags, toRespondActivityTaskCanceledByIDRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendRespondActivityTaskCanceledByIDScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RespondActivityTaskCanceledByID(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RespondActivityTaskCompleted(ctx context.Context, rp1 *types.RespondActivityTaskCompletedRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondActivityTaskCompleted")}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	tags = append(tags, tag.WorkflowDomainName(domainName), tag.WorkflowID(token.WorkflowID), tag.WorkflowRunID(token.RunID))
	scope := h.metricsClient.Scope(metrics.FrontendRespondActivityTaskCompletedScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(domainName))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RespondActivityTaskCompleted(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RespondActivityTaskCompletedByID(ctx context.Context, rp1 *types.RespondActivityTaskCompletedByIDRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondActivityTaskCompletedByID")}
	tags = append(tags, toRespondActivityTaskCompletedByIDRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendRespondActivityTaskCompletedByIDScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RespondActivityTaskCompletedByID(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RespondActivityTaskFailed(ctx context.Context, rp1 *types.RespondActivityTaskFailedRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondActivityTaskFailed")}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	tags = append(tags, tag.WorkflowDomainName(domainName), tag.WorkflowID(token.WorkflowID), tag.WorkflowRunID(token.RunID))
	scope := h.metricsClient.Scope(metrics.FrontendRespondActivityTaskFailedScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(domainName))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RespondActivityTaskFailed(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RespondActivityTaskFailedByID(ctx context.Context, rp1 *types.RespondActivityTaskFailedByIDRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondActivityTaskFailedByID")}
	tags = append(tags, toRespondActivityTaskFailedByIDRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendRespondActivityTaskFailedByIDScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RespondActivityTaskFailedByID(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RespondDecisionTaskCompleted(ctx context.Context, rp1 *types.RespondDecisionTaskCompletedRequest) (rp2 *types.RespondDecisionTaskCompletedResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondDecisionTaskCompleted")}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	tags = append(tags, tag.WorkflowDomainName(domainName), tag.WorkflowID(token.WorkflowID), tag.WorkflowRunID(token.RunID))
	scope := h.metricsClient.Scope(metrics.FrontendRespondDecisionTaskCompletedScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(domainName))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	rp2, err = h.handler.RespondDecisionTaskCompleted(ctx, rp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return rp2, err
}
func (h *apiHandler) RespondDecisionTaskFailed(ctx context.Context, rp1 *types.RespondDecisionTaskFailedRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondDecisionTaskFailed")}
	token, err := h.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	tags = append(tags, tag.WorkflowDomainName(domainName), tag.WorkflowID(token.WorkflowID), tag.WorkflowRunID(token.RunID))
	scope := h.metricsClient.Scope(metrics.FrontendRespondDecisionTaskFailedScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(domainName))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RespondDecisionTaskFailed(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RespondQueryTaskCompleted(ctx context.Context, rp1 *types.RespondQueryTaskCompletedRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RespondQueryTaskCompleted")}
	token, err := h.tokenSerializer.DeserializeQueryTaskToken(rp1.TaskToken)
	if err != nil {
		return
	}
	domainName, err := h.domainCache.GetDomainName(token.DomainID)
	if err != nil {
		return
	}
	tags = append(tags, tag.WorkflowDomainName(domainName))
	scope := h.metricsClient.Scope(metrics.FrontendRespondQueryTaskCompletedScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(domainName))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.RespondQueryTaskCompleted(ctx, rp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) RestartWorkflowExecution(ctx context.Context, rp1 *types.RestartWorkflowExecutionRequest) (rp2 *types.RestartWorkflowExecutionResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("RestartWorkflowExecution")}
	tags = append(tags, toRestartWorkflowExecutionRequestTags(rp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendRestartWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(rp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	rp2, err = h.handler.RestartWorkflowExecution(ctx, rp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return rp2, err
}
func (h *apiHandler) ScanWorkflowExecutions(ctx context.Context, lp1 *types.ListWorkflowExecutionsRequest) (lp2 *types.ListWorkflowExecutionsResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("ScanWorkflowExecutions")}
	tags = append(tags, toScanWorkflowExecutionsRequestTags(lp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendScanWorkflowExecutionsScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(lp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	lp2, err = h.handler.ScanWorkflowExecutions(ctx, lp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return lp2, err
}
func (h *apiHandler) SignalWithStartWorkflowExecution(ctx context.Context, sp1 *types.SignalWithStartWorkflowExecutionRequest) (sp2 *types.StartWorkflowExecutionResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("SignalWithStartWorkflowExecution")}
	tags = append(tags, toSignalWithStartWorkflowExecutionRequestTags(sp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendSignalWithStartWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(sp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	sp2, err = h.handler.SignalWithStartWorkflowExecution(ctx, sp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return sp2, err
}
func (h *apiHandler) SignalWithStartWorkflowExecutionAsync(ctx context.Context, sp1 *types.SignalWithStartWorkflowExecutionAsyncRequest) (sp2 *types.SignalWithStartWorkflowExecutionAsyncResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("SignalWithStartWorkflowExecutionAsync")}
	tags = append(tags, toSignalWithStartWorkflowExecutionAsyncRequestTags(sp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendSignalWithStartWorkflowExecutionAsyncScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(sp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	sp2, err = h.handler.SignalWithStartWorkflowExecutionAsync(ctx, sp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return sp2, err
}
func (h *apiHandler) SignalWorkflowExecution(ctx context.Context, sp1 *types.SignalWorkflowExecutionRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	ctx = h.withSignalName(ctx, sp1.GetDomain(), sp1.GetSignalName())
	tags := []tag.Tag{tag.WorkflowHandlerName("SignalWorkflowExecution")}
	tags = append(tags, toSignalWorkflowExecutionRequestTags(sp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendSignalWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(sp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.SignalWorkflowExecution(ctx, sp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) StartWorkflowExecution(ctx context.Context, sp1 *types.StartWorkflowExecutionRequest) (sp2 *types.StartWorkflowExecutionResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("StartWorkflowExecution")}
	tags = append(tags, toStartWorkflowExecutionRequestTags(sp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendStartWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(sp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	sp2, err = h.handler.StartWorkflowExecution(ctx, sp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return sp2, err
}
func (h *apiHandler) StartWorkflowExecutionAsync(ctx context.Context, sp1 *types.StartWorkflowExecutionAsyncRequest) (sp2 *types.StartWorkflowExecutionAsyncResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("StartWorkflowExecutionAsync")}
	tags = append(tags, toStartWorkflowExecutionAsyncRequestTags(sp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendStartWorkflowExecutionAsyncScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(sp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	sp2, err = h.handler.StartWorkflowExecutionAsync(ctx, sp1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return sp2, err
}
func (h *apiHandler) TerminateWorkflowExecution(ctx context.Context, tp1 *types.TerminateWorkflowExecutionRequest) (err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("TerminateWorkflowExecution")}
	tags = append(tags, toTerminateWorkflowExecutionRequestTags(tp1)...)
	scope := h.metricsClient.Scope(metrics.FrontendTerminateWorkflowExecutionScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainTag(tp1.GetDomain()))...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	err = h.handler.TerminateWorkflowExecution(ctx, tp1)
	if err != nil {
		return h.handleErr(err, scope, logger)
	}
	return err
}
func (h *apiHandler) UpdateDomain(ctx context.Context, up1 *types.UpdateDomainRequest) (up2 *types.UpdateDomainResponse, err error) {
	defer func() { log.CapturePanic(recover(), h.logger, &err) }()
	tags := []tag.Tag{tag.WorkflowHandlerName("UpdateDomain")}
	scope := h.metricsClient.Scope(metrics.FrontendUpdateDomainScope).Tagged(append(metrics.GetContextTags(ctx), metrics.DomainUnknownTag())...)
	scope.IncCounter(metrics.CadenceRequests)
	sw := scope.StartTimer(metrics.CadenceLatency)
	defer sw.Stop()
	logger := h.logger.WithTags(tags...)

	up2, err = h.handler.UpdateDomain(ctx, up1)
	if err != nil {
		return nil, h.handleErr(err, scope, logger)
	}
	return up2, err
}
