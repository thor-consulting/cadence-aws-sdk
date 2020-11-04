// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package macie2stub

import (
	"github.com/aws/aws-sdk-go/service/macie2"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptInvitation(ctx workflow.Context, input *macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error)
	AcceptInvitationAsync(ctx workflow.Context, input *macie2.AcceptInvitationInput) *Macie2AcceptInvitationFuture

	BatchGetCustomDataIdentifiers(ctx workflow.Context, input *macie2.BatchGetCustomDataIdentifiersInput) (*macie2.BatchGetCustomDataIdentifiersOutput, error)
	BatchGetCustomDataIdentifiersAsync(ctx workflow.Context, input *macie2.BatchGetCustomDataIdentifiersInput) *Macie2BatchGetCustomDataIdentifiersFuture

	CreateClassificationJob(ctx workflow.Context, input *macie2.CreateClassificationJobInput) (*macie2.CreateClassificationJobOutput, error)
	CreateClassificationJobAsync(ctx workflow.Context, input *macie2.CreateClassificationJobInput) *Macie2CreateClassificationJobFuture

	CreateCustomDataIdentifier(ctx workflow.Context, input *macie2.CreateCustomDataIdentifierInput) (*macie2.CreateCustomDataIdentifierOutput, error)
	CreateCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.CreateCustomDataIdentifierInput) *Macie2CreateCustomDataIdentifierFuture

	CreateFindingsFilter(ctx workflow.Context, input *macie2.CreateFindingsFilterInput) (*macie2.CreateFindingsFilterOutput, error)
	CreateFindingsFilterAsync(ctx workflow.Context, input *macie2.CreateFindingsFilterInput) *Macie2CreateFindingsFilterFuture

	CreateInvitations(ctx workflow.Context, input *macie2.CreateInvitationsInput) (*macie2.CreateInvitationsOutput, error)
	CreateInvitationsAsync(ctx workflow.Context, input *macie2.CreateInvitationsInput) *Macie2CreateInvitationsFuture

	CreateMember(ctx workflow.Context, input *macie2.CreateMemberInput) (*macie2.CreateMemberOutput, error)
	CreateMemberAsync(ctx workflow.Context, input *macie2.CreateMemberInput) *Macie2CreateMemberFuture

	CreateSampleFindings(ctx workflow.Context, input *macie2.CreateSampleFindingsInput) (*macie2.CreateSampleFindingsOutput, error)
	CreateSampleFindingsAsync(ctx workflow.Context, input *macie2.CreateSampleFindingsInput) *Macie2CreateSampleFindingsFuture

	DeclineInvitations(ctx workflow.Context, input *macie2.DeclineInvitationsInput) (*macie2.DeclineInvitationsOutput, error)
	DeclineInvitationsAsync(ctx workflow.Context, input *macie2.DeclineInvitationsInput) *Macie2DeclineInvitationsFuture

	DeleteCustomDataIdentifier(ctx workflow.Context, input *macie2.DeleteCustomDataIdentifierInput) (*macie2.DeleteCustomDataIdentifierOutput, error)
	DeleteCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.DeleteCustomDataIdentifierInput) *Macie2DeleteCustomDataIdentifierFuture

	DeleteFindingsFilter(ctx workflow.Context, input *macie2.DeleteFindingsFilterInput) (*macie2.DeleteFindingsFilterOutput, error)
	DeleteFindingsFilterAsync(ctx workflow.Context, input *macie2.DeleteFindingsFilterInput) *Macie2DeleteFindingsFilterFuture

	DeleteInvitations(ctx workflow.Context, input *macie2.DeleteInvitationsInput) (*macie2.DeleteInvitationsOutput, error)
	DeleteInvitationsAsync(ctx workflow.Context, input *macie2.DeleteInvitationsInput) *Macie2DeleteInvitationsFuture

	DeleteMember(ctx workflow.Context, input *macie2.DeleteMemberInput) (*macie2.DeleteMemberOutput, error)
	DeleteMemberAsync(ctx workflow.Context, input *macie2.DeleteMemberInput) *Macie2DeleteMemberFuture

	DescribeBuckets(ctx workflow.Context, input *macie2.DescribeBucketsInput) (*macie2.DescribeBucketsOutput, error)
	DescribeBucketsAsync(ctx workflow.Context, input *macie2.DescribeBucketsInput) *Macie2DescribeBucketsFuture

	DescribeClassificationJob(ctx workflow.Context, input *macie2.DescribeClassificationJobInput) (*macie2.DescribeClassificationJobOutput, error)
	DescribeClassificationJobAsync(ctx workflow.Context, input *macie2.DescribeClassificationJobInput) *Macie2DescribeClassificationJobFuture

	DescribeOrganizationConfiguration(ctx workflow.Context, input *macie2.DescribeOrganizationConfigurationInput) (*macie2.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationAsync(ctx workflow.Context, input *macie2.DescribeOrganizationConfigurationInput) *Macie2DescribeOrganizationConfigurationFuture

	DisableMacie(ctx workflow.Context, input *macie2.DisableMacieInput) (*macie2.DisableMacieOutput, error)
	DisableMacieAsync(ctx workflow.Context, input *macie2.DisableMacieInput) *Macie2DisableMacieFuture

	DisableOrganizationAdminAccount(ctx workflow.Context, input *macie2.DisableOrganizationAdminAccountInput) (*macie2.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountAsync(ctx workflow.Context, input *macie2.DisableOrganizationAdminAccountInput) *Macie2DisableOrganizationAdminAccountFuture

	DisassociateFromMasterAccount(ctx workflow.Context, input *macie2.DisassociateFromMasterAccountInput) (*macie2.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountAsync(ctx workflow.Context, input *macie2.DisassociateFromMasterAccountInput) *Macie2DisassociateFromMasterAccountFuture

	DisassociateMember(ctx workflow.Context, input *macie2.DisassociateMemberInput) (*macie2.DisassociateMemberOutput, error)
	DisassociateMemberAsync(ctx workflow.Context, input *macie2.DisassociateMemberInput) *Macie2DisassociateMemberFuture

	EnableMacie(ctx workflow.Context, input *macie2.EnableMacieInput) (*macie2.EnableMacieOutput, error)
	EnableMacieAsync(ctx workflow.Context, input *macie2.EnableMacieInput) *Macie2EnableMacieFuture

	EnableOrganizationAdminAccount(ctx workflow.Context, input *macie2.EnableOrganizationAdminAccountInput) (*macie2.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountAsync(ctx workflow.Context, input *macie2.EnableOrganizationAdminAccountInput) *Macie2EnableOrganizationAdminAccountFuture

	GetBucketStatistics(ctx workflow.Context, input *macie2.GetBucketStatisticsInput) (*macie2.GetBucketStatisticsOutput, error)
	GetBucketStatisticsAsync(ctx workflow.Context, input *macie2.GetBucketStatisticsInput) *Macie2GetBucketStatisticsFuture

	GetClassificationExportConfiguration(ctx workflow.Context, input *macie2.GetClassificationExportConfigurationInput) (*macie2.GetClassificationExportConfigurationOutput, error)
	GetClassificationExportConfigurationAsync(ctx workflow.Context, input *macie2.GetClassificationExportConfigurationInput) *Macie2GetClassificationExportConfigurationFuture

	GetCustomDataIdentifier(ctx workflow.Context, input *macie2.GetCustomDataIdentifierInput) (*macie2.GetCustomDataIdentifierOutput, error)
	GetCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.GetCustomDataIdentifierInput) *Macie2GetCustomDataIdentifierFuture

	GetFindingStatistics(ctx workflow.Context, input *macie2.GetFindingStatisticsInput) (*macie2.GetFindingStatisticsOutput, error)
	GetFindingStatisticsAsync(ctx workflow.Context, input *macie2.GetFindingStatisticsInput) *Macie2GetFindingStatisticsFuture

	GetFindings(ctx workflow.Context, input *macie2.GetFindingsInput) (*macie2.GetFindingsOutput, error)
	GetFindingsAsync(ctx workflow.Context, input *macie2.GetFindingsInput) *Macie2GetFindingsFuture

	GetFindingsFilter(ctx workflow.Context, input *macie2.GetFindingsFilterInput) (*macie2.GetFindingsFilterOutput, error)
	GetFindingsFilterAsync(ctx workflow.Context, input *macie2.GetFindingsFilterInput) *Macie2GetFindingsFilterFuture

	GetInvitationsCount(ctx workflow.Context, input *macie2.GetInvitationsCountInput) (*macie2.GetInvitationsCountOutput, error)
	GetInvitationsCountAsync(ctx workflow.Context, input *macie2.GetInvitationsCountInput) *Macie2GetInvitationsCountFuture

	GetMacieSession(ctx workflow.Context, input *macie2.GetMacieSessionInput) (*macie2.GetMacieSessionOutput, error)
	GetMacieSessionAsync(ctx workflow.Context, input *macie2.GetMacieSessionInput) *Macie2GetMacieSessionFuture

	GetMasterAccount(ctx workflow.Context, input *macie2.GetMasterAccountInput) (*macie2.GetMasterAccountOutput, error)
	GetMasterAccountAsync(ctx workflow.Context, input *macie2.GetMasterAccountInput) *Macie2GetMasterAccountFuture

	GetMember(ctx workflow.Context, input *macie2.GetMemberInput) (*macie2.GetMemberOutput, error)
	GetMemberAsync(ctx workflow.Context, input *macie2.GetMemberInput) *Macie2GetMemberFuture

	GetUsageStatistics(ctx workflow.Context, input *macie2.GetUsageStatisticsInput) (*macie2.GetUsageStatisticsOutput, error)
	GetUsageStatisticsAsync(ctx workflow.Context, input *macie2.GetUsageStatisticsInput) *Macie2GetUsageStatisticsFuture

	GetUsageTotals(ctx workflow.Context, input *macie2.GetUsageTotalsInput) (*macie2.GetUsageTotalsOutput, error)
	GetUsageTotalsAsync(ctx workflow.Context, input *macie2.GetUsageTotalsInput) *Macie2GetUsageTotalsFuture

	ListClassificationJobs(ctx workflow.Context, input *macie2.ListClassificationJobsInput) (*macie2.ListClassificationJobsOutput, error)
	ListClassificationJobsAsync(ctx workflow.Context, input *macie2.ListClassificationJobsInput) *Macie2ListClassificationJobsFuture

	ListCustomDataIdentifiers(ctx workflow.Context, input *macie2.ListCustomDataIdentifiersInput) (*macie2.ListCustomDataIdentifiersOutput, error)
	ListCustomDataIdentifiersAsync(ctx workflow.Context, input *macie2.ListCustomDataIdentifiersInput) *Macie2ListCustomDataIdentifiersFuture

	ListFindings(ctx workflow.Context, input *macie2.ListFindingsInput) (*macie2.ListFindingsOutput, error)
	ListFindingsAsync(ctx workflow.Context, input *macie2.ListFindingsInput) *Macie2ListFindingsFuture

	ListFindingsFilters(ctx workflow.Context, input *macie2.ListFindingsFiltersInput) (*macie2.ListFindingsFiltersOutput, error)
	ListFindingsFiltersAsync(ctx workflow.Context, input *macie2.ListFindingsFiltersInput) *Macie2ListFindingsFiltersFuture

	ListInvitations(ctx workflow.Context, input *macie2.ListInvitationsInput) (*macie2.ListInvitationsOutput, error)
	ListInvitationsAsync(ctx workflow.Context, input *macie2.ListInvitationsInput) *Macie2ListInvitationsFuture

	ListMembers(ctx workflow.Context, input *macie2.ListMembersInput) (*macie2.ListMembersOutput, error)
	ListMembersAsync(ctx workflow.Context, input *macie2.ListMembersInput) *Macie2ListMembersFuture

	ListOrganizationAdminAccounts(ctx workflow.Context, input *macie2.ListOrganizationAdminAccountsInput) (*macie2.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsAsync(ctx workflow.Context, input *macie2.ListOrganizationAdminAccountsInput) *Macie2ListOrganizationAdminAccountsFuture

	ListTagsForResource(ctx workflow.Context, input *macie2.ListTagsForResourceInput) (*macie2.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *macie2.ListTagsForResourceInput) *Macie2ListTagsForResourceFuture

	PutClassificationExportConfiguration(ctx workflow.Context, input *macie2.PutClassificationExportConfigurationInput) (*macie2.PutClassificationExportConfigurationOutput, error)
	PutClassificationExportConfigurationAsync(ctx workflow.Context, input *macie2.PutClassificationExportConfigurationInput) *Macie2PutClassificationExportConfigurationFuture

	TagResource(ctx workflow.Context, input *macie2.TagResourceInput) (*macie2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *macie2.TagResourceInput) *Macie2TagResourceFuture

	TestCustomDataIdentifier(ctx workflow.Context, input *macie2.TestCustomDataIdentifierInput) (*macie2.TestCustomDataIdentifierOutput, error)
	TestCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.TestCustomDataIdentifierInput) *Macie2TestCustomDataIdentifierFuture

	UntagResource(ctx workflow.Context, input *macie2.UntagResourceInput) (*macie2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *macie2.UntagResourceInput) *Macie2UntagResourceFuture

	UpdateClassificationJob(ctx workflow.Context, input *macie2.UpdateClassificationJobInput) (*macie2.UpdateClassificationJobOutput, error)
	UpdateClassificationJobAsync(ctx workflow.Context, input *macie2.UpdateClassificationJobInput) *Macie2UpdateClassificationJobFuture

	UpdateFindingsFilter(ctx workflow.Context, input *macie2.UpdateFindingsFilterInput) (*macie2.UpdateFindingsFilterOutput, error)
	UpdateFindingsFilterAsync(ctx workflow.Context, input *macie2.UpdateFindingsFilterInput) *Macie2UpdateFindingsFilterFuture

	UpdateMacieSession(ctx workflow.Context, input *macie2.UpdateMacieSessionInput) (*macie2.UpdateMacieSessionOutput, error)
	UpdateMacieSessionAsync(ctx workflow.Context, input *macie2.UpdateMacieSessionInput) *Macie2UpdateMacieSessionFuture

	UpdateMemberSession(ctx workflow.Context, input *macie2.UpdateMemberSessionInput) (*macie2.UpdateMemberSessionOutput, error)
	UpdateMemberSessionAsync(ctx workflow.Context, input *macie2.UpdateMemberSessionInput) *Macie2UpdateMemberSessionFuture

	UpdateOrganizationConfiguration(ctx workflow.Context, input *macie2.UpdateOrganizationConfigurationInput) (*macie2.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationAsync(ctx workflow.Context, input *macie2.UpdateOrganizationConfigurationInput) *Macie2UpdateOrganizationConfigurationFuture
}

func NewClient() Client {
	return &stub{}
}