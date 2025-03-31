// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -package queue -source interface.go -destination interface_mock.go -self_package github.com/uber/cadence/service/history/queue
//

// Package queue is a generated GoMock package.
package queue

import (
	context "context"
	reflect "reflect"

	common "github.com/uber/cadence/service/history/common"
	task "github.com/uber/cadence/service/history/task"
	gomock "go.uber.org/mock/gomock"
)

// MockProcessingQueueState is a mock of ProcessingQueueState interface.
type MockProcessingQueueState struct {
	ctrl     *gomock.Controller
	recorder *MockProcessingQueueStateMockRecorder
	isgomock struct{}
}

// MockProcessingQueueStateMockRecorder is the mock recorder for MockProcessingQueueState.
type MockProcessingQueueStateMockRecorder struct {
	mock *MockProcessingQueueState
}

// NewMockProcessingQueueState creates a new mock instance.
func NewMockProcessingQueueState(ctrl *gomock.Controller) *MockProcessingQueueState {
	mock := &MockProcessingQueueState{ctrl: ctrl}
	mock.recorder = &MockProcessingQueueStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessingQueueState) EXPECT() *MockProcessingQueueStateMockRecorder {
	return m.recorder
}

// AckLevel mocks base method.
func (m *MockProcessingQueueState) AckLevel() task.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AckLevel")
	ret0, _ := ret[0].(task.Key)
	return ret0
}

// AckLevel indicates an expected call of AckLevel.
func (mr *MockProcessingQueueStateMockRecorder) AckLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AckLevel", reflect.TypeOf((*MockProcessingQueueState)(nil).AckLevel))
}

// DomainFilter mocks base method.
func (m *MockProcessingQueueState) DomainFilter() DomainFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainFilter")
	ret0, _ := ret[0].(DomainFilter)
	return ret0
}

// DomainFilter indicates an expected call of DomainFilter.
func (mr *MockProcessingQueueStateMockRecorder) DomainFilter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainFilter", reflect.TypeOf((*MockProcessingQueueState)(nil).DomainFilter))
}

// Level mocks base method.
func (m *MockProcessingQueueState) Level() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Level")
	ret0, _ := ret[0].(int)
	return ret0
}

// Level indicates an expected call of Level.
func (mr *MockProcessingQueueStateMockRecorder) Level() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Level", reflect.TypeOf((*MockProcessingQueueState)(nil).Level))
}

// MaxLevel mocks base method.
func (m *MockProcessingQueueState) MaxLevel() task.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxLevel")
	ret0, _ := ret[0].(task.Key)
	return ret0
}

// MaxLevel indicates an expected call of MaxLevel.
func (mr *MockProcessingQueueStateMockRecorder) MaxLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxLevel", reflect.TypeOf((*MockProcessingQueueState)(nil).MaxLevel))
}

// ReadLevel mocks base method.
func (m *MockProcessingQueueState) ReadLevel() task.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLevel")
	ret0, _ := ret[0].(task.Key)
	return ret0
}

// ReadLevel indicates an expected call of ReadLevel.
func (mr *MockProcessingQueueStateMockRecorder) ReadLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLevel", reflect.TypeOf((*MockProcessingQueueState)(nil).ReadLevel))
}

// MockProcessingQueue is a mock of ProcessingQueue interface.
type MockProcessingQueue struct {
	ctrl     *gomock.Controller
	recorder *MockProcessingQueueMockRecorder
	isgomock struct{}
}

// MockProcessingQueueMockRecorder is the mock recorder for MockProcessingQueue.
type MockProcessingQueueMockRecorder struct {
	mock *MockProcessingQueue
}

// NewMockProcessingQueue creates a new mock instance.
func NewMockProcessingQueue(ctrl *gomock.Controller) *MockProcessingQueue {
	mock := &MockProcessingQueue{ctrl: ctrl}
	mock.recorder = &MockProcessingQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessingQueue) EXPECT() *MockProcessingQueueMockRecorder {
	return m.recorder
}

// AddTasks mocks base method.
func (m *MockProcessingQueue) AddTasks(arg0 map[task.Key]task.Task, arg1 task.Key) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTasks", arg0, arg1)
}

// AddTasks indicates an expected call of AddTasks.
func (mr *MockProcessingQueueMockRecorder) AddTasks(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTasks", reflect.TypeOf((*MockProcessingQueue)(nil).AddTasks), arg0, arg1)
}

// GetTask mocks base method.
func (m *MockProcessingQueue) GetTask(arg0 task.Key) (task.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].(task.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockProcessingQueueMockRecorder) GetTask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockProcessingQueue)(nil).GetTask), arg0)
}

// GetTasks mocks base method.
func (m *MockProcessingQueue) GetTasks() []task.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks")
	ret0, _ := ret[0].([]task.Task)
	return ret0
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockProcessingQueueMockRecorder) GetTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockProcessingQueue)(nil).GetTasks))
}

// Merge mocks base method.
func (m *MockProcessingQueue) Merge(arg0 ProcessingQueue) []ProcessingQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Merge", arg0)
	ret0, _ := ret[0].([]ProcessingQueue)
	return ret0
}

// Merge indicates an expected call of Merge.
func (mr *MockProcessingQueueMockRecorder) Merge(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockProcessingQueue)(nil).Merge), arg0)
}

// Split mocks base method.
func (m *MockProcessingQueue) Split(arg0 ProcessingQueueSplitPolicy) []ProcessingQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Split", arg0)
	ret0, _ := ret[0].([]ProcessingQueue)
	return ret0
}

// Split indicates an expected call of Split.
func (mr *MockProcessingQueueMockRecorder) Split(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Split", reflect.TypeOf((*MockProcessingQueue)(nil).Split), arg0)
}

// State mocks base method.
func (m *MockProcessingQueue) State() ProcessingQueueState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(ProcessingQueueState)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockProcessingQueueMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockProcessingQueue)(nil).State))
}

// UpdateAckLevel mocks base method.
func (m *MockProcessingQueue) UpdateAckLevel() (task.Key, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAckLevel")
	ret0, _ := ret[0].(task.Key)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// UpdateAckLevel indicates an expected call of UpdateAckLevel.
func (mr *MockProcessingQueueMockRecorder) UpdateAckLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAckLevel", reflect.TypeOf((*MockProcessingQueue)(nil).UpdateAckLevel))
}

// MockProcessingQueueSplitPolicy is a mock of ProcessingQueueSplitPolicy interface.
type MockProcessingQueueSplitPolicy struct {
	ctrl     *gomock.Controller
	recorder *MockProcessingQueueSplitPolicyMockRecorder
	isgomock struct{}
}

// MockProcessingQueueSplitPolicyMockRecorder is the mock recorder for MockProcessingQueueSplitPolicy.
type MockProcessingQueueSplitPolicyMockRecorder struct {
	mock *MockProcessingQueueSplitPolicy
}

// NewMockProcessingQueueSplitPolicy creates a new mock instance.
func NewMockProcessingQueueSplitPolicy(ctrl *gomock.Controller) *MockProcessingQueueSplitPolicy {
	mock := &MockProcessingQueueSplitPolicy{ctrl: ctrl}
	mock.recorder = &MockProcessingQueueSplitPolicyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessingQueueSplitPolicy) EXPECT() *MockProcessingQueueSplitPolicyMockRecorder {
	return m.recorder
}

// Evaluate mocks base method.
func (m *MockProcessingQueueSplitPolicy) Evaluate(arg0 ProcessingQueue) []ProcessingQueueState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Evaluate", arg0)
	ret0, _ := ret[0].([]ProcessingQueueState)
	return ret0
}

// Evaluate indicates an expected call of Evaluate.
func (mr *MockProcessingQueueSplitPolicyMockRecorder) Evaluate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Evaluate", reflect.TypeOf((*MockProcessingQueueSplitPolicy)(nil).Evaluate), arg0)
}

// MockProcessingQueueCollection is a mock of ProcessingQueueCollection interface.
type MockProcessingQueueCollection struct {
	ctrl     *gomock.Controller
	recorder *MockProcessingQueueCollectionMockRecorder
	isgomock struct{}
}

// MockProcessingQueueCollectionMockRecorder is the mock recorder for MockProcessingQueueCollection.
type MockProcessingQueueCollectionMockRecorder struct {
	mock *MockProcessingQueueCollection
}

// NewMockProcessingQueueCollection creates a new mock instance.
func NewMockProcessingQueueCollection(ctrl *gomock.Controller) *MockProcessingQueueCollection {
	mock := &MockProcessingQueueCollection{ctrl: ctrl}
	mock.recorder = &MockProcessingQueueCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessingQueueCollection) EXPECT() *MockProcessingQueueCollectionMockRecorder {
	return m.recorder
}

// ActiveQueue mocks base method.
func (m *MockProcessingQueueCollection) ActiveQueue() ProcessingQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveQueue")
	ret0, _ := ret[0].(ProcessingQueue)
	return ret0
}

// ActiveQueue indicates an expected call of ActiveQueue.
func (mr *MockProcessingQueueCollectionMockRecorder) ActiveQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveQueue", reflect.TypeOf((*MockProcessingQueueCollection)(nil).ActiveQueue))
}

// AddTasks mocks base method.
func (m *MockProcessingQueueCollection) AddTasks(arg0 map[task.Key]task.Task, arg1 task.Key) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTasks", arg0, arg1)
}

// AddTasks indicates an expected call of AddTasks.
func (mr *MockProcessingQueueCollectionMockRecorder) AddTasks(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTasks", reflect.TypeOf((*MockProcessingQueueCollection)(nil).AddTasks), arg0, arg1)
}

// GetTask mocks base method.
func (m *MockProcessingQueueCollection) GetTask(arg0 task.Key) (task.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].(task.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockProcessingQueueCollectionMockRecorder) GetTask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockProcessingQueueCollection)(nil).GetTask), arg0)
}

// GetTasks mocks base method.
func (m *MockProcessingQueueCollection) GetTasks() []task.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks")
	ret0, _ := ret[0].([]task.Task)
	return ret0
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockProcessingQueueCollectionMockRecorder) GetTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockProcessingQueueCollection)(nil).GetTasks))
}

// Level mocks base method.
func (m *MockProcessingQueueCollection) Level() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Level")
	ret0, _ := ret[0].(int)
	return ret0
}

// Level indicates an expected call of Level.
func (mr *MockProcessingQueueCollectionMockRecorder) Level() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Level", reflect.TypeOf((*MockProcessingQueueCollection)(nil).Level))
}

// Merge mocks base method.
func (m *MockProcessingQueueCollection) Merge(arg0 []ProcessingQueue) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Merge", arg0)
}

// Merge indicates an expected call of Merge.
func (mr *MockProcessingQueueCollectionMockRecorder) Merge(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockProcessingQueueCollection)(nil).Merge), arg0)
}

// Queues mocks base method.
func (m *MockProcessingQueueCollection) Queues() []ProcessingQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Queues")
	ret0, _ := ret[0].([]ProcessingQueue)
	return ret0
}

// Queues indicates an expected call of Queues.
func (mr *MockProcessingQueueCollectionMockRecorder) Queues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Queues", reflect.TypeOf((*MockProcessingQueueCollection)(nil).Queues))
}

// Split mocks base method.
func (m *MockProcessingQueueCollection) Split(arg0 ProcessingQueueSplitPolicy) []ProcessingQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Split", arg0)
	ret0, _ := ret[0].([]ProcessingQueue)
	return ret0
}

// Split indicates an expected call of Split.
func (mr *MockProcessingQueueCollectionMockRecorder) Split(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Split", reflect.TypeOf((*MockProcessingQueueCollection)(nil).Split), arg0)
}

// UpdateAckLevels mocks base method.
func (m *MockProcessingQueueCollection) UpdateAckLevels() (task.Key, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAckLevels")
	ret0, _ := ret[0].(task.Key)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// UpdateAckLevels indicates an expected call of UpdateAckLevels.
func (mr *MockProcessingQueueCollectionMockRecorder) UpdateAckLevels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAckLevels", reflect.TypeOf((*MockProcessingQueueCollection)(nil).UpdateAckLevels))
}

// MockProcessor is a mock of Processor interface.
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
	isgomock struct{}
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor.
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance.
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// FailoverDomain mocks base method.
func (m *MockProcessor) FailoverDomain(domainIDs map[string]struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailoverDomain", domainIDs)
}

// FailoverDomain indicates an expected call of FailoverDomain.
func (mr *MockProcessorMockRecorder) FailoverDomain(domainIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailoverDomain", reflect.TypeOf((*MockProcessor)(nil).FailoverDomain), domainIDs)
}

// HandleAction mocks base method.
func (m *MockProcessor) HandleAction(ctx context.Context, clusterName string, action *Action) (*ActionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleAction", ctx, clusterName, action)
	ret0, _ := ret[0].(*ActionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleAction indicates an expected call of HandleAction.
func (mr *MockProcessorMockRecorder) HandleAction(ctx, clusterName, action any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAction", reflect.TypeOf((*MockProcessor)(nil).HandleAction), ctx, clusterName, action)
}

// LockTaskProcessing mocks base method.
func (m *MockProcessor) LockTaskProcessing() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockTaskProcessing")
}

// LockTaskProcessing indicates an expected call of LockTaskProcessing.
func (mr *MockProcessorMockRecorder) LockTaskProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockTaskProcessing", reflect.TypeOf((*MockProcessor)(nil).LockTaskProcessing))
}

// NotifyNewTask mocks base method.
func (m *MockProcessor) NotifyNewTask(clusterName string, info *common.NotifyTaskInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewTask", clusterName, info)
}

// NotifyNewTask indicates an expected call of NotifyNewTask.
func (mr *MockProcessorMockRecorder) NotifyNewTask(clusterName, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewTask", reflect.TypeOf((*MockProcessor)(nil).NotifyNewTask), clusterName, info)
}

// Start mocks base method.
func (m *MockProcessor) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockProcessorMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockProcessor)(nil).Start))
}

// Stop mocks base method.
func (m *MockProcessor) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockProcessorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockProcessor)(nil).Stop))
}

// UnlockTaskProcessing mocks base method.
func (m *MockProcessor) UnlockTaskProcessing() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockTaskProcessing")
}

// UnlockTaskProcessing indicates an expected call of UnlockTaskProcessing.
func (mr *MockProcessorMockRecorder) UnlockTaskProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockTaskProcessing", reflect.TypeOf((*MockProcessor)(nil).UnlockTaskProcessing))
}
