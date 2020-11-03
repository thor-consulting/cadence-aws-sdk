// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package workspacesstub

import (
	"github.com/aws/aws-sdk-go/service/workspaces"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateConnectionAlias(ctx workflow.Context, input *workspaces.AssociateConnectionAliasInput) (*workspaces.AssociateConnectionAliasOutput, error)
	AssociateConnectionAliasAsync(ctx workflow.Context, input *workspaces.AssociateConnectionAliasInput) *WorkSpacesAssociateConnectionAliasFuture

	AssociateIpGroups(ctx workflow.Context, input *workspaces.AssociateIpGroupsInput) (*workspaces.AssociateIpGroupsOutput, error)
	AssociateIpGroupsAsync(ctx workflow.Context, input *workspaces.AssociateIpGroupsInput) *WorkSpacesAssociateIpGroupsFuture

	AuthorizeIpRules(ctx workflow.Context, input *workspaces.AuthorizeIpRulesInput) (*workspaces.AuthorizeIpRulesOutput, error)
	AuthorizeIpRulesAsync(ctx workflow.Context, input *workspaces.AuthorizeIpRulesInput) *WorkSpacesAuthorizeIpRulesFuture

	CopyWorkspaceImage(ctx workflow.Context, input *workspaces.CopyWorkspaceImageInput) (*workspaces.CopyWorkspaceImageOutput, error)
	CopyWorkspaceImageAsync(ctx workflow.Context, input *workspaces.CopyWorkspaceImageInput) *WorkSpacesCopyWorkspaceImageFuture

	CreateConnectionAlias(ctx workflow.Context, input *workspaces.CreateConnectionAliasInput) (*workspaces.CreateConnectionAliasOutput, error)
	CreateConnectionAliasAsync(ctx workflow.Context, input *workspaces.CreateConnectionAliasInput) *WorkSpacesCreateConnectionAliasFuture

	CreateIpGroup(ctx workflow.Context, input *workspaces.CreateIpGroupInput) (*workspaces.CreateIpGroupOutput, error)
	CreateIpGroupAsync(ctx workflow.Context, input *workspaces.CreateIpGroupInput) *WorkSpacesCreateIpGroupFuture

	CreateTags(ctx workflow.Context, input *workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error)
	CreateTagsAsync(ctx workflow.Context, input *workspaces.CreateTagsInput) *WorkSpacesCreateTagsFuture

	CreateWorkspaces(ctx workflow.Context, input *workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error)
	CreateWorkspacesAsync(ctx workflow.Context, input *workspaces.CreateWorkspacesInput) *WorkSpacesCreateWorkspacesFuture

	DeleteConnectionAlias(ctx workflow.Context, input *workspaces.DeleteConnectionAliasInput) (*workspaces.DeleteConnectionAliasOutput, error)
	DeleteConnectionAliasAsync(ctx workflow.Context, input *workspaces.DeleteConnectionAliasInput) *WorkSpacesDeleteConnectionAliasFuture

	DeleteIpGroup(ctx workflow.Context, input *workspaces.DeleteIpGroupInput) (*workspaces.DeleteIpGroupOutput, error)
	DeleteIpGroupAsync(ctx workflow.Context, input *workspaces.DeleteIpGroupInput) *WorkSpacesDeleteIpGroupFuture

	DeleteTags(ctx workflow.Context, input *workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *workspaces.DeleteTagsInput) *WorkSpacesDeleteTagsFuture

	DeleteWorkspaceImage(ctx workflow.Context, input *workspaces.DeleteWorkspaceImageInput) (*workspaces.DeleteWorkspaceImageOutput, error)
	DeleteWorkspaceImageAsync(ctx workflow.Context, input *workspaces.DeleteWorkspaceImageInput) *WorkSpacesDeleteWorkspaceImageFuture

	DeregisterWorkspaceDirectory(ctx workflow.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) (*workspaces.DeregisterWorkspaceDirectoryOutput, error)
	DeregisterWorkspaceDirectoryAsync(ctx workflow.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) *WorkSpacesDeregisterWorkspaceDirectoryFuture

	DescribeAccount(ctx workflow.Context, input *workspaces.DescribeAccountInput) (*workspaces.DescribeAccountOutput, error)
	DescribeAccountAsync(ctx workflow.Context, input *workspaces.DescribeAccountInput) *WorkSpacesDescribeAccountFuture

	DescribeAccountModifications(ctx workflow.Context, input *workspaces.DescribeAccountModificationsInput) (*workspaces.DescribeAccountModificationsOutput, error)
	DescribeAccountModificationsAsync(ctx workflow.Context, input *workspaces.DescribeAccountModificationsInput) *WorkSpacesDescribeAccountModificationsFuture

	DescribeClientProperties(ctx workflow.Context, input *workspaces.DescribeClientPropertiesInput) (*workspaces.DescribeClientPropertiesOutput, error)
	DescribeClientPropertiesAsync(ctx workflow.Context, input *workspaces.DescribeClientPropertiesInput) *WorkSpacesDescribeClientPropertiesFuture

	DescribeConnectionAliasPermissions(ctx workflow.Context, input *workspaces.DescribeConnectionAliasPermissionsInput) (*workspaces.DescribeConnectionAliasPermissionsOutput, error)
	DescribeConnectionAliasPermissionsAsync(ctx workflow.Context, input *workspaces.DescribeConnectionAliasPermissionsInput) *WorkSpacesDescribeConnectionAliasPermissionsFuture

	DescribeConnectionAliases(ctx workflow.Context, input *workspaces.DescribeConnectionAliasesInput) (*workspaces.DescribeConnectionAliasesOutput, error)
	DescribeConnectionAliasesAsync(ctx workflow.Context, input *workspaces.DescribeConnectionAliasesInput) *WorkSpacesDescribeConnectionAliasesFuture

	DescribeIpGroups(ctx workflow.Context, input *workspaces.DescribeIpGroupsInput) (*workspaces.DescribeIpGroupsOutput, error)
	DescribeIpGroupsAsync(ctx workflow.Context, input *workspaces.DescribeIpGroupsInput) *WorkSpacesDescribeIpGroupsFuture

	DescribeTags(ctx workflow.Context, input *workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *workspaces.DescribeTagsInput) *WorkSpacesDescribeTagsFuture

	DescribeWorkspaceBundles(ctx workflow.Context, input *workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error)
	DescribeWorkspaceBundlesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceBundlesInput) *WorkSpacesDescribeWorkspaceBundlesFuture

	DescribeWorkspaceDirectories(ctx workflow.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
	DescribeWorkspaceDirectoriesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) *WorkSpacesDescribeWorkspaceDirectoriesFuture

	DescribeWorkspaceImagePermissions(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error)
	DescribeWorkspaceImagePermissionsAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) *WorkSpacesDescribeWorkspaceImagePermissionsFuture

	DescribeWorkspaceImages(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagesInput) (*workspaces.DescribeWorkspaceImagesOutput, error)
	DescribeWorkspaceImagesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagesInput) *WorkSpacesDescribeWorkspaceImagesFuture

	DescribeWorkspaceSnapshots(ctx workflow.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) (*workspaces.DescribeWorkspaceSnapshotsOutput, error)
	DescribeWorkspaceSnapshotsAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) *WorkSpacesDescribeWorkspaceSnapshotsFuture

	DescribeWorkspaces(ctx workflow.Context, input *workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error)
	DescribeWorkspacesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspacesInput) *WorkSpacesDescribeWorkspacesFuture

	DescribeWorkspacesConnectionStatus(ctx workflow.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error)
	DescribeWorkspacesConnectionStatusAsync(ctx workflow.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) *WorkSpacesDescribeWorkspacesConnectionStatusFuture

	DisassociateConnectionAlias(ctx workflow.Context, input *workspaces.DisassociateConnectionAliasInput) (*workspaces.DisassociateConnectionAliasOutput, error)
	DisassociateConnectionAliasAsync(ctx workflow.Context, input *workspaces.DisassociateConnectionAliasInput) *WorkSpacesDisassociateConnectionAliasFuture

	DisassociateIpGroups(ctx workflow.Context, input *workspaces.DisassociateIpGroupsInput) (*workspaces.DisassociateIpGroupsOutput, error)
	DisassociateIpGroupsAsync(ctx workflow.Context, input *workspaces.DisassociateIpGroupsInput) *WorkSpacesDisassociateIpGroupsFuture

	ImportWorkspaceImage(ctx workflow.Context, input *workspaces.ImportWorkspaceImageInput) (*workspaces.ImportWorkspaceImageOutput, error)
	ImportWorkspaceImageAsync(ctx workflow.Context, input *workspaces.ImportWorkspaceImageInput) *WorkSpacesImportWorkspaceImageFuture

	ListAvailableManagementCidrRanges(ctx workflow.Context, input *workspaces.ListAvailableManagementCidrRangesInput) (*workspaces.ListAvailableManagementCidrRangesOutput, error)
	ListAvailableManagementCidrRangesAsync(ctx workflow.Context, input *workspaces.ListAvailableManagementCidrRangesInput) *WorkSpacesListAvailableManagementCidrRangesFuture

	MigrateWorkspace(ctx workflow.Context, input *workspaces.MigrateWorkspaceInput) (*workspaces.MigrateWorkspaceOutput, error)
	MigrateWorkspaceAsync(ctx workflow.Context, input *workspaces.MigrateWorkspaceInput) *WorkSpacesMigrateWorkspaceFuture

	ModifyAccount(ctx workflow.Context, input *workspaces.ModifyAccountInput) (*workspaces.ModifyAccountOutput, error)
	ModifyAccountAsync(ctx workflow.Context, input *workspaces.ModifyAccountInput) *WorkSpacesModifyAccountFuture

	ModifyClientProperties(ctx workflow.Context, input *workspaces.ModifyClientPropertiesInput) (*workspaces.ModifyClientPropertiesOutput, error)
	ModifyClientPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyClientPropertiesInput) *WorkSpacesModifyClientPropertiesFuture

	ModifySelfservicePermissions(ctx workflow.Context, input *workspaces.ModifySelfservicePermissionsInput) (*workspaces.ModifySelfservicePermissionsOutput, error)
	ModifySelfservicePermissionsAsync(ctx workflow.Context, input *workspaces.ModifySelfservicePermissionsInput) *WorkSpacesModifySelfservicePermissionsFuture

	ModifyWorkspaceAccessProperties(ctx workflow.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) (*workspaces.ModifyWorkspaceAccessPropertiesOutput, error)
	ModifyWorkspaceAccessPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) *WorkSpacesModifyWorkspaceAccessPropertiesFuture

	ModifyWorkspaceCreationProperties(ctx workflow.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) (*workspaces.ModifyWorkspaceCreationPropertiesOutput, error)
	ModifyWorkspaceCreationPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) *WorkSpacesModifyWorkspaceCreationPropertiesFuture

	ModifyWorkspaceProperties(ctx workflow.Context, input *workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error)
	ModifyWorkspacePropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspacePropertiesInput) *WorkSpacesModifyWorkspacePropertiesFuture

	ModifyWorkspaceState(ctx workflow.Context, input *workspaces.ModifyWorkspaceStateInput) (*workspaces.ModifyWorkspaceStateOutput, error)
	ModifyWorkspaceStateAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceStateInput) *WorkSpacesModifyWorkspaceStateFuture

	RebootWorkspaces(ctx workflow.Context, input *workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error)
	RebootWorkspacesAsync(ctx workflow.Context, input *workspaces.RebootWorkspacesInput) *WorkSpacesRebootWorkspacesFuture

	RebuildWorkspaces(ctx workflow.Context, input *workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error)
	RebuildWorkspacesAsync(ctx workflow.Context, input *workspaces.RebuildWorkspacesInput) *WorkSpacesRebuildWorkspacesFuture

	RegisterWorkspaceDirectory(ctx workflow.Context, input *workspaces.RegisterWorkspaceDirectoryInput) (*workspaces.RegisterWorkspaceDirectoryOutput, error)
	RegisterWorkspaceDirectoryAsync(ctx workflow.Context, input *workspaces.RegisterWorkspaceDirectoryInput) *WorkSpacesRegisterWorkspaceDirectoryFuture

	RestoreWorkspace(ctx workflow.Context, input *workspaces.RestoreWorkspaceInput) (*workspaces.RestoreWorkspaceOutput, error)
	RestoreWorkspaceAsync(ctx workflow.Context, input *workspaces.RestoreWorkspaceInput) *WorkSpacesRestoreWorkspaceFuture

	RevokeIpRules(ctx workflow.Context, input *workspaces.RevokeIpRulesInput) (*workspaces.RevokeIpRulesOutput, error)
	RevokeIpRulesAsync(ctx workflow.Context, input *workspaces.RevokeIpRulesInput) *WorkSpacesRevokeIpRulesFuture

	StartWorkspaces(ctx workflow.Context, input *workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error)
	StartWorkspacesAsync(ctx workflow.Context, input *workspaces.StartWorkspacesInput) *WorkSpacesStartWorkspacesFuture

	StopWorkspaces(ctx workflow.Context, input *workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error)
	StopWorkspacesAsync(ctx workflow.Context, input *workspaces.StopWorkspacesInput) *WorkSpacesStopWorkspacesFuture

	TerminateWorkspaces(ctx workflow.Context, input *workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error)
	TerminateWorkspacesAsync(ctx workflow.Context, input *workspaces.TerminateWorkspacesInput) *WorkSpacesTerminateWorkspacesFuture

	UpdateConnectionAliasPermission(ctx workflow.Context, input *workspaces.UpdateConnectionAliasPermissionInput) (*workspaces.UpdateConnectionAliasPermissionOutput, error)
	UpdateConnectionAliasPermissionAsync(ctx workflow.Context, input *workspaces.UpdateConnectionAliasPermissionInput) *WorkSpacesUpdateConnectionAliasPermissionFuture

	UpdateRulesOfIpGroup(ctx workflow.Context, input *workspaces.UpdateRulesOfIpGroupInput) (*workspaces.UpdateRulesOfIpGroupOutput, error)
	UpdateRulesOfIpGroupAsync(ctx workflow.Context, input *workspaces.UpdateRulesOfIpGroupInput) *WorkSpacesUpdateRulesOfIpGroupFuture

	UpdateWorkspaceImagePermission(ctx workflow.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) (*workspaces.UpdateWorkspaceImagePermissionOutput, error)
	UpdateWorkspaceImagePermissionAsync(ctx workflow.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) *WorkSpacesUpdateWorkspaceImagePermissionFuture
}

func NewClient() Client {
	return &stub{}
}
