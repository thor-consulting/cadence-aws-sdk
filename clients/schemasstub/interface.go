// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package schemasstub

import (
	"github.com/aws/aws-sdk-go/service/schemas"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateDiscoverer(ctx workflow.Context, input *schemas.CreateDiscovererInput) (*schemas.CreateDiscovererOutput, error)
	CreateDiscovererAsync(ctx workflow.Context, input *schemas.CreateDiscovererInput) *CreateDiscovererFuture

	CreateRegistry(ctx workflow.Context, input *schemas.CreateRegistryInput) (*schemas.CreateRegistryOutput, error)
	CreateRegistryAsync(ctx workflow.Context, input *schemas.CreateRegistryInput) *CreateRegistryFuture

	CreateSchema(ctx workflow.Context, input *schemas.CreateSchemaInput) (*schemas.CreateSchemaOutput, error)
	CreateSchemaAsync(ctx workflow.Context, input *schemas.CreateSchemaInput) *CreateSchemaFuture

	DeleteDiscoverer(ctx workflow.Context, input *schemas.DeleteDiscovererInput) (*schemas.DeleteDiscovererOutput, error)
	DeleteDiscovererAsync(ctx workflow.Context, input *schemas.DeleteDiscovererInput) *DeleteDiscovererFuture

	DeleteRegistry(ctx workflow.Context, input *schemas.DeleteRegistryInput) (*schemas.DeleteRegistryOutput, error)
	DeleteRegistryAsync(ctx workflow.Context, input *schemas.DeleteRegistryInput) *DeleteRegistryFuture

	DeleteResourcePolicy(ctx workflow.Context, input *schemas.DeleteResourcePolicyInput) (*schemas.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyAsync(ctx workflow.Context, input *schemas.DeleteResourcePolicyInput) *DeleteResourcePolicyFuture

	DeleteSchema(ctx workflow.Context, input *schemas.DeleteSchemaInput) (*schemas.DeleteSchemaOutput, error)
	DeleteSchemaAsync(ctx workflow.Context, input *schemas.DeleteSchemaInput) *DeleteSchemaFuture

	DeleteSchemaVersion(ctx workflow.Context, input *schemas.DeleteSchemaVersionInput) (*schemas.DeleteSchemaVersionOutput, error)
	DeleteSchemaVersionAsync(ctx workflow.Context, input *schemas.DeleteSchemaVersionInput) *DeleteSchemaVersionFuture

	DescribeCodeBinding(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) (*schemas.DescribeCodeBindingOutput, error)
	DescribeCodeBindingAsync(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) *DescribeCodeBindingFuture

	DescribeDiscoverer(ctx workflow.Context, input *schemas.DescribeDiscovererInput) (*schemas.DescribeDiscovererOutput, error)
	DescribeDiscovererAsync(ctx workflow.Context, input *schemas.DescribeDiscovererInput) *DescribeDiscovererFuture

	DescribeRegistry(ctx workflow.Context, input *schemas.DescribeRegistryInput) (*schemas.DescribeRegistryOutput, error)
	DescribeRegistryAsync(ctx workflow.Context, input *schemas.DescribeRegistryInput) *DescribeRegistryFuture

	DescribeSchema(ctx workflow.Context, input *schemas.DescribeSchemaInput) (*schemas.DescribeSchemaOutput, error)
	DescribeSchemaAsync(ctx workflow.Context, input *schemas.DescribeSchemaInput) *DescribeSchemaFuture

	ExportSchema(ctx workflow.Context, input *schemas.ExportSchemaInput) (*schemas.ExportSchemaOutput, error)
	ExportSchemaAsync(ctx workflow.Context, input *schemas.ExportSchemaInput) *ExportSchemaFuture

	GetCodeBindingSource(ctx workflow.Context, input *schemas.GetCodeBindingSourceInput) (*schemas.GetCodeBindingSourceOutput, error)
	GetCodeBindingSourceAsync(ctx workflow.Context, input *schemas.GetCodeBindingSourceInput) *GetCodeBindingSourceFuture

	GetDiscoveredSchema(ctx workflow.Context, input *schemas.GetDiscoveredSchemaInput) (*schemas.GetDiscoveredSchemaOutput, error)
	GetDiscoveredSchemaAsync(ctx workflow.Context, input *schemas.GetDiscoveredSchemaInput) *GetDiscoveredSchemaFuture

	GetResourcePolicy(ctx workflow.Context, input *schemas.GetResourcePolicyInput) (*schemas.GetResourcePolicyOutput, error)
	GetResourcePolicyAsync(ctx workflow.Context, input *schemas.GetResourcePolicyInput) *GetResourcePolicyFuture

	ListDiscoverers(ctx workflow.Context, input *schemas.ListDiscoverersInput) (*schemas.ListDiscoverersOutput, error)
	ListDiscoverersAsync(ctx workflow.Context, input *schemas.ListDiscoverersInput) *ListDiscoverersFuture

	ListRegistries(ctx workflow.Context, input *schemas.ListRegistriesInput) (*schemas.ListRegistriesOutput, error)
	ListRegistriesAsync(ctx workflow.Context, input *schemas.ListRegistriesInput) *ListRegistriesFuture

	ListSchemaVersions(ctx workflow.Context, input *schemas.ListSchemaVersionsInput) (*schemas.ListSchemaVersionsOutput, error)
	ListSchemaVersionsAsync(ctx workflow.Context, input *schemas.ListSchemaVersionsInput) *ListSchemaVersionsFuture

	ListSchemas(ctx workflow.Context, input *schemas.ListSchemasInput) (*schemas.ListSchemasOutput, error)
	ListSchemasAsync(ctx workflow.Context, input *schemas.ListSchemasInput) *ListSchemasFuture

	ListTagsForResource(ctx workflow.Context, input *schemas.ListTagsForResourceInput) (*schemas.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *schemas.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutCodeBinding(ctx workflow.Context, input *schemas.PutCodeBindingInput) (*schemas.PutCodeBindingOutput, error)
	PutCodeBindingAsync(ctx workflow.Context, input *schemas.PutCodeBindingInput) *PutCodeBindingFuture

	PutResourcePolicy(ctx workflow.Context, input *schemas.PutResourcePolicyInput) (*schemas.PutResourcePolicyOutput, error)
	PutResourcePolicyAsync(ctx workflow.Context, input *schemas.PutResourcePolicyInput) *PutResourcePolicyFuture

	SearchSchemas(ctx workflow.Context, input *schemas.SearchSchemasInput) (*schemas.SearchSchemasOutput, error)
	SearchSchemasAsync(ctx workflow.Context, input *schemas.SearchSchemasInput) *SearchSchemasFuture

	StartDiscoverer(ctx workflow.Context, input *schemas.StartDiscovererInput) (*schemas.StartDiscovererOutput, error)
	StartDiscovererAsync(ctx workflow.Context, input *schemas.StartDiscovererInput) *StartDiscovererFuture

	StopDiscoverer(ctx workflow.Context, input *schemas.StopDiscovererInput) (*schemas.StopDiscovererOutput, error)
	StopDiscovererAsync(ctx workflow.Context, input *schemas.StopDiscovererInput) *StopDiscovererFuture

	TagResource(ctx workflow.Context, input *schemas.TagResourceInput) (*schemas.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *schemas.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *schemas.UntagResourceInput) (*schemas.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *schemas.UntagResourceInput) *UntagResourceFuture

	UpdateDiscoverer(ctx workflow.Context, input *schemas.UpdateDiscovererInput) (*schemas.UpdateDiscovererOutput, error)
	UpdateDiscovererAsync(ctx workflow.Context, input *schemas.UpdateDiscovererInput) *UpdateDiscovererFuture

	UpdateRegistry(ctx workflow.Context, input *schemas.UpdateRegistryInput) (*schemas.UpdateRegistryOutput, error)
	UpdateRegistryAsync(ctx workflow.Context, input *schemas.UpdateRegistryInput) *UpdateRegistryFuture

	UpdateSchema(ctx workflow.Context, input *schemas.UpdateSchemaInput) (*schemas.UpdateSchemaOutput, error)
	UpdateSchemaAsync(ctx workflow.Context, input *schemas.UpdateSchemaInput) *UpdateSchemaFuture

	WaitUntilCodeBindingExists(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) error
	WaitUntilCodeBindingExistsAsync(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
