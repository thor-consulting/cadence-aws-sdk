// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package prometheusservice

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/prometheusservice"
	"github.com/aws/aws-sdk-go/service/prometheusservice/prometheusserviceiface"

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
	client prometheusserviceiface.PrometheusServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := prometheusservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (prometheusserviceiface.PrometheusServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return prometheusservice.New(sess), nil
}

func (a *Activities) CreateWorkspace(ctx context.Context, input *prometheusservice.CreateWorkspaceInput) (*prometheusservice.CreateWorkspaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateWorkspaceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteWorkspace(ctx context.Context, input *prometheusservice.DeleteWorkspaceInput) (*prometheusservice.DeleteWorkspaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.DeleteWorkspaceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWorkspace(ctx context.Context, input *prometheusservice.DescribeWorkspaceInput) (*prometheusservice.DescribeWorkspaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWorkspaceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListWorkspaces(ctx context.Context, input *prometheusservice.ListWorkspacesInput) (*prometheusservice.ListWorkspacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListWorkspacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateWorkspaceAlias(ctx context.Context, input *prometheusservice.UpdateWorkspaceAliasInput) (*prometheusservice.UpdateWorkspaceAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.UpdateWorkspaceAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
