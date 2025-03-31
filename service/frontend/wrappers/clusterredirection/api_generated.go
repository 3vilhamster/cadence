package clusterredirection

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/clusterredirection.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/config"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/common/resource"
	"github.com/uber/cadence/common/types"
	"github.com/uber/cadence/service/frontend/api"
	frontendcfg "github.com/uber/cadence/service/frontend/config"
)

type (
	// ClusterRedirectionHandlerImpl is simple wrapper over frontend service, doing redirection based on policy for global domains not being active in current cluster
	clusterRedirectionHandler struct {
		resource.Resource

		currentClusterName string
		redirectionPolicy  ClusterRedirectionPolicy
		tokenSerializer    common.TaskTokenSerializer
		frontendHandler    api.Handler
		callOptions        []yarpc.CallOption
	}
)

// NewAPIHandler creates a frontend handler to handle cluster redirection for global domains not being active in current cluster
func NewAPIHandler(
	wfHandler api.Handler,
	resource resource.Resource,
	config *frontendcfg.Config,
	policy config.ClusterRedirectionPolicy,
) api.Handler {
	dcRedirectionPolicy := RedirectionPolicyGenerator(
		resource.GetClusterMetadata(),
		config,
		resource.GetDomainCache(),
		policy,
	)

	return &clusterRedirectionHandler{
		Resource:           resource,
		currentClusterName: resource.GetClusterMetadata().GetCurrentClusterName(),
		redirectionPolicy:  dcRedirectionPolicy,
		tokenSerializer:    common.NewJSONTaskTokenSerializer(),
		frontendHandler:    wfHandler,
		callOptions:        []yarpc.CallOption{yarpc.WithHeader(common.AutoforwardingClusterHeaderName, resource.GetClusterMetadata().GetCurrentClusterName())},
	}
}

func (handler *clusterRedirectionHandler) CountWorkflowExecutions(ctx context.Context, cp1 *types.CountWorkflowExecutionsRequest) (cp2 *types.CountWorkflowExecutionsResponse, err error) {
	var apiName = "CountWorkflowExecutions"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionCountWorkflowExecutionsScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, cp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, cp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			cp2, err = handler.frontendHandler.CountWorkflowExecutions(ctx, cp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			cp2, err = remoteClient.CountWorkflowExecutions(ctx, cp1, handler.callOptions...)
		}
		return err
	})

	return cp2, err
}

func (handler *clusterRedirectionHandler) DeprecateDomain(ctx context.Context, dp1 *types.DeprecateDomainRequest) (err error) {
	return handler.frontendHandler.DeprecateDomain(ctx, dp1)
}

func (handler *clusterRedirectionHandler) DescribeDomain(ctx context.Context, dp1 *types.DescribeDomainRequest) (dp2 *types.DescribeDomainResponse, err error) {
	return handler.frontendHandler.DescribeDomain(ctx, dp1)
}

func (handler *clusterRedirectionHandler) DescribeTaskList(ctx context.Context, dp1 *types.DescribeTaskListRequest) (dp2 *types.DescribeTaskListResponse, err error) {
	var apiName = "DescribeTaskList"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionDescribeTaskListScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, dp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, dp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			dp2, err = handler.frontendHandler.DescribeTaskList(ctx, dp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			dp2, err = remoteClient.DescribeTaskList(ctx, dp1, handler.callOptions...)
		}
		return err
	})

	return dp2, err
}

func (handler *clusterRedirectionHandler) DescribeWorkflowExecution(ctx context.Context, dp1 *types.DescribeWorkflowExecutionRequest) (dp2 *types.DescribeWorkflowExecutionResponse, err error) {
	var apiName = "DescribeWorkflowExecution"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionDescribeWorkflowExecutionScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, dp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, dp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			dp2, err = handler.frontendHandler.DescribeWorkflowExecution(ctx, dp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			dp2, err = remoteClient.DescribeWorkflowExecution(ctx, dp1, handler.callOptions...)
		}
		return err
	})

	return dp2, err
}

func (handler *clusterRedirectionHandler) DiagnoseWorkflowExecution(ctx context.Context, dp1 *types.DiagnoseWorkflowExecutionRequest) (dp2 *types.DiagnoseWorkflowExecutionResponse, err error) {
	return handler.frontendHandler.DiagnoseWorkflowExecution(ctx, dp1)
}

func (handler *clusterRedirectionHandler) GetClusterInfo(ctx context.Context) (cp1 *types.ClusterInfo, err error) {
	return handler.frontendHandler.GetClusterInfo(ctx)
}

func (handler *clusterRedirectionHandler) GetSearchAttributes(ctx context.Context) (gp1 *types.GetSearchAttributesResponse, err error) {
	return handler.frontendHandler.GetSearchAttributes(ctx)
}

func (handler *clusterRedirectionHandler) GetTaskListsByDomain(ctx context.Context, gp1 *types.GetTaskListsByDomainRequest) (gp2 *types.GetTaskListsByDomainResponse, err error) {
	var apiName = "GetTaskListsByDomain"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionGetTaskListsByDomainScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, gp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, gp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			gp2, err = handler.frontendHandler.GetTaskListsByDomain(ctx, gp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			gp2, err = remoteClient.GetTaskListsByDomain(ctx, gp1, handler.callOptions...)
		}
		return err
	})

	return gp2, err
}

func (handler *clusterRedirectionHandler) GetWorkflowExecutionHistory(ctx context.Context, gp1 *types.GetWorkflowExecutionHistoryRequest) (gp2 *types.GetWorkflowExecutionHistoryResponse, err error) {
	var apiName = "GetWorkflowExecutionHistory"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionGetWorkflowExecutionHistoryScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, gp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, gp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			gp2, err = handler.frontendHandler.GetWorkflowExecutionHistory(ctx, gp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			gp2, err = remoteClient.GetWorkflowExecutionHistory(ctx, gp1, handler.callOptions...)
		}
		return err
	})

	return gp2, err
}

func (handler *clusterRedirectionHandler) Health(ctx context.Context) (hp1 *types.HealthStatus, err error) {
	return handler.frontendHandler.Health(ctx)
}

func (handler *clusterRedirectionHandler) ListArchivedWorkflowExecutions(ctx context.Context, lp1 *types.ListArchivedWorkflowExecutionsRequest) (lp2 *types.ListArchivedWorkflowExecutionsResponse, err error) {
	var apiName = "ListArchivedWorkflowExecutions"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionListArchivedWorkflowExecutionsScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, lp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, lp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			lp2, err = handler.frontendHandler.ListArchivedWorkflowExecutions(ctx, lp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			lp2, err = remoteClient.ListArchivedWorkflowExecutions(ctx, lp1, handler.callOptions...)
		}
		return err
	})

	return lp2, err
}

func (handler *clusterRedirectionHandler) ListClosedWorkflowExecutions(ctx context.Context, lp1 *types.ListClosedWorkflowExecutionsRequest) (lp2 *types.ListClosedWorkflowExecutionsResponse, err error) {
	var apiName = "ListClosedWorkflowExecutions"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionListClosedWorkflowExecutionsScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, lp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, lp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			lp2, err = handler.frontendHandler.ListClosedWorkflowExecutions(ctx, lp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			lp2, err = remoteClient.ListClosedWorkflowExecutions(ctx, lp1, handler.callOptions...)
		}
		return err
	})

	return lp2, err
}

func (handler *clusterRedirectionHandler) ListDomains(ctx context.Context, lp1 *types.ListDomainsRequest) (lp2 *types.ListDomainsResponse, err error) {
	return handler.frontendHandler.ListDomains(ctx, lp1)
}

func (handler *clusterRedirectionHandler) ListOpenWorkflowExecutions(ctx context.Context, lp1 *types.ListOpenWorkflowExecutionsRequest) (lp2 *types.ListOpenWorkflowExecutionsResponse, err error) {
	var apiName = "ListOpenWorkflowExecutions"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionListOpenWorkflowExecutionsScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, lp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, lp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			lp2, err = handler.frontendHandler.ListOpenWorkflowExecutions(ctx, lp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			lp2, err = remoteClient.ListOpenWorkflowExecutions(ctx, lp1, handler.callOptions...)
		}
		return err
	})

	return lp2, err
}

func (handler *clusterRedirectionHandler) ListTaskListPartitions(ctx context.Context, lp1 *types.ListTaskListPartitionsRequest) (lp2 *types.ListTaskListPartitionsResponse, err error) {
	var apiName = "ListTaskListPartitions"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionListTaskListPartitionsScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, lp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, lp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			lp2, err = handler.frontendHandler.ListTaskListPartitions(ctx, lp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			lp2, err = remoteClient.ListTaskListPartitions(ctx, lp1, handler.callOptions...)
		}
		return err
	})

	return lp2, err
}

func (handler *clusterRedirectionHandler) ListWorkflowExecutions(ctx context.Context, lp1 *types.ListWorkflowExecutionsRequest) (lp2 *types.ListWorkflowExecutionsResponse, err error) {
	var apiName = "ListWorkflowExecutions"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionListWorkflowExecutionsScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, lp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, lp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			lp2, err = handler.frontendHandler.ListWorkflowExecutions(ctx, lp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			lp2, err = remoteClient.ListWorkflowExecutions(ctx, lp1, handler.callOptions...)
		}
		return err
	})

	return lp2, err
}

func (handler *clusterRedirectionHandler) PollForActivityTask(ctx context.Context, pp1 *types.PollForActivityTaskRequest) (pp2 *types.PollForActivityTaskResponse, err error) {
	var apiName = "PollForActivityTask"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionPollForActivityTaskScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, pp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, pp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			pp2, err = handler.frontendHandler.PollForActivityTask(ctx, pp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			pp2, err = remoteClient.PollForActivityTask(ctx, pp1, handler.callOptions...)
		}
		return err
	})

	return pp2, err
}

func (handler *clusterRedirectionHandler) PollForDecisionTask(ctx context.Context, pp1 *types.PollForDecisionTaskRequest) (pp2 *types.PollForDecisionTaskResponse, err error) {
	var apiName = "PollForDecisionTask"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionPollForDecisionTaskScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, pp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, pp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			pp2, err = handler.frontendHandler.PollForDecisionTask(ctx, pp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			pp2, err = remoteClient.PollForDecisionTask(ctx, pp1, handler.callOptions...)
		}
		return err
	})

	return pp2, err
}

func (handler *clusterRedirectionHandler) RecordActivityTaskHeartbeat(ctx context.Context, rp1 *types.RecordActivityTaskHeartbeatRequest) (rp2 *types.RecordActivityTaskHeartbeatResponse, err error) {
	var apiName = "RecordActivityTaskHeartbeat"
	var cluster string

	token := domainIDGetter(noopdomainIDGetter{})
	scope, startTime := handler.beforeCall(metrics.DCRedirectionRecordActivityTaskHeartbeatScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, "", token.GetDomainID(), cluster, &err)
	}()

	token, err = handler.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return nil, err
	}

	err = handler.redirectionPolicy.WithDomainIDRedirect(ctx, token.GetDomainID(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			rp2, err = handler.frontendHandler.RecordActivityTaskHeartbeat(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			rp2, err = remoteClient.RecordActivityTaskHeartbeat(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return rp2, err
}

func (handler *clusterRedirectionHandler) RecordActivityTaskHeartbeatByID(ctx context.Context, rp1 *types.RecordActivityTaskHeartbeatByIDRequest) (rp2 *types.RecordActivityTaskHeartbeatResponse, err error) {
	var apiName = "RecordActivityTaskHeartbeatByID"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionRecordActivityTaskHeartbeatByIDScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			rp2, err = handler.frontendHandler.RecordActivityTaskHeartbeatByID(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			rp2, err = remoteClient.RecordActivityTaskHeartbeatByID(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return rp2, err
}

func (handler *clusterRedirectionHandler) RefreshWorkflowTasks(ctx context.Context, rp1 *types.RefreshWorkflowTasksRequest) (err error) {
	var apiName = "RefreshWorkflowTasks"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionRefreshWorkflowTasksScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RefreshWorkflowTasks(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RefreshWorkflowTasks(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RegisterDomain(ctx context.Context, rp1 *types.RegisterDomainRequest) (err error) {
	return handler.frontendHandler.RegisterDomain(ctx, rp1)
}

func (handler *clusterRedirectionHandler) RequestCancelWorkflowExecution(ctx context.Context, rp1 *types.RequestCancelWorkflowExecutionRequest) (err error) {
	var apiName = "RequestCancelWorkflowExecution"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionRequestCancelWorkflowExecutionScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RequestCancelWorkflowExecution(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RequestCancelWorkflowExecution(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) ResetStickyTaskList(ctx context.Context, rp1 *types.ResetStickyTaskListRequest) (rp2 *types.ResetStickyTaskListResponse, err error) {
	var apiName = "ResetStickyTaskList"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionResetStickyTaskListScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			rp2, err = handler.frontendHandler.ResetStickyTaskList(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			rp2, err = remoteClient.ResetStickyTaskList(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return rp2, err
}

func (handler *clusterRedirectionHandler) ResetWorkflowExecution(ctx context.Context, rp1 *types.ResetWorkflowExecutionRequest) (rp2 *types.ResetWorkflowExecutionResponse, err error) {
	var apiName = "ResetWorkflowExecution"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionResetWorkflowExecutionScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			rp2, err = handler.frontendHandler.ResetWorkflowExecution(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			rp2, err = remoteClient.ResetWorkflowExecution(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return rp2, err
}

func (handler *clusterRedirectionHandler) RespondActivityTaskCanceled(ctx context.Context, rp1 *types.RespondActivityTaskCanceledRequest) (err error) {
	var apiName = "RespondActivityTaskCanceled"
	var cluster string

	token := domainIDGetter(noopdomainIDGetter{})
	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondActivityTaskCanceledScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, "", token.GetDomainID(), cluster, &err)
	}()

	token, err = handler.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return err
	}

	err = handler.redirectionPolicy.WithDomainIDRedirect(ctx, token.GetDomainID(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RespondActivityTaskCanceled(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RespondActivityTaskCanceled(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RespondActivityTaskCanceledByID(ctx context.Context, rp1 *types.RespondActivityTaskCanceledByIDRequest) (err error) {
	var apiName = "RespondActivityTaskCanceledByID"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondActivityTaskCanceledByIDScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RespondActivityTaskCanceledByID(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RespondActivityTaskCanceledByID(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RespondActivityTaskCompleted(ctx context.Context, rp1 *types.RespondActivityTaskCompletedRequest) (err error) {
	var apiName = "RespondActivityTaskCompleted"
	var cluster string

	token := domainIDGetter(noopdomainIDGetter{})
	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondActivityTaskCompletedScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, "", token.GetDomainID(), cluster, &err)
	}()

	token, err = handler.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return err
	}

	err = handler.redirectionPolicy.WithDomainIDRedirect(ctx, token.GetDomainID(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RespondActivityTaskCompleted(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RespondActivityTaskCompleted(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RespondActivityTaskCompletedByID(ctx context.Context, rp1 *types.RespondActivityTaskCompletedByIDRequest) (err error) {
	var apiName = "RespondActivityTaskCompletedByID"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondActivityTaskCompletedByIDScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RespondActivityTaskCompletedByID(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RespondActivityTaskCompletedByID(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RespondActivityTaskFailed(ctx context.Context, rp1 *types.RespondActivityTaskFailedRequest) (err error) {
	var apiName = "RespondActivityTaskFailed"
	var cluster string

	token := domainIDGetter(noopdomainIDGetter{})
	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondActivityTaskFailedScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, "", token.GetDomainID(), cluster, &err)
	}()

	token, err = handler.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return err
	}

	err = handler.redirectionPolicy.WithDomainIDRedirect(ctx, token.GetDomainID(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RespondActivityTaskFailed(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RespondActivityTaskFailed(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RespondActivityTaskFailedByID(ctx context.Context, rp1 *types.RespondActivityTaskFailedByIDRequest) (err error) {
	var apiName = "RespondActivityTaskFailedByID"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondActivityTaskFailedByIDScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RespondActivityTaskFailedByID(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RespondActivityTaskFailedByID(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RespondDecisionTaskCompleted(ctx context.Context, rp1 *types.RespondDecisionTaskCompletedRequest) (rp2 *types.RespondDecisionTaskCompletedResponse, err error) {
	var apiName = "RespondDecisionTaskCompleted"
	var cluster string

	token := domainIDGetter(noopdomainIDGetter{})
	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondDecisionTaskCompletedScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, "", token.GetDomainID(), cluster, &err)
	}()

	token, err = handler.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return nil, err
	}

	err = handler.redirectionPolicy.WithDomainIDRedirect(ctx, token.GetDomainID(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			rp2, err = handler.frontendHandler.RespondDecisionTaskCompleted(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			rp2, err = remoteClient.RespondDecisionTaskCompleted(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return rp2, err
}

func (handler *clusterRedirectionHandler) RespondDecisionTaskFailed(ctx context.Context, rp1 *types.RespondDecisionTaskFailedRequest) (err error) {
	var apiName = "RespondDecisionTaskFailed"
	var cluster string

	token := domainIDGetter(noopdomainIDGetter{})
	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondDecisionTaskFailedScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, "", token.GetDomainID(), cluster, &err)
	}()

	token, err = handler.tokenSerializer.Deserialize(rp1.TaskToken)
	if err != nil {
		return err
	}

	err = handler.redirectionPolicy.WithDomainIDRedirect(ctx, token.GetDomainID(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RespondDecisionTaskFailed(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RespondDecisionTaskFailed(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RespondQueryTaskCompleted(ctx context.Context, rp1 *types.RespondQueryTaskCompletedRequest) (err error) {
	var apiName = "RespondQueryTaskCompleted"
	var cluster string

	token := domainIDGetter(noopdomainIDGetter{})
	scope, startTime := handler.beforeCall(metrics.DCRedirectionRespondQueryTaskCompletedScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, "", token.GetDomainID(), cluster, &err)
	}()

	token, err = handler.tokenSerializer.DeserializeQueryTaskToken(rp1.TaskToken)
	if err != nil {
		return err
	}

	err = handler.redirectionPolicy.WithDomainIDRedirect(ctx, token.GetDomainID(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.RespondQueryTaskCompleted(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.RespondQueryTaskCompleted(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) RestartWorkflowExecution(ctx context.Context, rp1 *types.RestartWorkflowExecutionRequest) (rp2 *types.RestartWorkflowExecutionResponse, err error) {
	var apiName = "RestartWorkflowExecution"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionRestartWorkflowExecutionScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, rp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, rp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			rp2, err = handler.frontendHandler.RestartWorkflowExecution(ctx, rp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			rp2, err = remoteClient.RestartWorkflowExecution(ctx, rp1, handler.callOptions...)
		}
		return err
	})

	return rp2, err
}

func (handler *clusterRedirectionHandler) ScanWorkflowExecutions(ctx context.Context, lp1 *types.ListWorkflowExecutionsRequest) (lp2 *types.ListWorkflowExecutionsResponse, err error) {
	var apiName = "ScanWorkflowExecutions"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionScanWorkflowExecutionsScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, lp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, lp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			lp2, err = handler.frontendHandler.ScanWorkflowExecutions(ctx, lp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			lp2, err = remoteClient.ScanWorkflowExecutions(ctx, lp1, handler.callOptions...)
		}
		return err
	})

	return lp2, err
}

func (handler *clusterRedirectionHandler) SignalWithStartWorkflowExecution(ctx context.Context, sp1 *types.SignalWithStartWorkflowExecutionRequest) (sp2 *types.StartWorkflowExecutionResponse, err error) {
	var apiName = "SignalWithStartWorkflowExecution"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionSignalWithStartWorkflowExecutionScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, sp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, sp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			sp2, err = handler.frontendHandler.SignalWithStartWorkflowExecution(ctx, sp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			sp2, err = remoteClient.SignalWithStartWorkflowExecution(ctx, sp1, handler.callOptions...)
		}
		return err
	})

	return sp2, err
}

func (handler *clusterRedirectionHandler) SignalWithStartWorkflowExecutionAsync(ctx context.Context, sp1 *types.SignalWithStartWorkflowExecutionAsyncRequest) (sp2 *types.SignalWithStartWorkflowExecutionAsyncResponse, err error) {
	var apiName = "SignalWithStartWorkflowExecutionAsync"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionSignalWithStartWorkflowExecutionAsyncScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, sp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, sp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			sp2, err = handler.frontendHandler.SignalWithStartWorkflowExecutionAsync(ctx, sp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			sp2, err = remoteClient.SignalWithStartWorkflowExecutionAsync(ctx, sp1, handler.callOptions...)
		}
		return err
	})

	return sp2, err
}

func (handler *clusterRedirectionHandler) SignalWorkflowExecution(ctx context.Context, sp1 *types.SignalWorkflowExecutionRequest) (err error) {
	var apiName = "SignalWorkflowExecution"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionSignalWorkflowExecutionScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, sp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, sp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.SignalWorkflowExecution(ctx, sp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.SignalWorkflowExecution(ctx, sp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) StartWorkflowExecution(ctx context.Context, sp1 *types.StartWorkflowExecutionRequest) (sp2 *types.StartWorkflowExecutionResponse, err error) {
	var apiName = "StartWorkflowExecution"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionStartWorkflowExecutionScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, sp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, sp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			sp2, err = handler.frontendHandler.StartWorkflowExecution(ctx, sp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			sp2, err = remoteClient.StartWorkflowExecution(ctx, sp1, handler.callOptions...)
		}
		return err
	})

	return sp2, err
}

func (handler *clusterRedirectionHandler) StartWorkflowExecutionAsync(ctx context.Context, sp1 *types.StartWorkflowExecutionAsyncRequest) (sp2 *types.StartWorkflowExecutionAsyncResponse, err error) {
	var apiName = "StartWorkflowExecutionAsync"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionStartWorkflowExecutionAsyncScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, sp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, sp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			sp2, err = handler.frontendHandler.StartWorkflowExecutionAsync(ctx, sp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			sp2, err = remoteClient.StartWorkflowExecutionAsync(ctx, sp1, handler.callOptions...)
		}
		return err
	})

	return sp2, err
}

func (handler *clusterRedirectionHandler) TerminateWorkflowExecution(ctx context.Context, tp1 *types.TerminateWorkflowExecutionRequest) (err error) {
	var apiName = "TerminateWorkflowExecution"
	var cluster string

	scope, startTime := handler.beforeCall(metrics.DCRedirectionTerminateWorkflowExecutionScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, tp1.GetDomain(), "", cluster, &err)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, tp1.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			err = handler.frontendHandler.TerminateWorkflowExecution(ctx, tp1)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			err = remoteClient.TerminateWorkflowExecution(ctx, tp1, handler.callOptions...)
		}
		return err
	})

	return err
}

func (handler *clusterRedirectionHandler) UpdateDomain(ctx context.Context, up1 *types.UpdateDomainRequest) (up2 *types.UpdateDomainResponse, err error) {
	return handler.frontendHandler.UpdateDomain(ctx, up1)
}

func (handler *clusterRedirectionHandler) QueryWorkflow(
	ctx context.Context,
	request *types.QueryWorkflowRequest,
) (resp *types.QueryWorkflowResponse, retError error) {
	var apiName = "QueryWorkflow"
	var err error
	var cluster string

	// Only autoforward strong consistent queries, this is done for two reasons:
	// 1. Query is meant to be fast, autoforwarding all queries will increase latency.
	// 2. If eventual consistency was requested then the results from running out of local dc will be fine.
	if request.GetQueryConsistencyLevel() == types.QueryConsistencyLevelStrong {
		apiName = "QueryWorkflowStrongConsistency"
	}
	scope, startTime := handler.beforeCall(metrics.DCRedirectionQueryWorkflowScope)
	defer func() {
		handler.afterCall(recover(), scope, startTime, request.GetDomain(), "", cluster, &retError)
	}()

	err = handler.redirectionPolicy.WithDomainNameRedirect(ctx, request.GetDomain(), apiName, func(targetDC string) error {
		cluster = targetDC
		switch {
		case targetDC == handler.currentClusterName:
			resp, err = handler.frontendHandler.QueryWorkflow(ctx, request)
		default:
			remoteClient := handler.GetRemoteFrontendClient(targetDC)
			resp, err = remoteClient.QueryWorkflow(ctx, request, handler.callOptions...)
		}
		return err
	})

	return resp, err
}
