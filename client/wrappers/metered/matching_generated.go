// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package metered

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/metered.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"
	"strings"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/client/matching"
	"github.com/uber/cadence/common/constants"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/common/types"
)

// matchingClient implements matching.Client interface instrumented with retries
type matchingClient struct {
	client        matching.Client
	metricsClient metrics.Client
}

// NewMatchingClient creates a new instance of matchingClient with retry policy
func NewMatchingClient(client matching.Client, metricsClient metrics.Client) matching.Client {
	return &matchingClient{
		client:        client,
		metricsClient: metricsClient,
	}
}

func (c *matchingClient) AddActivityTask(ctx context.Context, ap1 *types.AddActivityTaskRequest, p1 ...yarpc.CallOption) (ap2 *types.AddActivityTaskResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientAddActivityTaskScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientAddActivityTaskScope, ap1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientAddActivityTaskScope, metrics.CadenceClientLatency)
	ap2, err = c.client.AddActivityTask(ctx, ap1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientAddActivityTaskScope, metrics.CadenceClientFailures)
	}
	return ap2, err
}

func (c *matchingClient) AddDecisionTask(ctx context.Context, ap1 *types.AddDecisionTaskRequest, p1 ...yarpc.CallOption) (ap2 *types.AddDecisionTaskResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientAddDecisionTaskScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientAddDecisionTaskScope, ap1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientAddDecisionTaskScope, metrics.CadenceClientLatency)
	ap2, err = c.client.AddDecisionTask(ctx, ap1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientAddDecisionTaskScope, metrics.CadenceClientFailures)
	}
	return ap2, err
}

func (c *matchingClient) CancelOutstandingPoll(ctx context.Context, cp1 *types.CancelOutstandingPollRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientCancelOutstandingPollScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientCancelOutstandingPollScope, cp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientCancelOutstandingPollScope, metrics.CadenceClientLatency)
	err = c.client.CancelOutstandingPoll(ctx, cp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientCancelOutstandingPollScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *matchingClient) DescribeTaskList(ctx context.Context, mp1 *types.MatchingDescribeTaskListRequest, p1 ...yarpc.CallOption) (dp1 *types.DescribeTaskListResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientDescribeTaskListScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientDescribeTaskListScope, mp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientDescribeTaskListScope, metrics.CadenceClientLatency)
	dp1, err = c.client.DescribeTaskList(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientDescribeTaskListScope, metrics.CadenceClientFailures)
	}
	return dp1, err
}

func (c *matchingClient) GetTaskListsByDomain(ctx context.Context, gp1 *types.GetTaskListsByDomainRequest, p1 ...yarpc.CallOption) (gp2 *types.GetTaskListsByDomainResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientGetTaskListsByDomainScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientGetTaskListsByDomainScope, gp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientGetTaskListsByDomainScope, metrics.CadenceClientLatency)
	gp2, err = c.client.GetTaskListsByDomain(ctx, gp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientGetTaskListsByDomainScope, metrics.CadenceClientFailures)
	}
	return gp2, err
}

func (c *matchingClient) ListTaskListPartitions(ctx context.Context, mp1 *types.MatchingListTaskListPartitionsRequest, p1 ...yarpc.CallOption) (lp1 *types.ListTaskListPartitionsResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientListTaskListPartitionsScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientListTaskListPartitionsScope, mp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientListTaskListPartitionsScope, metrics.CadenceClientLatency)
	lp1, err = c.client.ListTaskListPartitions(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientListTaskListPartitionsScope, metrics.CadenceClientFailures)
	}
	return lp1, err
}

func (c *matchingClient) PollForActivityTask(ctx context.Context, mp1 *types.MatchingPollForActivityTaskRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingPollForActivityTaskResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientPollForActivityTaskScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientPollForActivityTaskScope, mp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientPollForActivityTaskScope, metrics.CadenceClientLatency)
	mp2, err = c.client.PollForActivityTask(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientPollForActivityTaskScope, metrics.CadenceClientFailures)
	}
	return mp2, err
}

func (c *matchingClient) PollForDecisionTask(ctx context.Context, mp1 *types.MatchingPollForDecisionTaskRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingPollForDecisionTaskResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientPollForDecisionTaskScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientPollForDecisionTaskScope, mp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientPollForDecisionTaskScope, metrics.CadenceClientLatency)
	mp2, err = c.client.PollForDecisionTask(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientPollForDecisionTaskScope, metrics.CadenceClientFailures)
	}
	return mp2, err
}

func (c *matchingClient) QueryWorkflow(ctx context.Context, mp1 *types.MatchingQueryWorkflowRequest, p1 ...yarpc.CallOption) (qp1 *types.QueryWorkflowResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientQueryWorkflowScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientQueryWorkflowScope, mp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientQueryWorkflowScope, metrics.CadenceClientLatency)
	qp1, err = c.client.QueryWorkflow(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientQueryWorkflowScope, metrics.CadenceClientFailures)
	}
	return qp1, err
}

func (c *matchingClient) RefreshTaskListPartitionConfig(ctx context.Context, mp1 *types.MatchingRefreshTaskListPartitionConfigRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingRefreshTaskListPartitionConfigResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientRefreshTaskListPartitionConfigScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientRefreshTaskListPartitionConfigScope, mp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientRefreshTaskListPartitionConfigScope, metrics.CadenceClientLatency)
	mp2, err = c.client.RefreshTaskListPartitionConfig(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientRefreshTaskListPartitionConfigScope, metrics.CadenceClientFailures)
	}
	return mp2, err
}

func (c *matchingClient) RespondQueryTaskCompleted(ctx context.Context, mp1 *types.MatchingRespondQueryTaskCompletedRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientRespondQueryTaskCompletedScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientRespondQueryTaskCompletedScope, mp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientRespondQueryTaskCompletedScope, metrics.CadenceClientLatency)
	err = c.client.RespondQueryTaskCompleted(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientRespondQueryTaskCompletedScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *matchingClient) UpdateTaskListPartitionConfig(ctx context.Context, mp1 *types.MatchingUpdateTaskListPartitionConfigRequest, p1 ...yarpc.CallOption) (mp2 *types.MatchingUpdateTaskListPartitionConfigResponse, err error) {
	c.metricsClient.IncCounter(metrics.MatchingClientUpdateTaskListPartitionConfigScope, metrics.CadenceClientRequests)
	c.emitForwardedFromStats(metrics.MatchingClientUpdateTaskListPartitionConfigScope, mp1)

	sw := c.metricsClient.StartTimer(metrics.MatchingClientUpdateTaskListPartitionConfigScope, metrics.CadenceClientLatency)
	mp2, err = c.client.UpdateTaskListPartitionConfig(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.MatchingClientUpdateTaskListPartitionConfigScope, metrics.CadenceClientFailures)
	}
	return mp2, err
}

type forwardedRequest interface {
	GetForwardedFrom() string
	GetTaskList() *types.TaskList
}

func (c *matchingClient) emitForwardedFromStats(scope int, req any) {
	p, ok := req.(forwardedRequest)
	if !ok || p.GetTaskList() == nil {
		return
	}

	taskList := p.GetTaskList()
	forwardedFrom := p.GetForwardedFrom()

	isChildPartition := strings.HasPrefix(taskList.GetName(), constants.ReservedTaskListPrefix)
	if forwardedFrom != "" {
		c.metricsClient.IncCounter(scope, metrics.MatchingClientForwardedCounter)
		return
	}

	if isChildPartition {
		c.metricsClient.IncCounter(scope, metrics.MatchingClientInvalidTaskListName)
	}
	return
}
