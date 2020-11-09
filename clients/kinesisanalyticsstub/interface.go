// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package kinesisanalyticsstub

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) *AddApplicationCloudWatchLoggingOptionFuture

	AddApplicationInput(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error)
	AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) *AddApplicationInputFuture

	AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) *AddApplicationInputProcessingConfigurationFuture

	AddApplicationOutput(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error)
	AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) *AddApplicationOutputFuture

	AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) *AddApplicationReferenceDataSourceFuture

	CreateApplication(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) *CreateApplicationFuture

	DeleteApplication(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) *DeleteApplicationFuture

	DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) *DeleteApplicationCloudWatchLoggingOptionFuture

	DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) *DeleteApplicationInputProcessingConfigurationFuture

	DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) *DeleteApplicationOutputFuture

	DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) *DeleteApplicationReferenceDataSourceFuture

	DescribeApplication(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error)
	DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) *DescribeApplicationFuture

	DiscoverInputSchema(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) *DiscoverInputSchemaFuture

	ListApplications(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) *ListApplicationsFuture

	ListTagsForResource(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) *ListTagsForResourceFuture

	StartApplication(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error)
	StartApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) *StartApplicationFuture

	StopApplication(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error)
	StopApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) *StopApplicationFuture

	TagResource(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) *UntagResourceFuture

	UpdateApplication(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) *UpdateApplicationFuture
}

func NewClient() Client {
	return &stub{}
}
