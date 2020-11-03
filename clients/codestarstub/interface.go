// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package codestarstub

import (
	"github.com/aws/aws-sdk-go/service/codestar"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateTeamMember(ctx workflow.Context, input *codestar.AssociateTeamMemberInput) (*codestar.AssociateTeamMemberOutput, error)
	AssociateTeamMemberAsync(ctx workflow.Context, input *codestar.AssociateTeamMemberInput) *CodeStarAssociateTeamMemberFuture

	CreateProject(ctx workflow.Context, input *codestar.CreateProjectInput) (*codestar.CreateProjectOutput, error)
	CreateProjectAsync(ctx workflow.Context, input *codestar.CreateProjectInput) *CodeStarCreateProjectFuture

	CreateUserProfile(ctx workflow.Context, input *codestar.CreateUserProfileInput) (*codestar.CreateUserProfileOutput, error)
	CreateUserProfileAsync(ctx workflow.Context, input *codestar.CreateUserProfileInput) *CodeStarCreateUserProfileFuture

	DeleteProject(ctx workflow.Context, input *codestar.DeleteProjectInput) (*codestar.DeleteProjectOutput, error)
	DeleteProjectAsync(ctx workflow.Context, input *codestar.DeleteProjectInput) *CodeStarDeleteProjectFuture

	DeleteUserProfile(ctx workflow.Context, input *codestar.DeleteUserProfileInput) (*codestar.DeleteUserProfileOutput, error)
	DeleteUserProfileAsync(ctx workflow.Context, input *codestar.DeleteUserProfileInput) *CodeStarDeleteUserProfileFuture

	DescribeProject(ctx workflow.Context, input *codestar.DescribeProjectInput) (*codestar.DescribeProjectOutput, error)
	DescribeProjectAsync(ctx workflow.Context, input *codestar.DescribeProjectInput) *CodeStarDescribeProjectFuture

	DescribeUserProfile(ctx workflow.Context, input *codestar.DescribeUserProfileInput) (*codestar.DescribeUserProfileOutput, error)
	DescribeUserProfileAsync(ctx workflow.Context, input *codestar.DescribeUserProfileInput) *CodeStarDescribeUserProfileFuture

	DisassociateTeamMember(ctx workflow.Context, input *codestar.DisassociateTeamMemberInput) (*codestar.DisassociateTeamMemberOutput, error)
	DisassociateTeamMemberAsync(ctx workflow.Context, input *codestar.DisassociateTeamMemberInput) *CodeStarDisassociateTeamMemberFuture

	ListProjects(ctx workflow.Context, input *codestar.ListProjectsInput) (*codestar.ListProjectsOutput, error)
	ListProjectsAsync(ctx workflow.Context, input *codestar.ListProjectsInput) *CodeStarListProjectsFuture

	ListResources(ctx workflow.Context, input *codestar.ListResourcesInput) (*codestar.ListResourcesOutput, error)
	ListResourcesAsync(ctx workflow.Context, input *codestar.ListResourcesInput) *CodeStarListResourcesFuture

	ListTagsForProject(ctx workflow.Context, input *codestar.ListTagsForProjectInput) (*codestar.ListTagsForProjectOutput, error)
	ListTagsForProjectAsync(ctx workflow.Context, input *codestar.ListTagsForProjectInput) *CodeStarListTagsForProjectFuture

	ListTeamMembers(ctx workflow.Context, input *codestar.ListTeamMembersInput) (*codestar.ListTeamMembersOutput, error)
	ListTeamMembersAsync(ctx workflow.Context, input *codestar.ListTeamMembersInput) *CodeStarListTeamMembersFuture

	ListUserProfiles(ctx workflow.Context, input *codestar.ListUserProfilesInput) (*codestar.ListUserProfilesOutput, error)
	ListUserProfilesAsync(ctx workflow.Context, input *codestar.ListUserProfilesInput) *CodeStarListUserProfilesFuture

	TagProject(ctx workflow.Context, input *codestar.TagProjectInput) (*codestar.TagProjectOutput, error)
	TagProjectAsync(ctx workflow.Context, input *codestar.TagProjectInput) *CodeStarTagProjectFuture

	UntagProject(ctx workflow.Context, input *codestar.UntagProjectInput) (*codestar.UntagProjectOutput, error)
	UntagProjectAsync(ctx workflow.Context, input *codestar.UntagProjectInput) *CodeStarUntagProjectFuture

	UpdateProject(ctx workflow.Context, input *codestar.UpdateProjectInput) (*codestar.UpdateProjectOutput, error)
	UpdateProjectAsync(ctx workflow.Context, input *codestar.UpdateProjectInput) *CodeStarUpdateProjectFuture

	UpdateTeamMember(ctx workflow.Context, input *codestar.UpdateTeamMemberInput) (*codestar.UpdateTeamMemberOutput, error)
	UpdateTeamMemberAsync(ctx workflow.Context, input *codestar.UpdateTeamMemberInput) *CodeStarUpdateTeamMemberFuture

	UpdateUserProfile(ctx workflow.Context, input *codestar.UpdateUserProfileInput) (*codestar.UpdateUserProfileOutput, error)
	UpdateUserProfileAsync(ctx workflow.Context, input *codestar.UpdateUserProfileInput) *CodeStarUpdateUserProfileFuture
}

func NewClient() Client {
	return &stub{}
}
