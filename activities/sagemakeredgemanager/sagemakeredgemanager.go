// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package sagemakeredgemanager

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemakeredgemanager"
	"github.com/aws/aws-sdk-go/service/sagemakeredgemanager/sagemakeredgemanageriface"

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
	client sagemakeredgemanageriface.SagemakerEdgeManagerAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := sagemakeredgemanager.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (sagemakeredgemanageriface.SagemakerEdgeManagerAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return sagemakeredgemanager.New(sess), nil
}

func (a *Activities) GetDeviceRegistration(ctx context.Context, input *sagemakeredgemanager.GetDeviceRegistrationInput) (*sagemakeredgemanager.GetDeviceRegistrationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDeviceRegistrationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendHeartbeat(ctx context.Context, input *sagemakeredgemanager.SendHeartbeatInput) (*sagemakeredgemanager.SendHeartbeatOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SendHeartbeatWithContext(ctx, input)

	return output, internal.EncodeError(err)
}