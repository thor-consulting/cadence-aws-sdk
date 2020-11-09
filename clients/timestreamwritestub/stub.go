// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package timestreamwritestub

import (
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDatabaseFuture) Get(ctx workflow.Context) (*timestreamwrite.CreateDatabaseOutput, error) {
	var output timestreamwrite.CreateDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateTableFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateTableFuture) Get(ctx workflow.Context) (*timestreamwrite.CreateTableOutput, error) {
	var output timestreamwrite.CreateTableOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDatabaseFuture) Get(ctx workflow.Context) (*timestreamwrite.DeleteDatabaseOutput, error) {
	var output timestreamwrite.DeleteDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteTableFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteTableFuture) Get(ctx workflow.Context) (*timestreamwrite.DeleteTableOutput, error) {
	var output timestreamwrite.DeleteTableOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeDatabaseFuture) Get(ctx workflow.Context) (*timestreamwrite.DescribeDatabaseOutput, error) {
	var output timestreamwrite.DescribeDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeEndpointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeEndpointsFuture) Get(ctx workflow.Context) (*timestreamwrite.DescribeEndpointsOutput, error) {
	var output timestreamwrite.DescribeEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeTableFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeTableFuture) Get(ctx workflow.Context) (*timestreamwrite.DescribeTableOutput, error) {
	var output timestreamwrite.DescribeTableOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDatabasesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDatabasesFuture) Get(ctx workflow.Context) (*timestreamwrite.ListDatabasesOutput, error) {
	var output timestreamwrite.ListDatabasesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTablesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTablesFuture) Get(ctx workflow.Context) (*timestreamwrite.ListTablesOutput, error) {
	var output timestreamwrite.ListTablesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*timestreamwrite.ListTagsForResourceOutput, error) {
	var output timestreamwrite.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*timestreamwrite.TagResourceOutput, error) {
	var output timestreamwrite.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*timestreamwrite.UntagResourceOutput, error) {
	var output timestreamwrite.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDatabaseFuture) Get(ctx workflow.Context) (*timestreamwrite.UpdateDatabaseOutput, error) {
	var output timestreamwrite.UpdateDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateTableFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateTableFuture) Get(ctx workflow.Context) (*timestreamwrite.UpdateTableOutput, error) {
	var output timestreamwrite.UpdateTableOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WriteRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WriteRecordsFuture) Get(ctx workflow.Context) (*timestreamwrite.WriteRecordsOutput, error) {
	var output timestreamwrite.WriteRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatabase(ctx workflow.Context, input *timestreamwrite.CreateDatabaseInput) (*timestreamwrite.CreateDatabaseOutput, error) {
	var output timestreamwrite.CreateDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-CreateDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatabaseAsync(ctx workflow.Context, input *timestreamwrite.CreateDatabaseInput) *CreateDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-CreateDatabase", input)
	return &CreateDatabaseFuture{Future: future}
}

func (a *stub) CreateTable(ctx workflow.Context, input *timestreamwrite.CreateTableInput) (*timestreamwrite.CreateTableOutput, error) {
	var output timestreamwrite.CreateTableOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-CreateTable", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTableAsync(ctx workflow.Context, input *timestreamwrite.CreateTableInput) *CreateTableFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-CreateTable", input)
	return &CreateTableFuture{Future: future}
}

func (a *stub) DeleteDatabase(ctx workflow.Context, input *timestreamwrite.DeleteDatabaseInput) (*timestreamwrite.DeleteDatabaseOutput, error) {
	var output timestreamwrite.DeleteDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DeleteDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDatabaseAsync(ctx workflow.Context, input *timestreamwrite.DeleteDatabaseInput) *DeleteDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DeleteDatabase", input)
	return &DeleteDatabaseFuture{Future: future}
}

func (a *stub) DeleteTable(ctx workflow.Context, input *timestreamwrite.DeleteTableInput) (*timestreamwrite.DeleteTableOutput, error) {
	var output timestreamwrite.DeleteTableOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DeleteTable", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTableAsync(ctx workflow.Context, input *timestreamwrite.DeleteTableInput) *DeleteTableFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DeleteTable", input)
	return &DeleteTableFuture{Future: future}
}

func (a *stub) DescribeDatabase(ctx workflow.Context, input *timestreamwrite.DescribeDatabaseInput) (*timestreamwrite.DescribeDatabaseOutput, error) {
	var output timestreamwrite.DescribeDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DescribeDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDatabaseAsync(ctx workflow.Context, input *timestreamwrite.DescribeDatabaseInput) *DescribeDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DescribeDatabase", input)
	return &DescribeDatabaseFuture{Future: future}
}

func (a *stub) DescribeEndpoints(ctx workflow.Context, input *timestreamwrite.DescribeEndpointsInput) (*timestreamwrite.DescribeEndpointsOutput, error) {
	var output timestreamwrite.DescribeEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DescribeEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEndpointsAsync(ctx workflow.Context, input *timestreamwrite.DescribeEndpointsInput) *DescribeEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DescribeEndpoints", input)
	return &DescribeEndpointsFuture{Future: future}
}

func (a *stub) DescribeTable(ctx workflow.Context, input *timestreamwrite.DescribeTableInput) (*timestreamwrite.DescribeTableOutput, error) {
	var output timestreamwrite.DescribeTableOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DescribeTable", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTableAsync(ctx workflow.Context, input *timestreamwrite.DescribeTableInput) *DescribeTableFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-DescribeTable", input)
	return &DescribeTableFuture{Future: future}
}

func (a *stub) ListDatabases(ctx workflow.Context, input *timestreamwrite.ListDatabasesInput) (*timestreamwrite.ListDatabasesOutput, error) {
	var output timestreamwrite.ListDatabasesOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-ListDatabases", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatabasesAsync(ctx workflow.Context, input *timestreamwrite.ListDatabasesInput) *ListDatabasesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-ListDatabases", input)
	return &ListDatabasesFuture{Future: future}
}

func (a *stub) ListTables(ctx workflow.Context, input *timestreamwrite.ListTablesInput) (*timestreamwrite.ListTablesOutput, error) {
	var output timestreamwrite.ListTablesOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-ListTables", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTablesAsync(ctx workflow.Context, input *timestreamwrite.ListTablesInput) *ListTablesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-ListTables", input)
	return &ListTablesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *timestreamwrite.ListTagsForResourceInput) (*timestreamwrite.ListTagsForResourceOutput, error) {
	var output timestreamwrite.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *timestreamwrite.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *timestreamwrite.TagResourceInput) (*timestreamwrite.TagResourceOutput, error) {
	var output timestreamwrite.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *timestreamwrite.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *timestreamwrite.UntagResourceInput) (*timestreamwrite.UntagResourceOutput, error) {
	var output timestreamwrite.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *timestreamwrite.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateDatabase(ctx workflow.Context, input *timestreamwrite.UpdateDatabaseInput) (*timestreamwrite.UpdateDatabaseOutput, error) {
	var output timestreamwrite.UpdateDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-UpdateDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDatabaseAsync(ctx workflow.Context, input *timestreamwrite.UpdateDatabaseInput) *UpdateDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-UpdateDatabase", input)
	return &UpdateDatabaseFuture{Future: future}
}

func (a *stub) UpdateTable(ctx workflow.Context, input *timestreamwrite.UpdateTableInput) (*timestreamwrite.UpdateTableOutput, error) {
	var output timestreamwrite.UpdateTableOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-UpdateTable", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateTableAsync(ctx workflow.Context, input *timestreamwrite.UpdateTableInput) *UpdateTableFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-UpdateTable", input)
	return &UpdateTableFuture{Future: future}
}

func (a *stub) WriteRecords(ctx workflow.Context, input *timestreamwrite.WriteRecordsInput) (*timestreamwrite.WriteRecordsOutput, error) {
	var output timestreamwrite.WriteRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-WriteRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) WriteRecordsAsync(ctx workflow.Context, input *timestreamwrite.WriteRecordsInput) *WriteRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-timestreamwrite-WriteRecords", input)
	return &WriteRecordsFuture{Future: future}
}
