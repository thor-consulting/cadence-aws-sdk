// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package augmentedairuntimestub

import (
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeleteHumanLoop(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error)
	DeleteHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) *AugmentedAIRuntimeDeleteHumanLoopFuture

	DescribeHumanLoop(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error)
	DescribeHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) *AugmentedAIRuntimeDescribeHumanLoopFuture

	ListHumanLoops(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error)
	ListHumanLoopsAsync(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) *AugmentedAIRuntimeListHumanLoopsFuture

	StartHumanLoop(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error)
	StartHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) *AugmentedAIRuntimeStartHumanLoopFuture

	StopHumanLoop(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error)
	StopHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) *AugmentedAIRuntimeStopHumanLoopFuture
}

func NewClient() Client {
	return &stub{}
}