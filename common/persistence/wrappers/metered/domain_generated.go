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

// meteredDomainManager implements persistence.DomainManager interface instrumented with rate limiter.
type meteredDomainManager struct {
	base
	wrapped persistence.DomainManager
}

// NewDomainManager creates a new instance of DomainManager with ratelimiter.
func NewDomainManager(
	wrapped persistence.DomainManager,
	metricClient metrics.Client,
	logger log.Logger,
	cfg *config.Persistence,
) persistence.DomainManager {
	return &meteredDomainManager{
		wrapped: wrapped,
		base: base{
			metricClient:                  metricClient,
			logger:                        logger,
			enableLatencyHistogramMetrics: cfg.EnablePersistenceLatencyHistogramMetrics,
		},
	}
}

func (c *meteredDomainManager) Close() {
	c.wrapped.Close()
	return
}

func (c *meteredDomainManager) CreateDomain(ctx context.Context, request *persistence.CreateDomainRequest) (cp1 *persistence.CreateDomainResponse, err error) {
	op := func() error {
		cp1, err = c.wrapped.CreateDomain(ctx, request)
		c.emptyMetric("DomainManager.CreateDomain", request, cp1, err)
		return err
	}

	err = c.call(metrics.PersistenceCreateDomainScope, op, getCustomMetricTags(request)...)
	return
}

func (c *meteredDomainManager) DeleteDomain(ctx context.Context, request *persistence.DeleteDomainRequest) (err error) {
	op := func() error {
		err = c.wrapped.DeleteDomain(ctx, request)
		c.emptyMetric("DomainManager.DeleteDomain", request, err, err)
		return err
	}

	err = c.call(metrics.PersistenceDeleteDomainScope, op, getCustomMetricTags(request)...)
	return
}

func (c *meteredDomainManager) DeleteDomainByName(ctx context.Context, request *persistence.DeleteDomainByNameRequest) (err error) {
	op := func() error {
		err = c.wrapped.DeleteDomainByName(ctx, request)
		c.emptyMetric("DomainManager.DeleteDomainByName", request, err, err)
		return err
	}

	err = c.call(metrics.PersistenceDeleteDomainByNameScope, op, getCustomMetricTags(request)...)
	return
}

func (c *meteredDomainManager) GetDomain(ctx context.Context, request *persistence.GetDomainRequest) (gp1 *persistence.GetDomainResponse, err error) {
	op := func() error {
		gp1, err = c.wrapped.GetDomain(ctx, request)
		c.emptyMetric("DomainManager.GetDomain", request, gp1, err)
		return err
	}

	err = c.call(metrics.PersistenceGetDomainScope, op, getCustomMetricTags(request)...)
	return
}

func (c *meteredDomainManager) GetMetadata(ctx context.Context) (gp1 *persistence.GetMetadataResponse, err error) {
	op := func() error {
		gp1, err = c.wrapped.GetMetadata(ctx)
		return err
	}

	err = c.call(metrics.PersistenceGetMetadataScope, op)
	return
}

func (c *meteredDomainManager) GetName() (s1 string) {
	return c.wrapped.GetName()
}

func (c *meteredDomainManager) ListDomains(ctx context.Context, request *persistence.ListDomainsRequest) (lp1 *persistence.ListDomainsResponse, err error) {
	op := func() error {
		lp1, err = c.wrapped.ListDomains(ctx, request)
		c.emptyMetric("DomainManager.ListDomains", request, lp1, err)
		return err
	}

	err = c.call(metrics.PersistenceListDomainsScope, op, getCustomMetricTags(request)...)
	return
}

func (c *meteredDomainManager) UpdateDomain(ctx context.Context, request *persistence.UpdateDomainRequest) (err error) {
	op := func() error {
		err = c.wrapped.UpdateDomain(ctx, request)
		c.emptyMetric("DomainManager.UpdateDomain", request, err, err)
		return err
	}

	err = c.call(metrics.PersistenceUpdateDomainScope, op, getCustomMetricTags(request)...)
	return
}
