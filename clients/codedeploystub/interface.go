// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package codedeploystub

import (
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTagsToOnPremisesInstances(ctx workflow.Context, input *codedeploy.AddTagsToOnPremisesInstancesInput) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error)
	AddTagsToOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.AddTagsToOnPremisesInstancesInput) *AddTagsToOnPremisesInstancesFuture

	BatchGetApplicationRevisions(ctx workflow.Context, input *codedeploy.BatchGetApplicationRevisionsInput) (*codedeploy.BatchGetApplicationRevisionsOutput, error)
	BatchGetApplicationRevisionsAsync(ctx workflow.Context, input *codedeploy.BatchGetApplicationRevisionsInput) *BatchGetApplicationRevisionsFuture

	BatchGetApplications(ctx workflow.Context, input *codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error)
	BatchGetApplicationsAsync(ctx workflow.Context, input *codedeploy.BatchGetApplicationsInput) *BatchGetApplicationsFuture

	BatchGetDeploymentGroups(ctx workflow.Context, input *codedeploy.BatchGetDeploymentGroupsInput) (*codedeploy.BatchGetDeploymentGroupsOutput, error)
	BatchGetDeploymentGroupsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentGroupsInput) *BatchGetDeploymentGroupsFuture

	BatchGetDeploymentInstances(ctx workflow.Context, input *codedeploy.BatchGetDeploymentInstancesInput) (*codedeploy.BatchGetDeploymentInstancesOutput, error)
	BatchGetDeploymentInstancesAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentInstancesInput) *BatchGetDeploymentInstancesFuture

	BatchGetDeploymentTargets(ctx workflow.Context, input *codedeploy.BatchGetDeploymentTargetsInput) (*codedeploy.BatchGetDeploymentTargetsOutput, error)
	BatchGetDeploymentTargetsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentTargetsInput) *BatchGetDeploymentTargetsFuture

	BatchGetDeployments(ctx workflow.Context, input *codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error)
	BatchGetDeploymentsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentsInput) *BatchGetDeploymentsFuture

	BatchGetOnPremisesInstances(ctx workflow.Context, input *codedeploy.BatchGetOnPremisesInstancesInput) (*codedeploy.BatchGetOnPremisesInstancesOutput, error)
	BatchGetOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.BatchGetOnPremisesInstancesInput) *BatchGetOnPremisesInstancesFuture

	ContinueDeployment(ctx workflow.Context, input *codedeploy.ContinueDeploymentInput) (*codedeploy.ContinueDeploymentOutput, error)
	ContinueDeploymentAsync(ctx workflow.Context, input *codedeploy.ContinueDeploymentInput) *ContinueDeploymentFuture

	CreateApplication(ctx workflow.Context, input *codedeploy.CreateApplicationInput) (*codedeploy.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *codedeploy.CreateApplicationInput) *CreateApplicationFuture

	CreateDeployment(ctx workflow.Context, input *codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error)
	CreateDeploymentAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentInput) *CreateDeploymentFuture

	CreateDeploymentConfig(ctx workflow.Context, input *codedeploy.CreateDeploymentConfigInput) (*codedeploy.CreateDeploymentConfigOutput, error)
	CreateDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentConfigInput) *CreateDeploymentConfigFuture

	CreateDeploymentGroup(ctx workflow.Context, input *codedeploy.CreateDeploymentGroupInput) (*codedeploy.CreateDeploymentGroupOutput, error)
	CreateDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentGroupInput) *CreateDeploymentGroupFuture

	DeleteApplication(ctx workflow.Context, input *codedeploy.DeleteApplicationInput) (*codedeploy.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *codedeploy.DeleteApplicationInput) *DeleteApplicationFuture

	DeleteDeploymentConfig(ctx workflow.Context, input *codedeploy.DeleteDeploymentConfigInput) (*codedeploy.DeleteDeploymentConfigOutput, error)
	DeleteDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.DeleteDeploymentConfigInput) *DeleteDeploymentConfigFuture

	DeleteDeploymentGroup(ctx workflow.Context, input *codedeploy.DeleteDeploymentGroupInput) (*codedeploy.DeleteDeploymentGroupOutput, error)
	DeleteDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.DeleteDeploymentGroupInput) *DeleteDeploymentGroupFuture

	DeleteGitHubAccountToken(ctx workflow.Context, input *codedeploy.DeleteGitHubAccountTokenInput) (*codedeploy.DeleteGitHubAccountTokenOutput, error)
	DeleteGitHubAccountTokenAsync(ctx workflow.Context, input *codedeploy.DeleteGitHubAccountTokenInput) *DeleteGitHubAccountTokenFuture

	DeleteResourcesByExternalId(ctx workflow.Context, input *codedeploy.DeleteResourcesByExternalIdInput) (*codedeploy.DeleteResourcesByExternalIdOutput, error)
	DeleteResourcesByExternalIdAsync(ctx workflow.Context, input *codedeploy.DeleteResourcesByExternalIdInput) *DeleteResourcesByExternalIdFuture

	DeregisterOnPremisesInstance(ctx workflow.Context, input *codedeploy.DeregisterOnPremisesInstanceInput) (*codedeploy.DeregisterOnPremisesInstanceOutput, error)
	DeregisterOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.DeregisterOnPremisesInstanceInput) *DeregisterOnPremisesInstanceFuture

	GetApplication(ctx workflow.Context, input *codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error)
	GetApplicationAsync(ctx workflow.Context, input *codedeploy.GetApplicationInput) *GetApplicationFuture

	GetApplicationRevision(ctx workflow.Context, input *codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error)
	GetApplicationRevisionAsync(ctx workflow.Context, input *codedeploy.GetApplicationRevisionInput) *GetApplicationRevisionFuture

	GetDeployment(ctx workflow.Context, input *codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error)
	GetDeploymentAsync(ctx workflow.Context, input *codedeploy.GetDeploymentInput) *GetDeploymentFuture

	GetDeploymentConfig(ctx workflow.Context, input *codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error)
	GetDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.GetDeploymentConfigInput) *GetDeploymentConfigFuture

	GetDeploymentGroup(ctx workflow.Context, input *codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error)
	GetDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.GetDeploymentGroupInput) *GetDeploymentGroupFuture

	GetDeploymentInstance(ctx workflow.Context, input *codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error)
	GetDeploymentInstanceAsync(ctx workflow.Context, input *codedeploy.GetDeploymentInstanceInput) *GetDeploymentInstanceFuture

	GetDeploymentTarget(ctx workflow.Context, input *codedeploy.GetDeploymentTargetInput) (*codedeploy.GetDeploymentTargetOutput, error)
	GetDeploymentTargetAsync(ctx workflow.Context, input *codedeploy.GetDeploymentTargetInput) *GetDeploymentTargetFuture

	GetOnPremisesInstance(ctx workflow.Context, input *codedeploy.GetOnPremisesInstanceInput) (*codedeploy.GetOnPremisesInstanceOutput, error)
	GetOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.GetOnPremisesInstanceInput) *GetOnPremisesInstanceFuture

	ListApplicationRevisions(ctx workflow.Context, input *codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error)
	ListApplicationRevisionsAsync(ctx workflow.Context, input *codedeploy.ListApplicationRevisionsInput) *ListApplicationRevisionsFuture

	ListApplications(ctx workflow.Context, input *codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *codedeploy.ListApplicationsInput) *ListApplicationsFuture

	ListDeploymentConfigs(ctx workflow.Context, input *codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error)
	ListDeploymentConfigsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentConfigsInput) *ListDeploymentConfigsFuture

	ListDeploymentGroups(ctx workflow.Context, input *codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error)
	ListDeploymentGroupsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentGroupsInput) *ListDeploymentGroupsFuture

	ListDeploymentInstances(ctx workflow.Context, input *codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error)
	ListDeploymentInstancesAsync(ctx workflow.Context, input *codedeploy.ListDeploymentInstancesInput) *ListDeploymentInstancesFuture

	ListDeploymentTargets(ctx workflow.Context, input *codedeploy.ListDeploymentTargetsInput) (*codedeploy.ListDeploymentTargetsOutput, error)
	ListDeploymentTargetsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentTargetsInput) *ListDeploymentTargetsFuture

	ListDeployments(ctx workflow.Context, input *codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error)
	ListDeploymentsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentsInput) *ListDeploymentsFuture

	ListGitHubAccountTokenNames(ctx workflow.Context, input *codedeploy.ListGitHubAccountTokenNamesInput) (*codedeploy.ListGitHubAccountTokenNamesOutput, error)
	ListGitHubAccountTokenNamesAsync(ctx workflow.Context, input *codedeploy.ListGitHubAccountTokenNamesInput) *ListGitHubAccountTokenNamesFuture

	ListOnPremisesInstances(ctx workflow.Context, input *codedeploy.ListOnPremisesInstancesInput) (*codedeploy.ListOnPremisesInstancesOutput, error)
	ListOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.ListOnPremisesInstancesInput) *ListOnPremisesInstancesFuture

	ListTagsForResource(ctx workflow.Context, input *codedeploy.ListTagsForResourceInput) (*codedeploy.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codedeploy.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutLifecycleEventHookExecutionStatus(ctx workflow.Context, input *codedeploy.PutLifecycleEventHookExecutionStatusInput) (*codedeploy.PutLifecycleEventHookExecutionStatusOutput, error)
	PutLifecycleEventHookExecutionStatusAsync(ctx workflow.Context, input *codedeploy.PutLifecycleEventHookExecutionStatusInput) *PutLifecycleEventHookExecutionStatusFuture

	RegisterApplicationRevision(ctx workflow.Context, input *codedeploy.RegisterApplicationRevisionInput) (*codedeploy.RegisterApplicationRevisionOutput, error)
	RegisterApplicationRevisionAsync(ctx workflow.Context, input *codedeploy.RegisterApplicationRevisionInput) *RegisterApplicationRevisionFuture

	RegisterOnPremisesInstance(ctx workflow.Context, input *codedeploy.RegisterOnPremisesInstanceInput) (*codedeploy.RegisterOnPremisesInstanceOutput, error)
	RegisterOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.RegisterOnPremisesInstanceInput) *RegisterOnPremisesInstanceFuture

	RemoveTagsFromOnPremisesInstances(ctx workflow.Context, input *codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error)
	RemoveTagsFromOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.RemoveTagsFromOnPremisesInstancesInput) *RemoveTagsFromOnPremisesInstancesFuture

	SkipWaitTimeForInstanceTermination(ctx workflow.Context, input *codedeploy.SkipWaitTimeForInstanceTerminationInput) (*codedeploy.SkipWaitTimeForInstanceTerminationOutput, error)
	SkipWaitTimeForInstanceTerminationAsync(ctx workflow.Context, input *codedeploy.SkipWaitTimeForInstanceTerminationInput) *SkipWaitTimeForInstanceTerminationFuture

	StopDeployment(ctx workflow.Context, input *codedeploy.StopDeploymentInput) (*codedeploy.StopDeploymentOutput, error)
	StopDeploymentAsync(ctx workflow.Context, input *codedeploy.StopDeploymentInput) *StopDeploymentFuture

	TagResource(ctx workflow.Context, input *codedeploy.TagResourceInput) (*codedeploy.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codedeploy.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *codedeploy.UntagResourceInput) (*codedeploy.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codedeploy.UntagResourceInput) *UntagResourceFuture

	UpdateApplication(ctx workflow.Context, input *codedeploy.UpdateApplicationInput) (*codedeploy.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *codedeploy.UpdateApplicationInput) *UpdateApplicationFuture

	UpdateDeploymentGroup(ctx workflow.Context, input *codedeploy.UpdateDeploymentGroupInput) (*codedeploy.UpdateDeploymentGroupOutput, error)
	UpdateDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.UpdateDeploymentGroupInput) *UpdateDeploymentGroupFuture

	WaitUntilDeploymentSuccessful(ctx workflow.Context, input *codedeploy.GetDeploymentInput) error
	WaitUntilDeploymentSuccessfulAsync(ctx workflow.Context, input *codedeploy.GetDeploymentInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
