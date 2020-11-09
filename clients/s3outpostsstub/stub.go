// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package s3outpostsstub

import (
	"github.com/aws/aws-sdk-go/service/s3outposts"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateEndpointFuture) Get(ctx workflow.Context) (*s3outposts.CreateEndpointOutput, error) {
	var output s3outposts.CreateEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteEndpointFuture) Get(ctx workflow.Context) (*s3outposts.DeleteEndpointOutput, error) {
	var output s3outposts.DeleteEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListEndpointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListEndpointsFuture) Get(ctx workflow.Context) (*s3outposts.ListEndpointsOutput, error) {
	var output s3outposts.ListEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEndpoint(ctx workflow.Context, input *s3outposts.CreateEndpointInput) (*s3outposts.CreateEndpointOutput, error) {
	var output s3outposts.CreateEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws-s3outposts-CreateEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEndpointAsync(ctx workflow.Context, input *s3outposts.CreateEndpointInput) *CreateEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws-s3outposts-CreateEndpoint", input)
	return &CreateEndpointFuture{Future: future}
}

func (a *stub) DeleteEndpoint(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) (*s3outposts.DeleteEndpointOutput, error) {
	var output s3outposts.DeleteEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws-s3outposts-DeleteEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEndpointAsync(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) *DeleteEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws-s3outposts-DeleteEndpoint", input)
	return &DeleteEndpointFuture{Future: future}
}

func (a *stub) ListEndpoints(ctx workflow.Context, input *s3outposts.ListEndpointsInput) (*s3outposts.ListEndpointsOutput, error) {
	var output s3outposts.ListEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws-s3outposts-ListEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEndpointsAsync(ctx workflow.Context, input *s3outposts.ListEndpointsInput) *ListEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-s3outposts-ListEndpoints", input)
	return &ListEndpointsFuture{Future: future}
}
