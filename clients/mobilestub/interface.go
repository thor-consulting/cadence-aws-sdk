// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package mobilestub

import (
	"github.com/aws/aws-sdk-go/service/mobile"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateProject(ctx workflow.Context, input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error)
	CreateProjectAsync(ctx workflow.Context, input *mobile.CreateProjectInput) *CreateProjectFuture

	DeleteProject(ctx workflow.Context, input *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error)
	DeleteProjectAsync(ctx workflow.Context, input *mobile.DeleteProjectInput) *DeleteProjectFuture

	DescribeBundle(ctx workflow.Context, input *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error)
	DescribeBundleAsync(ctx workflow.Context, input *mobile.DescribeBundleInput) *DescribeBundleFuture

	DescribeProject(ctx workflow.Context, input *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error)
	DescribeProjectAsync(ctx workflow.Context, input *mobile.DescribeProjectInput) *DescribeProjectFuture

	ExportBundle(ctx workflow.Context, input *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error)
	ExportBundleAsync(ctx workflow.Context, input *mobile.ExportBundleInput) *ExportBundleFuture

	ExportProject(ctx workflow.Context, input *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error)
	ExportProjectAsync(ctx workflow.Context, input *mobile.ExportProjectInput) *ExportProjectFuture

	ListBundles(ctx workflow.Context, input *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error)
	ListBundlesAsync(ctx workflow.Context, input *mobile.ListBundlesInput) *ListBundlesFuture

	ListProjects(ctx workflow.Context, input *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error)
	ListProjectsAsync(ctx workflow.Context, input *mobile.ListProjectsInput) *ListProjectsFuture

	UpdateProject(ctx workflow.Context, input *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error)
	UpdateProjectAsync(ctx workflow.Context, input *mobile.UpdateProjectInput) *UpdateProjectFuture
}

func NewClient() Client {
	return &stub{}
}
