// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package augmentedairuntimestub

import (
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DeleteHumanLoopFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteHumanLoopFuture) Get(ctx workflow.Context) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
	var output augmentedairuntime.DeleteHumanLoopOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeHumanLoopFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeHumanLoopFuture) Get(ctx workflow.Context) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
	var output augmentedairuntime.DescribeHumanLoopOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListHumanLoopsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListHumanLoopsFuture) Get(ctx workflow.Context) (*augmentedairuntime.ListHumanLoopsOutput, error) {
	var output augmentedairuntime.ListHumanLoopsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartHumanLoopFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartHumanLoopFuture) Get(ctx workflow.Context) (*augmentedairuntime.StartHumanLoopOutput, error) {
	var output augmentedairuntime.StartHumanLoopOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopHumanLoopFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopHumanLoopFuture) Get(ctx workflow.Context) (*augmentedairuntime.StopHumanLoopOutput, error) {
	var output augmentedairuntime.StopHumanLoopOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHumanLoop(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
	var output augmentedairuntime.DeleteHumanLoopOutput
	err := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-DeleteHumanLoop", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) *DeleteHumanLoopFuture {
	future := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-DeleteHumanLoop", input)
	return &DeleteHumanLoopFuture{Future: future}
}

func (a *stub) DescribeHumanLoop(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
	var output augmentedairuntime.DescribeHumanLoopOutput
	err := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-DescribeHumanLoop", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) *DescribeHumanLoopFuture {
	future := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-DescribeHumanLoop", input)
	return &DescribeHumanLoopFuture{Future: future}
}

func (a *stub) ListHumanLoops(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error) {
	var output augmentedairuntime.ListHumanLoopsOutput
	err := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-ListHumanLoops", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHumanLoopsAsync(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) *ListHumanLoopsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-ListHumanLoops", input)
	return &ListHumanLoopsFuture{Future: future}
}

func (a *stub) StartHumanLoop(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error) {
	var output augmentedairuntime.StartHumanLoopOutput
	err := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-StartHumanLoop", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) *StartHumanLoopFuture {
	future := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-StartHumanLoop", input)
	return &StartHumanLoopFuture{Future: future}
}

func (a *stub) StopHumanLoop(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error) {
	var output augmentedairuntime.StopHumanLoopOutput
	err := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-StopHumanLoop", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) *StopHumanLoopFuture {
	future := workflow.ExecuteActivity(ctx, "aws-augmentedairuntime-StopHumanLoop", input)
	return &StopHumanLoopFuture{Future: future}
}
