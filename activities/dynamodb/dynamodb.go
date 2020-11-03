// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

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
	client dynamodbiface.DynamoDBAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := dynamodb.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (dynamodbiface.DynamoDBAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return dynamodb.New(sess), nil
}

func (a *Activities) BatchGetItem(ctx context.Context, input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetItemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchWriteItem(ctx context.Context, input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchWriteItemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateBackup(ctx context.Context, input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateGlobalTable(ctx context.Context, input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGlobalTableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateTable(ctx context.Context, input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateTableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBackup(ctx context.Context, input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteItem(ctx context.Context, input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteItemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTable(ctx context.Context, input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeBackup(ctx context.Context, input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeContinuousBackups(ctx context.Context, input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeContinuousBackupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeContributorInsights(ctx context.Context, input *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeContributorInsightsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEndpoints(ctx context.Context, input *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEndpointsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGlobalTable(ctx context.Context, input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGlobalTableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGlobalTableSettings(ctx context.Context, input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGlobalTableSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLimits(ctx context.Context, input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLimitsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTable(ctx context.Context, input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTableReplicaAutoScaling(ctx context.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTableReplicaAutoScalingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTimeToLive(ctx context.Context, input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTimeToLiveWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetItem(ctx context.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetItemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBackups(ctx context.Context, input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBackupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListContributorInsights(ctx context.Context, input *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListContributorInsightsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGlobalTables(ctx context.Context, input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGlobalTablesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTables(ctx context.Context, input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTablesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsOfResource(ctx context.Context, input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsOfResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutItem(ctx context.Context, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutItemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) Query(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.QueryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestoreTableFromBackup(ctx context.Context, input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestoreTableFromBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestoreTableToPointInTime(ctx context.Context, input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestoreTableToPointInTimeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) Scan(ctx context.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ScanWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TransactGetItems(ctx context.Context, input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TransactGetItemsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TransactWriteItems(ctx context.Context, input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TransactWriteItemsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContinuousBackups(ctx context.Context, input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContinuousBackupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContributorInsights(ctx context.Context, input *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContributorInsightsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGlobalTable(ctx context.Context, input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGlobalTableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGlobalTableSettings(ctx context.Context, input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGlobalTableSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateItemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateTable(ctx context.Context, input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateTableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateTableReplicaAutoScaling(ctx context.Context, input *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateTableReplicaAutoScalingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateTimeToLive(ctx context.Context, input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateTimeToLiveWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilTableExists(ctx context.Context, input *dynamodb.DescribeTableInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilTableExistsWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilTableNotExists(ctx context.Context, input *dynamodb.DescribeTableInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilTableNotExistsWithContext(ctx, input, options...))
	})
}
