// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package configservicestub

import (
	"github.com/aws/aws-sdk-go/service/configservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchGetAggregateResourceConfig(ctx workflow.Context, input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error)
	BatchGetAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.BatchGetAggregateResourceConfigInput) *BatchGetAggregateResourceConfigFuture

	BatchGetResourceConfig(ctx workflow.Context, input *configservice.BatchGetResourceConfigInput) (*configservice.BatchGetResourceConfigOutput, error)
	BatchGetResourceConfigAsync(ctx workflow.Context, input *configservice.BatchGetResourceConfigInput) *BatchGetResourceConfigFuture

	DeleteAggregationAuthorization(ctx workflow.Context, input *configservice.DeleteAggregationAuthorizationInput) (*configservice.DeleteAggregationAuthorizationOutput, error)
	DeleteAggregationAuthorizationAsync(ctx workflow.Context, input *configservice.DeleteAggregationAuthorizationInput) *DeleteAggregationAuthorizationFuture

	DeleteConfigRule(ctx workflow.Context, input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error)
	DeleteConfigRuleAsync(ctx workflow.Context, input *configservice.DeleteConfigRuleInput) *DeleteConfigRuleFuture

	DeleteConfigurationAggregator(ctx workflow.Context, input *configservice.DeleteConfigurationAggregatorInput) (*configservice.DeleteConfigurationAggregatorOutput, error)
	DeleteConfigurationAggregatorAsync(ctx workflow.Context, input *configservice.DeleteConfigurationAggregatorInput) *DeleteConfigurationAggregatorFuture

	DeleteConfigurationRecorder(ctx workflow.Context, input *configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error)
	DeleteConfigurationRecorderAsync(ctx workflow.Context, input *configservice.DeleteConfigurationRecorderInput) *DeleteConfigurationRecorderFuture

	DeleteConformancePack(ctx workflow.Context, input *configservice.DeleteConformancePackInput) (*configservice.DeleteConformancePackOutput, error)
	DeleteConformancePackAsync(ctx workflow.Context, input *configservice.DeleteConformancePackInput) *DeleteConformancePackFuture

	DeleteDeliveryChannel(ctx workflow.Context, input *configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error)
	DeleteDeliveryChannelAsync(ctx workflow.Context, input *configservice.DeleteDeliveryChannelInput) *DeleteDeliveryChannelFuture

	DeleteEvaluationResults(ctx workflow.Context, input *configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error)
	DeleteEvaluationResultsAsync(ctx workflow.Context, input *configservice.DeleteEvaluationResultsInput) *DeleteEvaluationResultsFuture

	DeleteOrganizationConfigRule(ctx workflow.Context, input *configservice.DeleteOrganizationConfigRuleInput) (*configservice.DeleteOrganizationConfigRuleOutput, error)
	DeleteOrganizationConfigRuleAsync(ctx workflow.Context, input *configservice.DeleteOrganizationConfigRuleInput) *DeleteOrganizationConfigRuleFuture

	DeleteOrganizationConformancePack(ctx workflow.Context, input *configservice.DeleteOrganizationConformancePackInput) (*configservice.DeleteOrganizationConformancePackOutput, error)
	DeleteOrganizationConformancePackAsync(ctx workflow.Context, input *configservice.DeleteOrganizationConformancePackInput) *DeleteOrganizationConformancePackFuture

	DeleteRemediationConfiguration(ctx workflow.Context, input *configservice.DeleteRemediationConfigurationInput) (*configservice.DeleteRemediationConfigurationOutput, error)
	DeleteRemediationConfigurationAsync(ctx workflow.Context, input *configservice.DeleteRemediationConfigurationInput) *DeleteRemediationConfigurationFuture

	DeleteRemediationExceptions(ctx workflow.Context, input *configservice.DeleteRemediationExceptionsInput) (*configservice.DeleteRemediationExceptionsOutput, error)
	DeleteRemediationExceptionsAsync(ctx workflow.Context, input *configservice.DeleteRemediationExceptionsInput) *DeleteRemediationExceptionsFuture

	DeleteResourceConfig(ctx workflow.Context, input *configservice.DeleteResourceConfigInput) (*configservice.DeleteResourceConfigOutput, error)
	DeleteResourceConfigAsync(ctx workflow.Context, input *configservice.DeleteResourceConfigInput) *DeleteResourceConfigFuture

	DeleteRetentionConfiguration(ctx workflow.Context, input *configservice.DeleteRetentionConfigurationInput) (*configservice.DeleteRetentionConfigurationOutput, error)
	DeleteRetentionConfigurationAsync(ctx workflow.Context, input *configservice.DeleteRetentionConfigurationInput) *DeleteRetentionConfigurationFuture

	DeleteStoredQuery(ctx workflow.Context, input *configservice.DeleteStoredQueryInput) (*configservice.DeleteStoredQueryOutput, error)
	DeleteStoredQueryAsync(ctx workflow.Context, input *configservice.DeleteStoredQueryInput) *DeleteStoredQueryFuture

	DeliverConfigSnapshot(ctx workflow.Context, input *configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error)
	DeliverConfigSnapshotAsync(ctx workflow.Context, input *configservice.DeliverConfigSnapshotInput) *DeliverConfigSnapshotFuture

	DescribeAggregateComplianceByConfigRules(ctx workflow.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error)
	DescribeAggregateComplianceByConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) *DescribeAggregateComplianceByConfigRulesFuture

	DescribeAggregationAuthorizations(ctx workflow.Context, input *configservice.DescribeAggregationAuthorizationsInput) (*configservice.DescribeAggregationAuthorizationsOutput, error)
	DescribeAggregationAuthorizationsAsync(ctx workflow.Context, input *configservice.DescribeAggregationAuthorizationsInput) *DescribeAggregationAuthorizationsFuture

	DescribeComplianceByConfigRule(ctx workflow.Context, input *configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error)
	DescribeComplianceByConfigRuleAsync(ctx workflow.Context, input *configservice.DescribeComplianceByConfigRuleInput) *DescribeComplianceByConfigRuleFuture

	DescribeComplianceByResource(ctx workflow.Context, input *configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error)
	DescribeComplianceByResourceAsync(ctx workflow.Context, input *configservice.DescribeComplianceByResourceInput) *DescribeComplianceByResourceFuture

	DescribeConfigRuleEvaluationStatus(ctx workflow.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
	DescribeConfigRuleEvaluationStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) *DescribeConfigRuleEvaluationStatusFuture

	DescribeConfigRules(ctx workflow.Context, input *configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error)
	DescribeConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeConfigRulesInput) *DescribeConfigRulesFuture

	DescribeConfigurationAggregatorSourcesStatus(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error)
	DescribeConfigurationAggregatorSourcesStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) *DescribeConfigurationAggregatorSourcesStatusFuture

	DescribeConfigurationAggregators(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorsInput) (*configservice.DescribeConfigurationAggregatorsOutput, error)
	DescribeConfigurationAggregatorsAsync(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorsInput) *DescribeConfigurationAggregatorsFuture

	DescribeConfigurationRecorderStatus(ctx workflow.Context, input *configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	DescribeConfigurationRecorderStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigurationRecorderStatusInput) *DescribeConfigurationRecorderStatusFuture

	DescribeConfigurationRecorders(ctx workflow.Context, input *configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConfigurationRecordersAsync(ctx workflow.Context, input *configservice.DescribeConfigurationRecordersInput) *DescribeConfigurationRecordersFuture

	DescribeConformancePackCompliance(ctx workflow.Context, input *configservice.DescribeConformancePackComplianceInput) (*configservice.DescribeConformancePackComplianceOutput, error)
	DescribeConformancePackComplianceAsync(ctx workflow.Context, input *configservice.DescribeConformancePackComplianceInput) *DescribeConformancePackComplianceFuture

	DescribeConformancePackStatus(ctx workflow.Context, input *configservice.DescribeConformancePackStatusInput) (*configservice.DescribeConformancePackStatusOutput, error)
	DescribeConformancePackStatusAsync(ctx workflow.Context, input *configservice.DescribeConformancePackStatusInput) *DescribeConformancePackStatusFuture

	DescribeConformancePacks(ctx workflow.Context, input *configservice.DescribeConformancePacksInput) (*configservice.DescribeConformancePacksOutput, error)
	DescribeConformancePacksAsync(ctx workflow.Context, input *configservice.DescribeConformancePacksInput) *DescribeConformancePacksFuture

	DescribeDeliveryChannelStatus(ctx workflow.Context, input *configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error)
	DescribeDeliveryChannelStatusAsync(ctx workflow.Context, input *configservice.DescribeDeliveryChannelStatusInput) *DescribeDeliveryChannelStatusFuture

	DescribeDeliveryChannels(ctx workflow.Context, input *configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error)
	DescribeDeliveryChannelsAsync(ctx workflow.Context, input *configservice.DescribeDeliveryChannelsInput) *DescribeDeliveryChannelsFuture

	DescribeOrganizationConfigRuleStatuses(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error)
	DescribeOrganizationConfigRuleStatusesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) *DescribeOrganizationConfigRuleStatusesFuture

	DescribeOrganizationConfigRules(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRulesInput) (*configservice.DescribeOrganizationConfigRulesOutput, error)
	DescribeOrganizationConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRulesInput) *DescribeOrganizationConfigRulesFuture

	DescribeOrganizationConformancePackStatuses(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error)
	DescribeOrganizationConformancePackStatusesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) *DescribeOrganizationConformancePackStatusesFuture

	DescribeOrganizationConformancePacks(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePacksInput) (*configservice.DescribeOrganizationConformancePacksOutput, error)
	DescribeOrganizationConformancePacksAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePacksInput) *DescribeOrganizationConformancePacksFuture

	DescribePendingAggregationRequests(ctx workflow.Context, input *configservice.DescribePendingAggregationRequestsInput) (*configservice.DescribePendingAggregationRequestsOutput, error)
	DescribePendingAggregationRequestsAsync(ctx workflow.Context, input *configservice.DescribePendingAggregationRequestsInput) *DescribePendingAggregationRequestsFuture

	DescribeRemediationConfigurations(ctx workflow.Context, input *configservice.DescribeRemediationConfigurationsInput) (*configservice.DescribeRemediationConfigurationsOutput, error)
	DescribeRemediationConfigurationsAsync(ctx workflow.Context, input *configservice.DescribeRemediationConfigurationsInput) *DescribeRemediationConfigurationsFuture

	DescribeRemediationExceptions(ctx workflow.Context, input *configservice.DescribeRemediationExceptionsInput) (*configservice.DescribeRemediationExceptionsOutput, error)
	DescribeRemediationExceptionsAsync(ctx workflow.Context, input *configservice.DescribeRemediationExceptionsInput) *DescribeRemediationExceptionsFuture

	DescribeRemediationExecutionStatus(ctx workflow.Context, input *configservice.DescribeRemediationExecutionStatusInput) (*configservice.DescribeRemediationExecutionStatusOutput, error)
	DescribeRemediationExecutionStatusAsync(ctx workflow.Context, input *configservice.DescribeRemediationExecutionStatusInput) *DescribeRemediationExecutionStatusFuture

	DescribeRetentionConfigurations(ctx workflow.Context, input *configservice.DescribeRetentionConfigurationsInput) (*configservice.DescribeRetentionConfigurationsOutput, error)
	DescribeRetentionConfigurationsAsync(ctx workflow.Context, input *configservice.DescribeRetentionConfigurationsInput) *DescribeRetentionConfigurationsFuture

	GetAggregateComplianceDetailsByConfigRule(ctx workflow.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error)
	GetAggregateComplianceDetailsByConfigRuleAsync(ctx workflow.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) *GetAggregateComplianceDetailsByConfigRuleFuture

	GetAggregateConfigRuleComplianceSummary(ctx workflow.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error)
	GetAggregateConfigRuleComplianceSummaryAsync(ctx workflow.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) *GetAggregateConfigRuleComplianceSummaryFuture

	GetAggregateDiscoveredResourceCounts(ctx workflow.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error)
	GetAggregateDiscoveredResourceCountsAsync(ctx workflow.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) *GetAggregateDiscoveredResourceCountsFuture

	GetAggregateResourceConfig(ctx workflow.Context, input *configservice.GetAggregateResourceConfigInput) (*configservice.GetAggregateResourceConfigOutput, error)
	GetAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.GetAggregateResourceConfigInput) *GetAggregateResourceConfigFuture

	GetComplianceDetailsByConfigRule(ctx workflow.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
	GetComplianceDetailsByConfigRuleAsync(ctx workflow.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) *GetComplianceDetailsByConfigRuleFuture

	GetComplianceDetailsByResource(ctx workflow.Context, input *configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error)
	GetComplianceDetailsByResourceAsync(ctx workflow.Context, input *configservice.GetComplianceDetailsByResourceInput) *GetComplianceDetailsByResourceFuture

	GetComplianceSummaryByConfigRule(ctx workflow.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
	GetComplianceSummaryByConfigRuleAsync(ctx workflow.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) *GetComplianceSummaryByConfigRuleFuture

	GetComplianceSummaryByResourceType(ctx workflow.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
	GetComplianceSummaryByResourceTypeAsync(ctx workflow.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) *GetComplianceSummaryByResourceTypeFuture

	GetConformancePackComplianceDetails(ctx workflow.Context, input *configservice.GetConformancePackComplianceDetailsInput) (*configservice.GetConformancePackComplianceDetailsOutput, error)
	GetConformancePackComplianceDetailsAsync(ctx workflow.Context, input *configservice.GetConformancePackComplianceDetailsInput) *GetConformancePackComplianceDetailsFuture

	GetConformancePackComplianceSummary(ctx workflow.Context, input *configservice.GetConformancePackComplianceSummaryInput) (*configservice.GetConformancePackComplianceSummaryOutput, error)
	GetConformancePackComplianceSummaryAsync(ctx workflow.Context, input *configservice.GetConformancePackComplianceSummaryInput) *GetConformancePackComplianceSummaryFuture

	GetDiscoveredResourceCounts(ctx workflow.Context, input *configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error)
	GetDiscoveredResourceCountsAsync(ctx workflow.Context, input *configservice.GetDiscoveredResourceCountsInput) *GetDiscoveredResourceCountsFuture

	GetOrganizationConfigRuleDetailedStatus(ctx workflow.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error)
	GetOrganizationConfigRuleDetailedStatusAsync(ctx workflow.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) *GetOrganizationConfigRuleDetailedStatusFuture

	GetOrganizationConformancePackDetailedStatus(ctx workflow.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error)
	GetOrganizationConformancePackDetailedStatusAsync(ctx workflow.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) *GetOrganizationConformancePackDetailedStatusFuture

	GetResourceConfigHistory(ctx workflow.Context, input *configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error)
	GetResourceConfigHistoryAsync(ctx workflow.Context, input *configservice.GetResourceConfigHistoryInput) *GetResourceConfigHistoryFuture

	GetStoredQuery(ctx workflow.Context, input *configservice.GetStoredQueryInput) (*configservice.GetStoredQueryOutput, error)
	GetStoredQueryAsync(ctx workflow.Context, input *configservice.GetStoredQueryInput) *GetStoredQueryFuture

	ListAggregateDiscoveredResources(ctx workflow.Context, input *configservice.ListAggregateDiscoveredResourcesInput) (*configservice.ListAggregateDiscoveredResourcesOutput, error)
	ListAggregateDiscoveredResourcesAsync(ctx workflow.Context, input *configservice.ListAggregateDiscoveredResourcesInput) *ListAggregateDiscoveredResourcesFuture

	ListDiscoveredResources(ctx workflow.Context, input *configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesAsync(ctx workflow.Context, input *configservice.ListDiscoveredResourcesInput) *ListDiscoveredResourcesFuture

	ListStoredQueries(ctx workflow.Context, input *configservice.ListStoredQueriesInput) (*configservice.ListStoredQueriesOutput, error)
	ListStoredQueriesAsync(ctx workflow.Context, input *configservice.ListStoredQueriesInput) *ListStoredQueriesFuture

	ListTagsForResource(ctx workflow.Context, input *configservice.ListTagsForResourceInput) (*configservice.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *configservice.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutAggregationAuthorization(ctx workflow.Context, input *configservice.PutAggregationAuthorizationInput) (*configservice.PutAggregationAuthorizationOutput, error)
	PutAggregationAuthorizationAsync(ctx workflow.Context, input *configservice.PutAggregationAuthorizationInput) *PutAggregationAuthorizationFuture

	PutConfigRule(ctx workflow.Context, input *configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error)
	PutConfigRuleAsync(ctx workflow.Context, input *configservice.PutConfigRuleInput) *PutConfigRuleFuture

	PutConfigurationAggregator(ctx workflow.Context, input *configservice.PutConfigurationAggregatorInput) (*configservice.PutConfigurationAggregatorOutput, error)
	PutConfigurationAggregatorAsync(ctx workflow.Context, input *configservice.PutConfigurationAggregatorInput) *PutConfigurationAggregatorFuture

	PutConfigurationRecorder(ctx workflow.Context, input *configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error)
	PutConfigurationRecorderAsync(ctx workflow.Context, input *configservice.PutConfigurationRecorderInput) *PutConfigurationRecorderFuture

	PutConformancePack(ctx workflow.Context, input *configservice.PutConformancePackInput) (*configservice.PutConformancePackOutput, error)
	PutConformancePackAsync(ctx workflow.Context, input *configservice.PutConformancePackInput) *PutConformancePackFuture

	PutDeliveryChannel(ctx workflow.Context, input *configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error)
	PutDeliveryChannelAsync(ctx workflow.Context, input *configservice.PutDeliveryChannelInput) *PutDeliveryChannelFuture

	PutEvaluations(ctx workflow.Context, input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)
	PutEvaluationsAsync(ctx workflow.Context, input *configservice.PutEvaluationsInput) *PutEvaluationsFuture

	PutExternalEvaluation(ctx workflow.Context, input *configservice.PutExternalEvaluationInput) (*configservice.PutExternalEvaluationOutput, error)
	PutExternalEvaluationAsync(ctx workflow.Context, input *configservice.PutExternalEvaluationInput) *PutExternalEvaluationFuture

	PutOrganizationConfigRule(ctx workflow.Context, input *configservice.PutOrganizationConfigRuleInput) (*configservice.PutOrganizationConfigRuleOutput, error)
	PutOrganizationConfigRuleAsync(ctx workflow.Context, input *configservice.PutOrganizationConfigRuleInput) *PutOrganizationConfigRuleFuture

	PutOrganizationConformancePack(ctx workflow.Context, input *configservice.PutOrganizationConformancePackInput) (*configservice.PutOrganizationConformancePackOutput, error)
	PutOrganizationConformancePackAsync(ctx workflow.Context, input *configservice.PutOrganizationConformancePackInput) *PutOrganizationConformancePackFuture

	PutRemediationConfigurations(ctx workflow.Context, input *configservice.PutRemediationConfigurationsInput) (*configservice.PutRemediationConfigurationsOutput, error)
	PutRemediationConfigurationsAsync(ctx workflow.Context, input *configservice.PutRemediationConfigurationsInput) *PutRemediationConfigurationsFuture

	PutRemediationExceptions(ctx workflow.Context, input *configservice.PutRemediationExceptionsInput) (*configservice.PutRemediationExceptionsOutput, error)
	PutRemediationExceptionsAsync(ctx workflow.Context, input *configservice.PutRemediationExceptionsInput) *PutRemediationExceptionsFuture

	PutResourceConfig(ctx workflow.Context, input *configservice.PutResourceConfigInput) (*configservice.PutResourceConfigOutput, error)
	PutResourceConfigAsync(ctx workflow.Context, input *configservice.PutResourceConfigInput) *PutResourceConfigFuture

	PutRetentionConfiguration(ctx workflow.Context, input *configservice.PutRetentionConfigurationInput) (*configservice.PutRetentionConfigurationOutput, error)
	PutRetentionConfigurationAsync(ctx workflow.Context, input *configservice.PutRetentionConfigurationInput) *PutRetentionConfigurationFuture

	PutStoredQuery(ctx workflow.Context, input *configservice.PutStoredQueryInput) (*configservice.PutStoredQueryOutput, error)
	PutStoredQueryAsync(ctx workflow.Context, input *configservice.PutStoredQueryInput) *PutStoredQueryFuture

	SelectAggregateResourceConfig(ctx workflow.Context, input *configservice.SelectAggregateResourceConfigInput) (*configservice.SelectAggregateResourceConfigOutput, error)
	SelectAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.SelectAggregateResourceConfigInput) *SelectAggregateResourceConfigFuture

	SelectResourceConfig(ctx workflow.Context, input *configservice.SelectResourceConfigInput) (*configservice.SelectResourceConfigOutput, error)
	SelectResourceConfigAsync(ctx workflow.Context, input *configservice.SelectResourceConfigInput) *SelectResourceConfigFuture

	StartConfigRulesEvaluation(ctx workflow.Context, input *configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error)
	StartConfigRulesEvaluationAsync(ctx workflow.Context, input *configservice.StartConfigRulesEvaluationInput) *StartConfigRulesEvaluationFuture

	StartConfigurationRecorder(ctx workflow.Context, input *configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error)
	StartConfigurationRecorderAsync(ctx workflow.Context, input *configservice.StartConfigurationRecorderInput) *StartConfigurationRecorderFuture

	StartRemediationExecution(ctx workflow.Context, input *configservice.StartRemediationExecutionInput) (*configservice.StartRemediationExecutionOutput, error)
	StartRemediationExecutionAsync(ctx workflow.Context, input *configservice.StartRemediationExecutionInput) *StartRemediationExecutionFuture

	StopConfigurationRecorder(ctx workflow.Context, input *configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error)
	StopConfigurationRecorderAsync(ctx workflow.Context, input *configservice.StopConfigurationRecorderInput) *StopConfigurationRecorderFuture

	TagResource(ctx workflow.Context, input *configservice.TagResourceInput) (*configservice.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *configservice.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *configservice.UntagResourceInput) (*configservice.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *configservice.UntagResourceInput) *UntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
