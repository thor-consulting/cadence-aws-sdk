// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package healthlakestub

import (
	"github.com/aws/aws-sdk-go/service/healthlake"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateFHIRDatastore(ctx workflow.Context, input *healthlake.CreateFHIRDatastoreInput) (*healthlake.CreateFHIRDatastoreOutput, error)
	CreateFHIRDatastoreAsync(ctx workflow.Context, input *healthlake.CreateFHIRDatastoreInput) *CreateFHIRDatastoreFuture

	DeleteFHIRDatastore(ctx workflow.Context, input *healthlake.DeleteFHIRDatastoreInput) (*healthlake.DeleteFHIRDatastoreOutput, error)
	DeleteFHIRDatastoreAsync(ctx workflow.Context, input *healthlake.DeleteFHIRDatastoreInput) *DeleteFHIRDatastoreFuture

	DescribeFHIRDatastore(ctx workflow.Context, input *healthlake.DescribeFHIRDatastoreInput) (*healthlake.DescribeFHIRDatastoreOutput, error)
	DescribeFHIRDatastoreAsync(ctx workflow.Context, input *healthlake.DescribeFHIRDatastoreInput) *DescribeFHIRDatastoreFuture

	DescribeFHIRExportJob(ctx workflow.Context, input *healthlake.DescribeFHIRExportJobInput) (*healthlake.DescribeFHIRExportJobOutput, error)
	DescribeFHIRExportJobAsync(ctx workflow.Context, input *healthlake.DescribeFHIRExportJobInput) *DescribeFHIRExportJobFuture

	DescribeFHIRImportJob(ctx workflow.Context, input *healthlake.DescribeFHIRImportJobInput) (*healthlake.DescribeFHIRImportJobOutput, error)
	DescribeFHIRImportJobAsync(ctx workflow.Context, input *healthlake.DescribeFHIRImportJobInput) *DescribeFHIRImportJobFuture

	ListFHIRDatastores(ctx workflow.Context, input *healthlake.ListFHIRDatastoresInput) (*healthlake.ListFHIRDatastoresOutput, error)
	ListFHIRDatastoresAsync(ctx workflow.Context, input *healthlake.ListFHIRDatastoresInput) *ListFHIRDatastoresFuture

	StartFHIRExportJob(ctx workflow.Context, input *healthlake.StartFHIRExportJobInput) (*healthlake.StartFHIRExportJobOutput, error)
	StartFHIRExportJobAsync(ctx workflow.Context, input *healthlake.StartFHIRExportJobInput) *StartFHIRExportJobFuture

	StartFHIRImportJob(ctx workflow.Context, input *healthlake.StartFHIRImportJobInput) (*healthlake.StartFHIRImportJobOutput, error)
	StartFHIRImportJobAsync(ctx workflow.Context, input *healthlake.StartFHIRImportJobInput) *StartFHIRImportJobFuture
}

func NewClient() Client {
	return &stub{}
}
