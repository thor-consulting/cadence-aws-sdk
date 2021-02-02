// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package shieldstub

import (
	"github.com/aws/aws-sdk-go/service/shield"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AssociateDRTLogBucketFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateDRTLogBucketFuture) Get(ctx workflow.Context) (*shield.AssociateDRTLogBucketOutput, error) {
	var output shield.AssociateDRTLogBucketOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AssociateDRTRoleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateDRTRoleFuture) Get(ctx workflow.Context) (*shield.AssociateDRTRoleOutput, error) {
	var output shield.AssociateDRTRoleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AssociateHealthCheckFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateHealthCheckFuture) Get(ctx workflow.Context) (*shield.AssociateHealthCheckOutput, error) {
	var output shield.AssociateHealthCheckOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AssociateProactiveEngagementDetailsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateProactiveEngagementDetailsFuture) Get(ctx workflow.Context) (*shield.AssociateProactiveEngagementDetailsOutput, error) {
	var output shield.AssociateProactiveEngagementDetailsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateProtectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateProtectionFuture) Get(ctx workflow.Context) (*shield.CreateProtectionOutput, error) {
	var output shield.CreateProtectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateProtectionGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateProtectionGroupFuture) Get(ctx workflow.Context) (*shield.CreateProtectionGroupOutput, error) {
	var output shield.CreateProtectionGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateSubscriptionFuture) Get(ctx workflow.Context) (*shield.CreateSubscriptionOutput, error) {
	var output shield.CreateSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteProtectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteProtectionFuture) Get(ctx workflow.Context) (*shield.DeleteProtectionOutput, error) {
	var output shield.DeleteProtectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteProtectionGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteProtectionGroupFuture) Get(ctx workflow.Context) (*shield.DeleteProtectionGroupOutput, error) {
	var output shield.DeleteProtectionGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSubscriptionFuture) Get(ctx workflow.Context) (*shield.DeleteSubscriptionOutput, error) {
	var output shield.DeleteSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAttackFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAttackFuture) Get(ctx workflow.Context) (*shield.DescribeAttackOutput, error) {
	var output shield.DescribeAttackOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAttackStatisticsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAttackStatisticsFuture) Get(ctx workflow.Context) (*shield.DescribeAttackStatisticsOutput, error) {
	var output shield.DescribeAttackStatisticsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeDRTAccessFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeDRTAccessFuture) Get(ctx workflow.Context) (*shield.DescribeDRTAccessOutput, error) {
	var output shield.DescribeDRTAccessOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeEmergencyContactSettingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeEmergencyContactSettingsFuture) Get(ctx workflow.Context) (*shield.DescribeEmergencyContactSettingsOutput, error) {
	var output shield.DescribeEmergencyContactSettingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeProtectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeProtectionFuture) Get(ctx workflow.Context) (*shield.DescribeProtectionOutput, error) {
	var output shield.DescribeProtectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeProtectionGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeProtectionGroupFuture) Get(ctx workflow.Context) (*shield.DescribeProtectionGroupOutput, error) {
	var output shield.DescribeProtectionGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeSubscriptionFuture) Get(ctx workflow.Context) (*shield.DescribeSubscriptionOutput, error) {
	var output shield.DescribeSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisableProactiveEngagementFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisableProactiveEngagementFuture) Get(ctx workflow.Context) (*shield.DisableProactiveEngagementOutput, error) {
	var output shield.DisableProactiveEngagementOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateDRTLogBucketFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateDRTLogBucketFuture) Get(ctx workflow.Context) (*shield.DisassociateDRTLogBucketOutput, error) {
	var output shield.DisassociateDRTLogBucketOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateDRTRoleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateDRTRoleFuture) Get(ctx workflow.Context) (*shield.DisassociateDRTRoleOutput, error) {
	var output shield.DisassociateDRTRoleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateHealthCheckFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateHealthCheckFuture) Get(ctx workflow.Context) (*shield.DisassociateHealthCheckOutput, error) {
	var output shield.DisassociateHealthCheckOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EnableProactiveEngagementFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EnableProactiveEngagementFuture) Get(ctx workflow.Context) (*shield.EnableProactiveEngagementOutput, error) {
	var output shield.EnableProactiveEngagementOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetSubscriptionStateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetSubscriptionStateFuture) Get(ctx workflow.Context) (*shield.GetSubscriptionStateOutput, error) {
	var output shield.GetSubscriptionStateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAttacksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAttacksFuture) Get(ctx workflow.Context) (*shield.ListAttacksOutput, error) {
	var output shield.ListAttacksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListProtectionGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListProtectionGroupsFuture) Get(ctx workflow.Context) (*shield.ListProtectionGroupsOutput, error) {
	var output shield.ListProtectionGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListProtectionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListProtectionsFuture) Get(ctx workflow.Context) (*shield.ListProtectionsOutput, error) {
	var output shield.ListProtectionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListResourcesInProtectionGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListResourcesInProtectionGroupFuture) Get(ctx workflow.Context) (*shield.ListResourcesInProtectionGroupOutput, error) {
	var output shield.ListResourcesInProtectionGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateEmergencyContactSettingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateEmergencyContactSettingsFuture) Get(ctx workflow.Context) (*shield.UpdateEmergencyContactSettingsOutput, error) {
	var output shield.UpdateEmergencyContactSettingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateProtectionGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateProtectionGroupFuture) Get(ctx workflow.Context) (*shield.UpdateProtectionGroupOutput, error) {
	var output shield.UpdateProtectionGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateSubscriptionFuture) Get(ctx workflow.Context) (*shield.UpdateSubscriptionOutput, error) {
	var output shield.UpdateSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateDRTLogBucket(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error) {
	var output shield.AssociateDRTLogBucketOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-AssociateDRTLogBucket", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateDRTLogBucketAsync(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) *AssociateDRTLogBucketFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-AssociateDRTLogBucket", input)
	return &AssociateDRTLogBucketFuture{Future: future}
}

func (a *stub) AssociateDRTRole(ctx workflow.Context, input *shield.AssociateDRTRoleInput) (*shield.AssociateDRTRoleOutput, error) {
	var output shield.AssociateDRTRoleOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-AssociateDRTRole", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateDRTRoleAsync(ctx workflow.Context, input *shield.AssociateDRTRoleInput) *AssociateDRTRoleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-AssociateDRTRole", input)
	return &AssociateDRTRoleFuture{Future: future}
}

func (a *stub) AssociateHealthCheck(ctx workflow.Context, input *shield.AssociateHealthCheckInput) (*shield.AssociateHealthCheckOutput, error) {
	var output shield.AssociateHealthCheckOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-AssociateHealthCheck", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateHealthCheckAsync(ctx workflow.Context, input *shield.AssociateHealthCheckInput) *AssociateHealthCheckFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-AssociateHealthCheck", input)
	return &AssociateHealthCheckFuture{Future: future}
}

func (a *stub) AssociateProactiveEngagementDetails(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) (*shield.AssociateProactiveEngagementDetailsOutput, error) {
	var output shield.AssociateProactiveEngagementDetailsOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-AssociateProactiveEngagementDetails", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateProactiveEngagementDetailsAsync(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) *AssociateProactiveEngagementDetailsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-AssociateProactiveEngagementDetails", input)
	return &AssociateProactiveEngagementDetailsFuture{Future: future}
}

func (a *stub) CreateProtection(ctx workflow.Context, input *shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error) {
	var output shield.CreateProtectionOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-CreateProtection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProtectionAsync(ctx workflow.Context, input *shield.CreateProtectionInput) *CreateProtectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-CreateProtection", input)
	return &CreateProtectionFuture{Future: future}
}

func (a *stub) CreateProtectionGroup(ctx workflow.Context, input *shield.CreateProtectionGroupInput) (*shield.CreateProtectionGroupOutput, error) {
	var output shield.CreateProtectionGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-CreateProtectionGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProtectionGroupAsync(ctx workflow.Context, input *shield.CreateProtectionGroupInput) *CreateProtectionGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-CreateProtectionGroup", input)
	return &CreateProtectionGroupFuture{Future: future}
}

func (a *stub) CreateSubscription(ctx workflow.Context, input *shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error) {
	var output shield.CreateSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-CreateSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSubscriptionAsync(ctx workflow.Context, input *shield.CreateSubscriptionInput) *CreateSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-CreateSubscription", input)
	return &CreateSubscriptionFuture{Future: future}
}

func (a *stub) DeleteProtection(ctx workflow.Context, input *shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error) {
	var output shield.DeleteProtectionOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DeleteProtection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProtectionAsync(ctx workflow.Context, input *shield.DeleteProtectionInput) *DeleteProtectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DeleteProtection", input)
	return &DeleteProtectionFuture{Future: future}
}

func (a *stub) DeleteProtectionGroup(ctx workflow.Context, input *shield.DeleteProtectionGroupInput) (*shield.DeleteProtectionGroupOutput, error) {
	var output shield.DeleteProtectionGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DeleteProtectionGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProtectionGroupAsync(ctx workflow.Context, input *shield.DeleteProtectionGroupInput) *DeleteProtectionGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DeleteProtectionGroup", input)
	return &DeleteProtectionGroupFuture{Future: future}
}

func (a *stub) DeleteSubscription(ctx workflow.Context, input *shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error) {
	var output shield.DeleteSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DeleteSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSubscriptionAsync(ctx workflow.Context, input *shield.DeleteSubscriptionInput) *DeleteSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DeleteSubscription", input)
	return &DeleteSubscriptionFuture{Future: future}
}

func (a *stub) DescribeAttack(ctx workflow.Context, input *shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error) {
	var output shield.DescribeAttackOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DescribeAttack", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAttackAsync(ctx workflow.Context, input *shield.DescribeAttackInput) *DescribeAttackFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DescribeAttack", input)
	return &DescribeAttackFuture{Future: future}
}

func (a *stub) DescribeAttackStatistics(ctx workflow.Context, input *shield.DescribeAttackStatisticsInput) (*shield.DescribeAttackStatisticsOutput, error) {
	var output shield.DescribeAttackStatisticsOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DescribeAttackStatistics", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAttackStatisticsAsync(ctx workflow.Context, input *shield.DescribeAttackStatisticsInput) *DescribeAttackStatisticsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DescribeAttackStatistics", input)
	return &DescribeAttackStatisticsFuture{Future: future}
}

func (a *stub) DescribeDRTAccess(ctx workflow.Context, input *shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error) {
	var output shield.DescribeDRTAccessOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DescribeDRTAccess", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDRTAccessAsync(ctx workflow.Context, input *shield.DescribeDRTAccessInput) *DescribeDRTAccessFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DescribeDRTAccess", input)
	return &DescribeDRTAccessFuture{Future: future}
}

func (a *stub) DescribeEmergencyContactSettings(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error) {
	var output shield.DescribeEmergencyContactSettingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DescribeEmergencyContactSettings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) *DescribeEmergencyContactSettingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DescribeEmergencyContactSettings", input)
	return &DescribeEmergencyContactSettingsFuture{Future: future}
}

func (a *stub) DescribeProtection(ctx workflow.Context, input *shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error) {
	var output shield.DescribeProtectionOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DescribeProtection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProtectionAsync(ctx workflow.Context, input *shield.DescribeProtectionInput) *DescribeProtectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DescribeProtection", input)
	return &DescribeProtectionFuture{Future: future}
}

func (a *stub) DescribeProtectionGroup(ctx workflow.Context, input *shield.DescribeProtectionGroupInput) (*shield.DescribeProtectionGroupOutput, error) {
	var output shield.DescribeProtectionGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DescribeProtectionGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProtectionGroupAsync(ctx workflow.Context, input *shield.DescribeProtectionGroupInput) *DescribeProtectionGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DescribeProtectionGroup", input)
	return &DescribeProtectionGroupFuture{Future: future}
}

func (a *stub) DescribeSubscription(ctx workflow.Context, input *shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error) {
	var output shield.DescribeSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DescribeSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSubscriptionAsync(ctx workflow.Context, input *shield.DescribeSubscriptionInput) *DescribeSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DescribeSubscription", input)
	return &DescribeSubscriptionFuture{Future: future}
}

func (a *stub) DisableProactiveEngagement(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) (*shield.DisableProactiveEngagementOutput, error) {
	var output shield.DisableProactiveEngagementOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DisableProactiveEngagement", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableProactiveEngagementAsync(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) *DisableProactiveEngagementFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DisableProactiveEngagement", input)
	return &DisableProactiveEngagementFuture{Future: future}
}

func (a *stub) DisassociateDRTLogBucket(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) (*shield.DisassociateDRTLogBucketOutput, error) {
	var output shield.DisassociateDRTLogBucketOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DisassociateDRTLogBucket", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateDRTLogBucketAsync(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) *DisassociateDRTLogBucketFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DisassociateDRTLogBucket", input)
	return &DisassociateDRTLogBucketFuture{Future: future}
}

func (a *stub) DisassociateDRTRole(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) (*shield.DisassociateDRTRoleOutput, error) {
	var output shield.DisassociateDRTRoleOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DisassociateDRTRole", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateDRTRoleAsync(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) *DisassociateDRTRoleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DisassociateDRTRole", input)
	return &DisassociateDRTRoleFuture{Future: future}
}

func (a *stub) DisassociateHealthCheck(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) (*shield.DisassociateHealthCheckOutput, error) {
	var output shield.DisassociateHealthCheckOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-DisassociateHealthCheck", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateHealthCheckAsync(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) *DisassociateHealthCheckFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-DisassociateHealthCheck", input)
	return &DisassociateHealthCheckFuture{Future: future}
}

func (a *stub) EnableProactiveEngagement(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) (*shield.EnableProactiveEngagementOutput, error) {
	var output shield.EnableProactiveEngagementOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-EnableProactiveEngagement", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableProactiveEngagementAsync(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) *EnableProactiveEngagementFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-EnableProactiveEngagement", input)
	return &EnableProactiveEngagementFuture{Future: future}
}

func (a *stub) GetSubscriptionState(ctx workflow.Context, input *shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error) {
	var output shield.GetSubscriptionStateOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-GetSubscriptionState", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSubscriptionStateAsync(ctx workflow.Context, input *shield.GetSubscriptionStateInput) *GetSubscriptionStateFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-GetSubscriptionState", input)
	return &GetSubscriptionStateFuture{Future: future}
}

func (a *stub) ListAttacks(ctx workflow.Context, input *shield.ListAttacksInput) (*shield.ListAttacksOutput, error) {
	var output shield.ListAttacksOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-ListAttacks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAttacksAsync(ctx workflow.Context, input *shield.ListAttacksInput) *ListAttacksFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-ListAttacks", input)
	return &ListAttacksFuture{Future: future}
}

func (a *stub) ListProtectionGroups(ctx workflow.Context, input *shield.ListProtectionGroupsInput) (*shield.ListProtectionGroupsOutput, error) {
	var output shield.ListProtectionGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-ListProtectionGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProtectionGroupsAsync(ctx workflow.Context, input *shield.ListProtectionGroupsInput) *ListProtectionGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-ListProtectionGroups", input)
	return &ListProtectionGroupsFuture{Future: future}
}

func (a *stub) ListProtections(ctx workflow.Context, input *shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error) {
	var output shield.ListProtectionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-ListProtections", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProtectionsAsync(ctx workflow.Context, input *shield.ListProtectionsInput) *ListProtectionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-ListProtections", input)
	return &ListProtectionsFuture{Future: future}
}

func (a *stub) ListResourcesInProtectionGroup(ctx workflow.Context, input *shield.ListResourcesInProtectionGroupInput) (*shield.ListResourcesInProtectionGroupOutput, error) {
	var output shield.ListResourcesInProtectionGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-ListResourcesInProtectionGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResourcesInProtectionGroupAsync(ctx workflow.Context, input *shield.ListResourcesInProtectionGroupInput) *ListResourcesInProtectionGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-ListResourcesInProtectionGroup", input)
	return &ListResourcesInProtectionGroupFuture{Future: future}
}

func (a *stub) UpdateEmergencyContactSettings(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) (*shield.UpdateEmergencyContactSettingsOutput, error) {
	var output shield.UpdateEmergencyContactSettingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-UpdateEmergencyContactSettings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) *UpdateEmergencyContactSettingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-UpdateEmergencyContactSettings", input)
	return &UpdateEmergencyContactSettingsFuture{Future: future}
}

func (a *stub) UpdateProtectionGroup(ctx workflow.Context, input *shield.UpdateProtectionGroupInput) (*shield.UpdateProtectionGroupOutput, error) {
	var output shield.UpdateProtectionGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-UpdateProtectionGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateProtectionGroupAsync(ctx workflow.Context, input *shield.UpdateProtectionGroupInput) *UpdateProtectionGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-UpdateProtectionGroup", input)
	return &UpdateProtectionGroupFuture{Future: future}
}

func (a *stub) UpdateSubscription(ctx workflow.Context, input *shield.UpdateSubscriptionInput) (*shield.UpdateSubscriptionOutput, error) {
	var output shield.UpdateSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws-shield-UpdateSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateSubscriptionAsync(ctx workflow.Context, input *shield.UpdateSubscriptionInput) *UpdateSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-shield-UpdateSubscription", input)
	return &UpdateSubscriptionFuture{Future: future}
}
