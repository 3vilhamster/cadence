// Code generated by MockGen. DO NOT EDIT.
// Source: partition_config_provider.go
//
// Generated by this command:
//
//	mockgen -package matching -source partition_config_provider.go -destination partition_config_provider_mock.go -package matching github.com/uber/cadence/client/matching PartitionConfigProvider
//

// Package matching is a generated GoMock package.
package matching

import (
	reflect "reflect"

	types "github.com/uber/cadence/common/types"
	gomock "go.uber.org/mock/gomock"
)

// MockPartitionConfigProvider is a mock of PartitionConfigProvider interface.
type MockPartitionConfigProvider struct {
	ctrl     *gomock.Controller
	recorder *MockPartitionConfigProviderMockRecorder
	isgomock struct{}
}

// MockPartitionConfigProviderMockRecorder is the mock recorder for MockPartitionConfigProvider.
type MockPartitionConfigProviderMockRecorder struct {
	mock *MockPartitionConfigProvider
}

// NewMockPartitionConfigProvider creates a new mock instance.
func NewMockPartitionConfigProvider(ctrl *gomock.Controller) *MockPartitionConfigProvider {
	mock := &MockPartitionConfigProvider{ctrl: ctrl}
	mock.recorder = &MockPartitionConfigProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPartitionConfigProvider) EXPECT() *MockPartitionConfigProviderMockRecorder {
	return m.recorder
}

// GetNumberOfReadPartitions mocks base method.
func (m *MockPartitionConfigProvider) GetNumberOfReadPartitions(domainID string, taskList types.TaskList, taskListType int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumberOfReadPartitions", domainID, taskList, taskListType)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNumberOfReadPartitions indicates an expected call of GetNumberOfReadPartitions.
func (mr *MockPartitionConfigProviderMockRecorder) GetNumberOfReadPartitions(domainID, taskList, taskListType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumberOfReadPartitions", reflect.TypeOf((*MockPartitionConfigProvider)(nil).GetNumberOfReadPartitions), domainID, taskList, taskListType)
}

// GetNumberOfWritePartitions mocks base method.
func (m *MockPartitionConfigProvider) GetNumberOfWritePartitions(domainID string, taskList types.TaskList, taskListType int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumberOfWritePartitions", domainID, taskList, taskListType)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNumberOfWritePartitions indicates an expected call of GetNumberOfWritePartitions.
func (mr *MockPartitionConfigProviderMockRecorder) GetNumberOfWritePartitions(domainID, taskList, taskListType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumberOfWritePartitions", reflect.TypeOf((*MockPartitionConfigProvider)(nil).GetNumberOfWritePartitions), domainID, taskList, taskListType)
}

// GetPartitionConfig mocks base method.
func (m *MockPartitionConfigProvider) GetPartitionConfig(domainID string, taskList types.TaskList, taskListType int) *types.TaskListPartitionConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionConfig", domainID, taskList, taskListType)
	ret0, _ := ret[0].(*types.TaskListPartitionConfig)
	return ret0
}

// GetPartitionConfig indicates an expected call of GetPartitionConfig.
func (mr *MockPartitionConfigProviderMockRecorder) GetPartitionConfig(domainID, taskList, taskListType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionConfig", reflect.TypeOf((*MockPartitionConfigProvider)(nil).GetPartitionConfig), domainID, taskList, taskListType)
}

// UpdatePartitionConfig mocks base method.
func (m *MockPartitionConfigProvider) UpdatePartitionConfig(domainID string, taskList types.TaskList, taskListType int, config *types.TaskListPartitionConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePartitionConfig", domainID, taskList, taskListType, config)
}

// UpdatePartitionConfig indicates an expected call of UpdatePartitionConfig.
func (mr *MockPartitionConfigProviderMockRecorder) UpdatePartitionConfig(domainID, taskList, taskListType, config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePartitionConfig", reflect.TypeOf((*MockPartitionConfigProvider)(nil).UpdatePartitionConfig), domainID, taskList, taskListType, config)
}
