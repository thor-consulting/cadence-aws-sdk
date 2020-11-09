// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package amplifystub

import (
	"github.com/aws/aws-sdk-go/service/amplify"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateAppFuture) Get(ctx workflow.Context) (*amplify.CreateAppOutput, error) {
	var output amplify.CreateAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBackendEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBackendEnvironmentFuture) Get(ctx workflow.Context) (*amplify.CreateBackendEnvironmentOutput, error) {
	var output amplify.CreateBackendEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBranchFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBranchFuture) Get(ctx workflow.Context) (*amplify.CreateBranchOutput, error) {
	var output amplify.CreateBranchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDeploymentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDeploymentFuture) Get(ctx workflow.Context) (*amplify.CreateDeploymentOutput, error) {
	var output amplify.CreateDeploymentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDomainAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDomainAssociationFuture) Get(ctx workflow.Context) (*amplify.CreateDomainAssociationOutput, error) {
	var output amplify.CreateDomainAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateWebhookFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateWebhookFuture) Get(ctx workflow.Context) (*amplify.CreateWebhookOutput, error) {
	var output amplify.CreateWebhookOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteAppFuture) Get(ctx workflow.Context) (*amplify.DeleteAppOutput, error) {
	var output amplify.DeleteAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBackendEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBackendEnvironmentFuture) Get(ctx workflow.Context) (*amplify.DeleteBackendEnvironmentOutput, error) {
	var output amplify.DeleteBackendEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBranchFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBranchFuture) Get(ctx workflow.Context) (*amplify.DeleteBranchOutput, error) {
	var output amplify.DeleteBranchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDomainAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDomainAssociationFuture) Get(ctx workflow.Context) (*amplify.DeleteDomainAssociationOutput, error) {
	var output amplify.DeleteDomainAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteJobFuture) Get(ctx workflow.Context) (*amplify.DeleteJobOutput, error) {
	var output amplify.DeleteJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteWebhookFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteWebhookFuture) Get(ctx workflow.Context) (*amplify.DeleteWebhookOutput, error) {
	var output amplify.DeleteWebhookOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GenerateAccessLogsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GenerateAccessLogsFuture) Get(ctx workflow.Context) (*amplify.GenerateAccessLogsOutput, error) {
	var output amplify.GenerateAccessLogsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetAppFuture) Get(ctx workflow.Context) (*amplify.GetAppOutput, error) {
	var output amplify.GetAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetArtifactUrlFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetArtifactUrlFuture) Get(ctx workflow.Context) (*amplify.GetArtifactUrlOutput, error) {
	var output amplify.GetArtifactUrlOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBackendEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBackendEnvironmentFuture) Get(ctx workflow.Context) (*amplify.GetBackendEnvironmentOutput, error) {
	var output amplify.GetBackendEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBranchFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBranchFuture) Get(ctx workflow.Context) (*amplify.GetBranchOutput, error) {
	var output amplify.GetBranchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDomainAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDomainAssociationFuture) Get(ctx workflow.Context) (*amplify.GetDomainAssociationOutput, error) {
	var output amplify.GetDomainAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetJobFuture) Get(ctx workflow.Context) (*amplify.GetJobOutput, error) {
	var output amplify.GetJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetWebhookFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetWebhookFuture) Get(ctx workflow.Context) (*amplify.GetWebhookOutput, error) {
	var output amplify.GetWebhookOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAppsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAppsFuture) Get(ctx workflow.Context) (*amplify.ListAppsOutput, error) {
	var output amplify.ListAppsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListArtifactsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListArtifactsFuture) Get(ctx workflow.Context) (*amplify.ListArtifactsOutput, error) {
	var output amplify.ListArtifactsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBackendEnvironmentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBackendEnvironmentsFuture) Get(ctx workflow.Context) (*amplify.ListBackendEnvironmentsOutput, error) {
	var output amplify.ListBackendEnvironmentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBranchesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBranchesFuture) Get(ctx workflow.Context) (*amplify.ListBranchesOutput, error) {
	var output amplify.ListBranchesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDomainAssociationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDomainAssociationsFuture) Get(ctx workflow.Context) (*amplify.ListDomainAssociationsOutput, error) {
	var output amplify.ListDomainAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListJobsFuture) Get(ctx workflow.Context) (*amplify.ListJobsOutput, error) {
	var output amplify.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*amplify.ListTagsForResourceOutput, error) {
	var output amplify.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListWebhooksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListWebhooksFuture) Get(ctx workflow.Context) (*amplify.ListWebhooksOutput, error) {
	var output amplify.ListWebhooksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartDeploymentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartDeploymentFuture) Get(ctx workflow.Context) (*amplify.StartDeploymentOutput, error) {
	var output amplify.StartDeploymentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartJobFuture) Get(ctx workflow.Context) (*amplify.StartJobOutput, error) {
	var output amplify.StartJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopJobFuture) Get(ctx workflow.Context) (*amplify.StopJobOutput, error) {
	var output amplify.StopJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*amplify.TagResourceOutput, error) {
	var output amplify.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*amplify.UntagResourceOutput, error) {
	var output amplify.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateAppFuture) Get(ctx workflow.Context) (*amplify.UpdateAppOutput, error) {
	var output amplify.UpdateAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBranchFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBranchFuture) Get(ctx workflow.Context) (*amplify.UpdateBranchOutput, error) {
	var output amplify.UpdateBranchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDomainAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDomainAssociationFuture) Get(ctx workflow.Context) (*amplify.UpdateDomainAssociationOutput, error) {
	var output amplify.UpdateDomainAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateWebhookFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateWebhookFuture) Get(ctx workflow.Context) (*amplify.UpdateWebhookOutput, error) {
	var output amplify.UpdateWebhookOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApp(ctx workflow.Context, input *amplify.CreateAppInput) (*amplify.CreateAppOutput, error) {
	var output amplify.CreateAppOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-CreateApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAppAsync(ctx workflow.Context, input *amplify.CreateAppInput) *CreateAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-CreateApp", input)
	return &CreateAppFuture{Future: future}
}

func (a *stub) CreateBackendEnvironment(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) (*amplify.CreateBackendEnvironmentOutput, error) {
	var output amplify.CreateBackendEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-CreateBackendEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBackendEnvironmentAsync(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) *CreateBackendEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-CreateBackendEnvironment", input)
	return &CreateBackendEnvironmentFuture{Future: future}
}

func (a *stub) CreateBranch(ctx workflow.Context, input *amplify.CreateBranchInput) (*amplify.CreateBranchOutput, error) {
	var output amplify.CreateBranchOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-CreateBranch", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBranchAsync(ctx workflow.Context, input *amplify.CreateBranchInput) *CreateBranchFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-CreateBranch", input)
	return &CreateBranchFuture{Future: future}
}

func (a *stub) CreateDeployment(ctx workflow.Context, input *amplify.CreateDeploymentInput) (*amplify.CreateDeploymentOutput, error) {
	var output amplify.CreateDeploymentOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-CreateDeployment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDeploymentAsync(ctx workflow.Context, input *amplify.CreateDeploymentInput) *CreateDeploymentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-CreateDeployment", input)
	return &CreateDeploymentFuture{Future: future}
}

func (a *stub) CreateDomainAssociation(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) (*amplify.CreateDomainAssociationOutput, error) {
	var output amplify.CreateDomainAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-CreateDomainAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDomainAssociationAsync(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) *CreateDomainAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-CreateDomainAssociation", input)
	return &CreateDomainAssociationFuture{Future: future}
}

func (a *stub) CreateWebhook(ctx workflow.Context, input *amplify.CreateWebhookInput) (*amplify.CreateWebhookOutput, error) {
	var output amplify.CreateWebhookOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-CreateWebhook", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWebhookAsync(ctx workflow.Context, input *amplify.CreateWebhookInput) *CreateWebhookFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-CreateWebhook", input)
	return &CreateWebhookFuture{Future: future}
}

func (a *stub) DeleteApp(ctx workflow.Context, input *amplify.DeleteAppInput) (*amplify.DeleteAppOutput, error) {
	var output amplify.DeleteAppOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAppAsync(ctx workflow.Context, input *amplify.DeleteAppInput) *DeleteAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteApp", input)
	return &DeleteAppFuture{Future: future}
}

func (a *stub) DeleteBackendEnvironment(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) (*amplify.DeleteBackendEnvironmentOutput, error) {
	var output amplify.DeleteBackendEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteBackendEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBackendEnvironmentAsync(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) *DeleteBackendEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteBackendEnvironment", input)
	return &DeleteBackendEnvironmentFuture{Future: future}
}

func (a *stub) DeleteBranch(ctx workflow.Context, input *amplify.DeleteBranchInput) (*amplify.DeleteBranchOutput, error) {
	var output amplify.DeleteBranchOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteBranch", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBranchAsync(ctx workflow.Context, input *amplify.DeleteBranchInput) *DeleteBranchFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteBranch", input)
	return &DeleteBranchFuture{Future: future}
}

func (a *stub) DeleteDomainAssociation(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) (*amplify.DeleteDomainAssociationOutput, error) {
	var output amplify.DeleteDomainAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteDomainAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDomainAssociationAsync(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) *DeleteDomainAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteDomainAssociation", input)
	return &DeleteDomainAssociationFuture{Future: future}
}

func (a *stub) DeleteJob(ctx workflow.Context, input *amplify.DeleteJobInput) (*amplify.DeleteJobOutput, error) {
	var output amplify.DeleteJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteJobAsync(ctx workflow.Context, input *amplify.DeleteJobInput) *DeleteJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteJob", input)
	return &DeleteJobFuture{Future: future}
}

func (a *stub) DeleteWebhook(ctx workflow.Context, input *amplify.DeleteWebhookInput) (*amplify.DeleteWebhookOutput, error) {
	var output amplify.DeleteWebhookOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteWebhook", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteWebhookAsync(ctx workflow.Context, input *amplify.DeleteWebhookInput) *DeleteWebhookFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-DeleteWebhook", input)
	return &DeleteWebhookFuture{Future: future}
}

func (a *stub) GenerateAccessLogs(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) (*amplify.GenerateAccessLogsOutput, error) {
	var output amplify.GenerateAccessLogsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-GenerateAccessLogs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateAccessLogsAsync(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) *GenerateAccessLogsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-GenerateAccessLogs", input)
	return &GenerateAccessLogsFuture{Future: future}
}

func (a *stub) GetApp(ctx workflow.Context, input *amplify.GetAppInput) (*amplify.GetAppOutput, error) {
	var output amplify.GetAppOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-GetApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAppAsync(ctx workflow.Context, input *amplify.GetAppInput) *GetAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-GetApp", input)
	return &GetAppFuture{Future: future}
}

func (a *stub) GetArtifactUrl(ctx workflow.Context, input *amplify.GetArtifactUrlInput) (*amplify.GetArtifactUrlOutput, error) {
	var output amplify.GetArtifactUrlOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-GetArtifactUrl", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetArtifactUrlAsync(ctx workflow.Context, input *amplify.GetArtifactUrlInput) *GetArtifactUrlFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-GetArtifactUrl", input)
	return &GetArtifactUrlFuture{Future: future}
}

func (a *stub) GetBackendEnvironment(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) (*amplify.GetBackendEnvironmentOutput, error) {
	var output amplify.GetBackendEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-GetBackendEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBackendEnvironmentAsync(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) *GetBackendEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-GetBackendEnvironment", input)
	return &GetBackendEnvironmentFuture{Future: future}
}

func (a *stub) GetBranch(ctx workflow.Context, input *amplify.GetBranchInput) (*amplify.GetBranchOutput, error) {
	var output amplify.GetBranchOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-GetBranch", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBranchAsync(ctx workflow.Context, input *amplify.GetBranchInput) *GetBranchFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-GetBranch", input)
	return &GetBranchFuture{Future: future}
}

func (a *stub) GetDomainAssociation(ctx workflow.Context, input *amplify.GetDomainAssociationInput) (*amplify.GetDomainAssociationOutput, error) {
	var output amplify.GetDomainAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-GetDomainAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDomainAssociationAsync(ctx workflow.Context, input *amplify.GetDomainAssociationInput) *GetDomainAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-GetDomainAssociation", input)
	return &GetDomainAssociationFuture{Future: future}
}

func (a *stub) GetJob(ctx workflow.Context, input *amplify.GetJobInput) (*amplify.GetJobOutput, error) {
	var output amplify.GetJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-GetJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetJobAsync(ctx workflow.Context, input *amplify.GetJobInput) *GetJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-GetJob", input)
	return &GetJobFuture{Future: future}
}

func (a *stub) GetWebhook(ctx workflow.Context, input *amplify.GetWebhookInput) (*amplify.GetWebhookOutput, error) {
	var output amplify.GetWebhookOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-GetWebhook", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetWebhookAsync(ctx workflow.Context, input *amplify.GetWebhookInput) *GetWebhookFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-GetWebhook", input)
	return &GetWebhookFuture{Future: future}
}

func (a *stub) ListApps(ctx workflow.Context, input *amplify.ListAppsInput) (*amplify.ListAppsOutput, error) {
	var output amplify.ListAppsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-ListApps", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAppsAsync(ctx workflow.Context, input *amplify.ListAppsInput) *ListAppsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-ListApps", input)
	return &ListAppsFuture{Future: future}
}

func (a *stub) ListArtifacts(ctx workflow.Context, input *amplify.ListArtifactsInput) (*amplify.ListArtifactsOutput, error) {
	var output amplify.ListArtifactsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-ListArtifacts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListArtifactsAsync(ctx workflow.Context, input *amplify.ListArtifactsInput) *ListArtifactsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-ListArtifacts", input)
	return &ListArtifactsFuture{Future: future}
}

func (a *stub) ListBackendEnvironments(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) (*amplify.ListBackendEnvironmentsOutput, error) {
	var output amplify.ListBackendEnvironmentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-ListBackendEnvironments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBackendEnvironmentsAsync(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) *ListBackendEnvironmentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-ListBackendEnvironments", input)
	return &ListBackendEnvironmentsFuture{Future: future}
}

func (a *stub) ListBranches(ctx workflow.Context, input *amplify.ListBranchesInput) (*amplify.ListBranchesOutput, error) {
	var output amplify.ListBranchesOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-ListBranches", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBranchesAsync(ctx workflow.Context, input *amplify.ListBranchesInput) *ListBranchesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-ListBranches", input)
	return &ListBranchesFuture{Future: future}
}

func (a *stub) ListDomainAssociations(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) (*amplify.ListDomainAssociationsOutput, error) {
	var output amplify.ListDomainAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-ListDomainAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDomainAssociationsAsync(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) *ListDomainAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-ListDomainAssociations", input)
	return &ListDomainAssociationsFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *amplify.ListJobsInput) (*amplify.ListJobsOutput, error) {
	var output amplify.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *amplify.ListJobsInput) *ListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-ListJobs", input)
	return &ListJobsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *amplify.ListTagsForResourceInput) (*amplify.ListTagsForResourceOutput, error) {
	var output amplify.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *amplify.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ListWebhooks(ctx workflow.Context, input *amplify.ListWebhooksInput) (*amplify.ListWebhooksOutput, error) {
	var output amplify.ListWebhooksOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-ListWebhooks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWebhooksAsync(ctx workflow.Context, input *amplify.ListWebhooksInput) *ListWebhooksFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-ListWebhooks", input)
	return &ListWebhooksFuture{Future: future}
}

func (a *stub) StartDeployment(ctx workflow.Context, input *amplify.StartDeploymentInput) (*amplify.StartDeploymentOutput, error) {
	var output amplify.StartDeploymentOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-StartDeployment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartDeploymentAsync(ctx workflow.Context, input *amplify.StartDeploymentInput) *StartDeploymentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-StartDeployment", input)
	return &StartDeploymentFuture{Future: future}
}

func (a *stub) StartJob(ctx workflow.Context, input *amplify.StartJobInput) (*amplify.StartJobOutput, error) {
	var output amplify.StartJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-StartJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartJobAsync(ctx workflow.Context, input *amplify.StartJobInput) *StartJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-StartJob", input)
	return &StartJobFuture{Future: future}
}

func (a *stub) StopJob(ctx workflow.Context, input *amplify.StopJobInput) (*amplify.StopJobOutput, error) {
	var output amplify.StopJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-StopJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopJobAsync(ctx workflow.Context, input *amplify.StopJobInput) *StopJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-StopJob", input)
	return &StopJobFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *amplify.TagResourceInput) (*amplify.TagResourceOutput, error) {
	var output amplify.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *amplify.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *amplify.UntagResourceInput) (*amplify.UntagResourceOutput, error) {
	var output amplify.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *amplify.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateApp(ctx workflow.Context, input *amplify.UpdateAppInput) (*amplify.UpdateAppOutput, error) {
	var output amplify.UpdateAppOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-UpdateApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAppAsync(ctx workflow.Context, input *amplify.UpdateAppInput) *UpdateAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-UpdateApp", input)
	return &UpdateAppFuture{Future: future}
}

func (a *stub) UpdateBranch(ctx workflow.Context, input *amplify.UpdateBranchInput) (*amplify.UpdateBranchOutput, error) {
	var output amplify.UpdateBranchOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-UpdateBranch", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBranchAsync(ctx workflow.Context, input *amplify.UpdateBranchInput) *UpdateBranchFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-UpdateBranch", input)
	return &UpdateBranchFuture{Future: future}
}

func (a *stub) UpdateDomainAssociation(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) (*amplify.UpdateDomainAssociationOutput, error) {
	var output amplify.UpdateDomainAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-UpdateDomainAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDomainAssociationAsync(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) *UpdateDomainAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-UpdateDomainAssociation", input)
	return &UpdateDomainAssociationFuture{Future: future}
}

func (a *stub) UpdateWebhook(ctx workflow.Context, input *amplify.UpdateWebhookInput) (*amplify.UpdateWebhookOutput, error) {
	var output amplify.UpdateWebhookOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplify-UpdateWebhook", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateWebhookAsync(ctx workflow.Context, input *amplify.UpdateWebhookInput) *UpdateWebhookFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplify-UpdateWebhook", input)
	return &UpdateWebhookFuture{Future: future}
}
