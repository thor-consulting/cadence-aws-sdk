// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package mediapackagestub

import (
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	ConfigureLogs(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error)
	ConfigureLogsAsync(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) *ConfigureLogsFuture

	CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error)
	CreateChannelAsync(ctx workflow.Context, input *mediapackage.CreateChannelInput) *CreateChannelFuture

	CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error)
	CreateHarvestJobAsync(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) *CreateHarvestJobFuture

	CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error)
	CreateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) *CreateOriginEndpointFuture

	DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error)
	DeleteChannelAsync(ctx workflow.Context, input *mediapackage.DeleteChannelInput) *DeleteChannelFuture

	DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error)
	DeleteOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) *DeleteOriginEndpointFuture

	DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error)
	DescribeChannelAsync(ctx workflow.Context, input *mediapackage.DescribeChannelInput) *DescribeChannelFuture

	DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error)
	DescribeHarvestJobAsync(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) *DescribeHarvestJobFuture

	DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error)
	DescribeOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) *DescribeOriginEndpointFuture

	ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error)
	ListChannelsAsync(ctx workflow.Context, input *mediapackage.ListChannelsInput) *ListChannelsFuture

	ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error)
	ListHarvestJobsAsync(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) *ListHarvestJobsFuture

	ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error)
	ListOriginEndpointsAsync(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) *ListOriginEndpointsFuture

	ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) *ListTagsForResourceFuture

	RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error)
	RotateChannelCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) *RotateChannelCredentialsFuture

	RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error)
	RotateIngestEndpointCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) *RotateIngestEndpointCredentialsFuture

	TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediapackage.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediapackage.UntagResourceInput) *UntagResourceFuture

	UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error)
	UpdateChannelAsync(ctx workflow.Context, input *mediapackage.UpdateChannelInput) *UpdateChannelFuture

	UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error)
	UpdateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) *UpdateOriginEndpointFuture
}

func NewClient() Client {
	return &stub{}
}
