// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package resourcegroupstaggingapistub

import (
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DescribeReportCreationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeReportCreationFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	var output resourcegroupstaggingapi.DescribeReportCreationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetComplianceSummaryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetComplianceSummaryFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	var output resourcegroupstaggingapi.GetComplianceSummaryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetResourcesFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	var output resourcegroupstaggingapi.GetResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetTagKeysFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetTagKeysFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	var output resourcegroupstaggingapi.GetTagKeysOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetTagValuesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetTagValuesFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	var output resourcegroupstaggingapi.GetTagValuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartReportCreationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartReportCreationFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	var output resourcegroupstaggingapi.StartReportCreationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourcesFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	var output resourcegroupstaggingapi.TagResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourcesFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	var output resourcegroupstaggingapi.UntagResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	var output resourcegroupstaggingapi.DescribeReportCreationOutput
	err := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-DescribeReportCreation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) *DescribeReportCreationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-DescribeReportCreation", input)
	return &DescribeReportCreationFuture{Future: future}
}

func (a *stub) GetComplianceSummary(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	var output resourcegroupstaggingapi.GetComplianceSummaryOutput
	err := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-GetComplianceSummary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetComplianceSummaryAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) *GetComplianceSummaryFuture {
	future := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-GetComplianceSummary", input)
	return &GetComplianceSummaryFuture{Future: future}
}

func (a *stub) GetResources(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	var output resourcegroupstaggingapi.GetResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-GetResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) *GetResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-GetResources", input)
	return &GetResourcesFuture{Future: future}
}

func (a *stub) GetTagKeys(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	var output resourcegroupstaggingapi.GetTagKeysOutput
	err := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-GetTagKeys", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTagKeysAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) *GetTagKeysFuture {
	future := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-GetTagKeys", input)
	return &GetTagKeysFuture{Future: future}
}

func (a *stub) GetTagValues(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	var output resourcegroupstaggingapi.GetTagValuesOutput
	err := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-GetTagValues", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTagValuesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) *GetTagValuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-GetTagValues", input)
	return &GetTagValuesFuture{Future: future}
}

func (a *stub) StartReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	var output resourcegroupstaggingapi.StartReportCreationOutput
	err := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-StartReportCreation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) *StartReportCreationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-StartReportCreation", input)
	return &StartReportCreationFuture{Future: future}
}

func (a *stub) TagResources(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	var output resourcegroupstaggingapi.TagResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-TagResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) *TagResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-TagResources", input)
	return &TagResourcesFuture{Future: future}
}

func (a *stub) UntagResources(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	var output resourcegroupstaggingapi.UntagResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-UntagResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) *UntagResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-resourcegroupstaggingapi-UntagResources", input)
	return &UntagResourcesFuture{Future: future}
}
