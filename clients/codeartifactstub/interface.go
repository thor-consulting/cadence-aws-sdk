// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package codeartifactstub

import (
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateExternalConnection(ctx workflow.Context, input *codeartifact.AssociateExternalConnectionInput) (*codeartifact.AssociateExternalConnectionOutput, error)
	AssociateExternalConnectionAsync(ctx workflow.Context, input *codeartifact.AssociateExternalConnectionInput) *CodeArtifactAssociateExternalConnectionFuture

	CopyPackageVersions(ctx workflow.Context, input *codeartifact.CopyPackageVersionsInput) (*codeartifact.CopyPackageVersionsOutput, error)
	CopyPackageVersionsAsync(ctx workflow.Context, input *codeartifact.CopyPackageVersionsInput) *CodeArtifactCopyPackageVersionsFuture

	CreateDomain(ctx workflow.Context, input *codeartifact.CreateDomainInput) (*codeartifact.CreateDomainOutput, error)
	CreateDomainAsync(ctx workflow.Context, input *codeartifact.CreateDomainInput) *CodeArtifactCreateDomainFuture

	CreateRepository(ctx workflow.Context, input *codeartifact.CreateRepositoryInput) (*codeartifact.CreateRepositoryOutput, error)
	CreateRepositoryAsync(ctx workflow.Context, input *codeartifact.CreateRepositoryInput) *CodeArtifactCreateRepositoryFuture

	DeleteDomain(ctx workflow.Context, input *codeartifact.DeleteDomainInput) (*codeartifact.DeleteDomainOutput, error)
	DeleteDomainAsync(ctx workflow.Context, input *codeartifact.DeleteDomainInput) *CodeArtifactDeleteDomainFuture

	DeleteDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.DeleteDomainPermissionsPolicyInput) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error)
	DeleteDomainPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.DeleteDomainPermissionsPolicyInput) *CodeArtifactDeleteDomainPermissionsPolicyFuture

	DeletePackageVersions(ctx workflow.Context, input *codeartifact.DeletePackageVersionsInput) (*codeartifact.DeletePackageVersionsOutput, error)
	DeletePackageVersionsAsync(ctx workflow.Context, input *codeartifact.DeletePackageVersionsInput) *CodeArtifactDeletePackageVersionsFuture

	DeleteRepository(ctx workflow.Context, input *codeartifact.DeleteRepositoryInput) (*codeartifact.DeleteRepositoryOutput, error)
	DeleteRepositoryAsync(ctx workflow.Context, input *codeartifact.DeleteRepositoryInput) *CodeArtifactDeleteRepositoryFuture

	DeleteRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.DeleteRepositoryPermissionsPolicyInput) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error)
	DeleteRepositoryPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.DeleteRepositoryPermissionsPolicyInput) *CodeArtifactDeleteRepositoryPermissionsPolicyFuture

	DescribeDomain(ctx workflow.Context, input *codeartifact.DescribeDomainInput) (*codeartifact.DescribeDomainOutput, error)
	DescribeDomainAsync(ctx workflow.Context, input *codeartifact.DescribeDomainInput) *CodeArtifactDescribeDomainFuture

	DescribePackageVersion(ctx workflow.Context, input *codeartifact.DescribePackageVersionInput) (*codeartifact.DescribePackageVersionOutput, error)
	DescribePackageVersionAsync(ctx workflow.Context, input *codeartifact.DescribePackageVersionInput) *CodeArtifactDescribePackageVersionFuture

	DescribeRepository(ctx workflow.Context, input *codeartifact.DescribeRepositoryInput) (*codeartifact.DescribeRepositoryOutput, error)
	DescribeRepositoryAsync(ctx workflow.Context, input *codeartifact.DescribeRepositoryInput) *CodeArtifactDescribeRepositoryFuture

	DisassociateExternalConnection(ctx workflow.Context, input *codeartifact.DisassociateExternalConnectionInput) (*codeartifact.DisassociateExternalConnectionOutput, error)
	DisassociateExternalConnectionAsync(ctx workflow.Context, input *codeartifact.DisassociateExternalConnectionInput) *CodeArtifactDisassociateExternalConnectionFuture

	DisposePackageVersions(ctx workflow.Context, input *codeartifact.DisposePackageVersionsInput) (*codeartifact.DisposePackageVersionsOutput, error)
	DisposePackageVersionsAsync(ctx workflow.Context, input *codeartifact.DisposePackageVersionsInput) *CodeArtifactDisposePackageVersionsFuture

	GetAuthorizationToken(ctx workflow.Context, input *codeartifact.GetAuthorizationTokenInput) (*codeartifact.GetAuthorizationTokenOutput, error)
	GetAuthorizationTokenAsync(ctx workflow.Context, input *codeartifact.GetAuthorizationTokenInput) *CodeArtifactGetAuthorizationTokenFuture

	GetDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.GetDomainPermissionsPolicyInput) (*codeartifact.GetDomainPermissionsPolicyOutput, error)
	GetDomainPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.GetDomainPermissionsPolicyInput) *CodeArtifactGetDomainPermissionsPolicyFuture

	GetPackageVersionAsset(ctx workflow.Context, input *codeartifact.GetPackageVersionAssetInput) (*codeartifact.GetPackageVersionAssetOutput, error)
	GetPackageVersionAssetAsync(ctx workflow.Context, input *codeartifact.GetPackageVersionAssetInput) *CodeArtifactGetPackageVersionAssetFuture

	GetPackageVersionReadme(ctx workflow.Context, input *codeartifact.GetPackageVersionReadmeInput) (*codeartifact.GetPackageVersionReadmeOutput, error)
	GetPackageVersionReadmeAsync(ctx workflow.Context, input *codeartifact.GetPackageVersionReadmeInput) *CodeArtifactGetPackageVersionReadmeFuture

	GetRepositoryEndpoint(ctx workflow.Context, input *codeartifact.GetRepositoryEndpointInput) (*codeartifact.GetRepositoryEndpointOutput, error)
	GetRepositoryEndpointAsync(ctx workflow.Context, input *codeartifact.GetRepositoryEndpointInput) *CodeArtifactGetRepositoryEndpointFuture

	GetRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.GetRepositoryPermissionsPolicyInput) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error)
	GetRepositoryPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.GetRepositoryPermissionsPolicyInput) *CodeArtifactGetRepositoryPermissionsPolicyFuture

	ListDomains(ctx workflow.Context, input *codeartifact.ListDomainsInput) (*codeartifact.ListDomainsOutput, error)
	ListDomainsAsync(ctx workflow.Context, input *codeartifact.ListDomainsInput) *CodeArtifactListDomainsFuture

	ListPackageVersionAssets(ctx workflow.Context, input *codeartifact.ListPackageVersionAssetsInput) (*codeartifact.ListPackageVersionAssetsOutput, error)
	ListPackageVersionAssetsAsync(ctx workflow.Context, input *codeartifact.ListPackageVersionAssetsInput) *CodeArtifactListPackageVersionAssetsFuture

	ListPackageVersionDependencies(ctx workflow.Context, input *codeartifact.ListPackageVersionDependenciesInput) (*codeartifact.ListPackageVersionDependenciesOutput, error)
	ListPackageVersionDependenciesAsync(ctx workflow.Context, input *codeartifact.ListPackageVersionDependenciesInput) *CodeArtifactListPackageVersionDependenciesFuture

	ListPackageVersions(ctx workflow.Context, input *codeartifact.ListPackageVersionsInput) (*codeartifact.ListPackageVersionsOutput, error)
	ListPackageVersionsAsync(ctx workflow.Context, input *codeartifact.ListPackageVersionsInput) *CodeArtifactListPackageVersionsFuture

	ListPackages(ctx workflow.Context, input *codeartifact.ListPackagesInput) (*codeartifact.ListPackagesOutput, error)
	ListPackagesAsync(ctx workflow.Context, input *codeartifact.ListPackagesInput) *CodeArtifactListPackagesFuture

	ListRepositories(ctx workflow.Context, input *codeartifact.ListRepositoriesInput) (*codeartifact.ListRepositoriesOutput, error)
	ListRepositoriesAsync(ctx workflow.Context, input *codeartifact.ListRepositoriesInput) *CodeArtifactListRepositoriesFuture

	ListRepositoriesInDomain(ctx workflow.Context, input *codeartifact.ListRepositoriesInDomainInput) (*codeartifact.ListRepositoriesInDomainOutput, error)
	ListRepositoriesInDomainAsync(ctx workflow.Context, input *codeartifact.ListRepositoriesInDomainInput) *CodeArtifactListRepositoriesInDomainFuture

	ListTagsForResource(ctx workflow.Context, input *codeartifact.ListTagsForResourceInput) (*codeartifact.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codeartifact.ListTagsForResourceInput) *CodeArtifactListTagsForResourceFuture

	PutDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.PutDomainPermissionsPolicyInput) (*codeartifact.PutDomainPermissionsPolicyOutput, error)
	PutDomainPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.PutDomainPermissionsPolicyInput) *CodeArtifactPutDomainPermissionsPolicyFuture

	PutRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.PutRepositoryPermissionsPolicyInput) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error)
	PutRepositoryPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.PutRepositoryPermissionsPolicyInput) *CodeArtifactPutRepositoryPermissionsPolicyFuture

	TagResource(ctx workflow.Context, input *codeartifact.TagResourceInput) (*codeartifact.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codeartifact.TagResourceInput) *CodeArtifactTagResourceFuture

	UntagResource(ctx workflow.Context, input *codeartifact.UntagResourceInput) (*codeartifact.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codeartifact.UntagResourceInput) *CodeArtifactUntagResourceFuture

	UpdatePackageVersionsStatus(ctx workflow.Context, input *codeartifact.UpdatePackageVersionsStatusInput) (*codeartifact.UpdatePackageVersionsStatusOutput, error)
	UpdatePackageVersionsStatusAsync(ctx workflow.Context, input *codeartifact.UpdatePackageVersionsStatusInput) *CodeArtifactUpdatePackageVersionsStatusFuture

	UpdateRepository(ctx workflow.Context, input *codeartifact.UpdateRepositoryInput) (*codeartifact.UpdateRepositoryOutput, error)
	UpdateRepositoryAsync(ctx workflow.Context, input *codeartifact.UpdateRepositoryInput) *CodeArtifactUpdateRepositoryFuture
}

func NewClient() Client {
	return &stub{}
}
