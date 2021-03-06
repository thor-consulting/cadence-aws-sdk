// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package route53stub

import (
	"github.com/aws/aws-sdk-go/service/route53"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	ActivateKeySigningKey(ctx workflow.Context, input *route53.ActivateKeySigningKeyInput) (*route53.ActivateKeySigningKeyOutput, error)
	ActivateKeySigningKeyAsync(ctx workflow.Context, input *route53.ActivateKeySigningKeyInput) *ActivateKeySigningKeyFuture

	AssociateVPCWithHostedZone(ctx workflow.Context, input *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error)
	AssociateVPCWithHostedZoneAsync(ctx workflow.Context, input *route53.AssociateVPCWithHostedZoneInput) *AssociateVPCWithHostedZoneFuture

	ChangeResourceRecordSets(ctx workflow.Context, input *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)
	ChangeResourceRecordSetsAsync(ctx workflow.Context, input *route53.ChangeResourceRecordSetsInput) *ChangeResourceRecordSetsFuture

	ChangeTagsForResource(ctx workflow.Context, input *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error)
	ChangeTagsForResourceAsync(ctx workflow.Context, input *route53.ChangeTagsForResourceInput) *ChangeTagsForResourceFuture

	CreateHealthCheck(ctx workflow.Context, input *route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error)
	CreateHealthCheckAsync(ctx workflow.Context, input *route53.CreateHealthCheckInput) *CreateHealthCheckFuture

	CreateHostedZone(ctx workflow.Context, input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)
	CreateHostedZoneAsync(ctx workflow.Context, input *route53.CreateHostedZoneInput) *CreateHostedZoneFuture

	CreateKeySigningKey(ctx workflow.Context, input *route53.CreateKeySigningKeyInput) (*route53.CreateKeySigningKeyOutput, error)
	CreateKeySigningKeyAsync(ctx workflow.Context, input *route53.CreateKeySigningKeyInput) *CreateKeySigningKeyFuture

	CreateQueryLoggingConfig(ctx workflow.Context, input *route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error)
	CreateQueryLoggingConfigAsync(ctx workflow.Context, input *route53.CreateQueryLoggingConfigInput) *CreateQueryLoggingConfigFuture

	CreateReusableDelegationSet(ctx workflow.Context, input *route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error)
	CreateReusableDelegationSetAsync(ctx workflow.Context, input *route53.CreateReusableDelegationSetInput) *CreateReusableDelegationSetFuture

	CreateTrafficPolicy(ctx workflow.Context, input *route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error)
	CreateTrafficPolicyAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyInput) *CreateTrafficPolicyFuture

	CreateTrafficPolicyInstance(ctx workflow.Context, input *route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error)
	CreateTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyInstanceInput) *CreateTrafficPolicyInstanceFuture

	CreateTrafficPolicyVersion(ctx workflow.Context, input *route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error)
	CreateTrafficPolicyVersionAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyVersionInput) *CreateTrafficPolicyVersionFuture

	CreateVPCAssociationAuthorization(ctx workflow.Context, input *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error)
	CreateVPCAssociationAuthorizationAsync(ctx workflow.Context, input *route53.CreateVPCAssociationAuthorizationInput) *CreateVPCAssociationAuthorizationFuture

	DeactivateKeySigningKey(ctx workflow.Context, input *route53.DeactivateKeySigningKeyInput) (*route53.DeactivateKeySigningKeyOutput, error)
	DeactivateKeySigningKeyAsync(ctx workflow.Context, input *route53.DeactivateKeySigningKeyInput) *DeactivateKeySigningKeyFuture

	DeleteHealthCheck(ctx workflow.Context, input *route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error)
	DeleteHealthCheckAsync(ctx workflow.Context, input *route53.DeleteHealthCheckInput) *DeleteHealthCheckFuture

	DeleteHostedZone(ctx workflow.Context, input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error)
	DeleteHostedZoneAsync(ctx workflow.Context, input *route53.DeleteHostedZoneInput) *DeleteHostedZoneFuture

	DeleteKeySigningKey(ctx workflow.Context, input *route53.DeleteKeySigningKeyInput) (*route53.DeleteKeySigningKeyOutput, error)
	DeleteKeySigningKeyAsync(ctx workflow.Context, input *route53.DeleteKeySigningKeyInput) *DeleteKeySigningKeyFuture

	DeleteQueryLoggingConfig(ctx workflow.Context, input *route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error)
	DeleteQueryLoggingConfigAsync(ctx workflow.Context, input *route53.DeleteQueryLoggingConfigInput) *DeleteQueryLoggingConfigFuture

	DeleteReusableDelegationSet(ctx workflow.Context, input *route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error)
	DeleteReusableDelegationSetAsync(ctx workflow.Context, input *route53.DeleteReusableDelegationSetInput) *DeleteReusableDelegationSetFuture

	DeleteTrafficPolicy(ctx workflow.Context, input *route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error)
	DeleteTrafficPolicyAsync(ctx workflow.Context, input *route53.DeleteTrafficPolicyInput) *DeleteTrafficPolicyFuture

	DeleteTrafficPolicyInstance(ctx workflow.Context, input *route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error)
	DeleteTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.DeleteTrafficPolicyInstanceInput) *DeleteTrafficPolicyInstanceFuture

	DeleteVPCAssociationAuthorization(ctx workflow.Context, input *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error)
	DeleteVPCAssociationAuthorizationAsync(ctx workflow.Context, input *route53.DeleteVPCAssociationAuthorizationInput) *DeleteVPCAssociationAuthorizationFuture

	DisableHostedZoneDNSSEC(ctx workflow.Context, input *route53.DisableHostedZoneDNSSECInput) (*route53.DisableHostedZoneDNSSECOutput, error)
	DisableHostedZoneDNSSECAsync(ctx workflow.Context, input *route53.DisableHostedZoneDNSSECInput) *DisableHostedZoneDNSSECFuture

	DisassociateVPCFromHostedZone(ctx workflow.Context, input *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error)
	DisassociateVPCFromHostedZoneAsync(ctx workflow.Context, input *route53.DisassociateVPCFromHostedZoneInput) *DisassociateVPCFromHostedZoneFuture

	EnableHostedZoneDNSSEC(ctx workflow.Context, input *route53.EnableHostedZoneDNSSECInput) (*route53.EnableHostedZoneDNSSECOutput, error)
	EnableHostedZoneDNSSECAsync(ctx workflow.Context, input *route53.EnableHostedZoneDNSSECInput) *EnableHostedZoneDNSSECFuture

	GetAccountLimit(ctx workflow.Context, input *route53.GetAccountLimitInput) (*route53.GetAccountLimitOutput, error)
	GetAccountLimitAsync(ctx workflow.Context, input *route53.GetAccountLimitInput) *GetAccountLimitFuture

	GetChange(ctx workflow.Context, input *route53.GetChangeInput) (*route53.GetChangeOutput, error)
	GetChangeAsync(ctx workflow.Context, input *route53.GetChangeInput) *GetChangeFuture

	GetCheckerIpRanges(ctx workflow.Context, input *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error)
	GetCheckerIpRangesAsync(ctx workflow.Context, input *route53.GetCheckerIpRangesInput) *GetCheckerIpRangesFuture

	GetDNSSEC(ctx workflow.Context, input *route53.GetDNSSECInput) (*route53.GetDNSSECOutput, error)
	GetDNSSECAsync(ctx workflow.Context, input *route53.GetDNSSECInput) *GetDNSSECFuture

	GetGeoLocation(ctx workflow.Context, input *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error)
	GetGeoLocationAsync(ctx workflow.Context, input *route53.GetGeoLocationInput) *GetGeoLocationFuture

	GetHealthCheck(ctx workflow.Context, input *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error)
	GetHealthCheckAsync(ctx workflow.Context, input *route53.GetHealthCheckInput) *GetHealthCheckFuture

	GetHealthCheckCount(ctx workflow.Context, input *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error)
	GetHealthCheckCountAsync(ctx workflow.Context, input *route53.GetHealthCheckCountInput) *GetHealthCheckCountFuture

	GetHealthCheckLastFailureReason(ctx workflow.Context, input *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error)
	GetHealthCheckLastFailureReasonAsync(ctx workflow.Context, input *route53.GetHealthCheckLastFailureReasonInput) *GetHealthCheckLastFailureReasonFuture

	GetHealthCheckStatus(ctx workflow.Context, input *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error)
	GetHealthCheckStatusAsync(ctx workflow.Context, input *route53.GetHealthCheckStatusInput) *GetHealthCheckStatusFuture

	GetHostedZone(ctx workflow.Context, input *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error)
	GetHostedZoneAsync(ctx workflow.Context, input *route53.GetHostedZoneInput) *GetHostedZoneFuture

	GetHostedZoneCount(ctx workflow.Context, input *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error)
	GetHostedZoneCountAsync(ctx workflow.Context, input *route53.GetHostedZoneCountInput) *GetHostedZoneCountFuture

	GetHostedZoneLimit(ctx workflow.Context, input *route53.GetHostedZoneLimitInput) (*route53.GetHostedZoneLimitOutput, error)
	GetHostedZoneLimitAsync(ctx workflow.Context, input *route53.GetHostedZoneLimitInput) *GetHostedZoneLimitFuture

	GetQueryLoggingConfig(ctx workflow.Context, input *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error)
	GetQueryLoggingConfigAsync(ctx workflow.Context, input *route53.GetQueryLoggingConfigInput) *GetQueryLoggingConfigFuture

	GetReusableDelegationSet(ctx workflow.Context, input *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error)
	GetReusableDelegationSetAsync(ctx workflow.Context, input *route53.GetReusableDelegationSetInput) *GetReusableDelegationSetFuture

	GetReusableDelegationSetLimit(ctx workflow.Context, input *route53.GetReusableDelegationSetLimitInput) (*route53.GetReusableDelegationSetLimitOutput, error)
	GetReusableDelegationSetLimitAsync(ctx workflow.Context, input *route53.GetReusableDelegationSetLimitInput) *GetReusableDelegationSetLimitFuture

	GetTrafficPolicy(ctx workflow.Context, input *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error)
	GetTrafficPolicyAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInput) *GetTrafficPolicyFuture

	GetTrafficPolicyInstance(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error)
	GetTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceInput) *GetTrafficPolicyInstanceFuture

	GetTrafficPolicyInstanceCount(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error)
	GetTrafficPolicyInstanceCountAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceCountInput) *GetTrafficPolicyInstanceCountFuture

	ListGeoLocations(ctx workflow.Context, input *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error)
	ListGeoLocationsAsync(ctx workflow.Context, input *route53.ListGeoLocationsInput) *ListGeoLocationsFuture

	ListHealthChecks(ctx workflow.Context, input *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error)
	ListHealthChecksAsync(ctx workflow.Context, input *route53.ListHealthChecksInput) *ListHealthChecksFuture

	ListHostedZones(ctx workflow.Context, input *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error)
	ListHostedZonesAsync(ctx workflow.Context, input *route53.ListHostedZonesInput) *ListHostedZonesFuture

	ListHostedZonesByName(ctx workflow.Context, input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error)
	ListHostedZonesByNameAsync(ctx workflow.Context, input *route53.ListHostedZonesByNameInput) *ListHostedZonesByNameFuture

	ListHostedZonesByVPC(ctx workflow.Context, input *route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error)
	ListHostedZonesByVPCAsync(ctx workflow.Context, input *route53.ListHostedZonesByVPCInput) *ListHostedZonesByVPCFuture

	ListQueryLoggingConfigs(ctx workflow.Context, input *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error)
	ListQueryLoggingConfigsAsync(ctx workflow.Context, input *route53.ListQueryLoggingConfigsInput) *ListQueryLoggingConfigsFuture

	ListResourceRecordSets(ctx workflow.Context, input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error)
	ListResourceRecordSetsAsync(ctx workflow.Context, input *route53.ListResourceRecordSetsInput) *ListResourceRecordSetsFuture

	ListReusableDelegationSets(ctx workflow.Context, input *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error)
	ListReusableDelegationSetsAsync(ctx workflow.Context, input *route53.ListReusableDelegationSetsInput) *ListReusableDelegationSetsFuture

	ListTagsForResource(ctx workflow.Context, input *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *route53.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListTagsForResources(ctx workflow.Context, input *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error)
	ListTagsForResourcesAsync(ctx workflow.Context, input *route53.ListTagsForResourcesInput) *ListTagsForResourcesFuture

	ListTrafficPolicies(ctx workflow.Context, input *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error)
	ListTrafficPoliciesAsync(ctx workflow.Context, input *route53.ListTrafficPoliciesInput) *ListTrafficPoliciesFuture

	ListTrafficPolicyInstances(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error)
	ListTrafficPolicyInstancesAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesInput) *ListTrafficPolicyInstancesFuture

	ListTrafficPolicyInstancesByHostedZone(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
	ListTrafficPolicyInstancesByHostedZoneAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput) *ListTrafficPolicyInstancesByHostedZoneFuture

	ListTrafficPolicyInstancesByPolicy(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error)
	ListTrafficPolicyInstancesByPolicyAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput) *ListTrafficPolicyInstancesByPolicyFuture

	ListTrafficPolicyVersions(ctx workflow.Context, input *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error)
	ListTrafficPolicyVersionsAsync(ctx workflow.Context, input *route53.ListTrafficPolicyVersionsInput) *ListTrafficPolicyVersionsFuture

	ListVPCAssociationAuthorizations(ctx workflow.Context, input *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error)
	ListVPCAssociationAuthorizationsAsync(ctx workflow.Context, input *route53.ListVPCAssociationAuthorizationsInput) *ListVPCAssociationAuthorizationsFuture

	TestDNSAnswer(ctx workflow.Context, input *route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error)
	TestDNSAnswerAsync(ctx workflow.Context, input *route53.TestDNSAnswerInput) *TestDNSAnswerFuture

	UpdateHealthCheck(ctx workflow.Context, input *route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error)
	UpdateHealthCheckAsync(ctx workflow.Context, input *route53.UpdateHealthCheckInput) *UpdateHealthCheckFuture

	UpdateHostedZoneComment(ctx workflow.Context, input *route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error)
	UpdateHostedZoneCommentAsync(ctx workflow.Context, input *route53.UpdateHostedZoneCommentInput) *UpdateHostedZoneCommentFuture

	UpdateTrafficPolicyComment(ctx workflow.Context, input *route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error)
	UpdateTrafficPolicyCommentAsync(ctx workflow.Context, input *route53.UpdateTrafficPolicyCommentInput) *UpdateTrafficPolicyCommentFuture

	UpdateTrafficPolicyInstance(ctx workflow.Context, input *route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error)
	UpdateTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.UpdateTrafficPolicyInstanceInput) *UpdateTrafficPolicyInstanceFuture

	WaitUntilResourceRecordSetsChanged(ctx workflow.Context, input *route53.GetChangeInput) error
	WaitUntilResourceRecordSetsChangedAsync(ctx workflow.Context, input *route53.GetChangeInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
