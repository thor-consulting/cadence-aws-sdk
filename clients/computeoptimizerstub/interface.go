// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package computeoptimizerstub

import (
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeRecommendationExportJobs(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error)
	DescribeRecommendationExportJobsAsync(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) *DescribeRecommendationExportJobsFuture

	ExportAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error)
	ExportAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) *ExportAutoScalingGroupRecommendationsFuture

	ExportEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error)
	ExportEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) *ExportEC2InstanceRecommendationsFuture

	GetAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error)
	GetAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) *GetAutoScalingGroupRecommendationsFuture

	GetEBSVolumeRecommendations(ctx workflow.Context, input *computeoptimizer.GetEBSVolumeRecommendationsInput) (*computeoptimizer.GetEBSVolumeRecommendationsOutput, error)
	GetEBSVolumeRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetEBSVolumeRecommendationsInput) *GetEBSVolumeRecommendationsFuture

	GetEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error)
	GetEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) *GetEC2InstanceRecommendationsFuture

	GetEC2RecommendationProjectedMetrics(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error)
	GetEC2RecommendationProjectedMetricsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) *GetEC2RecommendationProjectedMetricsFuture

	GetEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error)
	GetEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) *GetEnrollmentStatusFuture

	GetLambdaFunctionRecommendations(ctx workflow.Context, input *computeoptimizer.GetLambdaFunctionRecommendationsInput) (*computeoptimizer.GetLambdaFunctionRecommendationsOutput, error)
	GetLambdaFunctionRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetLambdaFunctionRecommendationsInput) *GetLambdaFunctionRecommendationsFuture

	GetRecommendationSummaries(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error)
	GetRecommendationSummariesAsync(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) *GetRecommendationSummariesFuture

	UpdateEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error)
	UpdateEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) *UpdateEnrollmentStatusFuture
}

func NewClient() Client {
	return &stub{}
}
