// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/workspaces/workspacesiface"

	"github.com/banzaicloud/cadence-aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client workspacesiface.WorkSpacesAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := workspaces.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (workspacesiface.WorkSpacesAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return workspaces.New(sess), nil
}

func (a *Activities) AssociateConnectionAlias(ctx context.Context, input *workspaces.AssociateConnectionAliasInput) (*workspaces.AssociateConnectionAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateConnectionAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateIpGroups(ctx context.Context, input *workspaces.AssociateIpGroupsInput) (*workspaces.AssociateIpGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateIpGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AuthorizeIpRules(ctx context.Context, input *workspaces.AuthorizeIpRulesInput) (*workspaces.AuthorizeIpRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AuthorizeIpRulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CopyWorkspaceImage(ctx context.Context, input *workspaces.CopyWorkspaceImageInput) (*workspaces.CopyWorkspaceImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CopyWorkspaceImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateConnectionAlias(ctx context.Context, input *workspaces.CreateConnectionAliasInput) (*workspaces.CreateConnectionAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateConnectionAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateIpGroup(ctx context.Context, input *workspaces.CreateIpGroupInput) (*workspaces.CreateIpGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateIpGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateTags(ctx context.Context, input *workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateWorkspaces(ctx context.Context, input *workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateWorkspacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteConnectionAlias(ctx context.Context, input *workspaces.DeleteConnectionAliasInput) (*workspaces.DeleteConnectionAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteConnectionAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteIpGroup(ctx context.Context, input *workspaces.DeleteIpGroupInput) (*workspaces.DeleteIpGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteIpGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTags(ctx context.Context, input *workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteWorkspaceImage(ctx context.Context, input *workspaces.DeleteWorkspaceImageInput) (*workspaces.DeleteWorkspaceImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteWorkspaceImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeregisterWorkspaceDirectory(ctx context.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) (*workspaces.DeregisterWorkspaceDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeregisterWorkspaceDirectoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAccount(ctx context.Context, input *workspaces.DescribeAccountInput) (*workspaces.DescribeAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAccountWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAccountModifications(ctx context.Context, input *workspaces.DescribeAccountModificationsInput) (*workspaces.DescribeAccountModificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAccountModificationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeClientProperties(ctx context.Context, input *workspaces.DescribeClientPropertiesInput) (*workspaces.DescribeClientPropertiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeClientPropertiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeConnectionAliasPermissions(ctx context.Context, input *workspaces.DescribeConnectionAliasPermissionsInput) (*workspaces.DescribeConnectionAliasPermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeConnectionAliasPermissionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeConnectionAliases(ctx context.Context, input *workspaces.DescribeConnectionAliasesInput) (*workspaces.DescribeConnectionAliasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeConnectionAliasesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeIpGroups(ctx context.Context, input *workspaces.DescribeIpGroupsInput) (*workspaces.DescribeIpGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeIpGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTags(ctx context.Context, input *workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWorkspaceBundles(ctx context.Context, input *workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWorkspaceBundlesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWorkspaceDirectories(ctx context.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWorkspaceDirectoriesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWorkspaceImagePermissions(ctx context.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWorkspaceImagePermissionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWorkspaceImages(ctx context.Context, input *workspaces.DescribeWorkspaceImagesInput) (*workspaces.DescribeWorkspaceImagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWorkspaceImagesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWorkspaceSnapshots(ctx context.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWorkspaceSnapshotsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWorkspaces(ctx context.Context, input *workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWorkspacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWorkspacesConnectionStatus(ctx context.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWorkspacesConnectionStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateConnectionAlias(ctx context.Context, input *workspaces.DisassociateConnectionAliasInput) (*workspaces.DisassociateConnectionAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateConnectionAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateIpGroups(ctx context.Context, input *workspaces.DisassociateIpGroupsInput) (*workspaces.DisassociateIpGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateIpGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ImportWorkspaceImage(ctx context.Context, input *workspaces.ImportWorkspaceImageInput) (*workspaces.ImportWorkspaceImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ImportWorkspaceImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAvailableManagementCidrRanges(ctx context.Context, input *workspaces.ListAvailableManagementCidrRangesInput) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAvailableManagementCidrRangesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) MigrateWorkspace(ctx context.Context, input *workspaces.MigrateWorkspaceInput) (*workspaces.MigrateWorkspaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.MigrateWorkspaceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyAccount(ctx context.Context, input *workspaces.ModifyAccountInput) (*workspaces.ModifyAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyAccountWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyClientProperties(ctx context.Context, input *workspaces.ModifyClientPropertiesInput) (*workspaces.ModifyClientPropertiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyClientPropertiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifySelfservicePermissions(ctx context.Context, input *workspaces.ModifySelfservicePermissionsInput) (*workspaces.ModifySelfservicePermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifySelfservicePermissionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyWorkspaceAccessProperties(ctx context.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) (*workspaces.ModifyWorkspaceAccessPropertiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyWorkspaceAccessPropertiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyWorkspaceCreationProperties(ctx context.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) (*workspaces.ModifyWorkspaceCreationPropertiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyWorkspaceCreationPropertiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyWorkspaceProperties(ctx context.Context, input *workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyWorkspacePropertiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyWorkspaceState(ctx context.Context, input *workspaces.ModifyWorkspaceStateInput) (*workspaces.ModifyWorkspaceStateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyWorkspaceStateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RebootWorkspaces(ctx context.Context, input *workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RebootWorkspacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RebuildWorkspaces(ctx context.Context, input *workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RebuildWorkspacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterWorkspaceDirectory(ctx context.Context, input *workspaces.RegisterWorkspaceDirectoryInput) (*workspaces.RegisterWorkspaceDirectoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterWorkspaceDirectoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestoreWorkspace(ctx context.Context, input *workspaces.RestoreWorkspaceInput) (*workspaces.RestoreWorkspaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestoreWorkspaceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RevokeIpRules(ctx context.Context, input *workspaces.RevokeIpRulesInput) (*workspaces.RevokeIpRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RevokeIpRulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartWorkspaces(ctx context.Context, input *workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartWorkspacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopWorkspaces(ctx context.Context, input *workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopWorkspacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TerminateWorkspaces(ctx context.Context, input *workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TerminateWorkspacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateConnectionAliasPermission(ctx context.Context, input *workspaces.UpdateConnectionAliasPermissionInput) (*workspaces.UpdateConnectionAliasPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateConnectionAliasPermissionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRulesOfIpGroup(ctx context.Context, input *workspaces.UpdateRulesOfIpGroupInput) (*workspaces.UpdateRulesOfIpGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRulesOfIpGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateWorkspaceImagePermission(ctx context.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) (*workspaces.UpdateWorkspaceImagePermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateWorkspaceImagePermissionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
