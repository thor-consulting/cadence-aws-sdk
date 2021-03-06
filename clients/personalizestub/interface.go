// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package personalizestub

import (
	"github.com/aws/aws-sdk-go/service/personalize"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateBatchInferenceJob(ctx workflow.Context, input *personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error)
	CreateBatchInferenceJobAsync(ctx workflow.Context, input *personalize.CreateBatchInferenceJobInput) *CreateBatchInferenceJobFuture

	CreateCampaign(ctx workflow.Context, input *personalize.CreateCampaignInput) (*personalize.CreateCampaignOutput, error)
	CreateCampaignAsync(ctx workflow.Context, input *personalize.CreateCampaignInput) *CreateCampaignFuture

	CreateDataset(ctx workflow.Context, input *personalize.CreateDatasetInput) (*personalize.CreateDatasetOutput, error)
	CreateDatasetAsync(ctx workflow.Context, input *personalize.CreateDatasetInput) *CreateDatasetFuture

	CreateDatasetGroup(ctx workflow.Context, input *personalize.CreateDatasetGroupInput) (*personalize.CreateDatasetGroupOutput, error)
	CreateDatasetGroupAsync(ctx workflow.Context, input *personalize.CreateDatasetGroupInput) *CreateDatasetGroupFuture

	CreateDatasetImportJob(ctx workflow.Context, input *personalize.CreateDatasetImportJobInput) (*personalize.CreateDatasetImportJobOutput, error)
	CreateDatasetImportJobAsync(ctx workflow.Context, input *personalize.CreateDatasetImportJobInput) *CreateDatasetImportJobFuture

	CreateEventTracker(ctx workflow.Context, input *personalize.CreateEventTrackerInput) (*personalize.CreateEventTrackerOutput, error)
	CreateEventTrackerAsync(ctx workflow.Context, input *personalize.CreateEventTrackerInput) *CreateEventTrackerFuture

	CreateFilter(ctx workflow.Context, input *personalize.CreateFilterInput) (*personalize.CreateFilterOutput, error)
	CreateFilterAsync(ctx workflow.Context, input *personalize.CreateFilterInput) *CreateFilterFuture

	CreateSchema(ctx workflow.Context, input *personalize.CreateSchemaInput) (*personalize.CreateSchemaOutput, error)
	CreateSchemaAsync(ctx workflow.Context, input *personalize.CreateSchemaInput) *CreateSchemaFuture

	CreateSolution(ctx workflow.Context, input *personalize.CreateSolutionInput) (*personalize.CreateSolutionOutput, error)
	CreateSolutionAsync(ctx workflow.Context, input *personalize.CreateSolutionInput) *CreateSolutionFuture

	CreateSolutionVersion(ctx workflow.Context, input *personalize.CreateSolutionVersionInput) (*personalize.CreateSolutionVersionOutput, error)
	CreateSolutionVersionAsync(ctx workflow.Context, input *personalize.CreateSolutionVersionInput) *CreateSolutionVersionFuture

	DeleteCampaign(ctx workflow.Context, input *personalize.DeleteCampaignInput) (*personalize.DeleteCampaignOutput, error)
	DeleteCampaignAsync(ctx workflow.Context, input *personalize.DeleteCampaignInput) *DeleteCampaignFuture

	DeleteDataset(ctx workflow.Context, input *personalize.DeleteDatasetInput) (*personalize.DeleteDatasetOutput, error)
	DeleteDatasetAsync(ctx workflow.Context, input *personalize.DeleteDatasetInput) *DeleteDatasetFuture

	DeleteDatasetGroup(ctx workflow.Context, input *personalize.DeleteDatasetGroupInput) (*personalize.DeleteDatasetGroupOutput, error)
	DeleteDatasetGroupAsync(ctx workflow.Context, input *personalize.DeleteDatasetGroupInput) *DeleteDatasetGroupFuture

	DeleteEventTracker(ctx workflow.Context, input *personalize.DeleteEventTrackerInput) (*personalize.DeleteEventTrackerOutput, error)
	DeleteEventTrackerAsync(ctx workflow.Context, input *personalize.DeleteEventTrackerInput) *DeleteEventTrackerFuture

	DeleteFilter(ctx workflow.Context, input *personalize.DeleteFilterInput) (*personalize.DeleteFilterOutput, error)
	DeleteFilterAsync(ctx workflow.Context, input *personalize.DeleteFilterInput) *DeleteFilterFuture

	DeleteSchema(ctx workflow.Context, input *personalize.DeleteSchemaInput) (*personalize.DeleteSchemaOutput, error)
	DeleteSchemaAsync(ctx workflow.Context, input *personalize.DeleteSchemaInput) *DeleteSchemaFuture

	DeleteSolution(ctx workflow.Context, input *personalize.DeleteSolutionInput) (*personalize.DeleteSolutionOutput, error)
	DeleteSolutionAsync(ctx workflow.Context, input *personalize.DeleteSolutionInput) *DeleteSolutionFuture

	DescribeAlgorithm(ctx workflow.Context, input *personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error)
	DescribeAlgorithmAsync(ctx workflow.Context, input *personalize.DescribeAlgorithmInput) *DescribeAlgorithmFuture

	DescribeBatchInferenceJob(ctx workflow.Context, input *personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error)
	DescribeBatchInferenceJobAsync(ctx workflow.Context, input *personalize.DescribeBatchInferenceJobInput) *DescribeBatchInferenceJobFuture

	DescribeCampaign(ctx workflow.Context, input *personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error)
	DescribeCampaignAsync(ctx workflow.Context, input *personalize.DescribeCampaignInput) *DescribeCampaignFuture

	DescribeDataset(ctx workflow.Context, input *personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error)
	DescribeDatasetAsync(ctx workflow.Context, input *personalize.DescribeDatasetInput) *DescribeDatasetFuture

	DescribeDatasetGroup(ctx workflow.Context, input *personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error)
	DescribeDatasetGroupAsync(ctx workflow.Context, input *personalize.DescribeDatasetGroupInput) *DescribeDatasetGroupFuture

	DescribeDatasetImportJob(ctx workflow.Context, input *personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error)
	DescribeDatasetImportJobAsync(ctx workflow.Context, input *personalize.DescribeDatasetImportJobInput) *DescribeDatasetImportJobFuture

	DescribeEventTracker(ctx workflow.Context, input *personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error)
	DescribeEventTrackerAsync(ctx workflow.Context, input *personalize.DescribeEventTrackerInput) *DescribeEventTrackerFuture

	DescribeFeatureTransformation(ctx workflow.Context, input *personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error)
	DescribeFeatureTransformationAsync(ctx workflow.Context, input *personalize.DescribeFeatureTransformationInput) *DescribeFeatureTransformationFuture

	DescribeFilter(ctx workflow.Context, input *personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error)
	DescribeFilterAsync(ctx workflow.Context, input *personalize.DescribeFilterInput) *DescribeFilterFuture

	DescribeRecipe(ctx workflow.Context, input *personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error)
	DescribeRecipeAsync(ctx workflow.Context, input *personalize.DescribeRecipeInput) *DescribeRecipeFuture

	DescribeSchema(ctx workflow.Context, input *personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error)
	DescribeSchemaAsync(ctx workflow.Context, input *personalize.DescribeSchemaInput) *DescribeSchemaFuture

	DescribeSolution(ctx workflow.Context, input *personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error)
	DescribeSolutionAsync(ctx workflow.Context, input *personalize.DescribeSolutionInput) *DescribeSolutionFuture

	DescribeSolutionVersion(ctx workflow.Context, input *personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error)
	DescribeSolutionVersionAsync(ctx workflow.Context, input *personalize.DescribeSolutionVersionInput) *DescribeSolutionVersionFuture

	GetSolutionMetrics(ctx workflow.Context, input *personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error)
	GetSolutionMetricsAsync(ctx workflow.Context, input *personalize.GetSolutionMetricsInput) *GetSolutionMetricsFuture

	ListBatchInferenceJobs(ctx workflow.Context, input *personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error)
	ListBatchInferenceJobsAsync(ctx workflow.Context, input *personalize.ListBatchInferenceJobsInput) *ListBatchInferenceJobsFuture

	ListCampaigns(ctx workflow.Context, input *personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error)
	ListCampaignsAsync(ctx workflow.Context, input *personalize.ListCampaignsInput) *ListCampaignsFuture

	ListDatasetGroups(ctx workflow.Context, input *personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error)
	ListDatasetGroupsAsync(ctx workflow.Context, input *personalize.ListDatasetGroupsInput) *ListDatasetGroupsFuture

	ListDatasetImportJobs(ctx workflow.Context, input *personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error)
	ListDatasetImportJobsAsync(ctx workflow.Context, input *personalize.ListDatasetImportJobsInput) *ListDatasetImportJobsFuture

	ListDatasets(ctx workflow.Context, input *personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error)
	ListDatasetsAsync(ctx workflow.Context, input *personalize.ListDatasetsInput) *ListDatasetsFuture

	ListEventTrackers(ctx workflow.Context, input *personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error)
	ListEventTrackersAsync(ctx workflow.Context, input *personalize.ListEventTrackersInput) *ListEventTrackersFuture

	ListFilters(ctx workflow.Context, input *personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error)
	ListFiltersAsync(ctx workflow.Context, input *personalize.ListFiltersInput) *ListFiltersFuture

	ListRecipes(ctx workflow.Context, input *personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error)
	ListRecipesAsync(ctx workflow.Context, input *personalize.ListRecipesInput) *ListRecipesFuture

	ListSchemas(ctx workflow.Context, input *personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error)
	ListSchemasAsync(ctx workflow.Context, input *personalize.ListSchemasInput) *ListSchemasFuture

	ListSolutionVersions(ctx workflow.Context, input *personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error)
	ListSolutionVersionsAsync(ctx workflow.Context, input *personalize.ListSolutionVersionsInput) *ListSolutionVersionsFuture

	ListSolutions(ctx workflow.Context, input *personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error)
	ListSolutionsAsync(ctx workflow.Context, input *personalize.ListSolutionsInput) *ListSolutionsFuture

	UpdateCampaign(ctx workflow.Context, input *personalize.UpdateCampaignInput) (*personalize.UpdateCampaignOutput, error)
	UpdateCampaignAsync(ctx workflow.Context, input *personalize.UpdateCampaignInput) *UpdateCampaignFuture
}

func NewClient() Client {
	return &stub{}
}
