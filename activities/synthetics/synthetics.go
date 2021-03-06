// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package synthetics

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/synthetics"
	"github.com/aws/aws-sdk-go/service/synthetics/syntheticsiface"

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
	client syntheticsiface.SyntheticsAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := synthetics.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (syntheticsiface.SyntheticsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return synthetics.New(sess), nil
}

func (a *Activities) CreateCanary(ctx context.Context, input *synthetics.CreateCanaryInput) (*synthetics.CreateCanaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateCanaryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCanary(ctx context.Context, input *synthetics.DeleteCanaryInput) (*synthetics.DeleteCanaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteCanaryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCanaries(ctx context.Context, input *synthetics.DescribeCanariesInput) (*synthetics.DescribeCanariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCanariesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCanariesLastRun(ctx context.Context, input *synthetics.DescribeCanariesLastRunInput) (*synthetics.DescribeCanariesLastRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCanariesLastRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRuntimeVersions(ctx context.Context, input *synthetics.DescribeRuntimeVersionsInput) (*synthetics.DescribeRuntimeVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRuntimeVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCanary(ctx context.Context, input *synthetics.GetCanaryInput) (*synthetics.GetCanaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCanaryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCanaryRuns(ctx context.Context, input *synthetics.GetCanaryRunsInput) (*synthetics.GetCanaryRunsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCanaryRunsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *synthetics.ListTagsForResourceInput) (*synthetics.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartCanary(ctx context.Context, input *synthetics.StartCanaryInput) (*synthetics.StartCanaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartCanaryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopCanary(ctx context.Context, input *synthetics.StopCanaryInput) (*synthetics.StopCanaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopCanaryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *synthetics.TagResourceInput) (*synthetics.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *synthetics.UntagResourceInput) (*synthetics.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateCanary(ctx context.Context, input *synthetics.UpdateCanaryInput) (*synthetics.UpdateCanaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateCanaryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
