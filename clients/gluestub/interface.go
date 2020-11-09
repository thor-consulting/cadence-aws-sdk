// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package gluestub

import (
	"github.com/aws/aws-sdk-go/service/glue"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchCreatePartition(ctx workflow.Context, input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error)
	BatchCreatePartitionAsync(ctx workflow.Context, input *glue.BatchCreatePartitionInput) *BatchCreatePartitionFuture

	BatchDeleteConnection(ctx workflow.Context, input *glue.BatchDeleteConnectionInput) (*glue.BatchDeleteConnectionOutput, error)
	BatchDeleteConnectionAsync(ctx workflow.Context, input *glue.BatchDeleteConnectionInput) *BatchDeleteConnectionFuture

	BatchDeletePartition(ctx workflow.Context, input *glue.BatchDeletePartitionInput) (*glue.BatchDeletePartitionOutput, error)
	BatchDeletePartitionAsync(ctx workflow.Context, input *glue.BatchDeletePartitionInput) *BatchDeletePartitionFuture

	BatchDeleteTable(ctx workflow.Context, input *glue.BatchDeleteTableInput) (*glue.BatchDeleteTableOutput, error)
	BatchDeleteTableAsync(ctx workflow.Context, input *glue.BatchDeleteTableInput) *BatchDeleteTableFuture

	BatchDeleteTableVersion(ctx workflow.Context, input *glue.BatchDeleteTableVersionInput) (*glue.BatchDeleteTableVersionOutput, error)
	BatchDeleteTableVersionAsync(ctx workflow.Context, input *glue.BatchDeleteTableVersionInput) *BatchDeleteTableVersionFuture

	BatchGetCrawlers(ctx workflow.Context, input *glue.BatchGetCrawlersInput) (*glue.BatchGetCrawlersOutput, error)
	BatchGetCrawlersAsync(ctx workflow.Context, input *glue.BatchGetCrawlersInput) *BatchGetCrawlersFuture

	BatchGetDevEndpoints(ctx workflow.Context, input *glue.BatchGetDevEndpointsInput) (*glue.BatchGetDevEndpointsOutput, error)
	BatchGetDevEndpointsAsync(ctx workflow.Context, input *glue.BatchGetDevEndpointsInput) *BatchGetDevEndpointsFuture

	BatchGetJobs(ctx workflow.Context, input *glue.BatchGetJobsInput) (*glue.BatchGetJobsOutput, error)
	BatchGetJobsAsync(ctx workflow.Context, input *glue.BatchGetJobsInput) *BatchGetJobsFuture

	BatchGetPartition(ctx workflow.Context, input *glue.BatchGetPartitionInput) (*glue.BatchGetPartitionOutput, error)
	BatchGetPartitionAsync(ctx workflow.Context, input *glue.BatchGetPartitionInput) *BatchGetPartitionFuture

	BatchGetTriggers(ctx workflow.Context, input *glue.BatchGetTriggersInput) (*glue.BatchGetTriggersOutput, error)
	BatchGetTriggersAsync(ctx workflow.Context, input *glue.BatchGetTriggersInput) *BatchGetTriggersFuture

	BatchGetWorkflows(ctx workflow.Context, input *glue.BatchGetWorkflowsInput) (*glue.BatchGetWorkflowsOutput, error)
	BatchGetWorkflowsAsync(ctx workflow.Context, input *glue.BatchGetWorkflowsInput) *BatchGetWorkflowsFuture

	BatchStopJobRun(ctx workflow.Context, input *glue.BatchStopJobRunInput) (*glue.BatchStopJobRunOutput, error)
	BatchStopJobRunAsync(ctx workflow.Context, input *glue.BatchStopJobRunInput) *BatchStopJobRunFuture

	BatchUpdatePartition(ctx workflow.Context, input *glue.BatchUpdatePartitionInput) (*glue.BatchUpdatePartitionOutput, error)
	BatchUpdatePartitionAsync(ctx workflow.Context, input *glue.BatchUpdatePartitionInput) *BatchUpdatePartitionFuture

	CancelMLTaskRun(ctx workflow.Context, input *glue.CancelMLTaskRunInput) (*glue.CancelMLTaskRunOutput, error)
	CancelMLTaskRunAsync(ctx workflow.Context, input *glue.CancelMLTaskRunInput) *CancelMLTaskRunFuture

	CreateClassifier(ctx workflow.Context, input *glue.CreateClassifierInput) (*glue.CreateClassifierOutput, error)
	CreateClassifierAsync(ctx workflow.Context, input *glue.CreateClassifierInput) *CreateClassifierFuture

	CreateConnection(ctx workflow.Context, input *glue.CreateConnectionInput) (*glue.CreateConnectionOutput, error)
	CreateConnectionAsync(ctx workflow.Context, input *glue.CreateConnectionInput) *CreateConnectionFuture

	CreateCrawler(ctx workflow.Context, input *glue.CreateCrawlerInput) (*glue.CreateCrawlerOutput, error)
	CreateCrawlerAsync(ctx workflow.Context, input *glue.CreateCrawlerInput) *CreateCrawlerFuture

	CreateDatabase(ctx workflow.Context, input *glue.CreateDatabaseInput) (*glue.CreateDatabaseOutput, error)
	CreateDatabaseAsync(ctx workflow.Context, input *glue.CreateDatabaseInput) *CreateDatabaseFuture

	CreateDevEndpoint(ctx workflow.Context, input *glue.CreateDevEndpointInput) (*glue.CreateDevEndpointOutput, error)
	CreateDevEndpointAsync(ctx workflow.Context, input *glue.CreateDevEndpointInput) *CreateDevEndpointFuture

	CreateJob(ctx workflow.Context, input *glue.CreateJobInput) (*glue.CreateJobOutput, error)
	CreateJobAsync(ctx workflow.Context, input *glue.CreateJobInput) *CreateJobFuture

	CreateMLTransform(ctx workflow.Context, input *glue.CreateMLTransformInput) (*glue.CreateMLTransformOutput, error)
	CreateMLTransformAsync(ctx workflow.Context, input *glue.CreateMLTransformInput) *CreateMLTransformFuture

	CreatePartition(ctx workflow.Context, input *glue.CreatePartitionInput) (*glue.CreatePartitionOutput, error)
	CreatePartitionAsync(ctx workflow.Context, input *glue.CreatePartitionInput) *CreatePartitionFuture

	CreateScript(ctx workflow.Context, input *glue.CreateScriptInput) (*glue.CreateScriptOutput, error)
	CreateScriptAsync(ctx workflow.Context, input *glue.CreateScriptInput) *CreateScriptFuture

	CreateSecurityConfiguration(ctx workflow.Context, input *glue.CreateSecurityConfigurationInput) (*glue.CreateSecurityConfigurationOutput, error)
	CreateSecurityConfigurationAsync(ctx workflow.Context, input *glue.CreateSecurityConfigurationInput) *CreateSecurityConfigurationFuture

	CreateTable(ctx workflow.Context, input *glue.CreateTableInput) (*glue.CreateTableOutput, error)
	CreateTableAsync(ctx workflow.Context, input *glue.CreateTableInput) *CreateTableFuture

	CreateTrigger(ctx workflow.Context, input *glue.CreateTriggerInput) (*glue.CreateTriggerOutput, error)
	CreateTriggerAsync(ctx workflow.Context, input *glue.CreateTriggerInput) *CreateTriggerFuture

	CreateUserDefinedFunction(ctx workflow.Context, input *glue.CreateUserDefinedFunctionInput) (*glue.CreateUserDefinedFunctionOutput, error)
	CreateUserDefinedFunctionAsync(ctx workflow.Context, input *glue.CreateUserDefinedFunctionInput) *CreateUserDefinedFunctionFuture

	CreateWorkflow(ctx workflow.Context, input *glue.CreateWorkflowInput) (*glue.CreateWorkflowOutput, error)
	CreateWorkflowAsync(ctx workflow.Context, input *glue.CreateWorkflowInput) *CreateWorkflowFuture

	DeleteClassifier(ctx workflow.Context, input *glue.DeleteClassifierInput) (*glue.DeleteClassifierOutput, error)
	DeleteClassifierAsync(ctx workflow.Context, input *glue.DeleteClassifierInput) *DeleteClassifierFuture

	DeleteColumnStatisticsForPartition(ctx workflow.Context, input *glue.DeleteColumnStatisticsForPartitionInput) (*glue.DeleteColumnStatisticsForPartitionOutput, error)
	DeleteColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.DeleteColumnStatisticsForPartitionInput) *DeleteColumnStatisticsForPartitionFuture

	DeleteColumnStatisticsForTable(ctx workflow.Context, input *glue.DeleteColumnStatisticsForTableInput) (*glue.DeleteColumnStatisticsForTableOutput, error)
	DeleteColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.DeleteColumnStatisticsForTableInput) *DeleteColumnStatisticsForTableFuture

	DeleteConnection(ctx workflow.Context, input *glue.DeleteConnectionInput) (*glue.DeleteConnectionOutput, error)
	DeleteConnectionAsync(ctx workflow.Context, input *glue.DeleteConnectionInput) *DeleteConnectionFuture

	DeleteCrawler(ctx workflow.Context, input *glue.DeleteCrawlerInput) (*glue.DeleteCrawlerOutput, error)
	DeleteCrawlerAsync(ctx workflow.Context, input *glue.DeleteCrawlerInput) *DeleteCrawlerFuture

	DeleteDatabase(ctx workflow.Context, input *glue.DeleteDatabaseInput) (*glue.DeleteDatabaseOutput, error)
	DeleteDatabaseAsync(ctx workflow.Context, input *glue.DeleteDatabaseInput) *DeleteDatabaseFuture

	DeleteDevEndpoint(ctx workflow.Context, input *glue.DeleteDevEndpointInput) (*glue.DeleteDevEndpointOutput, error)
	DeleteDevEndpointAsync(ctx workflow.Context, input *glue.DeleteDevEndpointInput) *DeleteDevEndpointFuture

	DeleteJob(ctx workflow.Context, input *glue.DeleteJobInput) (*glue.DeleteJobOutput, error)
	DeleteJobAsync(ctx workflow.Context, input *glue.DeleteJobInput) *DeleteJobFuture

	DeleteMLTransform(ctx workflow.Context, input *glue.DeleteMLTransformInput) (*glue.DeleteMLTransformOutput, error)
	DeleteMLTransformAsync(ctx workflow.Context, input *glue.DeleteMLTransformInput) *DeleteMLTransformFuture

	DeletePartition(ctx workflow.Context, input *glue.DeletePartitionInput) (*glue.DeletePartitionOutput, error)
	DeletePartitionAsync(ctx workflow.Context, input *glue.DeletePartitionInput) *DeletePartitionFuture

	DeleteResourcePolicy(ctx workflow.Context, input *glue.DeleteResourcePolicyInput) (*glue.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyAsync(ctx workflow.Context, input *glue.DeleteResourcePolicyInput) *DeleteResourcePolicyFuture

	DeleteSecurityConfiguration(ctx workflow.Context, input *glue.DeleteSecurityConfigurationInput) (*glue.DeleteSecurityConfigurationOutput, error)
	DeleteSecurityConfigurationAsync(ctx workflow.Context, input *glue.DeleteSecurityConfigurationInput) *DeleteSecurityConfigurationFuture

	DeleteTable(ctx workflow.Context, input *glue.DeleteTableInput) (*glue.DeleteTableOutput, error)
	DeleteTableAsync(ctx workflow.Context, input *glue.DeleteTableInput) *DeleteTableFuture

	DeleteTableVersion(ctx workflow.Context, input *glue.DeleteTableVersionInput) (*glue.DeleteTableVersionOutput, error)
	DeleteTableVersionAsync(ctx workflow.Context, input *glue.DeleteTableVersionInput) *DeleteTableVersionFuture

	DeleteTrigger(ctx workflow.Context, input *glue.DeleteTriggerInput) (*glue.DeleteTriggerOutput, error)
	DeleteTriggerAsync(ctx workflow.Context, input *glue.DeleteTriggerInput) *DeleteTriggerFuture

	DeleteUserDefinedFunction(ctx workflow.Context, input *glue.DeleteUserDefinedFunctionInput) (*glue.DeleteUserDefinedFunctionOutput, error)
	DeleteUserDefinedFunctionAsync(ctx workflow.Context, input *glue.DeleteUserDefinedFunctionInput) *DeleteUserDefinedFunctionFuture

	DeleteWorkflow(ctx workflow.Context, input *glue.DeleteWorkflowInput) (*glue.DeleteWorkflowOutput, error)
	DeleteWorkflowAsync(ctx workflow.Context, input *glue.DeleteWorkflowInput) *DeleteWorkflowFuture

	GetCatalogImportStatus(ctx workflow.Context, input *glue.GetCatalogImportStatusInput) (*glue.GetCatalogImportStatusOutput, error)
	GetCatalogImportStatusAsync(ctx workflow.Context, input *glue.GetCatalogImportStatusInput) *GetCatalogImportStatusFuture

	GetClassifier(ctx workflow.Context, input *glue.GetClassifierInput) (*glue.GetClassifierOutput, error)
	GetClassifierAsync(ctx workflow.Context, input *glue.GetClassifierInput) *GetClassifierFuture

	GetClassifiers(ctx workflow.Context, input *glue.GetClassifiersInput) (*glue.GetClassifiersOutput, error)
	GetClassifiersAsync(ctx workflow.Context, input *glue.GetClassifiersInput) *GetClassifiersFuture

	GetColumnStatisticsForPartition(ctx workflow.Context, input *glue.GetColumnStatisticsForPartitionInput) (*glue.GetColumnStatisticsForPartitionOutput, error)
	GetColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.GetColumnStatisticsForPartitionInput) *GetColumnStatisticsForPartitionFuture

	GetColumnStatisticsForTable(ctx workflow.Context, input *glue.GetColumnStatisticsForTableInput) (*glue.GetColumnStatisticsForTableOutput, error)
	GetColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.GetColumnStatisticsForTableInput) *GetColumnStatisticsForTableFuture

	GetConnection(ctx workflow.Context, input *glue.GetConnectionInput) (*glue.GetConnectionOutput, error)
	GetConnectionAsync(ctx workflow.Context, input *glue.GetConnectionInput) *GetConnectionFuture

	GetConnections(ctx workflow.Context, input *glue.GetConnectionsInput) (*glue.GetConnectionsOutput, error)
	GetConnectionsAsync(ctx workflow.Context, input *glue.GetConnectionsInput) *GetConnectionsFuture

	GetCrawler(ctx workflow.Context, input *glue.GetCrawlerInput) (*glue.GetCrawlerOutput, error)
	GetCrawlerAsync(ctx workflow.Context, input *glue.GetCrawlerInput) *GetCrawlerFuture

	GetCrawlerMetrics(ctx workflow.Context, input *glue.GetCrawlerMetricsInput) (*glue.GetCrawlerMetricsOutput, error)
	GetCrawlerMetricsAsync(ctx workflow.Context, input *glue.GetCrawlerMetricsInput) *GetCrawlerMetricsFuture

	GetCrawlers(ctx workflow.Context, input *glue.GetCrawlersInput) (*glue.GetCrawlersOutput, error)
	GetCrawlersAsync(ctx workflow.Context, input *glue.GetCrawlersInput) *GetCrawlersFuture

	GetDataCatalogEncryptionSettings(ctx workflow.Context, input *glue.GetDataCatalogEncryptionSettingsInput) (*glue.GetDataCatalogEncryptionSettingsOutput, error)
	GetDataCatalogEncryptionSettingsAsync(ctx workflow.Context, input *glue.GetDataCatalogEncryptionSettingsInput) *GetDataCatalogEncryptionSettingsFuture

	GetDatabase(ctx workflow.Context, input *glue.GetDatabaseInput) (*glue.GetDatabaseOutput, error)
	GetDatabaseAsync(ctx workflow.Context, input *glue.GetDatabaseInput) *GetDatabaseFuture

	GetDatabases(ctx workflow.Context, input *glue.GetDatabasesInput) (*glue.GetDatabasesOutput, error)
	GetDatabasesAsync(ctx workflow.Context, input *glue.GetDatabasesInput) *GetDatabasesFuture

	GetDataflowGraph(ctx workflow.Context, input *glue.GetDataflowGraphInput) (*glue.GetDataflowGraphOutput, error)
	GetDataflowGraphAsync(ctx workflow.Context, input *glue.GetDataflowGraphInput) *GetDataflowGraphFuture

	GetDevEndpoint(ctx workflow.Context, input *glue.GetDevEndpointInput) (*glue.GetDevEndpointOutput, error)
	GetDevEndpointAsync(ctx workflow.Context, input *glue.GetDevEndpointInput) *GetDevEndpointFuture

	GetDevEndpoints(ctx workflow.Context, input *glue.GetDevEndpointsInput) (*glue.GetDevEndpointsOutput, error)
	GetDevEndpointsAsync(ctx workflow.Context, input *glue.GetDevEndpointsInput) *GetDevEndpointsFuture

	GetJob(ctx workflow.Context, input *glue.GetJobInput) (*glue.GetJobOutput, error)
	GetJobAsync(ctx workflow.Context, input *glue.GetJobInput) *GetJobFuture

	GetJobBookmark(ctx workflow.Context, input *glue.GetJobBookmarkInput) (*glue.GetJobBookmarkOutput, error)
	GetJobBookmarkAsync(ctx workflow.Context, input *glue.GetJobBookmarkInput) *GetJobBookmarkFuture

	GetJobRun(ctx workflow.Context, input *glue.GetJobRunInput) (*glue.GetJobRunOutput, error)
	GetJobRunAsync(ctx workflow.Context, input *glue.GetJobRunInput) *GetJobRunFuture

	GetJobRuns(ctx workflow.Context, input *glue.GetJobRunsInput) (*glue.GetJobRunsOutput, error)
	GetJobRunsAsync(ctx workflow.Context, input *glue.GetJobRunsInput) *GetJobRunsFuture

	GetJobs(ctx workflow.Context, input *glue.GetJobsInput) (*glue.GetJobsOutput, error)
	GetJobsAsync(ctx workflow.Context, input *glue.GetJobsInput) *GetJobsFuture

	GetMLTaskRun(ctx workflow.Context, input *glue.GetMLTaskRunInput) (*glue.GetMLTaskRunOutput, error)
	GetMLTaskRunAsync(ctx workflow.Context, input *glue.GetMLTaskRunInput) *GetMLTaskRunFuture

	GetMLTaskRuns(ctx workflow.Context, input *glue.GetMLTaskRunsInput) (*glue.GetMLTaskRunsOutput, error)
	GetMLTaskRunsAsync(ctx workflow.Context, input *glue.GetMLTaskRunsInput) *GetMLTaskRunsFuture

	GetMLTransform(ctx workflow.Context, input *glue.GetMLTransformInput) (*glue.GetMLTransformOutput, error)
	GetMLTransformAsync(ctx workflow.Context, input *glue.GetMLTransformInput) *GetMLTransformFuture

	GetMLTransforms(ctx workflow.Context, input *glue.GetMLTransformsInput) (*glue.GetMLTransformsOutput, error)
	GetMLTransformsAsync(ctx workflow.Context, input *glue.GetMLTransformsInput) *GetMLTransformsFuture

	GetMapping(ctx workflow.Context, input *glue.GetMappingInput) (*glue.GetMappingOutput, error)
	GetMappingAsync(ctx workflow.Context, input *glue.GetMappingInput) *GetMappingFuture

	GetPartition(ctx workflow.Context, input *glue.GetPartitionInput) (*glue.GetPartitionOutput, error)
	GetPartitionAsync(ctx workflow.Context, input *glue.GetPartitionInput) *GetPartitionFuture

	GetPartitionIndexes(ctx workflow.Context, input *glue.GetPartitionIndexesInput) (*glue.GetPartitionIndexesOutput, error)
	GetPartitionIndexesAsync(ctx workflow.Context, input *glue.GetPartitionIndexesInput) *GetPartitionIndexesFuture

	GetPartitions(ctx workflow.Context, input *glue.GetPartitionsInput) (*glue.GetPartitionsOutput, error)
	GetPartitionsAsync(ctx workflow.Context, input *glue.GetPartitionsInput) *GetPartitionsFuture

	GetPlan(ctx workflow.Context, input *glue.GetPlanInput) (*glue.GetPlanOutput, error)
	GetPlanAsync(ctx workflow.Context, input *glue.GetPlanInput) *GetPlanFuture

	GetResourcePolicies(ctx workflow.Context, input *glue.GetResourcePoliciesInput) (*glue.GetResourcePoliciesOutput, error)
	GetResourcePoliciesAsync(ctx workflow.Context, input *glue.GetResourcePoliciesInput) *GetResourcePoliciesFuture

	GetResourcePolicy(ctx workflow.Context, input *glue.GetResourcePolicyInput) (*glue.GetResourcePolicyOutput, error)
	GetResourcePolicyAsync(ctx workflow.Context, input *glue.GetResourcePolicyInput) *GetResourcePolicyFuture

	GetSecurityConfiguration(ctx workflow.Context, input *glue.GetSecurityConfigurationInput) (*glue.GetSecurityConfigurationOutput, error)
	GetSecurityConfigurationAsync(ctx workflow.Context, input *glue.GetSecurityConfigurationInput) *GetSecurityConfigurationFuture

	GetSecurityConfigurations(ctx workflow.Context, input *glue.GetSecurityConfigurationsInput) (*glue.GetSecurityConfigurationsOutput, error)
	GetSecurityConfigurationsAsync(ctx workflow.Context, input *glue.GetSecurityConfigurationsInput) *GetSecurityConfigurationsFuture

	GetTable(ctx workflow.Context, input *glue.GetTableInput) (*glue.GetTableOutput, error)
	GetTableAsync(ctx workflow.Context, input *glue.GetTableInput) *GetTableFuture

	GetTableVersion(ctx workflow.Context, input *glue.GetTableVersionInput) (*glue.GetTableVersionOutput, error)
	GetTableVersionAsync(ctx workflow.Context, input *glue.GetTableVersionInput) *GetTableVersionFuture

	GetTableVersions(ctx workflow.Context, input *glue.GetTableVersionsInput) (*glue.GetTableVersionsOutput, error)
	GetTableVersionsAsync(ctx workflow.Context, input *glue.GetTableVersionsInput) *GetTableVersionsFuture

	GetTables(ctx workflow.Context, input *glue.GetTablesInput) (*glue.GetTablesOutput, error)
	GetTablesAsync(ctx workflow.Context, input *glue.GetTablesInput) *GetTablesFuture

	GetTags(ctx workflow.Context, input *glue.GetTagsInput) (*glue.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *glue.GetTagsInput) *GetTagsFuture

	GetTrigger(ctx workflow.Context, input *glue.GetTriggerInput) (*glue.GetTriggerOutput, error)
	GetTriggerAsync(ctx workflow.Context, input *glue.GetTriggerInput) *GetTriggerFuture

	GetTriggers(ctx workflow.Context, input *glue.GetTriggersInput) (*glue.GetTriggersOutput, error)
	GetTriggersAsync(ctx workflow.Context, input *glue.GetTriggersInput) *GetTriggersFuture

	GetUserDefinedFunction(ctx workflow.Context, input *glue.GetUserDefinedFunctionInput) (*glue.GetUserDefinedFunctionOutput, error)
	GetUserDefinedFunctionAsync(ctx workflow.Context, input *glue.GetUserDefinedFunctionInput) *GetUserDefinedFunctionFuture

	GetUserDefinedFunctions(ctx workflow.Context, input *glue.GetUserDefinedFunctionsInput) (*glue.GetUserDefinedFunctionsOutput, error)
	GetUserDefinedFunctionsAsync(ctx workflow.Context, input *glue.GetUserDefinedFunctionsInput) *GetUserDefinedFunctionsFuture

	GetWorkflow(ctx workflow.Context, input *glue.GetWorkflowInput) (*glue.GetWorkflowOutput, error)
	GetWorkflowAsync(ctx workflow.Context, input *glue.GetWorkflowInput) *GetWorkflowFuture

	GetWorkflowRun(ctx workflow.Context, input *glue.GetWorkflowRunInput) (*glue.GetWorkflowRunOutput, error)
	GetWorkflowRunAsync(ctx workflow.Context, input *glue.GetWorkflowRunInput) *GetWorkflowRunFuture

	GetWorkflowRunProperties(ctx workflow.Context, input *glue.GetWorkflowRunPropertiesInput) (*glue.GetWorkflowRunPropertiesOutput, error)
	GetWorkflowRunPropertiesAsync(ctx workflow.Context, input *glue.GetWorkflowRunPropertiesInput) *GetWorkflowRunPropertiesFuture

	GetWorkflowRuns(ctx workflow.Context, input *glue.GetWorkflowRunsInput) (*glue.GetWorkflowRunsOutput, error)
	GetWorkflowRunsAsync(ctx workflow.Context, input *glue.GetWorkflowRunsInput) *GetWorkflowRunsFuture

	ImportCatalogToGlue(ctx workflow.Context, input *glue.ImportCatalogToGlueInput) (*glue.ImportCatalogToGlueOutput, error)
	ImportCatalogToGlueAsync(ctx workflow.Context, input *glue.ImportCatalogToGlueInput) *ImportCatalogToGlueFuture

	ListCrawlers(ctx workflow.Context, input *glue.ListCrawlersInput) (*glue.ListCrawlersOutput, error)
	ListCrawlersAsync(ctx workflow.Context, input *glue.ListCrawlersInput) *ListCrawlersFuture

	ListDevEndpoints(ctx workflow.Context, input *glue.ListDevEndpointsInput) (*glue.ListDevEndpointsOutput, error)
	ListDevEndpointsAsync(ctx workflow.Context, input *glue.ListDevEndpointsInput) *ListDevEndpointsFuture

	ListJobs(ctx workflow.Context, input *glue.ListJobsInput) (*glue.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *glue.ListJobsInput) *ListJobsFuture

	ListMLTransforms(ctx workflow.Context, input *glue.ListMLTransformsInput) (*glue.ListMLTransformsOutput, error)
	ListMLTransformsAsync(ctx workflow.Context, input *glue.ListMLTransformsInput) *ListMLTransformsFuture

	ListTriggers(ctx workflow.Context, input *glue.ListTriggersInput) (*glue.ListTriggersOutput, error)
	ListTriggersAsync(ctx workflow.Context, input *glue.ListTriggersInput) *ListTriggersFuture

	ListWorkflows(ctx workflow.Context, input *glue.ListWorkflowsInput) (*glue.ListWorkflowsOutput, error)
	ListWorkflowsAsync(ctx workflow.Context, input *glue.ListWorkflowsInput) *ListWorkflowsFuture

	PutDataCatalogEncryptionSettings(ctx workflow.Context, input *glue.PutDataCatalogEncryptionSettingsInput) (*glue.PutDataCatalogEncryptionSettingsOutput, error)
	PutDataCatalogEncryptionSettingsAsync(ctx workflow.Context, input *glue.PutDataCatalogEncryptionSettingsInput) *PutDataCatalogEncryptionSettingsFuture

	PutResourcePolicy(ctx workflow.Context, input *glue.PutResourcePolicyInput) (*glue.PutResourcePolicyOutput, error)
	PutResourcePolicyAsync(ctx workflow.Context, input *glue.PutResourcePolicyInput) *PutResourcePolicyFuture

	PutWorkflowRunProperties(ctx workflow.Context, input *glue.PutWorkflowRunPropertiesInput) (*glue.PutWorkflowRunPropertiesOutput, error)
	PutWorkflowRunPropertiesAsync(ctx workflow.Context, input *glue.PutWorkflowRunPropertiesInput) *PutWorkflowRunPropertiesFuture

	ResetJobBookmark(ctx workflow.Context, input *glue.ResetJobBookmarkInput) (*glue.ResetJobBookmarkOutput, error)
	ResetJobBookmarkAsync(ctx workflow.Context, input *glue.ResetJobBookmarkInput) *ResetJobBookmarkFuture

	ResumeWorkflowRun(ctx workflow.Context, input *glue.ResumeWorkflowRunInput) (*glue.ResumeWorkflowRunOutput, error)
	ResumeWorkflowRunAsync(ctx workflow.Context, input *glue.ResumeWorkflowRunInput) *ResumeWorkflowRunFuture

	SearchTables(ctx workflow.Context, input *glue.SearchTablesInput) (*glue.SearchTablesOutput, error)
	SearchTablesAsync(ctx workflow.Context, input *glue.SearchTablesInput) *SearchTablesFuture

	StartCrawler(ctx workflow.Context, input *glue.StartCrawlerInput) (*glue.StartCrawlerOutput, error)
	StartCrawlerAsync(ctx workflow.Context, input *glue.StartCrawlerInput) *StartCrawlerFuture

	StartCrawlerSchedule(ctx workflow.Context, input *glue.StartCrawlerScheduleInput) (*glue.StartCrawlerScheduleOutput, error)
	StartCrawlerScheduleAsync(ctx workflow.Context, input *glue.StartCrawlerScheduleInput) *StartCrawlerScheduleFuture

	StartExportLabelsTaskRun(ctx workflow.Context, input *glue.StartExportLabelsTaskRunInput) (*glue.StartExportLabelsTaskRunOutput, error)
	StartExportLabelsTaskRunAsync(ctx workflow.Context, input *glue.StartExportLabelsTaskRunInput) *StartExportLabelsTaskRunFuture

	StartImportLabelsTaskRun(ctx workflow.Context, input *glue.StartImportLabelsTaskRunInput) (*glue.StartImportLabelsTaskRunOutput, error)
	StartImportLabelsTaskRunAsync(ctx workflow.Context, input *glue.StartImportLabelsTaskRunInput) *StartImportLabelsTaskRunFuture

	StartJobRun(ctx workflow.Context, input *glue.StartJobRunInput) (*glue.StartJobRunOutput, error)
	StartJobRunAsync(ctx workflow.Context, input *glue.StartJobRunInput) *StartJobRunFuture

	StartMLEvaluationTaskRun(ctx workflow.Context, input *glue.StartMLEvaluationTaskRunInput) (*glue.StartMLEvaluationTaskRunOutput, error)
	StartMLEvaluationTaskRunAsync(ctx workflow.Context, input *glue.StartMLEvaluationTaskRunInput) *StartMLEvaluationTaskRunFuture

	StartMLLabelingSetGenerationTaskRun(ctx workflow.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) (*glue.StartMLLabelingSetGenerationTaskRunOutput, error)
	StartMLLabelingSetGenerationTaskRunAsync(ctx workflow.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) *StartMLLabelingSetGenerationTaskRunFuture

	StartTrigger(ctx workflow.Context, input *glue.StartTriggerInput) (*glue.StartTriggerOutput, error)
	StartTriggerAsync(ctx workflow.Context, input *glue.StartTriggerInput) *StartTriggerFuture

	StartWorkflowRun(ctx workflow.Context, input *glue.StartWorkflowRunInput) (*glue.StartWorkflowRunOutput, error)
	StartWorkflowRunAsync(ctx workflow.Context, input *glue.StartWorkflowRunInput) *StartWorkflowRunFuture

	StopCrawler(ctx workflow.Context, input *glue.StopCrawlerInput) (*glue.StopCrawlerOutput, error)
	StopCrawlerAsync(ctx workflow.Context, input *glue.StopCrawlerInput) *StopCrawlerFuture

	StopCrawlerSchedule(ctx workflow.Context, input *glue.StopCrawlerScheduleInput) (*glue.StopCrawlerScheduleOutput, error)
	StopCrawlerScheduleAsync(ctx workflow.Context, input *glue.StopCrawlerScheduleInput) *StopCrawlerScheduleFuture

	StopTrigger(ctx workflow.Context, input *glue.StopTriggerInput) (*glue.StopTriggerOutput, error)
	StopTriggerAsync(ctx workflow.Context, input *glue.StopTriggerInput) *StopTriggerFuture

	StopWorkflowRun(ctx workflow.Context, input *glue.StopWorkflowRunInput) (*glue.StopWorkflowRunOutput, error)
	StopWorkflowRunAsync(ctx workflow.Context, input *glue.StopWorkflowRunInput) *StopWorkflowRunFuture

	TagResource(ctx workflow.Context, input *glue.TagResourceInput) (*glue.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *glue.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *glue.UntagResourceInput) (*glue.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *glue.UntagResourceInput) *UntagResourceFuture

	UpdateClassifier(ctx workflow.Context, input *glue.UpdateClassifierInput) (*glue.UpdateClassifierOutput, error)
	UpdateClassifierAsync(ctx workflow.Context, input *glue.UpdateClassifierInput) *UpdateClassifierFuture

	UpdateColumnStatisticsForPartition(ctx workflow.Context, input *glue.UpdateColumnStatisticsForPartitionInput) (*glue.UpdateColumnStatisticsForPartitionOutput, error)
	UpdateColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.UpdateColumnStatisticsForPartitionInput) *UpdateColumnStatisticsForPartitionFuture

	UpdateColumnStatisticsForTable(ctx workflow.Context, input *glue.UpdateColumnStatisticsForTableInput) (*glue.UpdateColumnStatisticsForTableOutput, error)
	UpdateColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.UpdateColumnStatisticsForTableInput) *UpdateColumnStatisticsForTableFuture

	UpdateConnection(ctx workflow.Context, input *glue.UpdateConnectionInput) (*glue.UpdateConnectionOutput, error)
	UpdateConnectionAsync(ctx workflow.Context, input *glue.UpdateConnectionInput) *UpdateConnectionFuture

	UpdateCrawler(ctx workflow.Context, input *glue.UpdateCrawlerInput) (*glue.UpdateCrawlerOutput, error)
	UpdateCrawlerAsync(ctx workflow.Context, input *glue.UpdateCrawlerInput) *UpdateCrawlerFuture

	UpdateCrawlerSchedule(ctx workflow.Context, input *glue.UpdateCrawlerScheduleInput) (*glue.UpdateCrawlerScheduleOutput, error)
	UpdateCrawlerScheduleAsync(ctx workflow.Context, input *glue.UpdateCrawlerScheduleInput) *UpdateCrawlerScheduleFuture

	UpdateDatabase(ctx workflow.Context, input *glue.UpdateDatabaseInput) (*glue.UpdateDatabaseOutput, error)
	UpdateDatabaseAsync(ctx workflow.Context, input *glue.UpdateDatabaseInput) *UpdateDatabaseFuture

	UpdateDevEndpoint(ctx workflow.Context, input *glue.UpdateDevEndpointInput) (*glue.UpdateDevEndpointOutput, error)
	UpdateDevEndpointAsync(ctx workflow.Context, input *glue.UpdateDevEndpointInput) *UpdateDevEndpointFuture

	UpdateJob(ctx workflow.Context, input *glue.UpdateJobInput) (*glue.UpdateJobOutput, error)
	UpdateJobAsync(ctx workflow.Context, input *glue.UpdateJobInput) *UpdateJobFuture

	UpdateMLTransform(ctx workflow.Context, input *glue.UpdateMLTransformInput) (*glue.UpdateMLTransformOutput, error)
	UpdateMLTransformAsync(ctx workflow.Context, input *glue.UpdateMLTransformInput) *UpdateMLTransformFuture

	UpdatePartition(ctx workflow.Context, input *glue.UpdatePartitionInput) (*glue.UpdatePartitionOutput, error)
	UpdatePartitionAsync(ctx workflow.Context, input *glue.UpdatePartitionInput) *UpdatePartitionFuture

	UpdateTable(ctx workflow.Context, input *glue.UpdateTableInput) (*glue.UpdateTableOutput, error)
	UpdateTableAsync(ctx workflow.Context, input *glue.UpdateTableInput) *UpdateTableFuture

	UpdateTrigger(ctx workflow.Context, input *glue.UpdateTriggerInput) (*glue.UpdateTriggerOutput, error)
	UpdateTriggerAsync(ctx workflow.Context, input *glue.UpdateTriggerInput) *UpdateTriggerFuture

	UpdateUserDefinedFunction(ctx workflow.Context, input *glue.UpdateUserDefinedFunctionInput) (*glue.UpdateUserDefinedFunctionOutput, error)
	UpdateUserDefinedFunctionAsync(ctx workflow.Context, input *glue.UpdateUserDefinedFunctionInput) *UpdateUserDefinedFunctionFuture

	UpdateWorkflow(ctx workflow.Context, input *glue.UpdateWorkflowInput) (*glue.UpdateWorkflowOutput, error)
	UpdateWorkflowAsync(ctx workflow.Context, input *glue.UpdateWorkflowInput) *UpdateWorkflowFuture
}

func NewClient() Client {
	return &stub{}
}
