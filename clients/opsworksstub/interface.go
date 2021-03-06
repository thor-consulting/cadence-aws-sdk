// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package opsworksstub

import (
	"github.com/aws/aws-sdk-go/service/opsworks"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssignInstance(ctx workflow.Context, input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error)
	AssignInstanceAsync(ctx workflow.Context, input *opsworks.AssignInstanceInput) *AssignInstanceFuture

	AssignVolume(ctx workflow.Context, input *opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error)
	AssignVolumeAsync(ctx workflow.Context, input *opsworks.AssignVolumeInput) *AssignVolumeFuture

	AssociateElasticIp(ctx workflow.Context, input *opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error)
	AssociateElasticIpAsync(ctx workflow.Context, input *opsworks.AssociateElasticIpInput) *AssociateElasticIpFuture

	AttachElasticLoadBalancer(ctx workflow.Context, input *opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error)
	AttachElasticLoadBalancerAsync(ctx workflow.Context, input *opsworks.AttachElasticLoadBalancerInput) *AttachElasticLoadBalancerFuture

	CloneStack(ctx workflow.Context, input *opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error)
	CloneStackAsync(ctx workflow.Context, input *opsworks.CloneStackInput) *CloneStackFuture

	CreateApp(ctx workflow.Context, input *opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error)
	CreateAppAsync(ctx workflow.Context, input *opsworks.CreateAppInput) *CreateAppFuture

	CreateDeployment(ctx workflow.Context, input *opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error)
	CreateDeploymentAsync(ctx workflow.Context, input *opsworks.CreateDeploymentInput) *CreateDeploymentFuture

	CreateInstance(ctx workflow.Context, input *opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error)
	CreateInstanceAsync(ctx workflow.Context, input *opsworks.CreateInstanceInput) *CreateInstanceFuture

	CreateLayer(ctx workflow.Context, input *opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error)
	CreateLayerAsync(ctx workflow.Context, input *opsworks.CreateLayerInput) *CreateLayerFuture

	CreateStack(ctx workflow.Context, input *opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error)
	CreateStackAsync(ctx workflow.Context, input *opsworks.CreateStackInput) *CreateStackFuture

	CreateUserProfile(ctx workflow.Context, input *opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error)
	CreateUserProfileAsync(ctx workflow.Context, input *opsworks.CreateUserProfileInput) *CreateUserProfileFuture

	DeleteApp(ctx workflow.Context, input *opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error)
	DeleteAppAsync(ctx workflow.Context, input *opsworks.DeleteAppInput) *DeleteAppFuture

	DeleteInstance(ctx workflow.Context, input *opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error)
	DeleteInstanceAsync(ctx workflow.Context, input *opsworks.DeleteInstanceInput) *DeleteInstanceFuture

	DeleteLayer(ctx workflow.Context, input *opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error)
	DeleteLayerAsync(ctx workflow.Context, input *opsworks.DeleteLayerInput) *DeleteLayerFuture

	DeleteStack(ctx workflow.Context, input *opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error)
	DeleteStackAsync(ctx workflow.Context, input *opsworks.DeleteStackInput) *DeleteStackFuture

	DeleteUserProfile(ctx workflow.Context, input *opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error)
	DeleteUserProfileAsync(ctx workflow.Context, input *opsworks.DeleteUserProfileInput) *DeleteUserProfileFuture

	DeregisterEcsCluster(ctx workflow.Context, input *opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error)
	DeregisterEcsClusterAsync(ctx workflow.Context, input *opsworks.DeregisterEcsClusterInput) *DeregisterEcsClusterFuture

	DeregisterElasticIp(ctx workflow.Context, input *opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error)
	DeregisterElasticIpAsync(ctx workflow.Context, input *opsworks.DeregisterElasticIpInput) *DeregisterElasticIpFuture

	DeregisterInstance(ctx workflow.Context, input *opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error)
	DeregisterInstanceAsync(ctx workflow.Context, input *opsworks.DeregisterInstanceInput) *DeregisterInstanceFuture

	DeregisterRdsDbInstance(ctx workflow.Context, input *opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error)
	DeregisterRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.DeregisterRdsDbInstanceInput) *DeregisterRdsDbInstanceFuture

	DeregisterVolume(ctx workflow.Context, input *opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error)
	DeregisterVolumeAsync(ctx workflow.Context, input *opsworks.DeregisterVolumeInput) *DeregisterVolumeFuture

	DescribeAgentVersions(ctx workflow.Context, input *opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error)
	DescribeAgentVersionsAsync(ctx workflow.Context, input *opsworks.DescribeAgentVersionsInput) *DescribeAgentVersionsFuture

	DescribeApps(ctx workflow.Context, input *opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error)
	DescribeAppsAsync(ctx workflow.Context, input *opsworks.DescribeAppsInput) *DescribeAppsFuture

	DescribeCommands(ctx workflow.Context, input *opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error)
	DescribeCommandsAsync(ctx workflow.Context, input *opsworks.DescribeCommandsInput) *DescribeCommandsFuture

	DescribeDeployments(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error)
	DescribeDeploymentsAsync(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) *DescribeDeploymentsFuture

	DescribeEcsClusters(ctx workflow.Context, input *opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error)
	DescribeEcsClustersAsync(ctx workflow.Context, input *opsworks.DescribeEcsClustersInput) *DescribeEcsClustersFuture

	DescribeElasticIps(ctx workflow.Context, input *opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error)
	DescribeElasticIpsAsync(ctx workflow.Context, input *opsworks.DescribeElasticIpsInput) *DescribeElasticIpsFuture

	DescribeElasticLoadBalancers(ctx workflow.Context, input *opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error)
	DescribeElasticLoadBalancersAsync(ctx workflow.Context, input *opsworks.DescribeElasticLoadBalancersInput) *DescribeElasticLoadBalancersFuture

	DescribeInstances(ctx workflow.Context, input *opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error)
	DescribeInstancesAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *DescribeInstancesFuture

	DescribeLayers(ctx workflow.Context, input *opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error)
	DescribeLayersAsync(ctx workflow.Context, input *opsworks.DescribeLayersInput) *DescribeLayersFuture

	DescribeLoadBasedAutoScaling(ctx workflow.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)
	DescribeLoadBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) *DescribeLoadBasedAutoScalingFuture

	DescribeMyUserProfile(ctx workflow.Context, input *opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error)
	DescribeMyUserProfileAsync(ctx workflow.Context, input *opsworks.DescribeMyUserProfileInput) *DescribeMyUserProfileFuture

	DescribeOperatingSystems(ctx workflow.Context, input *opsworks.DescribeOperatingSystemsInput) (*opsworks.DescribeOperatingSystemsOutput, error)
	DescribeOperatingSystemsAsync(ctx workflow.Context, input *opsworks.DescribeOperatingSystemsInput) *DescribeOperatingSystemsFuture

	DescribePermissions(ctx workflow.Context, input *opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error)
	DescribePermissionsAsync(ctx workflow.Context, input *opsworks.DescribePermissionsInput) *DescribePermissionsFuture

	DescribeRaidArrays(ctx workflow.Context, input *opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error)
	DescribeRaidArraysAsync(ctx workflow.Context, input *opsworks.DescribeRaidArraysInput) *DescribeRaidArraysFuture

	DescribeRdsDbInstances(ctx workflow.Context, input *opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error)
	DescribeRdsDbInstancesAsync(ctx workflow.Context, input *opsworks.DescribeRdsDbInstancesInput) *DescribeRdsDbInstancesFuture

	DescribeServiceErrors(ctx workflow.Context, input *opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error)
	DescribeServiceErrorsAsync(ctx workflow.Context, input *opsworks.DescribeServiceErrorsInput) *DescribeServiceErrorsFuture

	DescribeStackProvisioningParameters(ctx workflow.Context, input *opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error)
	DescribeStackProvisioningParametersAsync(ctx workflow.Context, input *opsworks.DescribeStackProvisioningParametersInput) *DescribeStackProvisioningParametersFuture

	DescribeStackSummary(ctx workflow.Context, input *opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error)
	DescribeStackSummaryAsync(ctx workflow.Context, input *opsworks.DescribeStackSummaryInput) *DescribeStackSummaryFuture

	DescribeStacks(ctx workflow.Context, input *opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error)
	DescribeStacksAsync(ctx workflow.Context, input *opsworks.DescribeStacksInput) *DescribeStacksFuture

	DescribeTimeBasedAutoScaling(ctx workflow.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)
	DescribeTimeBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) *DescribeTimeBasedAutoScalingFuture

	DescribeUserProfiles(ctx workflow.Context, input *opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error)
	DescribeUserProfilesAsync(ctx workflow.Context, input *opsworks.DescribeUserProfilesInput) *DescribeUserProfilesFuture

	DescribeVolumes(ctx workflow.Context, input *opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error)
	DescribeVolumesAsync(ctx workflow.Context, input *opsworks.DescribeVolumesInput) *DescribeVolumesFuture

	DetachElasticLoadBalancer(ctx workflow.Context, input *opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error)
	DetachElasticLoadBalancerAsync(ctx workflow.Context, input *opsworks.DetachElasticLoadBalancerInput) *DetachElasticLoadBalancerFuture

	DisassociateElasticIp(ctx workflow.Context, input *opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error)
	DisassociateElasticIpAsync(ctx workflow.Context, input *opsworks.DisassociateElasticIpInput) *DisassociateElasticIpFuture

	GetHostnameSuggestion(ctx workflow.Context, input *opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error)
	GetHostnameSuggestionAsync(ctx workflow.Context, input *opsworks.GetHostnameSuggestionInput) *GetHostnameSuggestionFuture

	GrantAccess(ctx workflow.Context, input *opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error)
	GrantAccessAsync(ctx workflow.Context, input *opsworks.GrantAccessInput) *GrantAccessFuture

	ListTags(ctx workflow.Context, input *opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *opsworks.ListTagsInput) *ListTagsFuture

	RebootInstance(ctx workflow.Context, input *opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error)
	RebootInstanceAsync(ctx workflow.Context, input *opsworks.RebootInstanceInput) *RebootInstanceFuture

	RegisterEcsCluster(ctx workflow.Context, input *opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error)
	RegisterEcsClusterAsync(ctx workflow.Context, input *opsworks.RegisterEcsClusterInput) *RegisterEcsClusterFuture

	RegisterElasticIp(ctx workflow.Context, input *opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error)
	RegisterElasticIpAsync(ctx workflow.Context, input *opsworks.RegisterElasticIpInput) *RegisterElasticIpFuture

	RegisterInstance(ctx workflow.Context, input *opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error)
	RegisterInstanceAsync(ctx workflow.Context, input *opsworks.RegisterInstanceInput) *RegisterInstanceFuture

	RegisterRdsDbInstance(ctx workflow.Context, input *opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error)
	RegisterRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.RegisterRdsDbInstanceInput) *RegisterRdsDbInstanceFuture

	RegisterVolume(ctx workflow.Context, input *opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error)
	RegisterVolumeAsync(ctx workflow.Context, input *opsworks.RegisterVolumeInput) *RegisterVolumeFuture

	SetLoadBasedAutoScaling(ctx workflow.Context, input *opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error)
	SetLoadBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.SetLoadBasedAutoScalingInput) *SetLoadBasedAutoScalingFuture

	SetPermission(ctx workflow.Context, input *opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error)
	SetPermissionAsync(ctx workflow.Context, input *opsworks.SetPermissionInput) *SetPermissionFuture

	SetTimeBasedAutoScaling(ctx workflow.Context, input *opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error)
	SetTimeBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.SetTimeBasedAutoScalingInput) *SetTimeBasedAutoScalingFuture

	StartInstance(ctx workflow.Context, input *opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error)
	StartInstanceAsync(ctx workflow.Context, input *opsworks.StartInstanceInput) *StartInstanceFuture

	StartStack(ctx workflow.Context, input *opsworks.StartStackInput) (*opsworks.StartStackOutput, error)
	StartStackAsync(ctx workflow.Context, input *opsworks.StartStackInput) *StartStackFuture

	StopInstance(ctx workflow.Context, input *opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error)
	StopInstanceAsync(ctx workflow.Context, input *opsworks.StopInstanceInput) *StopInstanceFuture

	StopStack(ctx workflow.Context, input *opsworks.StopStackInput) (*opsworks.StopStackOutput, error)
	StopStackAsync(ctx workflow.Context, input *opsworks.StopStackInput) *StopStackFuture

	TagResource(ctx workflow.Context, input *opsworks.TagResourceInput) (*opsworks.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *opsworks.TagResourceInput) *TagResourceFuture

	UnassignInstance(ctx workflow.Context, input *opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error)
	UnassignInstanceAsync(ctx workflow.Context, input *opsworks.UnassignInstanceInput) *UnassignInstanceFuture

	UnassignVolume(ctx workflow.Context, input *opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error)
	UnassignVolumeAsync(ctx workflow.Context, input *opsworks.UnassignVolumeInput) *UnassignVolumeFuture

	UntagResource(ctx workflow.Context, input *opsworks.UntagResourceInput) (*opsworks.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *opsworks.UntagResourceInput) *UntagResourceFuture

	UpdateApp(ctx workflow.Context, input *opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error)
	UpdateAppAsync(ctx workflow.Context, input *opsworks.UpdateAppInput) *UpdateAppFuture

	UpdateElasticIp(ctx workflow.Context, input *opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error)
	UpdateElasticIpAsync(ctx workflow.Context, input *opsworks.UpdateElasticIpInput) *UpdateElasticIpFuture

	UpdateInstance(ctx workflow.Context, input *opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error)
	UpdateInstanceAsync(ctx workflow.Context, input *opsworks.UpdateInstanceInput) *UpdateInstanceFuture

	UpdateLayer(ctx workflow.Context, input *opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error)
	UpdateLayerAsync(ctx workflow.Context, input *opsworks.UpdateLayerInput) *UpdateLayerFuture

	UpdateMyUserProfile(ctx workflow.Context, input *opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error)
	UpdateMyUserProfileAsync(ctx workflow.Context, input *opsworks.UpdateMyUserProfileInput) *UpdateMyUserProfileFuture

	UpdateRdsDbInstance(ctx workflow.Context, input *opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error)
	UpdateRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.UpdateRdsDbInstanceInput) *UpdateRdsDbInstanceFuture

	UpdateStack(ctx workflow.Context, input *opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error)
	UpdateStackAsync(ctx workflow.Context, input *opsworks.UpdateStackInput) *UpdateStackFuture

	UpdateUserProfile(ctx workflow.Context, input *opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error)
	UpdateUserProfileAsync(ctx workflow.Context, input *opsworks.UpdateUserProfileInput) *UpdateUserProfileFuture

	UpdateVolume(ctx workflow.Context, input *opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error)
	UpdateVolumeAsync(ctx workflow.Context, input *opsworks.UpdateVolumeInput) *UpdateVolumeFuture

	WaitUntilAppExists(ctx workflow.Context, input *opsworks.DescribeAppsInput) error
	WaitUntilAppExistsAsync(ctx workflow.Context, input *opsworks.DescribeAppsInput) *clients.VoidFuture

	WaitUntilDeploymentSuccessful(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) error
	WaitUntilDeploymentSuccessfulAsync(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) *clients.VoidFuture

	WaitUntilInstanceOnline(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
	WaitUntilInstanceOnlineAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *clients.VoidFuture

	WaitUntilInstanceRegistered(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
	WaitUntilInstanceRegisteredAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *clients.VoidFuture

	WaitUntilInstanceStopped(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
	WaitUntilInstanceStoppedAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *clients.VoidFuture

	WaitUntilInstanceTerminated(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
	WaitUntilInstanceTerminatedAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
