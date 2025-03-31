package errorinjectors

// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/errorinjector.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/persistence"
)

// injectorConfigStoreManager implements persistence.ConfigStoreManager interface instrumented with error injection.
type injectorConfigStoreManager struct {
	wrapped   persistence.ConfigStoreManager
	errorRate float64
	logger    log.Logger
}

// NewConfigStoreManager creates a new instance of ConfigStoreManager with error injection.
func NewConfigStoreManager(
	wrapped persistence.ConfigStoreManager,
	errorRate float64,
	logger log.Logger,
) persistence.ConfigStoreManager {
	return &injectorConfigStoreManager{
		wrapped:   wrapped,
		errorRate: errorRate,
		logger:    logger,
	}
}

func (c *injectorConfigStoreManager) Close() {
	c.wrapped.Close()
	return
}

func (c *injectorConfigStoreManager) FetchDynamicConfig(ctx context.Context, cfgType persistence.ConfigType) (fp1 *persistence.FetchDynamicConfigResponse, err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		fp1, err = c.wrapped.FetchDynamicConfig(ctx, cfgType)
	}

	if fakeErr != nil {
		logErr(c.logger, "ConfigStoreManager.FetchDynamicConfig", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}

func (c *injectorConfigStoreManager) UpdateDynamicConfig(ctx context.Context, request *persistence.UpdateDynamicConfigRequest, cfgType persistence.ConfigType) (err error) {
	fakeErr := generateFakeError(c.errorRate)
	var forwardCall bool
	if forwardCall = shouldForwardCallToPersistence(fakeErr); forwardCall {
		err = c.wrapped.UpdateDynamicConfig(ctx, request, cfgType)
	}

	if fakeErr != nil {
		logErr(c.logger, "ConfigStoreManager.UpdateDynamicConfig", fakeErr, forwardCall, err)
		err = fakeErr
		return
	}
	return
}
