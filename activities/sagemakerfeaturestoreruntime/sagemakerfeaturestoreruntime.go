// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package sagemakerfeaturestoreruntime

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime"
	"github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime/sagemakerfeaturestoreruntimeiface"

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
	client sagemakerfeaturestoreruntimeiface.SageMakerFeatureStoreRuntimeAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := sagemakerfeaturestoreruntime.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (sagemakerfeaturestoreruntimeiface.SageMakerFeatureStoreRuntimeAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return sagemakerfeaturestoreruntime.New(sess), nil
}

func (a *Activities) DeleteRecord(ctx context.Context, input *sagemakerfeaturestoreruntime.DeleteRecordInput) (*sagemakerfeaturestoreruntime.DeleteRecordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRecordWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRecord(ctx context.Context, input *sagemakerfeaturestoreruntime.GetRecordInput) (*sagemakerfeaturestoreruntime.GetRecordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRecordWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRecord(ctx context.Context, input *sagemakerfeaturestoreruntime.PutRecordInput) (*sagemakerfeaturestoreruntime.PutRecordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRecordWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
