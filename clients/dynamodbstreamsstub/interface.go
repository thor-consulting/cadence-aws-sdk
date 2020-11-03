// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package dynamodbstreamsstub

import (
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeStream(ctx workflow.Context, input *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error)
	DescribeStreamAsync(ctx workflow.Context, input *dynamodbstreams.DescribeStreamInput) *DynamoDBStreamsDescribeStreamFuture

	GetRecords(ctx workflow.Context, input *dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error)
	GetRecordsAsync(ctx workflow.Context, input *dynamodbstreams.GetRecordsInput) *DynamoDBStreamsGetRecordsFuture

	GetShardIterator(ctx workflow.Context, input *dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error)
	GetShardIteratorAsync(ctx workflow.Context, input *dynamodbstreams.GetShardIteratorInput) *DynamoDBStreamsGetShardIteratorFuture

	ListStreams(ctx workflow.Context, input *dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error)
	ListStreamsAsync(ctx workflow.Context, input *dynamodbstreams.ListStreamsInput) *DynamoDBStreamsListStreamsFuture
}

func NewClient() Client {
	return &stub{}
}
