// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package greengrassv2stub

import (
	"github.com/aws/aws-sdk-go/service/greengrassv2"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CancelDeploymentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CancelDeploymentFuture) Get(ctx workflow.Context) (*greengrassv2.CancelDeploymentOutput, error) {
	var output greengrassv2.CancelDeploymentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateComponentVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateComponentVersionFuture) Get(ctx workflow.Context) (*greengrassv2.CreateComponentVersionOutput, error) {
	var output greengrassv2.CreateComponentVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDeploymentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDeploymentFuture) Get(ctx workflow.Context) (*greengrassv2.CreateDeploymentOutput, error) {
	var output greengrassv2.CreateDeploymentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteComponentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteComponentFuture) Get(ctx workflow.Context) (*greengrassv2.DeleteComponentOutput, error) {
	var output greengrassv2.DeleteComponentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteCoreDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteCoreDeviceFuture) Get(ctx workflow.Context) (*greengrassv2.DeleteCoreDeviceOutput, error) {
	var output greengrassv2.DeleteCoreDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeComponentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeComponentFuture) Get(ctx workflow.Context) (*greengrassv2.DescribeComponentOutput, error) {
	var output greengrassv2.DescribeComponentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetComponentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetComponentFuture) Get(ctx workflow.Context) (*greengrassv2.GetComponentOutput, error) {
	var output greengrassv2.GetComponentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetComponentVersionArtifactFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetComponentVersionArtifactFuture) Get(ctx workflow.Context) (*greengrassv2.GetComponentVersionArtifactOutput, error) {
	var output greengrassv2.GetComponentVersionArtifactOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetCoreDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetCoreDeviceFuture) Get(ctx workflow.Context) (*greengrassv2.GetCoreDeviceOutput, error) {
	var output greengrassv2.GetCoreDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDeploymentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDeploymentFuture) Get(ctx workflow.Context) (*greengrassv2.GetDeploymentOutput, error) {
	var output greengrassv2.GetDeploymentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListComponentVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListComponentVersionsFuture) Get(ctx workflow.Context) (*greengrassv2.ListComponentVersionsOutput, error) {
	var output greengrassv2.ListComponentVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListComponentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListComponentsFuture) Get(ctx workflow.Context) (*greengrassv2.ListComponentsOutput, error) {
	var output greengrassv2.ListComponentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListCoreDevicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListCoreDevicesFuture) Get(ctx workflow.Context) (*greengrassv2.ListCoreDevicesOutput, error) {
	var output greengrassv2.ListCoreDevicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDeploymentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDeploymentsFuture) Get(ctx workflow.Context) (*greengrassv2.ListDeploymentsOutput, error) {
	var output greengrassv2.ListDeploymentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListEffectiveDeploymentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListEffectiveDeploymentsFuture) Get(ctx workflow.Context) (*greengrassv2.ListEffectiveDeploymentsOutput, error) {
	var output greengrassv2.ListEffectiveDeploymentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListInstalledComponentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListInstalledComponentsFuture) Get(ctx workflow.Context) (*greengrassv2.ListInstalledComponentsOutput, error) {
	var output greengrassv2.ListInstalledComponentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*greengrassv2.ListTagsForResourceOutput, error) {
	var output greengrassv2.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResolveComponentCandidatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResolveComponentCandidatesFuture) Get(ctx workflow.Context) (*greengrassv2.ResolveComponentCandidatesOutput, error) {
	var output greengrassv2.ResolveComponentCandidatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*greengrassv2.TagResourceOutput, error) {
	var output greengrassv2.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*greengrassv2.UntagResourceOutput, error) {
	var output greengrassv2.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelDeployment(ctx workflow.Context, input *greengrassv2.CancelDeploymentInput) (*greengrassv2.CancelDeploymentOutput, error) {
	var output greengrassv2.CancelDeploymentOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-CancelDeployment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelDeploymentAsync(ctx workflow.Context, input *greengrassv2.CancelDeploymentInput) *CancelDeploymentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-CancelDeployment", input)
	return &CancelDeploymentFuture{Future: future}
}

func (a *stub) CreateComponentVersion(ctx workflow.Context, input *greengrassv2.CreateComponentVersionInput) (*greengrassv2.CreateComponentVersionOutput, error) {
	var output greengrassv2.CreateComponentVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-CreateComponentVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateComponentVersionAsync(ctx workflow.Context, input *greengrassv2.CreateComponentVersionInput) *CreateComponentVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-CreateComponentVersion", input)
	return &CreateComponentVersionFuture{Future: future}
}

func (a *stub) CreateDeployment(ctx workflow.Context, input *greengrassv2.CreateDeploymentInput) (*greengrassv2.CreateDeploymentOutput, error) {
	var output greengrassv2.CreateDeploymentOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-CreateDeployment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDeploymentAsync(ctx workflow.Context, input *greengrassv2.CreateDeploymentInput) *CreateDeploymentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-CreateDeployment", input)
	return &CreateDeploymentFuture{Future: future}
}

func (a *stub) DeleteComponent(ctx workflow.Context, input *greengrassv2.DeleteComponentInput) (*greengrassv2.DeleteComponentOutput, error) {
	var output greengrassv2.DeleteComponentOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-DeleteComponent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteComponentAsync(ctx workflow.Context, input *greengrassv2.DeleteComponentInput) *DeleteComponentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-DeleteComponent", input)
	return &DeleteComponentFuture{Future: future}
}

func (a *stub) DeleteCoreDevice(ctx workflow.Context, input *greengrassv2.DeleteCoreDeviceInput) (*greengrassv2.DeleteCoreDeviceOutput, error) {
	var output greengrassv2.DeleteCoreDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-DeleteCoreDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteCoreDeviceAsync(ctx workflow.Context, input *greengrassv2.DeleteCoreDeviceInput) *DeleteCoreDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-DeleteCoreDevice", input)
	return &DeleteCoreDeviceFuture{Future: future}
}

func (a *stub) DescribeComponent(ctx workflow.Context, input *greengrassv2.DescribeComponentInput) (*greengrassv2.DescribeComponentOutput, error) {
	var output greengrassv2.DescribeComponentOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-DescribeComponent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeComponentAsync(ctx workflow.Context, input *greengrassv2.DescribeComponentInput) *DescribeComponentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-DescribeComponent", input)
	return &DescribeComponentFuture{Future: future}
}

func (a *stub) GetComponent(ctx workflow.Context, input *greengrassv2.GetComponentInput) (*greengrassv2.GetComponentOutput, error) {
	var output greengrassv2.GetComponentOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-GetComponent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetComponentAsync(ctx workflow.Context, input *greengrassv2.GetComponentInput) *GetComponentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-GetComponent", input)
	return &GetComponentFuture{Future: future}
}

func (a *stub) GetComponentVersionArtifact(ctx workflow.Context, input *greengrassv2.GetComponentVersionArtifactInput) (*greengrassv2.GetComponentVersionArtifactOutput, error) {
	var output greengrassv2.GetComponentVersionArtifactOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-GetComponentVersionArtifact", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetComponentVersionArtifactAsync(ctx workflow.Context, input *greengrassv2.GetComponentVersionArtifactInput) *GetComponentVersionArtifactFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-GetComponentVersionArtifact", input)
	return &GetComponentVersionArtifactFuture{Future: future}
}

func (a *stub) GetCoreDevice(ctx workflow.Context, input *greengrassv2.GetCoreDeviceInput) (*greengrassv2.GetCoreDeviceOutput, error) {
	var output greengrassv2.GetCoreDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-GetCoreDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCoreDeviceAsync(ctx workflow.Context, input *greengrassv2.GetCoreDeviceInput) *GetCoreDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-GetCoreDevice", input)
	return &GetCoreDeviceFuture{Future: future}
}

func (a *stub) GetDeployment(ctx workflow.Context, input *greengrassv2.GetDeploymentInput) (*greengrassv2.GetDeploymentOutput, error) {
	var output greengrassv2.GetDeploymentOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-GetDeployment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDeploymentAsync(ctx workflow.Context, input *greengrassv2.GetDeploymentInput) *GetDeploymentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-GetDeployment", input)
	return &GetDeploymentFuture{Future: future}
}

func (a *stub) ListComponentVersions(ctx workflow.Context, input *greengrassv2.ListComponentVersionsInput) (*greengrassv2.ListComponentVersionsOutput, error) {
	var output greengrassv2.ListComponentVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListComponentVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListComponentVersionsAsync(ctx workflow.Context, input *greengrassv2.ListComponentVersionsInput) *ListComponentVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListComponentVersions", input)
	return &ListComponentVersionsFuture{Future: future}
}

func (a *stub) ListComponents(ctx workflow.Context, input *greengrassv2.ListComponentsInput) (*greengrassv2.ListComponentsOutput, error) {
	var output greengrassv2.ListComponentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListComponents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListComponentsAsync(ctx workflow.Context, input *greengrassv2.ListComponentsInput) *ListComponentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListComponents", input)
	return &ListComponentsFuture{Future: future}
}

func (a *stub) ListCoreDevices(ctx workflow.Context, input *greengrassv2.ListCoreDevicesInput) (*greengrassv2.ListCoreDevicesOutput, error) {
	var output greengrassv2.ListCoreDevicesOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListCoreDevices", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListCoreDevicesAsync(ctx workflow.Context, input *greengrassv2.ListCoreDevicesInput) *ListCoreDevicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListCoreDevices", input)
	return &ListCoreDevicesFuture{Future: future}
}

func (a *stub) ListDeployments(ctx workflow.Context, input *greengrassv2.ListDeploymentsInput) (*greengrassv2.ListDeploymentsOutput, error) {
	var output greengrassv2.ListDeploymentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListDeployments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDeploymentsAsync(ctx workflow.Context, input *greengrassv2.ListDeploymentsInput) *ListDeploymentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListDeployments", input)
	return &ListDeploymentsFuture{Future: future}
}

func (a *stub) ListEffectiveDeployments(ctx workflow.Context, input *greengrassv2.ListEffectiveDeploymentsInput) (*greengrassv2.ListEffectiveDeploymentsOutput, error) {
	var output greengrassv2.ListEffectiveDeploymentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListEffectiveDeployments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEffectiveDeploymentsAsync(ctx workflow.Context, input *greengrassv2.ListEffectiveDeploymentsInput) *ListEffectiveDeploymentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListEffectiveDeployments", input)
	return &ListEffectiveDeploymentsFuture{Future: future}
}

func (a *stub) ListInstalledComponents(ctx workflow.Context, input *greengrassv2.ListInstalledComponentsInput) (*greengrassv2.ListInstalledComponentsOutput, error) {
	var output greengrassv2.ListInstalledComponentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListInstalledComponents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInstalledComponentsAsync(ctx workflow.Context, input *greengrassv2.ListInstalledComponentsInput) *ListInstalledComponentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListInstalledComponents", input)
	return &ListInstalledComponentsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *greengrassv2.ListTagsForResourceInput) (*greengrassv2.ListTagsForResourceOutput, error) {
	var output greengrassv2.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *greengrassv2.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ResolveComponentCandidates(ctx workflow.Context, input *greengrassv2.ResolveComponentCandidatesInput) (*greengrassv2.ResolveComponentCandidatesOutput, error) {
	var output greengrassv2.ResolveComponentCandidatesOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ResolveComponentCandidates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ResolveComponentCandidatesAsync(ctx workflow.Context, input *greengrassv2.ResolveComponentCandidatesInput) *ResolveComponentCandidatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-ResolveComponentCandidates", input)
	return &ResolveComponentCandidatesFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *greengrassv2.TagResourceInput) (*greengrassv2.TagResourceOutput, error) {
	var output greengrassv2.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *greengrassv2.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *greengrassv2.UntagResourceInput) (*greengrassv2.UntagResourceOutput, error) {
	var output greengrassv2.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-greengrassv2-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *greengrassv2.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-greengrassv2-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}
