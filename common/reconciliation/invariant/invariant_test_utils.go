// The MIT License (MIT)
//
// Copyright (c) 2017-2020 Uber Technologies Inc.
//
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

package invariant

import (
	"github.com/uber/cadence/common/persistence"
	"github.com/uber/cadence/common/reconciliation/entity"
)

const (
	domainID     = "test-domain-id"
	domainName   = "test-domain-name"
	workflowID   = "test-workflow-id"
	runID        = "test-run-id"
	shardID      = 0
	treeID       = "test-tree-id"
	branchID     = "test-branch-id"
	openState    = persistence.WorkflowStateCreated
	closedState  = persistence.WorkflowStateCompleted
	currentRunID = "test-current-run-id"
)

var (
	branchToken = []byte{1, 2, 3}
)

func getOpenConcreteExecution() *entity.ConcreteExecution {
	return &entity.ConcreteExecution{
		Execution: entity.Execution{
			ShardID:    shardID,
			DomainID:   domainID,
			WorkflowID: workflowID,
			RunID:      runID,
			State:      openState,
		},
		BranchToken: branchToken,
		TreeID:      treeID,
		BranchID:    branchID,
	}
}

func getClosedConcreteExecution() *entity.ConcreteExecution {
	return &entity.ConcreteExecution{
		Execution: entity.Execution{
			ShardID:    shardID,
			DomainID:   domainID,
			WorkflowID: workflowID,
			RunID:      runID,
			State:      closedState,
		},
		BranchToken: branchToken,
		TreeID:      treeID,
		BranchID:    branchID,
	}
}

func getOpenCurrentExecution() *entity.CurrentExecution {
	return &entity.CurrentExecution{
		Execution: entity.Execution{
			ShardID:    shardID,
			DomainID:   domainID,
			WorkflowID: workflowID,
			RunID:      runID,
			State:      openState,
		},
		CurrentRunID: currentRunID,
	}
}

func getClosedCurrentExecution() *entity.CurrentExecution {
	return &entity.CurrentExecution{
		Execution: entity.Execution{
			ShardID:    shardID,
			DomainID:   domainID,
			WorkflowID: workflowID,
			RunID:      runID,
			State:      closedState,
		},
		CurrentRunID: currentRunID,
	}
}
