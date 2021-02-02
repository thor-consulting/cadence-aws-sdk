// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package databasemigrationservicestub

import (
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTagsToResource(ctx workflow.Context, input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *databasemigrationservice.AddTagsToResourceInput) *AddTagsToResourceFuture

	ApplyPendingMaintenanceAction(ctx workflow.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error)
	ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) *ApplyPendingMaintenanceActionFuture

	CancelReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error)
	CancelReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) *CancelReplicationTaskAssessmentRunFuture

	CreateEndpoint(ctx workflow.Context, input *databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error)
	CreateEndpointAsync(ctx workflow.Context, input *databasemigrationservice.CreateEndpointInput) *CreateEndpointFuture

	CreateEventSubscription(ctx workflow.Context, input *databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.CreateEventSubscriptionInput) *CreateEventSubscriptionFuture

	CreateReplicationInstance(ctx workflow.Context, input *databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error)
	CreateReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationInstanceInput) *CreateReplicationInstanceFuture

	CreateReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error)
	CreateReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) *CreateReplicationSubnetGroupFuture

	CreateReplicationTask(ctx workflow.Context, input *databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error)
	CreateReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationTaskInput) *CreateReplicationTaskFuture

	DeleteCertificate(ctx workflow.Context, input *databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error)
	DeleteCertificateAsync(ctx workflow.Context, input *databasemigrationservice.DeleteCertificateInput) *DeleteCertificateFuture

	DeleteConnection(ctx workflow.Context, input *databasemigrationservice.DeleteConnectionInput) (*databasemigrationservice.DeleteConnectionOutput, error)
	DeleteConnectionAsync(ctx workflow.Context, input *databasemigrationservice.DeleteConnectionInput) *DeleteConnectionFuture

	DeleteEndpoint(ctx workflow.Context, input *databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *databasemigrationservice.DeleteEndpointInput) *DeleteEndpointFuture

	DeleteEventSubscription(ctx workflow.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) *DeleteEventSubscriptionFuture

	DeleteReplicationInstance(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error)
	DeleteReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) *DeleteReplicationInstanceFuture

	DeleteReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error)
	DeleteReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) *DeleteReplicationSubnetGroupFuture

	DeleteReplicationTask(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error)
	DeleteReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskInput) *DeleteReplicationTaskFuture

	DeleteReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error)
	DeleteReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) *DeleteReplicationTaskAssessmentRunFuture

	DescribeAccountAttributes(ctx workflow.Context, input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeAccountAttributesInput) *DescribeAccountAttributesFuture

	DescribeApplicableIndividualAssessments(ctx workflow.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error)
	DescribeApplicableIndividualAssessmentsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) *DescribeApplicableIndividualAssessmentsFuture

	DescribeCertificates(ctx workflow.Context, input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error)
	DescribeCertificatesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeCertificatesInput) *DescribeCertificatesFuture

	DescribeConnections(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error)
	DescribeConnectionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) *DescribeConnectionsFuture

	DescribeEndpointTypes(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error)
	DescribeEndpointTypesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointTypesInput) *DescribeEndpointTypesFuture

	DescribeEndpoints(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error)
	DescribeEndpointsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) *DescribeEndpointsFuture

	DescribeEventCategories(ctx workflow.Context, input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventCategoriesInput) *DescribeEventCategoriesFuture

	DescribeEventSubscriptions(ctx workflow.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) *DescribeEventSubscriptionsFuture

	DescribeEvents(ctx workflow.Context, input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventsInput) *DescribeEventsFuture

	DescribeOrderableReplicationInstances(ctx workflow.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error)
	DescribeOrderableReplicationInstancesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) *DescribeOrderableReplicationInstancesFuture

	DescribePendingMaintenanceActions(ctx workflow.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error)
	DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) *DescribePendingMaintenanceActionsFuture

	DescribeRefreshSchemasStatus(ctx workflow.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error)
	DescribeRefreshSchemasStatusAsync(ctx workflow.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) *DescribeRefreshSchemasStatusFuture

	DescribeReplicationInstanceTaskLogs(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error)
	DescribeReplicationInstanceTaskLogsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) *DescribeReplicationInstanceTaskLogsFuture

	DescribeReplicationInstances(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
	DescribeReplicationInstancesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) *DescribeReplicationInstancesFuture

	DescribeReplicationSubnetGroups(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error)
	DescribeReplicationSubnetGroupsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) *DescribeReplicationSubnetGroupsFuture

	DescribeReplicationTaskAssessmentResults(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error)
	DescribeReplicationTaskAssessmentResultsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) *DescribeReplicationTaskAssessmentResultsFuture

	DescribeReplicationTaskAssessmentRuns(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error)
	DescribeReplicationTaskAssessmentRunsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) *DescribeReplicationTaskAssessmentRunsFuture

	DescribeReplicationTaskIndividualAssessments(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error)
	DescribeReplicationTaskIndividualAssessmentsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) *DescribeReplicationTaskIndividualAssessmentsFuture

	DescribeReplicationTasks(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error)
	DescribeReplicationTasksAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *DescribeReplicationTasksFuture

	DescribeSchemas(ctx workflow.Context, input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error)
	DescribeSchemasAsync(ctx workflow.Context, input *databasemigrationservice.DescribeSchemasInput) *DescribeSchemasFuture

	DescribeTableStatistics(ctx workflow.Context, input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error)
	DescribeTableStatisticsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeTableStatisticsInput) *DescribeTableStatisticsFuture

	ImportCertificate(ctx workflow.Context, input *databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error)
	ImportCertificateAsync(ctx workflow.Context, input *databasemigrationservice.ImportCertificateInput) *ImportCertificateFuture

	ListTagsForResource(ctx workflow.Context, input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *databasemigrationservice.ListTagsForResourceInput) *ListTagsForResourceFuture

	ModifyEndpoint(ctx workflow.Context, input *databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error)
	ModifyEndpointAsync(ctx workflow.Context, input *databasemigrationservice.ModifyEndpointInput) *ModifyEndpointFuture

	ModifyEventSubscription(ctx workflow.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) *ModifyEventSubscriptionFuture

	ModifyReplicationInstance(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error)
	ModifyReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) *ModifyReplicationInstanceFuture

	ModifyReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error)
	ModifyReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) *ModifyReplicationSubnetGroupFuture

	ModifyReplicationTask(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error)
	ModifyReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationTaskInput) *ModifyReplicationTaskFuture

	MoveReplicationTask(ctx workflow.Context, input *databasemigrationservice.MoveReplicationTaskInput) (*databasemigrationservice.MoveReplicationTaskOutput, error)
	MoveReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.MoveReplicationTaskInput) *MoveReplicationTaskFuture

	RebootReplicationInstance(ctx workflow.Context, input *databasemigrationservice.RebootReplicationInstanceInput) (*databasemigrationservice.RebootReplicationInstanceOutput, error)
	RebootReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.RebootReplicationInstanceInput) *RebootReplicationInstanceFuture

	RefreshSchemas(ctx workflow.Context, input *databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error)
	RefreshSchemasAsync(ctx workflow.Context, input *databasemigrationservice.RefreshSchemasInput) *RefreshSchemasFuture

	ReloadTables(ctx workflow.Context, input *databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error)
	ReloadTablesAsync(ctx workflow.Context, input *databasemigrationservice.ReloadTablesInput) *ReloadTablesFuture

	RemoveTagsFromResource(ctx workflow.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) *RemoveTagsFromResourceFuture

	StartReplicationTask(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error)
	StartReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskInput) *StartReplicationTaskFuture

	StartReplicationTaskAssessment(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error)
	StartReplicationTaskAssessmentAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) *StartReplicationTaskAssessmentFuture

	StartReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error)
	StartReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) *StartReplicationTaskAssessmentRunFuture

	StopReplicationTask(ctx workflow.Context, input *databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error)
	StopReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.StopReplicationTaskInput) *StopReplicationTaskFuture

	TestConnection(ctx workflow.Context, input *databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error)
	TestConnectionAsync(ctx workflow.Context, input *databasemigrationservice.TestConnectionInput) *TestConnectionFuture

	WaitUntilEndpointDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) error
	WaitUntilEndpointDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) *clients.VoidFuture

	WaitUntilReplicationInstanceAvailable(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error
	WaitUntilReplicationInstanceAvailableAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) *clients.VoidFuture

	WaitUntilReplicationInstanceDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error
	WaitUntilReplicationInstanceDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) *clients.VoidFuture

	WaitUntilReplicationTaskDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
	WaitUntilReplicationTaskDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *clients.VoidFuture

	WaitUntilReplicationTaskReady(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
	WaitUntilReplicationTaskReadyAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *clients.VoidFuture

	WaitUntilReplicationTaskRunning(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
	WaitUntilReplicationTaskRunningAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *clients.VoidFuture

	WaitUntilReplicationTaskStopped(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
	WaitUntilReplicationTaskStoppedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *clients.VoidFuture

	WaitUntilTestConnectionSucceeds(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) error
	WaitUntilTestConnectionSucceedsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
