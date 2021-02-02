// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package iotfleethubstub

import (
	"github.com/aws/aws-sdk-go/service/iotfleethub"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateApplication(ctx workflow.Context, input *iotfleethub.CreateApplicationInput) (*iotfleethub.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *iotfleethub.CreateApplicationInput) *CreateApplicationFuture

	DeleteApplication(ctx workflow.Context, input *iotfleethub.DeleteApplicationInput) (*iotfleethub.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *iotfleethub.DeleteApplicationInput) *DeleteApplicationFuture

	DescribeApplication(ctx workflow.Context, input *iotfleethub.DescribeApplicationInput) (*iotfleethub.DescribeApplicationOutput, error)
	DescribeApplicationAsync(ctx workflow.Context, input *iotfleethub.DescribeApplicationInput) *DescribeApplicationFuture

	ListApplications(ctx workflow.Context, input *iotfleethub.ListApplicationsInput) (*iotfleethub.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *iotfleethub.ListApplicationsInput) *ListApplicationsFuture

	ListTagsForResource(ctx workflow.Context, input *iotfleethub.ListTagsForResourceInput) (*iotfleethub.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *iotfleethub.ListTagsForResourceInput) *ListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *iotfleethub.TagResourceInput) (*iotfleethub.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *iotfleethub.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *iotfleethub.UntagResourceInput) (*iotfleethub.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *iotfleethub.UntagResourceInput) *UntagResourceFuture

	UpdateApplication(ctx workflow.Context, input *iotfleethub.UpdateApplicationInput) (*iotfleethub.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *iotfleethub.UpdateApplicationInput) *UpdateApplicationFuture
}

func NewClient() Client {
	return &stub{}
}