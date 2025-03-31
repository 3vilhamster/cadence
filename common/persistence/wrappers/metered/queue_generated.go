package metered

// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/metered.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"github.com/uber/cadence/common/config"
	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/common/persistence"
)

// meteredQueueManager implements persistence.QueueManager interface instrumented with rate limiter.
type meteredQueueManager struct {
	base
	wrapped persistence.QueueManager
}

// NewQueueManager creates a new instance of QueueManager with ratelimiter.
func NewQueueManager(
	wrapped persistence.QueueManager,
	metricClient metrics.Client,
	logger log.Logger,
	cfg *config.Persistence,
) persistence.QueueManager {
	return &meteredQueueManager{
		wrapped: wrapped,
		base: base{
			metricClient:                  metricClient,
			logger:                        logger,
			enableLatencyHistogramMetrics: cfg.EnablePersistenceLatencyHistogramMetrics,
		},
	}
}

func (c *meteredQueueManager) Close() {
	c.wrapped.Close()
	return
}

func (c *meteredQueueManager) DeleteMessageFromDLQ(ctx context.Context, messageID int64) (err error) {
	op := func() error {
		err = c.wrapped.DeleteMessageFromDLQ(ctx, messageID)
		c.emptyMetric("QueueManager.DeleteMessageFromDLQ", messageID, err, err)
		return err
	}

	err = c.call(metrics.PersistenceDeleteMessageFromDLQScope, op, getCustomMetricTags(messageID)...)
	return
}

func (c *meteredQueueManager) DeleteMessagesBefore(ctx context.Context, messageID int64) (err error) {
	op := func() error {
		err = c.wrapped.DeleteMessagesBefore(ctx, messageID)
		c.emptyMetric("QueueManager.DeleteMessagesBefore", messageID, err, err)
		return err
	}

	err = c.call(metrics.PersistenceDeleteMessagesBeforeScope, op, getCustomMetricTags(messageID)...)
	return
}

func (c *meteredQueueManager) EnqueueMessage(ctx context.Context, messagePayload []byte) (err error) {
	op := func() error {
		err = c.wrapped.EnqueueMessage(ctx, messagePayload)
		c.emptyMetric("QueueManager.EnqueueMessage", messagePayload, err, err)
		return err
	}

	err = c.call(metrics.PersistenceEnqueueMessageScope, op, getCustomMetricTags(messagePayload)...)
	return
}

func (c *meteredQueueManager) EnqueueMessageToDLQ(ctx context.Context, messagePayload []byte) (err error) {
	op := func() error {
		err = c.wrapped.EnqueueMessageToDLQ(ctx, messagePayload)
		c.emptyMetric("QueueManager.EnqueueMessageToDLQ", messagePayload, err, err)
		return err
	}

	err = c.call(metrics.PersistenceEnqueueMessageToDLQScope, op, getCustomMetricTags(messagePayload)...)
	return
}

func (c *meteredQueueManager) GetAckLevels(ctx context.Context) (m1 map[string]int64, err error) {
	op := func() error {
		m1, err = c.wrapped.GetAckLevels(ctx)
		return err
	}

	err = c.call(metrics.PersistenceGetAckLevelsScope, op)
	return
}

func (c *meteredQueueManager) GetDLQAckLevels(ctx context.Context) (m1 map[string]int64, err error) {
	op := func() error {
		m1, err = c.wrapped.GetDLQAckLevels(ctx)
		return err
	}

	err = c.call(metrics.PersistenceGetDLQAckLevelsScope, op)
	return
}

func (c *meteredQueueManager) GetDLQSize(ctx context.Context) (i1 int64, err error) {
	op := func() error {
		i1, err = c.wrapped.GetDLQSize(ctx)
		return err
	}

	err = c.call(metrics.PersistenceGetDLQSizeScope, op)
	return
}

func (c *meteredQueueManager) RangeDeleteMessagesFromDLQ(ctx context.Context, firstMessageID int64, lastMessageID int64) (err error) {
	op := func() error {
		err = c.wrapped.RangeDeleteMessagesFromDLQ(ctx, firstMessageID, lastMessageID)
		c.emptyMetric("QueueManager.RangeDeleteMessagesFromDLQ", firstMessageID, err, err)
		return err
	}

	err = c.call(metrics.PersistenceRangeDeleteMessagesFromDLQScope, op, getCustomMetricTags(firstMessageID)...)
	return
}

func (c *meteredQueueManager) ReadMessages(ctx context.Context, lastMessageID int64, maxCount int) (q1 persistence.QueueMessageList, err error) {
	op := func() error {
		q1, err = c.wrapped.ReadMessages(ctx, lastMessageID, maxCount)
		c.emptyMetric("QueueManager.ReadMessages", lastMessageID, q1, err)
		return err
	}

	err = c.call(metrics.PersistenceReadMessagesScope, op, getCustomMetricTags(lastMessageID)...)
	return
}

func (c *meteredQueueManager) ReadMessagesFromDLQ(ctx context.Context, firstMessageID int64, lastMessageID int64, pageSize int, pageToken []byte) (qpa1 []*persistence.QueueMessage, ba1 []byte, err error) {
	op := func() error {
		qpa1, ba1, err = c.wrapped.ReadMessagesFromDLQ(ctx, firstMessageID, lastMessageID, pageSize, pageToken)
		c.emptyMetric("QueueManager.ReadMessagesFromDLQ", firstMessageID, qpa1, err)
		return err
	}

	err = c.call(metrics.PersistenceReadMessagesFromDLQScope, op, getCustomMetricTags(firstMessageID)...)
	return
}

func (c *meteredQueueManager) UpdateAckLevel(ctx context.Context, messageID int64, clusterName string) (err error) {
	op := func() error {
		err = c.wrapped.UpdateAckLevel(ctx, messageID, clusterName)
		c.emptyMetric("QueueManager.UpdateAckLevel", messageID, err, err)
		return err
	}

	err = c.call(metrics.PersistenceUpdateAckLevelScope, op, getCustomMetricTags(messageID)...)
	return
}

func (c *meteredQueueManager) UpdateDLQAckLevel(ctx context.Context, messageID int64, clusterName string) (err error) {
	op := func() error {
		err = c.wrapped.UpdateDLQAckLevel(ctx, messageID, clusterName)
		c.emptyMetric("QueueManager.UpdateDLQAckLevel", messageID, err, err)
		return err
	}

	err = c.call(metrics.PersistenceUpdateDLQAckLevelScope, op, getCustomMetricTags(messageID)...)
	return
}
