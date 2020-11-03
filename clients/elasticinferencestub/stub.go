// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package elasticinferencestub

import (
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ElasticInferenceDescribeAcceleratorOfferingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ElasticInferenceDescribeAcceleratorOfferingsFuture) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
	var output elasticinference.DescribeAcceleratorOfferingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceDescribeAcceleratorTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ElasticInferenceDescribeAcceleratorTypesFuture) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
	var output elasticinference.DescribeAcceleratorTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceDescribeAcceleratorsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ElasticInferenceDescribeAcceleratorsFuture) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorsOutput, error) {
	var output elasticinference.DescribeAcceleratorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ElasticInferenceListTagsForResourceFuture) Get(ctx workflow.Context) (*elasticinference.ListTagsForResourceOutput, error) {
	var output elasticinference.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ElasticInferenceTagResourceFuture) Get(ctx workflow.Context) (*elasticinference.TagResourceOutput, error) {
	var output elasticinference.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ElasticInferenceUntagResourceFuture) Get(ctx workflow.Context) (*elasticinference.UntagResourceOutput, error) {
	var output elasticinference.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAcceleratorOfferings(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
	var output elasticinference.DescribeAcceleratorOfferingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elasticinference-DescribeAcceleratorOfferings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAcceleratorOfferingsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) *ElasticInferenceDescribeAcceleratorOfferingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elasticinference-DescribeAcceleratorOfferings", input)
	return &ElasticInferenceDescribeAcceleratorOfferingsFuture{Future: future}
}

func (a *stub) DescribeAcceleratorTypes(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
	var output elasticinference.DescribeAcceleratorTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elasticinference-DescribeAcceleratorTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAcceleratorTypesAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) *ElasticInferenceDescribeAcceleratorTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elasticinference-DescribeAcceleratorTypes", input)
	return &ElasticInferenceDescribeAcceleratorTypesFuture{Future: future}
}

func (a *stub) DescribeAccelerators(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error) {
	var output elasticinference.DescribeAcceleratorsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elasticinference-DescribeAccelerators", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAcceleratorsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) *ElasticInferenceDescribeAcceleratorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elasticinference-DescribeAccelerators", input)
	return &ElasticInferenceDescribeAcceleratorsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error) {
	var output elasticinference.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-elasticinference-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) *ElasticInferenceListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elasticinference-ListTagsForResource", input)
	return &ElasticInferenceListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error) {
	var output elasticinference.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-elasticinference-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *elasticinference.TagResourceInput) *ElasticInferenceTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elasticinference-TagResource", input)
	return &ElasticInferenceTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error) {
	var output elasticinference.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-elasticinference-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *elasticinference.UntagResourceInput) *ElasticInferenceUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elasticinference-UntagResource", input)
	return &ElasticInferenceUntagResourceFuture{Future: future}
}
