// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package sfnstub

import (
	"github.com/aws/aws-sdk-go/service/sfn"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateActivityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateActivityFuture) Get(ctx workflow.Context) (*sfn.CreateActivityOutput, error) {
	var output sfn.CreateActivityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateStateMachineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateStateMachineFuture) Get(ctx workflow.Context) (*sfn.CreateStateMachineOutput, error) {
	var output sfn.CreateStateMachineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteActivityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteActivityFuture) Get(ctx workflow.Context) (*sfn.DeleteActivityOutput, error) {
	var output sfn.DeleteActivityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteStateMachineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteStateMachineFuture) Get(ctx workflow.Context) (*sfn.DeleteStateMachineOutput, error) {
	var output sfn.DeleteStateMachineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeActivityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeActivityFuture) Get(ctx workflow.Context) (*sfn.DescribeActivityOutput, error) {
	var output sfn.DescribeActivityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeExecutionFuture) Get(ctx workflow.Context) (*sfn.DescribeExecutionOutput, error) {
	var output sfn.DescribeExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeStateMachineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeStateMachineFuture) Get(ctx workflow.Context) (*sfn.DescribeStateMachineOutput, error) {
	var output sfn.DescribeStateMachineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeStateMachineForExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeStateMachineForExecutionFuture) Get(ctx workflow.Context) (*sfn.DescribeStateMachineForExecutionOutput, error) {
	var output sfn.DescribeStateMachineForExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetActivityTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetActivityTaskFuture) Get(ctx workflow.Context) (*sfn.GetActivityTaskOutput, error) {
	var output sfn.GetActivityTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetExecutionHistoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetExecutionHistoryFuture) Get(ctx workflow.Context) (*sfn.GetExecutionHistoryOutput, error) {
	var output sfn.GetExecutionHistoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListActivitiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListActivitiesFuture) Get(ctx workflow.Context) (*sfn.ListActivitiesOutput, error) {
	var output sfn.ListActivitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListExecutionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListExecutionsFuture) Get(ctx workflow.Context) (*sfn.ListExecutionsOutput, error) {
	var output sfn.ListExecutionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListStateMachinesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListStateMachinesFuture) Get(ctx workflow.Context) (*sfn.ListStateMachinesOutput, error) {
	var output sfn.ListStateMachinesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*sfn.ListTagsForResourceOutput, error) {
	var output sfn.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SendTaskFailureFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SendTaskFailureFuture) Get(ctx workflow.Context) (*sfn.SendTaskFailureOutput, error) {
	var output sfn.SendTaskFailureOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SendTaskHeartbeatFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SendTaskHeartbeatFuture) Get(ctx workflow.Context) (*sfn.SendTaskHeartbeatOutput, error) {
	var output sfn.SendTaskHeartbeatOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SendTaskSuccessFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SendTaskSuccessFuture) Get(ctx workflow.Context) (*sfn.SendTaskSuccessOutput, error) {
	var output sfn.SendTaskSuccessOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartExecutionFuture) Get(ctx workflow.Context) (*sfn.StartExecutionOutput, error) {
	var output sfn.StartExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartSyncExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartSyncExecutionFuture) Get(ctx workflow.Context) (*sfn.StartSyncExecutionOutput, error) {
	var output sfn.StartSyncExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopExecutionFuture) Get(ctx workflow.Context) (*sfn.StopExecutionOutput, error) {
	var output sfn.StopExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*sfn.TagResourceOutput, error) {
	var output sfn.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*sfn.UntagResourceOutput, error) {
	var output sfn.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateStateMachineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateStateMachineFuture) Get(ctx workflow.Context) (*sfn.UpdateStateMachineOutput, error) {
	var output sfn.UpdateStateMachineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateActivity(ctx workflow.Context, input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
	var output sfn.CreateActivityOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-CreateActivity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateActivityAsync(ctx workflow.Context, input *sfn.CreateActivityInput) *CreateActivityFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-CreateActivity", input)
	return &CreateActivityFuture{Future: future}
}

func (a *stub) CreateStateMachine(ctx workflow.Context, input *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error) {
	var output sfn.CreateStateMachineOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-CreateStateMachine", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateStateMachineAsync(ctx workflow.Context, input *sfn.CreateStateMachineInput) *CreateStateMachineFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-CreateStateMachine", input)
	return &CreateStateMachineFuture{Future: future}
}

func (a *stub) DeleteActivity(ctx workflow.Context, input *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error) {
	var output sfn.DeleteActivityOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-DeleteActivity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteActivityAsync(ctx workflow.Context, input *sfn.DeleteActivityInput) *DeleteActivityFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-DeleteActivity", input)
	return &DeleteActivityFuture{Future: future}
}

func (a *stub) DeleteStateMachine(ctx workflow.Context, input *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error) {
	var output sfn.DeleteStateMachineOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-DeleteStateMachine", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteStateMachineAsync(ctx workflow.Context, input *sfn.DeleteStateMachineInput) *DeleteStateMachineFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-DeleteStateMachine", input)
	return &DeleteStateMachineFuture{Future: future}
}

func (a *stub) DescribeActivity(ctx workflow.Context, input *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error) {
	var output sfn.DescribeActivityOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-DescribeActivity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeActivityAsync(ctx workflow.Context, input *sfn.DescribeActivityInput) *DescribeActivityFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-DescribeActivity", input)
	return &DescribeActivityFuture{Future: future}
}

func (a *stub) DescribeExecution(ctx workflow.Context, input *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error) {
	var output sfn.DescribeExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-DescribeExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeExecutionAsync(ctx workflow.Context, input *sfn.DescribeExecutionInput) *DescribeExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-DescribeExecution", input)
	return &DescribeExecutionFuture{Future: future}
}

func (a *stub) DescribeStateMachine(ctx workflow.Context, input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
	var output sfn.DescribeStateMachineOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-DescribeStateMachine", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStateMachineAsync(ctx workflow.Context, input *sfn.DescribeStateMachineInput) *DescribeStateMachineFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-DescribeStateMachine", input)
	return &DescribeStateMachineFuture{Future: future}
}

func (a *stub) DescribeStateMachineForExecution(ctx workflow.Context, input *sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error) {
	var output sfn.DescribeStateMachineForExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-DescribeStateMachineForExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStateMachineForExecutionAsync(ctx workflow.Context, input *sfn.DescribeStateMachineForExecutionInput) *DescribeStateMachineForExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-DescribeStateMachineForExecution", input)
	return &DescribeStateMachineForExecutionFuture{Future: future}
}

func (a *stub) GetActivityTask(ctx workflow.Context, input *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error) {
	var output sfn.GetActivityTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-GetActivityTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetActivityTaskAsync(ctx workflow.Context, input *sfn.GetActivityTaskInput) *GetActivityTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-GetActivityTask", input)
	return &GetActivityTaskFuture{Future: future}
}

func (a *stub) GetExecutionHistory(ctx workflow.Context, input *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error) {
	var output sfn.GetExecutionHistoryOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-GetExecutionHistory", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetExecutionHistoryAsync(ctx workflow.Context, input *sfn.GetExecutionHistoryInput) *GetExecutionHistoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-GetExecutionHistory", input)
	return &GetExecutionHistoryFuture{Future: future}
}

func (a *stub) ListActivities(ctx workflow.Context, input *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error) {
	var output sfn.ListActivitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-ListActivities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListActivitiesAsync(ctx workflow.Context, input *sfn.ListActivitiesInput) *ListActivitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-ListActivities", input)
	return &ListActivitiesFuture{Future: future}
}

func (a *stub) ListExecutions(ctx workflow.Context, input *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error) {
	var output sfn.ListExecutionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-ListExecutions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListExecutionsAsync(ctx workflow.Context, input *sfn.ListExecutionsInput) *ListExecutionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-ListExecutions", input)
	return &ListExecutionsFuture{Future: future}
}

func (a *stub) ListStateMachines(ctx workflow.Context, input *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error) {
	var output sfn.ListStateMachinesOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-ListStateMachines", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStateMachinesAsync(ctx workflow.Context, input *sfn.ListStateMachinesInput) *ListStateMachinesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-ListStateMachines", input)
	return &ListStateMachinesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *sfn.ListTagsForResourceInput) (*sfn.ListTagsForResourceOutput, error) {
	var output sfn.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *sfn.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) SendTaskFailure(ctx workflow.Context, input *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error) {
	var output sfn.SendTaskFailureOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-SendTaskFailure", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendTaskFailureAsync(ctx workflow.Context, input *sfn.SendTaskFailureInput) *SendTaskFailureFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-SendTaskFailure", input)
	return &SendTaskFailureFuture{Future: future}
}

func (a *stub) SendTaskHeartbeat(ctx workflow.Context, input *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error) {
	var output sfn.SendTaskHeartbeatOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-SendTaskHeartbeat", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendTaskHeartbeatAsync(ctx workflow.Context, input *sfn.SendTaskHeartbeatInput) *SendTaskHeartbeatFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-SendTaskHeartbeat", input)
	return &SendTaskHeartbeatFuture{Future: future}
}

func (a *stub) SendTaskSuccess(ctx workflow.Context, input *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error) {
	var output sfn.SendTaskSuccessOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-SendTaskSuccess", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendTaskSuccessAsync(ctx workflow.Context, input *sfn.SendTaskSuccessInput) *SendTaskSuccessFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-SendTaskSuccess", input)
	return &SendTaskSuccessFuture{Future: future}
}

func (a *stub) StartExecution(ctx workflow.Context, input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
	var output sfn.StartExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-StartExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartExecutionAsync(ctx workflow.Context, input *sfn.StartExecutionInput) *StartExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-StartExecution", input)
	return &StartExecutionFuture{Future: future}
}

func (a *stub) StartSyncExecution(ctx workflow.Context, input *sfn.StartSyncExecutionInput) (*sfn.StartSyncExecutionOutput, error) {
	var output sfn.StartSyncExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-StartSyncExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartSyncExecutionAsync(ctx workflow.Context, input *sfn.StartSyncExecutionInput) *StartSyncExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-StartSyncExecution", input)
	return &StartSyncExecutionFuture{Future: future}
}

func (a *stub) StopExecution(ctx workflow.Context, input *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error) {
	var output sfn.StopExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-StopExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopExecutionAsync(ctx workflow.Context, input *sfn.StopExecutionInput) *StopExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-StopExecution", input)
	return &StopExecutionFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *sfn.TagResourceInput) (*sfn.TagResourceOutput, error) {
	var output sfn.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *sfn.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *sfn.UntagResourceInput) (*sfn.UntagResourceOutput, error) {
	var output sfn.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *sfn.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateStateMachine(ctx workflow.Context, input *sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error) {
	var output sfn.UpdateStateMachineOutput
	err := workflow.ExecuteActivity(ctx, "aws-sfn-UpdateStateMachine", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateStateMachineAsync(ctx workflow.Context, input *sfn.UpdateStateMachineInput) *UpdateStateMachineFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sfn-UpdateStateMachine", input)
	return &UpdateStateMachineFuture{Future: future}
}
