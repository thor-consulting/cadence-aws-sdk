// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package batchstub

import (
	"github.com/aws/aws-sdk-go/service/batch"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CancelJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CancelJobFuture) Get(ctx workflow.Context) (*batch.CancelJobOutput, error) {
	var output batch.CancelJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateComputeEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateComputeEnvironmentFuture) Get(ctx workflow.Context) (*batch.CreateComputeEnvironmentOutput, error) {
	var output batch.CreateComputeEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateJobQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateJobQueueFuture) Get(ctx workflow.Context) (*batch.CreateJobQueueOutput, error) {
	var output batch.CreateJobQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteComputeEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteComputeEnvironmentFuture) Get(ctx workflow.Context) (*batch.DeleteComputeEnvironmentOutput, error) {
	var output batch.DeleteComputeEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteJobQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteJobQueueFuture) Get(ctx workflow.Context) (*batch.DeleteJobQueueOutput, error) {
	var output batch.DeleteJobQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeregisterJobDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeregisterJobDefinitionFuture) Get(ctx workflow.Context) (*batch.DeregisterJobDefinitionOutput, error) {
	var output batch.DeregisterJobDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeComputeEnvironmentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeComputeEnvironmentsFuture) Get(ctx workflow.Context) (*batch.DescribeComputeEnvironmentsOutput, error) {
	var output batch.DescribeComputeEnvironmentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeJobDefinitionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeJobDefinitionsFuture) Get(ctx workflow.Context) (*batch.DescribeJobDefinitionsOutput, error) {
	var output batch.DescribeJobDefinitionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeJobQueuesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeJobQueuesFuture) Get(ctx workflow.Context) (*batch.DescribeJobQueuesOutput, error) {
	var output batch.DescribeJobQueuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeJobsFuture) Get(ctx workflow.Context) (*batch.DescribeJobsOutput, error) {
	var output batch.DescribeJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListJobsFuture) Get(ctx workflow.Context) (*batch.ListJobsOutput, error) {
	var output batch.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*batch.ListTagsForResourceOutput, error) {
	var output batch.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RegisterJobDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RegisterJobDefinitionFuture) Get(ctx workflow.Context) (*batch.RegisterJobDefinitionOutput, error) {
	var output batch.RegisterJobDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SubmitJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SubmitJobFuture) Get(ctx workflow.Context) (*batch.SubmitJobOutput, error) {
	var output batch.SubmitJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*batch.TagResourceOutput, error) {
	var output batch.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TerminateJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TerminateJobFuture) Get(ctx workflow.Context) (*batch.TerminateJobOutput, error) {
	var output batch.TerminateJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*batch.UntagResourceOutput, error) {
	var output batch.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateComputeEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateComputeEnvironmentFuture) Get(ctx workflow.Context) (*batch.UpdateComputeEnvironmentOutput, error) {
	var output batch.UpdateComputeEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateJobQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateJobQueueFuture) Get(ctx workflow.Context) (*batch.UpdateJobQueueOutput, error) {
	var output batch.UpdateJobQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJob(ctx workflow.Context, input *batch.CancelJobInput) (*batch.CancelJobOutput, error) {
	var output batch.CancelJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-CancelJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJobAsync(ctx workflow.Context, input *batch.CancelJobInput) *CancelJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-CancelJob", input)
	return &CancelJobFuture{Future: future}
}

func (a *stub) CreateComputeEnvironment(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error) {
	var output batch.CreateComputeEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-CreateComputeEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateComputeEnvironmentAsync(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) *CreateComputeEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-CreateComputeEnvironment", input)
	return &CreateComputeEnvironmentFuture{Future: future}
}

func (a *stub) CreateJobQueue(ctx workflow.Context, input *batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error) {
	var output batch.CreateJobQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-CreateJobQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateJobQueueAsync(ctx workflow.Context, input *batch.CreateJobQueueInput) *CreateJobQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-CreateJobQueue", input)
	return &CreateJobQueueFuture{Future: future}
}

func (a *stub) DeleteComputeEnvironment(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error) {
	var output batch.DeleteComputeEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DeleteComputeEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteComputeEnvironmentAsync(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) *DeleteComputeEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DeleteComputeEnvironment", input)
	return &DeleteComputeEnvironmentFuture{Future: future}
}

func (a *stub) DeleteJobQueue(ctx workflow.Context, input *batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error) {
	var output batch.DeleteJobQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DeleteJobQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteJobQueueAsync(ctx workflow.Context, input *batch.DeleteJobQueueInput) *DeleteJobQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DeleteJobQueue", input)
	return &DeleteJobQueueFuture{Future: future}
}

func (a *stub) DeregisterJobDefinition(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error) {
	var output batch.DeregisterJobDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DeregisterJobDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterJobDefinitionAsync(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) *DeregisterJobDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DeregisterJobDefinition", input)
	return &DeregisterJobDefinitionFuture{Future: future}
}

func (a *stub) DescribeComputeEnvironments(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error) {
	var output batch.DescribeComputeEnvironmentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DescribeComputeEnvironments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeComputeEnvironmentsAsync(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) *DescribeComputeEnvironmentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DescribeComputeEnvironments", input)
	return &DescribeComputeEnvironmentsFuture{Future: future}
}

func (a *stub) DescribeJobDefinitions(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error) {
	var output batch.DescribeJobDefinitionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobDefinitions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobDefinitionsAsync(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) *DescribeJobDefinitionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobDefinitions", input)
	return &DescribeJobDefinitionsFuture{Future: future}
}

func (a *stub) DescribeJobQueues(ctx workflow.Context, input *batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error) {
	var output batch.DescribeJobQueuesOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobQueues", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobQueuesAsync(ctx workflow.Context, input *batch.DescribeJobQueuesInput) *DescribeJobQueuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobQueues", input)
	return &DescribeJobQueuesFuture{Future: future}
}

func (a *stub) DescribeJobs(ctx workflow.Context, input *batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error) {
	var output batch.DescribeJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobsAsync(ctx workflow.Context, input *batch.DescribeJobsInput) *DescribeJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobs", input)
	return &DescribeJobsFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *batch.ListJobsInput) (*batch.ListJobsOutput, error) {
	var output batch.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *batch.ListJobsInput) *ListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-ListJobs", input)
	return &ListJobsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *batch.ListTagsForResourceInput) (*batch.ListTagsForResourceOutput, error) {
	var output batch.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *batch.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) RegisterJobDefinition(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error) {
	var output batch.RegisterJobDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-RegisterJobDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterJobDefinitionAsync(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) *RegisterJobDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-RegisterJobDefinition", input)
	return &RegisterJobDefinitionFuture{Future: future}
}

func (a *stub) SubmitJob(ctx workflow.Context, input *batch.SubmitJobInput) (*batch.SubmitJobOutput, error) {
	var output batch.SubmitJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-SubmitJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubmitJobAsync(ctx workflow.Context, input *batch.SubmitJobInput) *SubmitJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-SubmitJob", input)
	return &SubmitJobFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *batch.TagResourceInput) (*batch.TagResourceOutput, error) {
	var output batch.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *batch.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) TerminateJob(ctx workflow.Context, input *batch.TerminateJobInput) (*batch.TerminateJobOutput, error) {
	var output batch.TerminateJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-TerminateJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TerminateJobAsync(ctx workflow.Context, input *batch.TerminateJobInput) *TerminateJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-TerminateJob", input)
	return &TerminateJobFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *batch.UntagResourceInput) (*batch.UntagResourceOutput, error) {
	var output batch.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *batch.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateComputeEnvironment(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error) {
	var output batch.UpdateComputeEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-UpdateComputeEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateComputeEnvironmentAsync(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) *UpdateComputeEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-UpdateComputeEnvironment", input)
	return &UpdateComputeEnvironmentFuture{Future: future}
}

func (a *stub) UpdateJobQueue(ctx workflow.Context, input *batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error) {
	var output batch.UpdateJobQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-UpdateJobQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateJobQueueAsync(ctx workflow.Context, input *batch.UpdateJobQueueInput) *UpdateJobQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-UpdateJobQueue", input)
	return &UpdateJobQueueFuture{Future: future}
}
