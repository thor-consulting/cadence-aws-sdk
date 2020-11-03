// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package forecastservicestub

import (
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateDataset(ctx workflow.Context, input *forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error)
	CreateDatasetAsync(ctx workflow.Context, input *forecastservice.CreateDatasetInput) *ForecastServiceCreateDatasetFuture

	CreateDatasetGroup(ctx workflow.Context, input *forecastservice.CreateDatasetGroupInput) (*forecastservice.CreateDatasetGroupOutput, error)
	CreateDatasetGroupAsync(ctx workflow.Context, input *forecastservice.CreateDatasetGroupInput) *ForecastServiceCreateDatasetGroupFuture

	CreateDatasetImportJob(ctx workflow.Context, input *forecastservice.CreateDatasetImportJobInput) (*forecastservice.CreateDatasetImportJobOutput, error)
	CreateDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.CreateDatasetImportJobInput) *ForecastServiceCreateDatasetImportJobFuture

	CreateForecast(ctx workflow.Context, input *forecastservice.CreateForecastInput) (*forecastservice.CreateForecastOutput, error)
	CreateForecastAsync(ctx workflow.Context, input *forecastservice.CreateForecastInput) *ForecastServiceCreateForecastFuture

	CreateForecastExportJob(ctx workflow.Context, input *forecastservice.CreateForecastExportJobInput) (*forecastservice.CreateForecastExportJobOutput, error)
	CreateForecastExportJobAsync(ctx workflow.Context, input *forecastservice.CreateForecastExportJobInput) *ForecastServiceCreateForecastExportJobFuture

	CreatePredictor(ctx workflow.Context, input *forecastservice.CreatePredictorInput) (*forecastservice.CreatePredictorOutput, error)
	CreatePredictorAsync(ctx workflow.Context, input *forecastservice.CreatePredictorInput) *ForecastServiceCreatePredictorFuture

	DeleteDataset(ctx workflow.Context, input *forecastservice.DeleteDatasetInput) (*forecastservice.DeleteDatasetOutput, error)
	DeleteDatasetAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetInput) *ForecastServiceDeleteDatasetFuture

	DeleteDatasetGroup(ctx workflow.Context, input *forecastservice.DeleteDatasetGroupInput) (*forecastservice.DeleteDatasetGroupOutput, error)
	DeleteDatasetGroupAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetGroupInput) *ForecastServiceDeleteDatasetGroupFuture

	DeleteDatasetImportJob(ctx workflow.Context, input *forecastservice.DeleteDatasetImportJobInput) (*forecastservice.DeleteDatasetImportJobOutput, error)
	DeleteDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetImportJobInput) *ForecastServiceDeleteDatasetImportJobFuture

	DeleteForecast(ctx workflow.Context, input *forecastservice.DeleteForecastInput) (*forecastservice.DeleteForecastOutput, error)
	DeleteForecastAsync(ctx workflow.Context, input *forecastservice.DeleteForecastInput) *ForecastServiceDeleteForecastFuture

	DeleteForecastExportJob(ctx workflow.Context, input *forecastservice.DeleteForecastExportJobInput) (*forecastservice.DeleteForecastExportJobOutput, error)
	DeleteForecastExportJobAsync(ctx workflow.Context, input *forecastservice.DeleteForecastExportJobInput) *ForecastServiceDeleteForecastExportJobFuture

	DeletePredictor(ctx workflow.Context, input *forecastservice.DeletePredictorInput) (*forecastservice.DeletePredictorOutput, error)
	DeletePredictorAsync(ctx workflow.Context, input *forecastservice.DeletePredictorInput) *ForecastServiceDeletePredictorFuture

	DescribeDataset(ctx workflow.Context, input *forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error)
	DescribeDatasetAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetInput) *ForecastServiceDescribeDatasetFuture

	DescribeDatasetGroup(ctx workflow.Context, input *forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error)
	DescribeDatasetGroupAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetGroupInput) *ForecastServiceDescribeDatasetGroupFuture

	DescribeDatasetImportJob(ctx workflow.Context, input *forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error)
	DescribeDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetImportJobInput) *ForecastServiceDescribeDatasetImportJobFuture

	DescribeForecast(ctx workflow.Context, input *forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error)
	DescribeForecastAsync(ctx workflow.Context, input *forecastservice.DescribeForecastInput) *ForecastServiceDescribeForecastFuture

	DescribeForecastExportJob(ctx workflow.Context, input *forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error)
	DescribeForecastExportJobAsync(ctx workflow.Context, input *forecastservice.DescribeForecastExportJobInput) *ForecastServiceDescribeForecastExportJobFuture

	DescribePredictor(ctx workflow.Context, input *forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error)
	DescribePredictorAsync(ctx workflow.Context, input *forecastservice.DescribePredictorInput) *ForecastServiceDescribePredictorFuture

	GetAccuracyMetrics(ctx workflow.Context, input *forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error)
	GetAccuracyMetricsAsync(ctx workflow.Context, input *forecastservice.GetAccuracyMetricsInput) *ForecastServiceGetAccuracyMetricsFuture

	ListDatasetGroups(ctx workflow.Context, input *forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error)
	ListDatasetGroupsAsync(ctx workflow.Context, input *forecastservice.ListDatasetGroupsInput) *ForecastServiceListDatasetGroupsFuture

	ListDatasetImportJobs(ctx workflow.Context, input *forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error)
	ListDatasetImportJobsAsync(ctx workflow.Context, input *forecastservice.ListDatasetImportJobsInput) *ForecastServiceListDatasetImportJobsFuture

	ListDatasets(ctx workflow.Context, input *forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error)
	ListDatasetsAsync(ctx workflow.Context, input *forecastservice.ListDatasetsInput) *ForecastServiceListDatasetsFuture

	ListForecastExportJobs(ctx workflow.Context, input *forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error)
	ListForecastExportJobsAsync(ctx workflow.Context, input *forecastservice.ListForecastExportJobsInput) *ForecastServiceListForecastExportJobsFuture

	ListForecasts(ctx workflow.Context, input *forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error)
	ListForecastsAsync(ctx workflow.Context, input *forecastservice.ListForecastsInput) *ForecastServiceListForecastsFuture

	ListPredictors(ctx workflow.Context, input *forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error)
	ListPredictorsAsync(ctx workflow.Context, input *forecastservice.ListPredictorsInput) *ForecastServiceListPredictorsFuture

	ListTagsForResource(ctx workflow.Context, input *forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *forecastservice.ListTagsForResourceInput) *ForecastServiceListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *forecastservice.TagResourceInput) (*forecastservice.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *forecastservice.TagResourceInput) *ForecastServiceTagResourceFuture

	UntagResource(ctx workflow.Context, input *forecastservice.UntagResourceInput) (*forecastservice.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *forecastservice.UntagResourceInput) *ForecastServiceUntagResourceFuture

	UpdateDatasetGroup(ctx workflow.Context, input *forecastservice.UpdateDatasetGroupInput) (*forecastservice.UpdateDatasetGroupOutput, error)
	UpdateDatasetGroupAsync(ctx workflow.Context, input *forecastservice.UpdateDatasetGroupInput) *ForecastServiceUpdateDatasetGroupFuture
}

func NewClient() Client {
	return &stub{}
}
